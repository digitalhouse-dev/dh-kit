package preload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cast_01_Ok(t *testing.T) {
	preload := "{product{cluster{},path{}},user{lala{}},s{}, b{s{},a{}}}"

	m, err := Cast(preload)

	assert.Nil(t, err)
	assert.EqualValues(t, len(m), 4)
	assert.EqualValues(t, m["product"], "{cluster{}path{}}")
	assert.EqualValues(t, m["user"], "{lala{}}")
	assert.EqualValues(t, m["s"], "{}")
	assert.EqualValues(t, m["b"], "{s{}a{}}")
}

func Test_Cast_02_Ok(t *testing.T) {
	preload := "{product{cluster{}, path{}}}"

	m, err := Cast(preload)

	assert.Nil(t, err)
	assert.EqualValues(t, len(m), 1)
	assert.EqualValues(t, m["product"], "{cluster{}path{}}")
}

func Test_Cast_03_Ok(t *testing.T) {
	preload := "{product{value{}}, test{value{v1{}v2{}}} }"

	m, err := Cast(preload)

	assert.Nil(t, err)
	assert.EqualValues(t, len(m), 2)
	assert.EqualValues(t, m["product"], "{value{}}")
	assert.EqualValues(t, m["test"], "{value{v1{}v2{}}}")
}

func Test_Cast_04_Ok(t *testing.T) {
	preload := "{product{value{a{b{}}c{}}}test{value{v1{}v2{}}}t{}}"

	m, err := Cast(preload)

	assert.Nil(t, err)
	assert.EqualValues(t, len(m), 3)
	assert.EqualValues(t, m["product"], "{value{a{b{}}c{}}}")
	assert.EqualValues(t, m["test"], "{value{v1{}v2{}}}")
	assert.EqualValues(t, m["t"], "{}")
}

func Test_Cast_05_Ok(t *testing.T) {
	preload := "{}"

	m, err := Cast(preload)

	assert.Nil(t, err)
	assert.EqualValues(t, len(m), 0)
}

func Test_Cast_06_Ok(t *testing.T) {
	m, err := Cast("")

	assert.EqualValues(t, len(m), 0)
	assert.Nil(t, err, "")
}

func Test_Cast_01_Err(t *testing.T) {
	preload := "{product{cluster{{}, path{}},user{lala{}},s{},b{s{},a{}}}"

	m, err := Cast(preload)

	assert.EqualValues(t, len(m), 0)
	assert.EqualError(t, err, "invalid preload format {product{cluster{{}, path{}},user{lala{}},s{},b{s{},a{}}}")
}

func Test_Cast_02_Err(t *testing.T) {
	m, err := Cast("test")

	assert.Nil(t, m)
	assert.EqualError(t, err, "invalid preload format test")
}

func BenchmarkCast(b *testing.B) {
	preload := "{product{cluster{{}, path{}},user{lala{}},s{},b{s{},a{}}}"

	for i := 0; i < b.N; i++ {
		_, _ = Cast(preload)
	}

}
