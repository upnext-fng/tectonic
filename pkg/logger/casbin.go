package logger

import (
	"fmt"
	"strings"
)

func (l *Logger) EnableLog(bool) {}

func (l *Logger) IsEnabled() bool {
	return true
}

func (l *Logger) LogModel(model [][]string) {
	for _, v := range model {
		l.WithFields(map[string]any{"model": v}).Info(" RBAC models")
	}
}

func (l *Logger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	var reqStr strings.Builder

	for i, val := range request {
		if i != len(request)-1 {
			reqStr.WriteString(fmt.Sprintf("%v, ", val))
		} else {
			reqStr.WriteString(fmt.Sprintf("%v", val))
		}
	}

	l.WithFields(map[string]any{
		//"matcher":    matcher,
		"explains":   explains,
		"request":    reqStr.String(),
		"hit_policy": result,
	}).Info("RBAC enforce policy")
}

func (l *Logger) LogPolicy(policy map[string][][]string) {
	for k, v := range policy {
		l.WithFields(map[string]any{k: v}).Info("RBAC policies")
	}
}

func (l *Logger) LogRole(roles []string) {
	l.WithFields(map[string]any{"roles": roles}).Info("RBAC roles")
}

func (l *Logger) LogError(err error, msg ...string) {
	l.WithErr(err).Error("RBAC error", msg)
}
