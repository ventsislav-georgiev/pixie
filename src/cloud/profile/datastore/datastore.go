package datastore

import (
	"errors"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// TODO(zasgar): Move these to models ?

// UserInfo tracks information about a specific end-user.
type UserInfo struct {
	ID        uuid.UUID `db:"id"`
	OrgID     uuid.UUID `db:"org_id"`
	Username  string    `db:"username"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
}

// OrgInfo tracks information about an organization.
type OrgInfo struct {
	ID         uuid.UUID `db:"id"`
	OrgName    string    `db:"org_name"`
	DomainName string    `db:"domain_name"`
}

// Datastore is a postgres backed storage for entities.
type Datastore struct {
	db *sqlx.DB
}

// NewDatastore creates a Datastore.
func NewDatastore(db *sqlx.DB) *Datastore {
	return &Datastore{db: db}
}

// CreateUser creates a new user.
func (d *Datastore) CreateUser(userInfo *UserInfo) (uuid.UUID, error) {
	txn, err := d.db.Beginx()
	if err != nil {
		return uuid.Nil, err
	}
	defer txn.Rollback()

	u, err := d.createUserUsingTxn(txn, userInfo)
	if err != nil {
		return uuid.Nil, err
	}
	return u, txn.Commit()
}

// GetUser gets user information by user ID.
func (d *Datastore) GetUser(id uuid.UUID) (*UserInfo, error) {
	query := `SELECT * from users WHERE id=$1`
	rows, err := d.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var userInfo UserInfo
		err := rows.StructScan(&userInfo)
		return &userInfo, err
	}
	return nil, errors.New("failed to read user from database")
}

// CreateOrg creates a new organization, returning the created org ID.
func (d *Datastore) CreateOrg(orgInfo *OrgInfo) (uuid.UUID, error) {
	txn, err := d.db.Beginx()
	if err != nil {
		return uuid.Nil, err
	}
	defer txn.Rollback()

	u, err := d.createOrgUsingTxn(txn, orgInfo)
	if err != nil {
		return uuid.Nil, err
	}
	return u, txn.Commit()
}

// CreateUserAndOrg creates a new user and organization as needed for initial user/org creation.
func (d *Datastore) CreateUserAndOrg(orgInfo *OrgInfo, userInfo *UserInfo) (orgID uuid.UUID, userID uuid.UUID, err error) {
	txn, err := d.db.Beginx()
	if err != nil {
		return
	}
	defer txn.Rollback()

	orgID, err = d.createOrgUsingTxn(txn, orgInfo)
	if err != nil {
		return
	}
	userInfo.OrgID = orgID
	userID, err = d.createUserUsingTxn(txn, userInfo)

	orgInfo.ID = orgID
	userInfo.ID = userID
	userInfo.OrgID = orgID

	return orgID, userID, txn.Commit()
}

// GetOrg gets org information by ID.
func (d *Datastore) GetOrg(id uuid.UUID) (*OrgInfo, error) {
	query := `SELECT * from orgs WHERE id=$1`
	rows, err := d.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var orgInfo OrgInfo
		err := rows.StructScan(&orgInfo)
		return &orgInfo, err
	}
	return nil, errors.New("failed to get org info from database")
}

// GetOrgByDomain gets org information by domain.
func (d *Datastore) GetOrgByDomain(domainName string) (*OrgInfo, error) {
	query := `SELECT * from orgs WHERE domain_name=$1`
	rows, err := d.db.Queryx(query, domainName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var orgInfo OrgInfo
		err := rows.StructScan(&orgInfo)
		return &orgInfo, err
	}
	return nil, errors.New("failed org info from database")
}

// GetUserByEmail gets org information by domain.
func (d *Datastore) GetUserByEmail(email string) (*UserInfo, error) {
	query := `SELECT * from users WHERE email=$1`
	rows, err := d.db.Queryx(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var userInfo UserInfo
		err := rows.StructScan(&userInfo)
		return &userInfo, err
	}
	return nil, errors.New("failed to get user info from database")
}

func (d *Datastore) createUserUsingTxn(txn *sqlx.Tx, userInfo *UserInfo) (uuid.UUID, error) {
	query := `INSERT INTO users (org_id, username, first_name, last_name, email) VALUES (:org_id, :username, :first_name, :last_name, :email) RETURNING id`
	row, err := txn.NamedQuery(query, userInfo)
	if err != nil {
		return uuid.Nil, err
	}
	defer row.Close()

	if row.Next() {
		var id uuid.UUID
		if err := row.Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil
	}
	return uuid.Nil, errors.New("failed to read user id from the database")
}

func (d *Datastore) createOrgUsingTxn(txn *sqlx.Tx, orgInfo *OrgInfo) (uuid.UUID, error) {
	query := `INSERT INTO orgs (org_name, domain_name) VALUES (:org_name, :domain_name) RETURNING id`
	row, err := txn.NamedQuery(query, orgInfo)
	if err != nil {
		return uuid.Nil, err
	}
	defer row.Close()

	if row.Next() {
		var id uuid.UUID
		if err := row.Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil
	}
	return uuid.Nil, errors.New("failed to read org id from the database")
}
