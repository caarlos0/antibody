package project_test

import (
	"os"
	"testing"

	"github.com/getantibody/antibody/project"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	assert.NoError(project.NewGit(home, "caarlos0/jvm", "gh-pages").Download())
	list, err := project.List(home)
	assert.NoError(err)
	assert.Len(list, 1)
}

func TestListEmptyFolder(t *testing.T) {
	assert := assert.New(t)
	home := home()
	defer os.RemoveAll(home)
	list, err := project.List(home)
	assert.NoError(err)
	assert.Len(list, 0)
}

func TestListNonExistentFolder(t *testing.T) {
	assert := assert.New(t)
	list, err := project.List("/tmp/asdasdadadwhateverwtff")
	assert.Error(err)
	assert.Len(list, 0)
}

func TestUpdateUpdate(t *testing.T) {
	assert.Nil(t, project.Update(""))
}