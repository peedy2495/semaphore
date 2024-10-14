package bolt

import (
	"encoding/base64"

	"github.com/ansible-semaphore/semaphore/db"
	"github.com/gorilla/securecookie"
	"go.etcd.io/bbolt"
)

func (d *BoltDb) GetRunner(projectID int, runnerID int) (runner db.Runner, err error) {
	return
}

func (d *BoltDb) GetRunners(projectID int) (runners []db.Runner, err error) {
	return
}

func (d *BoltDb) DeleteRunner(projectID int, runnerID int) (err error) {
	return
}

func (d *BoltDb) GetGlobalRunnerByToken(token string) (runner db.Runner, err error) {

	runners := make([]db.Runner, 0)

	err = d.getObjects(0, db.GlobalRunnerProps, db.RetrieveQueryParams{}, func(i interface{}) bool {
		r := i.(db.Runner)
		return r.Token == token
	}, &runners)

	if err != nil {
		return
	}

	if len(runners) == 0 {
		err = db.ErrNotFound
		return
	}

	runner = runners[0]
	return
}

func (d *BoltDb) GetGlobalRunner(runnerID int) (runner db.Runner, err error) {
	err = d.getObject(0, db.GlobalRunnerProps, intObjectID(runnerID), &runner)

	return
}

func (d *BoltDb) GetGlobalRunners(activeOnly bool) (runners []db.Runner, err error) {
	err = d.getObjects(0, db.GlobalRunnerProps, db.RetrieveQueryParams{}, func(i interface{}) bool {
		runner := i.(db.Runner)
		if activeOnly {
			return runner.Active
		}
		return true
	}, &runners)
	return
}

func (d *BoltDb) DeleteGlobalRunner(runnerID int) (err error) {
	return d.db.Update(func(tx *bbolt.Tx) error {
		return d.deleteObject(0, db.GlobalRunnerProps, intObjectID(runnerID), tx)
	})
}

func (d *BoltDb) UpdateRunner(runner db.Runner) (err error) {
	return d.updateObject(0, db.GlobalRunnerProps, runner)
}

func (d *BoltDb) CreateRunner(runner db.Runner) (newRunner db.Runner, err error) {
	runner.Token = base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32))

	res, err := d.createObject(0, db.GlobalRunnerProps, runner)
	if err != nil {
		return
	}
	newRunner = res.(db.Runner)
	return
}
