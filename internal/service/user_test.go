package service

import (
    "context"
    "testing"
)

func TestUserCreate_HashesPasswordAndSetsRole(t *testing.T) {
    fu := &fakeUserRepo{}
    fr := &fakeRoleRepo{name: "warga"}
    svc := NewUserService(fu, fr)

    u := dummyUser()
    plain := u.Password
    if err := svc.Create(context.Background(), &u); err != nil {
        t.Fatalf("Create error: %v", err)
    }
    if u.RoleID == 0 { t.Fatalf("expected RoleID set, got 0") }
    if u.Password == plain { t.Fatalf("expected password hashed, stayed plain") }
}

func dummyUser() (u struct{ ID uint; Name, Email, Password string; RoleID uint }) {
    u.Name = "Tester"
    u.Email = "tester@example.com"
    u.Password = "secret"
    return
}

