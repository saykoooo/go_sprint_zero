package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
        cases := []struct {
            name string
            value string
            want float64
            err error
        }{
            {
                name: "10",
                value: "4 + 3 * 2 / 1",
                want: 10.0,
                err: nil,
            },
            {
                name: "-27",
                value: "3 + 5 * (2 - 8)",
                want: -27.0,
                err: nil,
            },
            {
                name: "div by zero",
                value: "1/0",
                want: 0,
                err: ErrDivByZero,
            },
            {
                name: "bad expression",
                value: "()1)/(2)",
                want: 0,
                err: ErrInvalidExpression,
            },
        }
        for _, tc := range cases {
            tc := tc
            t.Run(tc.name, func(t *testing.T) {
                got, err := Calc(tc.value)
                if got != tc.want {
                    t.Errorf("Expected Calc(%v) = %v; want %v", tc.value, got, tc.want)
                }
                if err != tc.err {
                    t.Errorf("Expected error to be %v, but got %v for input %q", tc.err, err, tc.value)
                }
            })
        }
}
