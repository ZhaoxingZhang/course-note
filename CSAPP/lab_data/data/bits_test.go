package bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func TestBits(t *testing.T) {
	t.Run("tmin", func(t *testing.T) {
		got := tmin()
		want := MinInt
		assert.Equal(t, want, got)
	} )
	t.Run("bitXor", func(t *testing.T) {
		got := bitXor(4, 5)
		want := 0x00000001
		assert.Equal(t, want, got)
	} )
	t.Run("tmax", func(t *testing.T) {
		got := tmax()
		want := MaxInt
		assert.Equal(t, want, got)
	} )
	t.Run("isPositive", func(t *testing.T) {
		got := isPositive(-1)
		want := false
		assert.Equal(t, want, got)
	} )
	t.Run("allOddBits", func(t *testing.T) {
		got := allOddBits(0xFFFFFFFD)
		want := 0
		assert.Equal(t, want, got)
		got = allOddBits(0xAAAAAAAA)
		want = 1
		assert.Equal(t, want, got)
	} )
	t.Run("bang", func(t *testing.T) {
		got := bang(0)
		want := 1
		assert.Equal(t, want, got)
		got = bang(1)
		want = 0
		assert.Equal(t, want, got)
		got = bang(-2)
		want = 0
		assert.Equal(t, want, got)
	} )
	t.Run("isLessOrEqual", func(t *testing.T) {
		got := isLessOrEqual(4,5)
		want := 1
		assert.Equal(t, want, got)
		got = isLessOrEqual(-1,0)
		want = 1
		assert.Equal(t, want, got)
	} )
	t.Run("ilog2", func(t *testing.T) {
		got := ilog2(6)
		want := 2
		assert.Equal(t, want, got)
		got = ilog2(8)
		want = 3
		assert.Equal(t, want, got)
		got = ilog2(1)
		want = 0
		assert.Equal(t, want, got)
	} )
	t.Run("isAsciiDigit", func(t *testing.T) {
		got := isAsciiDigit(0x35)
		want := 1
		assert.Equal(t, want, got)
		got = isAsciiDigit(0x3A)
		want = 0
		assert.Equal(t, want, got)
		got = isAsciiDigit(0x05)
		want = 0
		assert.Equal(t, want, got)
	} )
	t.Run("conditional", func(t *testing.T) {
		got := conditional(2,4,5)
		want := 4
		assert.Equal(t, want, got)
		got = conditional(-2,4,5)
		want = 5
		assert.Equal(t, want, got)
		got = conditional(0,4,5)
		want = 5
		assert.Equal(t, want, got)
	} )
	t.Run("logicalNeg", func(t *testing.T) {
		got := logicalNeg(0)
		want := 1
		assert.Equal(t, want, got)
		got = logicalNeg(1)
		want = 0
		assert.Equal(t, want, got)
		got = logicalNeg(tmin())
		want = 0
		assert.Equal(t, want, got)
	} )
}

func TestValue(t *testing.T){
	t.Run("bitvalue", func(t *testing.T) {
		for i:=0;i<3;i++{t.Log(-1>>i)}
		for i:=0;i<3;i++{t.Log(-1<<i)}
	})
}