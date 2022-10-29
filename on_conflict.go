package qbuilder

import (
	"strings"
)

type OnConflict struct {
	Insert           *InsertBuilder
	Target           string
	TargetConstraint string
	doUpdate         map[string]string
	doNothing        bool
}

func (o *OnConflict) DoUpdate(row map[string]string) *InsertBuilder {
	o.doNothing = false
	o.doUpdate = row
	return o.Insert
}

func (o *OnConflict) OnConstraint(c string) *OnConflict {
	o.Target = ""
	o.TargetConstraint = c
	return o
}

func (o *OnConflict) OnTarget(t string) *OnConflict {
	o.Target = t
	o.TargetConstraint = ""
	return o
}

func (o *OnConflict) DoNothing() *InsertBuilder {
	o.doNothing = true
	o.doUpdate = nil
	return o.Insert
}

func (o *OnConflict) String() string {
	if o == nil {
		return ""
	}
	if o.doNothing {
		return "ON CONFLICT DO NOTHING"
	}
	if len(o.doUpdate) > 0 && o.Target != "" {
		var set string
		for f, v := range o.doUpdate {
			set += "SET \"" + f + "\" = " + v + ","
		}
		return "ON CONFLICT " + o.Target + " DO UPDATE " + strings.Trim(set, ",")
	}
	if len(o.doUpdate) > 0 && o.TargetConstraint != "" {
		var set = "SET "
		for f, v := range o.doUpdate {
			set += "" + f + "\" = " + v + ", "
		}
		return "ON CONFLICT ON CONSTRAINT " + o.TargetConstraint + " DO UPDATE " + strings.Trim(set, ",")
	}
	return ""
}
