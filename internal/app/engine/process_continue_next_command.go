package engine

import (
	"fmt"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/geom"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const distanceToBallDuringPenalty = 1.0

func (e *Engine) createNextCommandContinueAction(actionType ContinueAction_Type) *ContinueAction {
	var actionState ContinueAction_State
	var continuationIssues = e.findIssuesForContinuation(*e.currentState.Command.Type)

	var lastReadyAt *timestamppb.Timestamp
	var lastContinueAction = e.LastContinueAction(actionType)
	if lastContinueAction != nil {
		lastReadyAt = lastContinueAction.ReadyAt
	}

	var readyAt *timestamppb.Timestamp
	if len(continuationIssues) > 0 {
		actionState = ContinueAction_BLOCKED
		readyAt = nil
	} else {
		if lastReadyAt == nil {
			readyAt = timestamppb.New(e.timeProvider().Add(e.gameConfig.PreparationTimeBeforeResume))
		} else {
			readyAt = lastReadyAt
		}
		readyAt = maxTime(readyAt, e.currentState.ReadyContinueTime)
		preparationTimeLeft := readyAt.AsTime().Sub(e.timeProvider())
		if preparationTimeLeft > 0 {
			actionState = ContinueAction_WAITING
			continuationIssues = append(continuationIssues, fmt.Sprintf("%.1f s left for preparation", preparationTimeLeft.Seconds()))
		} else if actionType == ContinueAction_NEXT_COMMAND {
			actionState = ContinueAction_READY_AUTO
		} else {
			actionState = ContinueAction_READY_MANUAL
		}
	}

	var forTeam *state.Team
	if e.currentState.NextCommand != nil {
		forTeam = e.currentState.NextCommand.ForTeam
	} else {
		forTeam = state.NewTeam(state.Team_UNKNOWN)
	}

	return &ContinueAction{
		Type:               &actionType,
		ForTeam:            forTeam,
		State:              &actionState,
		ContinuationIssues: continuationIssues,
		ReadyAt:            readyAt,
	}
}

func (e *Engine) LastContinueAction(actionType ContinueAction_Type) *ContinueAction {
	for _, action := range e.gcState.ContinueActions {
		if *action.Type == actionType {
			return action
		}
	}
	return nil
}

func maxTime(t1, t2 *timestamppb.Timestamp) *timestamppb.Timestamp {
	if t1.AsTime().After(t2.AsTime()) {
		return t1
	}
	return t2
}

func (e *Engine) findIssuesForContinuation(command state.Command_Type) []string {

	if e.trackerStateGc.Ball == nil {
		// No tracking data, can not check for issues
		return []string{}
	}

	switch command {
	case state.Command_KICKOFF:
		return e.readyToContinueKickoff()
	case state.Command_PENALTY:
		return e.readyToContinuePenalty()
	case state.Command_STOP:
		return e.readyToContinueFromStop()
	}

	return []string{}
}

func (e *Engine) readyToContinueKickoff() (issues []string) {
	ballToCenterCircleDist := e.trackerStateGc.Ball.Pos.ToVector2().Length() - e.getGeometry().CenterCircleRadius
	if ballToCenterCircleDist > 0 {
		issues = append(issues, fmt.Sprintf("Ball is %.1f m away from center circle", ballToCenterCircleDist))
	}

	var badRobots []*Robot
	for _, robot := range e.trackerStateGc.Robots {
		if *e.currentState.TeamState[robot.Id.Team.String()].OnPositiveHalf {
			if *robot.Pos.X < robotRadius {
				badRobots = append(badRobots, robot)
			}
		} else if *robot.Pos.X > -robotRadius {
			badRobots = append(badRobots, robot)
		}
	}

	if len(badRobots) < 3 {
		for _, robot := range badRobots {
			issues = append(issues, fmt.Sprintf("Robot %s is not in its own half", robot.Id.PrettyString()))
		}
	} else {
		issues = append(issues, "Multiple robots are not in their own half")
	}

	return issues
}

func (e *Engine) readyToContinuePenalty() (issues []string) {
	keeperId := e.penaltyKeeperId()
	if keeperId == nil {
		issues = append(issues, "There is no keeper")
		return
	}
	keeperPos := e.robotPos(keeperId)
	if keeperPos == nil {
		issues = append(issues, "Keeper position is unknown")
	} else if !e.posInsideGoal(keeperPos) {
		issues = append(issues, "Keeper is not inside goal")
	}

	keeperTeamInfo := e.currentState.TeamState[keeperId.Team.String()]
	ballPos := e.trackerStateGc.Ball.Pos

	numAttackersInFrontOfBall := 0
	var badRobots []*Robot
	for _, robot := range e.trackerStateGc.Robots {
		if *robot.Id.Id == *keeperId.Id && *robot.Id.Team == *keeperId.Team {
			// it's the keeper
			continue
		}
		if *keeperTeamInfo.OnPositiveHalf &&
			*robot.Pos.X < *ballPos.X-distanceToBallDuringPenalty {
			continue
		} else if !*keeperTeamInfo.OnPositiveHalf &&
			*robot.Pos.X > *ballPos.X+distanceToBallDuringPenalty {
			continue
		}
		if *robot.Id.Team == *keeperId.Team {
			badRobots = append(badRobots, robot)
		} else if numAttackersInFrontOfBall >= 1 {
			badRobots = append(badRobots, robot)
		} else {
			numAttackersInFrontOfBall++
		}
	}

	if len(badRobots) < 3 {
		for _, robot := range badRobots {
			issues = append(issues, fmt.Sprintf(
				"Robot %s does not keep required distance to ball", robot.Id.PrettyString()))
		}
	} else {
		issues = append(issues, "Multiple robots do not keep required distance to ball")
	}
	return
}

func (e *Engine) readyToContinueFromStop() (issues []string) {
	if e.tooManyRobots(state.Team_YELLOW) {
		issues = append(issues, "Yellow team has too many robots")
	}
	if e.tooManyRobots(state.Team_BLUE) {
		issues = append(issues, "Blue team has too many robots")
	}
	radius := e.gameConfig.DistanceToBallInStop + robotRadius - distanceThreshold
	robotNearBall := e.findRobotInsideRadius(e.trackerStateGc.Robots, e.trackerStateGc.Ball.Pos.ToVector2(), radius)
	if robotNearBall != nil {
		issues = append(issues, fmt.Sprintf("Robot %s is too close to ball", robotNearBall.Id.PrettyString()))
	}
	if e.currentState.PlacementPos != nil {
		ballToPlacementPosDist := e.currentState.PlacementPos.DistanceTo(e.trackerStateGc.Ball.Pos.ToVector2())
		if ballToPlacementPosDist > e.gameConfig.BallPlacementRequiredDistance {
			issues = append(issues, fmt.Sprintf("Ball is %.2fm (>%.2fm) away from placement pos",
				ballToPlacementPosDist, e.gameConfig.BallPlacementRequiredDistance))
		}
	}
	if !e.trackerStateGc.Ball.IsSteady() {
		issues = append(issues, "Ball position is not steady")
	}
	return
}

func (e *Engine) penaltyKeeperId() *state.RobotId {
	forTeam := e.currentState.Command.ForTeam.Opposite()
	teamInfo := e.currentState.TeamState[forTeam.String()]
	keeperId := uint32(*teamInfo.Goalkeeper)
	return &state.RobotId{
		Id:   &keeperId,
		Team: &forTeam,
	}
}

func (e *Engine) robotPos(robotId *state.RobotId) *geom.Vector2 {
	for _, robot := range e.trackerStateGc.Robots {
		if *robot.Id.Id == *robotId.Id && *robot.Id.Team == *robotId.Team {
			return robot.Pos
		}
	}
	return nil
}

func (e *Engine) posInsideGoal(pos *geom.Vector2) bool {
	goalTeam := e.currentState.Command.ForTeam.Opposite()
	teamInfo := e.currentState.TeamState[goalTeam.String()]
	goalCenter := geom.GoalCenter(e.getGeometry(), *teamInfo.OnPositiveHalf)
	goalArea := geom.NewRectangleFromCenter(goalCenter, robotRadius*2, e.getGeometry().GoalWidth)
	return goalArea.IsPointInside(pos)
}

func (e *Engine) tooManyRobots(team state.Team) bool {
	maxAllowed := *e.currentState.TeamState[team.String()].MaxAllowedBots
	current := e.trackerStateGc.NumTeamRobots(team)
	return current > maxAllowed
}
