// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createLeaveStmt, err = db.PrepareContext(ctx, CreateLeave); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLeave: %w", err)
	}
	if q.createLeaveSuccesssStmt, err = db.PrepareContext(ctx, CreateLeaveSuccesss); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLeaveSuccesss: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, CreateUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteLeaveStmt, err = db.PrepareContext(ctx, DeleteLeave); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteLeave: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, DeleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getAttendanceRecordStmt, err = db.PrepareContext(ctx, GetAttendanceRecord); err != nil {
		return nil, fmt.Errorf("error preparing query GetAttendanceRecord: %w", err)
	}
	if q.getEmployeeListByDepartmentStmt, err = db.PrepareContext(ctx, GetEmployeeListByDepartment); err != nil {
		return nil, fmt.Errorf("error preparing query GetEmployeeListByDepartment: %w", err)
	}
	if q.getLeaveByStartDateStmt, err = db.PrepareContext(ctx, GetLeaveByStartDate); err != nil {
		return nil, fmt.Errorf("error preparing query GetLeaveByStartDate: %w", err)
	}
	if q.getLeaveListByDirectorIdStmt, err = db.PrepareContext(ctx, GetLeaveListByDirectorId); err != nil {
		return nil, fmt.Errorf("error preparing query GetLeaveListByDirectorId: %w", err)
	}
	if q.getLeaveListByUserIdStmt, err = db.PrepareContext(ctx, GetLeaveListByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetLeaveListByUserId: %w", err)
	}
	if q.getLeaveRequisitionByIdStmt, err = db.PrepareContext(ctx, GetLeaveRequisitionById); err != nil {
		return nil, fmt.Errorf("error preparing query GetLeaveRequisitionById: %w", err)
	}
	if q.getLeaveTypesStmt, err = db.PrepareContext(ctx, GetLeaveTypes); err != nil {
		return nil, fmt.Errorf("error preparing query GetLeaveTypes: %w", err)
	}
	if q.getLeavesIsSuccessStmt, err = db.PrepareContext(ctx, GetLeavesIsSuccess); err != nil {
		return nil, fmt.Errorf("error preparing query GetLeavesIsSuccess: %w", err)
	}
	if q.getNationalHolidayStmt, err = db.PrepareContext(ctx, GetNationalHoliday); err != nil {
		return nil, fmt.Errorf("error preparing query GetNationalHoliday: %w", err)
	}
	if q.getTotalLeaveListByDirectorIdStmt, err = db.PrepareContext(ctx, GetTotalLeaveListByDirectorId); err != nil {
		return nil, fmt.Errorf("error preparing query GetTotalLeaveListByDirectorId: %w", err)
	}
	if q.getTotalLeaveListByEmployeeIdStmt, err = db.PrepareContext(ctx, GetTotalLeaveListByEmployeeId); err != nil {
		return nil, fmt.Errorf("error preparing query GetTotalLeaveListByEmployeeId: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, GetUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserInfoStmt, err = db.PrepareContext(ctx, GetUserInfo); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserInfo: %w", err)
	}
	if q.listUserBydDepartmentIdStmt, err = db.PrepareContext(ctx, ListUserBydDepartmentId); err != nil {
		return nil, fmt.Errorf("error preparing query ListUserBydDepartmentId: %w", err)
	}
	if q.sreachLeaveListByStartAtStmt, err = db.PrepareContext(ctx, SreachLeaveListByStartAt); err != nil {
		return nil, fmt.Errorf("error preparing query SreachLeaveListByStartAt: %w", err)
	}
	if q.updateLeaveStmt, err = db.PrepareContext(ctx, UpdateLeave); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateLeave: %w", err)
	}
	if q.updateNameStmt, err = db.PrepareContext(ctx, UpdateName); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateName: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createLeaveStmt != nil {
		if cerr := q.createLeaveStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLeaveStmt: %w", cerr)
		}
	}
	if q.createLeaveSuccesssStmt != nil {
		if cerr := q.createLeaveSuccesssStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLeaveSuccesssStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteLeaveStmt != nil {
		if cerr := q.deleteLeaveStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteLeaveStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getAttendanceRecordStmt != nil {
		if cerr := q.getAttendanceRecordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAttendanceRecordStmt: %w", cerr)
		}
	}
	if q.getEmployeeListByDepartmentStmt != nil {
		if cerr := q.getEmployeeListByDepartmentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEmployeeListByDepartmentStmt: %w", cerr)
		}
	}
	if q.getLeaveByStartDateStmt != nil {
		if cerr := q.getLeaveByStartDateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLeaveByStartDateStmt: %w", cerr)
		}
	}
	if q.getLeaveListByDirectorIdStmt != nil {
		if cerr := q.getLeaveListByDirectorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLeaveListByDirectorIdStmt: %w", cerr)
		}
	}
	if q.getLeaveListByUserIdStmt != nil {
		if cerr := q.getLeaveListByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLeaveListByUserIdStmt: %w", cerr)
		}
	}
	if q.getLeaveRequisitionByIdStmt != nil {
		if cerr := q.getLeaveRequisitionByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLeaveRequisitionByIdStmt: %w", cerr)
		}
	}
	if q.getLeaveTypesStmt != nil {
		if cerr := q.getLeaveTypesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLeaveTypesStmt: %w", cerr)
		}
	}
	if q.getLeavesIsSuccessStmt != nil {
		if cerr := q.getLeavesIsSuccessStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLeavesIsSuccessStmt: %w", cerr)
		}
	}
	if q.getNationalHolidayStmt != nil {
		if cerr := q.getNationalHolidayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNationalHolidayStmt: %w", cerr)
		}
	}
	if q.getTotalLeaveListByDirectorIdStmt != nil {
		if cerr := q.getTotalLeaveListByDirectorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTotalLeaveListByDirectorIdStmt: %w", cerr)
		}
	}
	if q.getTotalLeaveListByEmployeeIdStmt != nil {
		if cerr := q.getTotalLeaveListByEmployeeIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTotalLeaveListByEmployeeIdStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserInfoStmt != nil {
		if cerr := q.getUserInfoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserInfoStmt: %w", cerr)
		}
	}
	if q.listUserBydDepartmentIdStmt != nil {
		if cerr := q.listUserBydDepartmentIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUserBydDepartmentIdStmt: %w", cerr)
		}
	}
	if q.sreachLeaveListByStartAtStmt != nil {
		if cerr := q.sreachLeaveListByStartAtStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing sreachLeaveListByStartAtStmt: %w", cerr)
		}
	}
	if q.updateLeaveStmt != nil {
		if cerr := q.updateLeaveStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateLeaveStmt: %w", cerr)
		}
	}
	if q.updateNameStmt != nil {
		if cerr := q.updateNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNameStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                DBTX
	tx                                *sql.Tx
	createLeaveStmt                   *sql.Stmt
	createLeaveSuccesssStmt           *sql.Stmt
	createUserStmt                    *sql.Stmt
	deleteLeaveStmt                   *sql.Stmt
	deleteUserStmt                    *sql.Stmt
	getAttendanceRecordStmt           *sql.Stmt
	getEmployeeListByDepartmentStmt   *sql.Stmt
	getLeaveByStartDateStmt           *sql.Stmt
	getLeaveListByDirectorIdStmt      *sql.Stmt
	getLeaveListByUserIdStmt          *sql.Stmt
	getLeaveRequisitionByIdStmt       *sql.Stmt
	getLeaveTypesStmt                 *sql.Stmt
	getLeavesIsSuccessStmt            *sql.Stmt
	getNationalHolidayStmt            *sql.Stmt
	getTotalLeaveListByDirectorIdStmt *sql.Stmt
	getTotalLeaveListByEmployeeIdStmt *sql.Stmt
	getUserByEmailStmt                *sql.Stmt
	getUserInfoStmt                   *sql.Stmt
	listUserBydDepartmentIdStmt       *sql.Stmt
	sreachLeaveListByStartAtStmt      *sql.Stmt
	updateLeaveStmt                   *sql.Stmt
	updateNameStmt                    *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                tx,
		tx:                                tx,
		createLeaveStmt:                   q.createLeaveStmt,
		createLeaveSuccesssStmt:           q.createLeaveSuccesssStmt,
		createUserStmt:                    q.createUserStmt,
		deleteLeaveStmt:                   q.deleteLeaveStmt,
		deleteUserStmt:                    q.deleteUserStmt,
		getAttendanceRecordStmt:           q.getAttendanceRecordStmt,
		getEmployeeListByDepartmentStmt:   q.getEmployeeListByDepartmentStmt,
		getLeaveByStartDateStmt:           q.getLeaveByStartDateStmt,
		getLeaveListByDirectorIdStmt:      q.getLeaveListByDirectorIdStmt,
		getLeaveListByUserIdStmt:          q.getLeaveListByUserIdStmt,
		getLeaveRequisitionByIdStmt:       q.getLeaveRequisitionByIdStmt,
		getLeaveTypesStmt:                 q.getLeaveTypesStmt,
		getLeavesIsSuccessStmt:            q.getLeavesIsSuccessStmt,
		getNationalHolidayStmt:            q.getNationalHolidayStmt,
		getTotalLeaveListByDirectorIdStmt: q.getTotalLeaveListByDirectorIdStmt,
		getTotalLeaveListByEmployeeIdStmt: q.getTotalLeaveListByEmployeeIdStmt,
		getUserByEmailStmt:                q.getUserByEmailStmt,
		getUserInfoStmt:                   q.getUserInfoStmt,
		listUserBydDepartmentIdStmt:       q.listUserBydDepartmentIdStmt,
		sreachLeaveListByStartAtStmt:      q.sreachLeaveListByStartAtStmt,
		updateLeaveStmt:                   q.updateLeaveStmt,
		updateNameStmt:                    q.updateNameStmt,
	}
}
