// Copyright © 2018 Douglas Chimento <dchimento@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//
//
// Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
// Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
// Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
// Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
// Vestibulum commodo. Ut rhoncus gravida arcu.

package zapz

import (
	"go.uber.org/zap/zapcore"
	"fmt"
	"os"
	"time"
	"go.uber.org/zap"
	"github.com/dougEfresh/logzio-go"
)


// LogzTimeEncoder format to time.RFC3339Nano
func LogzTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(time.RFC3339Nano))
}

// Message needs to be the message key
var enCfg = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     LogzTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
}

const defaultType = "zap-logger"

// Zapz struc
type Zapz struct {
	lz *logzio.LogzioSender
	level zapcore.Level
	enCfg zapcore.EncoderConfig
	typ string
}

// New will create a zap logger compatible with logz
func New(token string, opts ...Option) (*zap.Logger, error)  {
	logz, err := logzio.New(token)
	if err != nil {
		return nil ,err
	}
	z := &Zapz{
		lz:    logz,
		level: zap.InfoLevel,
		enCfg: enCfg,
		typ: defaultType,
	}

	if len(opts) > 0 {
		for _, v := range opts {
			v.apply(z)
		}
	}

	en := zapcore.NewJSONEncoder(enCfg)
	return zap.New(zapcore.NewCore(en, z.lz, z.level)).With(zap.String("type", z.typ)), nil
}

// An Option configures a Logger.
type Option interface {
	apply(z *Zapz)
}

// SetLevel set the log level
func SetLevel(l  zapcore.Level) Option {
	return optionFunc(func(z *Zapz) {
		z.level = l
	})
}

// SetEncodeConfig set the encoder
func SetEncodeConfig(c  zapcore.EncoderConfig) Option {
	return optionFunc(func(z *Zapz) {
		z.enCfg = c
	})
}

// SetLogz use this logzsender
func SetLogz(c  *logzio.LogzioSender) Option {
	return optionFunc(func(z *Zapz) {
		z.lz = c
	})
}

// SetType setting log type zap.Field
func SetType(ty string) Option {
	return optionFunc(func(z *Zapz) {
		z.typ = ty
	})
}

// WithDebug enables
func WithDebug() Option {
	return optionFunc(func (z *Zapz){
		if e  :=  logzio.SetDebug(os.Stderr)(z.lz) ; e != nil {
			fmt.Fprintf(os.Stderr, "Error setting debug mode %s", e)
		}
	})
}

type optionFunc func(z *Zapz)

func (f optionFunc) apply(z *Zapz) {
	f(z)
}
