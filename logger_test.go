// Very simple drop-in logger with log level control
//
// Only supports log levels: ERROR, WARNING, INFO, TRACE.
// Default log level is ERROR, so only Error, Errorf, Errorln will print to stdout by default.
//
// Only print to stdout.
//
// Call 'InitLogger()' in top level function(normally it is 'main()') to change the default value.

package lgl

import (
	"bytes"
	"strings"
	"testing"
)

func TestTrace(t *testing.T) {
	type args struct {
		level  LogLevel
		input  string
		output string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{LTrace, "aa", "aa\n"}},
		{"2", args{LDebug, "aa", ""}},
		{"3", args{LInfo, "aa", ""}},
		{"4", args{LWarning, "aa", ""}},
		{"5", args{LError, "aa", ""}},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		Set(tt.args.level, &buf, &buf)
		Trace(tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Trace() output %s, want %s", o, tt.args.output)
		}
		Tracef("%s", tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Tracef() output %s, want %s", o, tt.args.output)
		}
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		level  LogLevel
		input  string
		output string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{LTrace, "aa", "aa\n"}},
		{"2", args{LDebug, "aa", "aa\n"}},
		{"3", args{LInfo, "aa", ""}},
		{"4", args{LWarning, "aa", ""}},
		{"5", args{LError, "aa", ""}},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		Set(tt.args.level, &buf, &buf)
		Debug(tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Debug() output %s, want %s", o, tt.args.output)
		}
		Debugf("%s", tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Debugf() output %s, want %s", o, tt.args.output)
		}
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		level  LogLevel
		input  string
		output string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{LTrace, "aa", "aa\n"}},
		{"2", args{LDebug, "aa", "aa\n"}},
		{"3", args{LInfo, "aa", "aa\n"}},
		{"4", args{LWarning, "aa", ""}},
		{"5", args{LError, "aa", ""}},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		Set(tt.args.level, &buf, &buf)
		Info(tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Info() output %s, want %s", o, tt.args.output)
		}
		Infof("%s", tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Infof() output %s, want %s", o, tt.args.output)
		}
	}
}

func TestWarning(t *testing.T) {
	type args struct {
		level  LogLevel
		input  string
		output string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{LTrace, "aa", "aa\n"}},
		{"2", args{LDebug, "aa", "aa\n"}},
		{"3", args{LInfo, "aa", "aa\n"}},
		{"4", args{LWarning, "aa", "aa\n"}},
		{"5", args{LError, "aa", ""}},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		Set(tt.args.level, &buf, &buf)
		Warning(tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Warning() output %s, want %s", o, tt.args.output)
		}
		Warningf("%s", tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Warningf() output %s, want %s", o, tt.args.output)
		}
	}
}

func TestError(t *testing.T) {
	type args struct {
		level  LogLevel
		input  string
		output string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{LTrace, "aa", "aa\n"}},
		{"2", args{LDebug, "aa", "aa\n"}},
		{"3", args{LInfo, "aa", "aa\n"}},
		{"4", args{LWarning, "aa", "aa\n"}},
		{"5", args{LError, "aa", "aa\n"}},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		Set(tt.args.level, &buf, &buf)
		Error(tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Error() output %s, want %s", o, tt.args.output)
		}
		Errorf("%s", tt.args.input)
		if o := buf.String(); !strings.HasSuffix(o, tt.args.output) {
			t.Errorf("Errorf() output %s, want %s", o, tt.args.output)
		}
	}
}
