package logicCore

import (
	"fmt"
	"errors"
	"github.com/seheee/PDK/logic-core/domain/model"
)

type logicCore struct {
	mux
}

type mux struct {
	chTable    map[string]map[string]chan model.LogicData
	logicTable map[string]string
}

func NewLogicCore() *logicCore {
	return &logicCore{
		mux{
			chTable:    make(map[string]map[string]chan model.LogicData),
			logicTable: make(map[string]string),
		},
	}
}

/*
func (m *mux) CreateAndStartLogic(r *model.ChainRequest) {
	listen := make(chan model.LogicData, 100)
	lchs, ok := m.chTable[r.SID]
	if !ok {
		m.chTable[r.SID] = make(map[string]chan model.LogicData)
		lchs, _ = m.chTable[r.SID]
	}
	lchs[r.Name] = listen

	chain := chainFactory(r.Rings)
	for d := range listen {
		chain.exec(&d)
	}
}
*/

func (m *mux) CreateAndStartLogic(r *model.RingRequest, id string) {
	fmt.Println("id : ", id)
	listen := make(chan model.LogicData, 100)
	lchs, ok := m.chTable[r.Sensor]
	if !ok {
		m.chTable[r.Sensor] = make(map[string]chan model.LogicData)
		lchs, _ = m.chTable[r.Sensor]
	}
	lchs[id] = listen
	m.logicTable[id] = r.Sensor
	/*for r.Logic[0].Elem == "empty" {
		r.Logic = r.Logic[1:]
	}*/
	chain := chainFactory(r)
	fmt.Println(chain)
	fmt.Println(chain.next)
	for d := range listen {
		chain.exec(&d)
	}
}

func (m *mux) GetLogicChans(key string) map[string]chan model.LogicData {
	lchs, ok := m.chTable[key]
	if !ok {
		return nil
	}
	return lchs
}

func (m *mux) RemoveLogic(id string) (err error) {
	sid, ok := m.logicTable[id]
	if !ok {
		err = errors.New("cannot find logicChain " + id)
	}
	ch, _ := m.chTable[sid][id]
	close(ch)
	return err
}

func (m *mux) RemoveLogicsBySID(sid string) (err error) {
	lchs, ok := m.chTable[sid]
	if !ok {
		err = errors.New("there is no sensor " + sid)
	}
	for _, ch := range lchs {
		close(ch)
	}
	return err
}
