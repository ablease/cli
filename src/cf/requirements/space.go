package requirements

import (
	"cf"
	"cf/api"
	"cf/net"
	"cf/terminal"
)

type SpaceRequirement interface {
	Requirement
	GetSpace() cf.Space
}

type SpaceApiRequirement struct {
	name      string
	ui        terminal.UI
	spaceRepo api.SpaceRepository
	space     cf.Space
}

func NewSpaceRequirement(name string, ui terminal.UI, sR api.SpaceRepository) (req *SpaceApiRequirement) {
	req = new(SpaceApiRequirement)
	req.name = name
	req.ui = ui
	req.spaceRepo = sR
	return
}

func (req *SpaceApiRequirement) Execute() (success bool) {
	var (
		apiErr *net.ApiError
		found  bool
	)

	req.space, found, apiErr = req.spaceRepo.FindByName(req.name)

	if apiErr != nil {
		req.ui.Failed(apiErr.Error())
		return false
	}

	return found
}

func (req *SpaceApiRequirement) GetSpace() cf.Space {
	return req.space
}
