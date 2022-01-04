package client

import (
	"os"
	"testing"
)

func TestPermissionGroups(t *testing.T) {
	baseUrl := os.Getenv("BASE_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if baseUrl == "" || username == "" || password == "" {
		t.Fatal("baseUrl, username, password must be set")
	}
	c, _ := NewClient(baseUrl, username, password)
	var createdGroupId int

	t.Run("Get PermissionGroups", func(t *testing.T) {
		groups, err := c.GetPermissionGroups()

		if err != nil {
			t.Fatal(err)
		}

		if len(groups) == 0 {
			t.Fatal("no groups returned")
		}
	})

	t.Run("Create PermissionGroup", func(t *testing.T) {
		group, err := c.CreatePermissionGroup("test-client")

		if err != nil {
			t.Fatal(err)
		}

		if group.ID == 0 {
			t.Fatal("group id is zero")
		}
		createdGroupId = group.ID
	})

	t.Run("Get PermissionGroup", func(t *testing.T) {
		group, err := c.GetPermissionGroup(createdGroupId)

		if err != nil {
			t.Fatal(err)
		}

		if group.ID == 0 || group.Name == "" {
			t.Fatal("Get PermissionGroup failed")
		}
	})

	t.Run("Delete PermissionGroup", func(t *testing.T) {
		group, _ := c.GetPermissionGroup(createdGroupId)

		err := c.DeletePermissionGroup(group.ID)

		if err != nil {
			t.Fatal(err)
		}

		group, _ = c.GetPermissionGroup(createdGroupId)

		if group.ID != 0 {
			t.Fatal("Delete PermissionGroup failed")
		}
	})
}
