package service

import (
	"Valentine-Proposal-Site-Backend/internals/models"
	"Valentine-Proposal-Site-Backend/internals/store"
	"Valentine-Proposal-Site-Backend/internals/util"
)

type PropsalService struct {
	store *store.MemoryStore
}

func NewPropsalService(store *store.MemoryStore) *PropsalService {
	return &PropsalService{
		store: store,
	}
}

func (p *PropsalService) CreateURL(proposal models.ProposalDTO) string {
	code, err := util.GenerateShortCode()
	if err != nil {
		return "There was an error while Generating the ShortCode" + err.Error()
	}

	p.store.Save(code, proposal)
	return code
}

func (p *PropsalService) GetProposal(short string) (models.ProposalDTO, bool) {
	proposal, ok := p.store.Get(short)
	if !ok {
		return proposal, false
	}

	return proposal, true
}
