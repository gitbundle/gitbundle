// Code generated by ent, DO NOT EDIT.

package models

import (
	"time"

	"github.com/gitbundle/server/pkg/store/models/repo"
	"github.com/gitbundle/server/pkg/store/models/user"
	"github.com/gitbundle/server/pkg/store/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	repoFields := schema.Repo{}.Fields()
	_ = repoFields
	// repoDescName is the schema descriptor for name field.
	repoDescName := repoFields[2].Descriptor()
	// repo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	repo.NameValidator = repoDescName.Validators[0].(func(string) error)
	// repoDescOriginalURL is the schema descriptor for original_url field.
	repoDescOriginalURL := repoFields[7].Descriptor()
	// repo.OriginalURLValidator is a validator for the "original_url" field. It is called by the builders before save.
	repo.OriginalURLValidator = repoDescOriginalURL.Validators[0].(func(string) error)
	// repoDescStatus is the schema descriptor for status field.
	repoDescStatus := repoFields[24].Descriptor()
	// repo.DefaultStatus holds the default value on creation for the status field.
	repo.DefaultStatus = repoDescStatus.Default.(int)
	// repoDescIsForked is the schema descriptor for is_forked field.
	repoDescIsForked := repoFields[25].Descriptor()
	// repo.DefaultIsForked holds the default value on creation for the is_forked field.
	repo.DefaultIsForked = repoDescIsForked.Default.(bool)
	// repoDescIsTemplate is the schema descriptor for is_template field.
	repoDescIsTemplate := repoFields[27].Descriptor()
	// repo.DefaultIsTemplate holds the default value on creation for the is_template field.
	repo.DefaultIsTemplate = repoDescIsTemplate.Default.(bool)
	// repoDescSize is the schema descriptor for size field.
	repoDescSize := repoFields[29].Descriptor()
	// repo.DefaultSize holds the default value on creation for the size field.
	repo.DefaultSize = repoDescSize.Default.(int64)
	// repoDescIsFsckEnabled is the schema descriptor for is_fsck_enabled field.
	repoDescIsFsckEnabled := repoFields[30].Descriptor()
	// repo.DefaultIsFsckEnabled holds the default value on creation for the is_fsck_enabled field.
	repo.DefaultIsFsckEnabled = repoDescIsFsckEnabled.Default.(bool)
	// repoDescCloseIssuesViaCommitInAnyBranch is the schema descriptor for close_issues_via_commit_in_any_branch field.
	repoDescCloseIssuesViaCommitInAnyBranch := repoFields[31].Descriptor()
	// repo.DefaultCloseIssuesViaCommitInAnyBranch holds the default value on creation for the close_issues_via_commit_in_any_branch field.
	repo.DefaultCloseIssuesViaCommitInAnyBranch = repoDescCloseIssuesViaCommitInAnyBranch.Default.(bool)
	// repoDescAvatar is the schema descriptor for avatar field.
	repoDescAvatar := repoFields[34].Descriptor()
	// repo.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	repo.AvatarValidator = repoDescAvatar.Validators[0].(func(string) error)
	// repoDescCreatedAt is the schema descriptor for created_at field.
	repoDescCreatedAt := repoFields[35].Descriptor()
	// repo.DefaultCreatedAt holds the default value on creation for the created_at field.
	repo.DefaultCreatedAt = repoDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPasswordHashAlgo is the schema descriptor for password_hash_algo field.
	userDescPasswordHashAlgo := userFields[7].Descriptor()
	// user.DefaultPasswordHashAlgo holds the default value on creation for the password_hash_algo field.
	user.DefaultPasswordHashAlgo = userDescPasswordHashAlgo.Default.(string)
	// userDescMustChangePassword is the schema descriptor for must_change_password field.
	userDescMustChangePassword := userFields[8].Descriptor()
	// user.DefaultMustChangePassword holds the default value on creation for the must_change_password field.
	user.DefaultMustChangePassword = userDescMustChangePassword.Default.(bool)
	// userDescLoginSource is the schema descriptor for login_source field.
	userDescLoginSource := userFields[10].Descriptor()
	// user.DefaultLoginSource holds the default value on creation for the login_source field.
	user.DefaultLoginSource = userDescLoginSource.Default.(int64)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[19].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescMaxRepoCreation is the schema descriptor for max_repo_creation field.
	userDescMaxRepoCreation := userFields[23].Descriptor()
	// user.DefaultMaxRepoCreation holds the default value on creation for the max_repo_creation field.
	user.DefaultMaxRepoCreation = userDescMaxRepoCreation.Default.(int)
	// userDescIsRestricted is the schema descriptor for is_restricted field.
	userDescIsRestricted := userFields[26].Descriptor()
	// user.DefaultIsRestricted holds the default value on creation for the is_restricted field.
	user.DefaultIsRestricted = userDescIsRestricted.Default.(bool)
	// userDescAllowCreateOrganization is the schema descriptor for allow_create_organization field.
	userDescAllowCreateOrganization := userFields[27].Descriptor()
	// user.DefaultAllowCreateOrganization holds the default value on creation for the allow_create_organization field.
	user.DefaultAllowCreateOrganization = userDescAllowCreateOrganization.Default.(bool)
	// userDescProhibitLogin is the schema descriptor for prohibit_login field.
	userDescProhibitLogin := userFields[28].Descriptor()
	// user.DefaultProhibitLogin holds the default value on creation for the prohibit_login field.
	user.DefaultProhibitLogin = userDescProhibitLogin.Default.(bool)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[29].Descriptor()
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescNumFollowing is the schema descriptor for num_following field.
	userDescNumFollowing := userFields[32].Descriptor()
	// user.DefaultNumFollowing holds the default value on creation for the num_following field.
	user.DefaultNumFollowing = userDescNumFollowing.Default.(int)
	// userDescVisibility is the schema descriptor for visibility field.
	userDescVisibility := userFields[37].Descriptor()
	// user.DefaultVisibility holds the default value on creation for the visibility field.
	user.DefaultVisibility = userDescVisibility.Default.(int)
	// userDescRepoAdminChangeTeamAccess is the schema descriptor for repo_admin_change_team_access field.
	userDescRepoAdminChangeTeamAccess := userFields[38].Descriptor()
	// user.DefaultRepoAdminChangeTeamAccess holds the default value on creation for the repo_admin_change_team_access field.
	user.DefaultRepoAdminChangeTeamAccess = userDescRepoAdminChangeTeamAccess.Default.(bool)
	// userDescKeepActivityPrivate is the schema descriptor for keep_activity_private field.
	userDescKeepActivityPrivate := userFields[41].Descriptor()
	// user.DefaultKeepActivityPrivate holds the default value on creation for the keep_activity_private field.
	user.DefaultKeepActivityPrivate = userDescKeepActivityPrivate.Default.(bool)
}