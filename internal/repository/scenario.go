package repository

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository/dao"
	"time"
)

type ScenarioRepo interface {
	GetScenarioByCode(ctx context.Context, code model.ScenarioCode) (model.Scenario, error)
	GetTriggerRuleByScenarioId(ctx context.Context, sid int64) (model.TriggerRule, error)
	GetSendingFrequencyByScenarioId(ctx context.Context, sid int64) (model.SendingFrequency, error)
}

type scenarioRepo struct {
	scenarioDAO dao.ScenarioDAO
}

func (repo *scenarioRepo) GetScenarioByCode(ctx context.Context,
	code model.ScenarioCode) (model.Scenario, error) {
	scenario, err := repo.scenarioDAO.GetScenarioByCode(ctx, string(code))
	if err != nil {
		return model.Scenario{}, err
	}

	return repo.ScenarioToModel(scenario), nil
}

func (repo *scenarioRepo) GetTriggerRuleByScenarioId(ctx context.Context, sid int64) (model.TriggerRule, error) {
	triggerRule, err := repo.scenarioDAO.GetTriggerRuleByScenarioId(ctx, sid)
	if err != nil {
		return model.TriggerRule{}, err
	}

	return repo.TriggerRuleToModel(triggerRule), nil
}

func (repo *scenarioRepo) GetSendingFrequencyByScenarioId(ctx context.Context, sid int64) (model.SendingFrequency, error) {
	sendingFrequency, err := repo.scenarioDAO.GetFrequencyByScenarioId(ctx, sid)
	if err != nil {
		return model.SendingFrequency{}, err
	}

	return repo.FrequencyToModel(sendingFrequency), nil
}

func NewScenarioRepo(scenarioDAO dao.ScenarioDAO) ScenarioRepo {
	return &scenarioRepo{
		scenarioDAO: scenarioDAO,
	}
}

func (repo *scenarioRepo) ScenarioToModel(scenario dao.Scenario) model.Scenario {
	return model.Scenario{
		Id:        scenario.Id,
		App:       scenario.App,
		Code:      model.ScenarioCode(scenario.Code),
		Name:      scenario.Name,
		Objective: scenario.Objective,
		IsActive:  scenario.IsActive,
		CreatedAt: time.UnixMilli(scenario.CreatedAt),
		UpdatedAt: time.UnixMilli(scenario.UpdatedAt),
	}
}

func (repo *scenarioRepo) TriggerRuleToModel(triggerRule dao.TriggerRule) model.TriggerRule {
	return model.TriggerRule{
		Id:         triggerRule.Id,
		ScenarioId: triggerRule.ScenarioId,
		Type:       model.TriggerRuleType(triggerRule.Type),
		Params:     triggerRule.Params,
		CreatedAt:  time.UnixMilli(triggerRule.CreatedAt),
		UpdatedAt:  time.UnixMilli(triggerRule.UpdatedAt),
	}
}

func (repo *scenarioRepo) FrequencyToModel(frequency dao.SendingFrequency) model.SendingFrequency {
	return model.SendingFrequency{
		Id:         frequency.Id,
		ScenarioId: frequency.ScenarioId,
		Type:       model.SendingFrequencyType(frequency.Type),
		Params:     frequency.Params,
		CreatedAt:  time.UnixMilli(frequency.CreatedAt),
		UpdatedAt:  time.UnixMilli(frequency.UpdatedAt),
	}
}
