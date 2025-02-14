// SPDX-License-Identifier: GPL-3.0-or-later

package khulnasoftapi

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_CHART(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.CHART(
		"",
		"id",
		"name",
		"title",
		"units",
		"family",
		"context",
		"line",
		1,
		1,
		"",
		"plugin",
		"module",
	)

	assert.Equal(
		t,
		"CHART '.id' 'name' 'title' 'units' 'family' 'context' 'line' '1' '1' '' 'plugin' 'module'\n",
		buf.String(),
	)
}

func TestAPI_DIMENSION(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.DIMENSION(
		"id",
		"name",
		"absolute",
		1,
		1,
		"",
	)

	assert.Equal(
		t,
		"DIMENSION 'id' 'name' 'absolute' '1' '1' ''\n",
		buf.String(),
	)
}

func TestAPI_BEGIN(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.BEGIN(
		"typeID",
		"id",
		0,
	)

	assert.Equal(
		t,
		"BEGIN 'typeID.id'\n",
		buf.String(),
	)

	buf.Reset()

	_ = a.BEGIN(
		"typeID",
		"id",
		1,
	)

	assert.Equal(
		t,
		"BEGIN 'typeID.id' 1\n",
		buf.String(),
	)
}

func TestAPI_SET(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.SET("id", 100)

	assert.Equal(
		t,
		"SET 'id' = 100\n",
		buf.String(),
	)
}

func TestAPI_SETEMPTY(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.SETEMPTY("id")

	assert.Equal(
		t,
		"SET 'id' = \n",
		buf.String(),
	)
}

func TestAPI_VARIABLE(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.VARIABLE("id", 100)

	assert.Equal(
		t,
		"VARIABLE CHART 'id' = 100\n",
		buf.String(),
	)
}

func TestAPI_END(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.END()

	assert.Equal(
		t,
		"END\n\n",
		buf.String(),
	)
}

func TestAPI_CLABEL(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.CLABEL("key", "value", 1)

	assert.Equal(
		t,
		"CLABEL 'key' 'value' '1'\n",
		buf.String(),
	)
}

func TestAPI_CLABELCOMMIT(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.CLABELCOMMIT()

	assert.Equal(
		t,
		"CLABEL_COMMIT\n",
		buf.String(),
	)
}

func TestAPI_DISABLE(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.DISABLE()

	assert.Equal(
		t,
		"DISABLE\n",
		buf.String(),
	)
}

func TestAPI_EMPTYLINE(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.EMPTYLINE()

	assert.Equal(
		t,
		"\n",
		buf.String(),
	)
}

func TestAPI_HOST(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.HOST("guid")

	assert.Equal(
		t,
		"HOST 'guid'\n\n",
		buf.String(),
	)
}

func TestAPI_HOSTDEFINE(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.HOSTDEFINE("guid", "hostname")

	assert.Equal(
		t,
		"HOST_DEFINE 'guid' 'hostname'\n",
		buf.String(),
	)
}

func TestAPI_HOSTLABEL(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.HOSTLABEL("name", "value")

	assert.Equal(
		t,
		"HOST_LABEL 'name' 'value'\n",
		buf.String(),
	)
}

func TestAPI_HOSTDEFINEEND(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.HOSTDEFINEEND()

	assert.Equal(
		t,
		"HOST_DEFINE_END\n\n",
		buf.String(),
	)
}

func TestAPI_HOSTINFO(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.HOSTINFO("guid", "hostname", map[string]string{"label1": "value1"})

	assert.Equal(
		t,
		`HOST_DEFINE 'guid' 'hostname'
HOST_LABEL 'label1' 'value1'
HOST_DEFINE_END

`,
		buf.String(),
	)
}

func TestAPI_DynCfgEnable(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.DynCfgEnable("plugin")

	assert.Equal(
		t,
		"DYNCFG_ENABLE 'plugin'\n\n",
		buf.String(),
	)
}

func TestAPI_DynCfgReset(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.DynCfgReset()

	assert.Equal(
		t,
		"DYNCFG_RESET\n",
		buf.String(),
	)
}

func TestAPI_DyncCfgRegisterModule(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.DyncCfgRegisterModule("module")

	assert.Equal(
		t,
		"DYNCFG_REGISTER_MODULE 'module' job_array\n\n",
		buf.String(),
	)
}

func TestAPI_DynCfgRegisterJob(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.DynCfgRegisterJob("module", "job", "type")

	assert.Equal(
		t,
		"DYNCFG_REGISTER_JOB 'module' 'job' 'type' 0\n\n",
		buf.String(),
	)
}

func TestAPI_DynCfgReportJobStatus(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.DynCfgReportJobStatus("module", "job", "status", "reason")

	assert.Equal(
		t,
		"REPORT_JOB_STATUS 'module' 'job' 'status' 0 'reason'\n\n",
		buf.String(),
	)
}

func TestAPI_FunctionResultSuccess(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.FunctionResultSuccess("uid", "contentType", "payload")

	assert.Equal(
		t,
		`FUNCTION_RESULT_BEGIN uid 1 contentType 0
payload
FUNCTION_RESULT_END

`,
		buf.String(),
	)
}

func TestAPI_FunctionResultReject(t *testing.T) {
	buf := &bytes.Buffer{}
	a := API{Writer: buf}

	_ = a.FunctionResultReject("uid", "contentType", "payload")

	assert.Equal(
		t,
		`FUNCTION_RESULT_BEGIN uid 0 contentType 0
payload
FUNCTION_RESULT_END

`,
		buf.String(),
	)
}
