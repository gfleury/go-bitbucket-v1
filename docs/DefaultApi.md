# \DefaultApi

All URIs are relative to *http://example.com/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupToUser**](DefaultApi.md#AddGroupToUser) | **Post** /api/1.0/admin/users/add-group | 
[**AddUserToGroup**](DefaultApi.md#AddUserToGroup) | **Post** /api/1.0/admin/groups/add-user | 
[**AddUserToGroups**](DefaultApi.md#AddUserToGroups) | **Post** /api/1.0/admin/users/add-groups | 
[**AddUsersToGroup**](DefaultApi.md#AddUsersToGroup) | **Post** /api/1.0/admin/groups/add-users | 
[**Approve**](DefaultApi.md#Approve) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve | 
[**AssignParticipantRole**](DefaultApi.md#AssignParticipantRole) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants | 
[**CanMerge**](DefaultApi.md#CanMerge) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/merge | 
[**ClearSenderAddress**](DefaultApi.md#ClearSenderAddress) | **Delete** /api/1.0/admin/mail-server/sender-address | 
[**ClearUserCaptchaChallenge**](DefaultApi.md#ClearUserCaptchaChallenge) | **Delete** /api/1.0/admin/users/captcha | 
[**CountPullRequestTasks**](DefaultApi.md#CountPullRequestTasks) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/tasks/count | 
[**Create**](DefaultApi.md#Create) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests | 
[**CreateBranch**](DefaultApi.md#CreateBranch) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | 
[**CreateComment**](DefaultApi.md#CreateComment) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments | 
[**CreateComment_0**](DefaultApi.md#CreateComment_0) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments | 
[**CreateGroup**](DefaultApi.md#CreateGroup) | **Post** /api/1.0/admin/groups | 
[**CreateProject**](DefaultApi.md#CreateProject) | **Post** /api/1.0/projects | 
[**CreateRepository**](DefaultApi.md#CreateRepository) | **Post** /api/1.0/projects/{projectKey}/repos | 
[**CreateTag**](DefaultApi.md#CreateTag) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/tags | 
[**CreateTask**](DefaultApi.md#CreateTask) | **Post** /api/1.0/tasks | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /api/1.0/admin/users | 
[**CreateWebhook**](DefaultApi.md#CreateWebhook) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | 
[**Decline**](DefaultApi.md#Decline) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/decline | 
[**Delete**](DefaultApi.md#Delete) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | 
[**DeleteAvatar**](DefaultApi.md#DeleteAvatar) | **Delete** /api/1.0/users/{userSlug}/avatar.png | 
[**DeleteComment**](DefaultApi.md#DeleteComment) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId} | 
[**DeleteComment_0**](DefaultApi.md#DeleteComment_0) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId} | 
[**DeleteGroup**](DefaultApi.md#DeleteGroup) | **Delete** /api/1.0/admin/groups | 
[**DeleteMailConfig**](DefaultApi.md#DeleteMailConfig) | **Delete** /api/1.0/admin/mail-server | 
[**DeleteProject**](DefaultApi.md#DeleteProject) | **Delete** /api/1.0/projects/{projectKey} | 
[**DeleteRepository**](DefaultApi.md#DeleteRepository) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug} | 
[**DeleteRepositoryHook**](DefaultApi.md#DeleteRepositoryHook) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey} | 
[**DeleteTask**](DefaultApi.md#DeleteTask) | **Delete** /api/1.0/tasks/{taskId} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /api/1.0/admin/users | 
[**DeleteWebhook**](DefaultApi.md#DeleteWebhook) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | 
[**DisableHook**](DefaultApi.md#DisableHook) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/enabled | 
[**DisableHook_0**](DefaultApi.md#DisableHook_0) | **Delete** /api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/enabled | 
[**EditFile**](DefaultApi.md#EditFile) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | 
[**EnableHook**](DefaultApi.md#EnableHook) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/enabled | 
[**EnableHook_0**](DefaultApi.md#EnableHook_0) | **Put** /api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/enabled | 
[**FindGroupsForUser**](DefaultApi.md#FindGroupsForUser) | **Get** /api/1.0/admin/users/more-members | 
[**FindOtherGroupsForUser**](DefaultApi.md#FindOtherGroupsForUser) | **Get** /api/1.0/admin/users/more-non-members | 
[**FindUsersInGroup**](DefaultApi.md#FindUsersInGroup) | **Get** /api/1.0/admin/groups/more-members | 
[**FindUsersNotInGroup**](DefaultApi.md#FindUsersNotInGroup) | **Get** /api/1.0/admin/groups/more-non-members | 
[**FindWebhooks**](DefaultApi.md#FindWebhooks) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | 
[**ForkRepository**](DefaultApi.md#ForkRepository) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug} | 
[**Get**](DefaultApi.md#Get) | **Get** /api/1.0/admin/license | 
[**GetActivities**](DefaultApi.md#GetActivities) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/activities | 
[**GetApplicationProperties**](DefaultApi.md#GetApplicationProperties) | **Get** /api/1.0/application-properties | 
[**GetArchive**](DefaultApi.md#GetArchive) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/archive | 
[**GetAvatar**](DefaultApi.md#GetAvatar) | **Get** /api/1.0/hooks/{hookKey}/avatar | 
[**GetBranches**](DefaultApi.md#GetBranches) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | 
[**GetChanges**](DefaultApi.md#GetChanges) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/changes | 
[**GetChanges_0**](DefaultApi.md#GetChanges_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/changes | 
[**GetComment**](DefaultApi.md#GetComment) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId} | 
[**GetComment_0**](DefaultApi.md#GetComment_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId} | 
[**GetComments**](DefaultApi.md#GetComments) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments | 
[**GetComments_0**](DefaultApi.md#GetComments_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments | 
[**GetCommit**](DefaultApi.md#GetCommit) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId} | 
[**GetCommits**](DefaultApi.md#GetCommits) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits | 
[**GetCommits_0**](DefaultApi.md#GetCommits_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/commits | 
[**GetContent**](DefaultApi.md#GetContent) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse | 
[**GetContent_0**](DefaultApi.md#GetContent_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | 
[**GetContent_1**](DefaultApi.md#GetContent_1) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/raw | 
[**GetContent_2**](DefaultApi.md#GetContent_2) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/raw/{path} | 
[**GetDefaultBranch**](DefaultApi.md#GetDefaultBranch) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches/default | 
[**GetForkedRepositories**](DefaultApi.md#GetForkedRepositories) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/forks | 
[**GetGroups**](DefaultApi.md#GetGroups) | **Get** /api/1.0/admin/groups | 
[**GetGroupsWithAnyPermission**](DefaultApi.md#GetGroupsWithAnyPermission) | **Get** /api/1.0/admin/permissions/groups | 
[**GetGroupsWithAnyPermission_0**](DefaultApi.md#GetGroupsWithAnyPermission_0) | **Get** /api/1.0/projects/{projectKey}/permissions/groups | 
[**GetGroupsWithAnyPermission_1**](DefaultApi.md#GetGroupsWithAnyPermission_1) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups | 
[**GetGroupsWithoutAnyPermission**](DefaultApi.md#GetGroupsWithoutAnyPermission) | **Get** /api/1.0/admin/permissions/groups/none | 
[**GetGroupsWithoutAnyPermission_0**](DefaultApi.md#GetGroupsWithoutAnyPermission_0) | **Get** /api/1.0/projects/{projectKey}/permissions/groups/none | 
[**GetGroupsWithoutAnyPermission_1**](DefaultApi.md#GetGroupsWithoutAnyPermission_1) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups/none | 
[**GetGroups_0**](DefaultApi.md#GetGroups_0) | **Get** /api/1.0/groups | 
[**GetInformation**](DefaultApi.md#GetInformation) | **Get** /api/1.0/admin/cluster | 
[**GetLatestInvocation**](DefaultApi.md#GetLatestInvocation) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/latest | 
[**GetLevel**](DefaultApi.md#GetLevel) | **Get** /api/1.0/logs/logger/{loggerName} | 
[**GetMailConfig**](DefaultApi.md#GetMailConfig) | **Get** /api/1.0/admin/mail-server | 
[**GetMergeConfig**](DefaultApi.md#GetMergeConfig) | **Get** /api/1.0/admin/pull-requests/{scmId} | 
[**GetPage**](DefaultApi.md#GetPage) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests | 
[**GetProject**](DefaultApi.md#GetProject) | **Get** /api/1.0/projects/{projectKey} | 
[**GetProjectAvatar**](DefaultApi.md#GetProjectAvatar) | **Get** /api/1.0/projects/{projectKey}/avatar.png | 
[**GetProjects**](DefaultApi.md#GetProjects) | **Get** /api/1.0/projects | 
[**GetPullRequestCount**](DefaultApi.md#GetPullRequestCount) | **Get** /api/1.0/inbox/pull-requests/count | 
[**GetPullRequestSettings**](DefaultApi.md#GetPullRequestSettings) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/pull-requests | 
[**GetPullRequestSettings_0**](DefaultApi.md#GetPullRequestSettings_0) | **Get** /api/1.0/projects/{projectKey}/settings/pull-requests/{scmId} | 
[**GetPullRequestSuggestions**](DefaultApi.md#GetPullRequestSuggestions) | **Get** /api/1.0/dashboard/pull-request-suggestions | 
[**GetPullRequestTasks**](DefaultApi.md#GetPullRequestTasks) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/tasks | 
[**GetPullRequests**](DefaultApi.md#GetPullRequests) | **Get** /api/1.0/dashboard/pull-requests | 
[**GetPullRequests_0**](DefaultApi.md#GetPullRequests_0) | **Get** /api/1.0/inbox/pull-requests | 
[**GetRelatedRepositories**](DefaultApi.md#GetRelatedRepositories) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/related | 
[**GetRepositories**](DefaultApi.md#GetRepositories) | **Get** /api/1.0/projects/{projectKey}/repos | 
[**GetRepositoriesRecentlyAccessed**](DefaultApi.md#GetRepositoriesRecentlyAccessed) | **Get** /api/1.0/profile/recent/repos | 
[**GetRepositories_0**](DefaultApi.md#GetRepositories_0) | **Get** /api/1.0/repos | 
[**GetRepository**](DefaultApi.md#GetRepository) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug} | 
[**GetRepositoryHook**](DefaultApi.md#GetRepositoryHook) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey} | 
[**GetRepositoryHook_0**](DefaultApi.md#GetRepositoryHook_0) | **Get** /api/1.0/projects/{projectKey}/settings/hooks/{hookKey} | 
[**GetRepositoryHooks**](DefaultApi.md#GetRepositoryHooks) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks | 
[**GetRepositoryHooks_0**](DefaultApi.md#GetRepositoryHooks_0) | **Get** /api/1.0/projects/{projectKey}/settings/hooks | 
[**GetRootLevel**](DefaultApi.md#GetRootLevel) | **Get** /api/1.0/logs/rootLogger | 
[**GetSenderAddress**](DefaultApi.md#GetSenderAddress) | **Get** /api/1.0/admin/mail-server/sender-address | 
[**GetSettings**](DefaultApi.md#GetSettings) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/settings | 
[**GetSettings_0**](DefaultApi.md#GetSettings_0) | **Get** /api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/settings | 
[**GetStatistics**](DefaultApi.md#GetStatistics) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/statistics | 
[**GetStatisticsSummary**](DefaultApi.md#GetStatisticsSummary) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/statistics/summary | 
[**GetTag**](DefaultApi.md#GetTag) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/tags/{name} | 
[**GetTags**](DefaultApi.md#GetTags) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/tags | 
[**GetTask**](DefaultApi.md#GetTask) | **Get** /api/1.0/tasks/{taskId} | 
[**GetUser**](DefaultApi.md#GetUser) | **Get** /api/1.0/users/{userSlug} | 
[**GetUserSettings**](DefaultApi.md#GetUserSettings) | **Get** /api/1.0/users/{userSlug}/settings | 
[**GetUsers**](DefaultApi.md#GetUsers) | **Get** /api/1.0/admin/users | 
[**GetUsersWithAnyPermission**](DefaultApi.md#GetUsersWithAnyPermission) | **Get** /api/1.0/admin/permissions/users | 
[**GetUsersWithAnyPermission_0**](DefaultApi.md#GetUsersWithAnyPermission_0) | **Get** /api/1.0/projects/{projectKey}/permissions/users | 
[**GetUsersWithAnyPermission_1**](DefaultApi.md#GetUsersWithAnyPermission_1) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users | 
[**GetUsersWithoutAnyPermission**](DefaultApi.md#GetUsersWithoutAnyPermission) | **Get** /api/1.0/admin/permissions/users/none | 
[**GetUsersWithoutPermission**](DefaultApi.md#GetUsersWithoutPermission) | **Get** /api/1.0/projects/{projectKey}/permissions/users/none | 
[**GetUsersWithoutPermission_0**](DefaultApi.md#GetUsersWithoutPermission_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users/none | 
[**GetUsers_0**](DefaultApi.md#GetUsers_0) | **Get** /api/1.0/users | 
[**GetWebhook**](DefaultApi.md#GetWebhook) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | 
[**Get_0**](DefaultApi.md#Get_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | 
[**HasAllUserPermission**](DefaultApi.md#HasAllUserPermission) | **Get** /api/1.0/projects/{projectKey}/permissions/{permission}/all | 
[**ListParticipants**](DefaultApi.md#ListParticipants) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants | 
[**Merge**](DefaultApi.md#Merge) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/merge | 
[**ModifyAllUserPermission**](DefaultApi.md#ModifyAllUserPermission) | **Post** /api/1.0/projects/{projectKey}/permissions/{permission}/all | 
[**Preview**](DefaultApi.md#Preview) | **Post** /api/1.0/markup/preview | 
[**RemoveGroupFromUser**](DefaultApi.md#RemoveGroupFromUser) | **Post** /api/1.0/admin/users/remove-group | 
[**RemoveUserFromGroup**](DefaultApi.md#RemoveUserFromGroup) | **Post** /api/1.0/admin/groups/remove-user | 
[**RenameUser**](DefaultApi.md#RenameUser) | **Post** /api/1.0/admin/users/rename | 
[**Reopen**](DefaultApi.md#Reopen) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/reopen | 
[**RetryCreateRepository**](DefaultApi.md#RetryCreateRepository) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/recreate | 
[**RevokePermissionsForGroup**](DefaultApi.md#RevokePermissionsForGroup) | **Delete** /api/1.0/admin/permissions/groups | 
[**RevokePermissionsForGroup_0**](DefaultApi.md#RevokePermissionsForGroup_0) | **Delete** /api/1.0/projects/{projectKey}/permissions/groups | 
[**RevokePermissionsForGroup_1**](DefaultApi.md#RevokePermissionsForGroup_1) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups | 
[**RevokePermissionsForUser**](DefaultApi.md#RevokePermissionsForUser) | **Delete** /api/1.0/admin/permissions/users | 
[**RevokePermissionsForUser_0**](DefaultApi.md#RevokePermissionsForUser_0) | **Delete** /api/1.0/projects/{projectKey}/permissions/users | 
[**RevokePermissionsForUser_1**](DefaultApi.md#RevokePermissionsForUser_1) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users | 
[**Search**](DefaultApi.md#Search) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/participants | 
[**SetDefaultBranch**](DefaultApi.md#SetDefaultBranch) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches/default | 
[**SetLevel**](DefaultApi.md#SetLevel) | **Put** /api/1.0/logs/logger/{loggerName}/{levelName} | 
[**SetMailConfig**](DefaultApi.md#SetMailConfig) | **Put** /api/1.0/admin/mail-server | 
[**SetMergeConfig**](DefaultApi.md#SetMergeConfig) | **Post** /api/1.0/admin/pull-requests/{scmId} | 
[**SetPermissionForGroup**](DefaultApi.md#SetPermissionForGroup) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups | 
[**SetPermissionForGroups**](DefaultApi.md#SetPermissionForGroups) | **Put** /api/1.0/admin/permissions/groups | 
[**SetPermissionForGroups_0**](DefaultApi.md#SetPermissionForGroups_0) | **Put** /api/1.0/projects/{projectKey}/permissions/groups | 
[**SetPermissionForUser**](DefaultApi.md#SetPermissionForUser) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users | 
[**SetPermissionForUsers**](DefaultApi.md#SetPermissionForUsers) | **Put** /api/1.0/admin/permissions/users | 
[**SetPermissionForUsers_0**](DefaultApi.md#SetPermissionForUsers_0) | **Put** /api/1.0/projects/{projectKey}/permissions/users | 
[**SetRootLevel**](DefaultApi.md#SetRootLevel) | **Put** /api/1.0/logs/rootLogger/{levelName} | 
[**SetSenderAddress**](DefaultApi.md#SetSenderAddress) | **Put** /api/1.0/admin/mail-server/sender-address | 
[**SetSettings**](DefaultApi.md#SetSettings) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/settings | 
[**SetSettings_0**](DefaultApi.md#SetSettings_0) | **Put** /api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/settings | 
[**Stream**](DefaultApi.md#Stream) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/last-modified | 
[**StreamChanges**](DefaultApi.md#StreamChanges) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/compare/changes | 
[**StreamChanges_0**](DefaultApi.md#StreamChanges_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/changes | 
[**StreamCommits**](DefaultApi.md#StreamCommits) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/compare/commits | 
[**StreamDiff**](DefaultApi.md#StreamDiff) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/diff | 
[**StreamDiff_0**](DefaultApi.md#StreamDiff_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/diff/{path} | 
[**StreamDiff_1**](DefaultApi.md#StreamDiff_1) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/compare/diff{path} | 
[**StreamDiff_2**](DefaultApi.md#StreamDiff_2) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/diff | 
[**StreamDiff_3**](DefaultApi.md#StreamDiff_3) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/diff/{path} | 
[**StreamDiff_4**](DefaultApi.md#StreamDiff_4) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/diff | 
[**StreamDiff_5**](DefaultApi.md#StreamDiff_5) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/diff/{path} | 
[**StreamFiles**](DefaultApi.md#StreamFiles) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/files | 
[**StreamFiles_0**](DefaultApi.md#StreamFiles_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/files/{path} | 
[**Stream_0**](DefaultApi.md#Stream_0) | **Get** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/last-modified/{path} | 
[**TestWebhook**](DefaultApi.md#TestWebhook) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/test | 
[**UnassignParticipantRole**](DefaultApi.md#UnassignParticipantRole) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants | 
[**UnassignParticipantRole_0**](DefaultApi.md#UnassignParticipantRole_0) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug} | 
[**Unwatch**](DefaultApi.md#Unwatch) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/watch | 
[**Unwatch_0**](DefaultApi.md#Unwatch_0) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/watch | 
[**Update**](DefaultApi.md#Update) | **Post** /api/1.0/admin/license | 
[**UpdateComment**](DefaultApi.md#UpdateComment) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId} | 
[**UpdateComment_0**](DefaultApi.md#UpdateComment_0) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId} | 
[**UpdateProject**](DefaultApi.md#UpdateProject) | **Put** /api/1.0/projects/{projectKey} | 
[**UpdatePullRequestSettings**](DefaultApi.md#UpdatePullRequestSettings) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/pull-requests | 
[**UpdatePullRequestSettings_0**](DefaultApi.md#UpdatePullRequestSettings_0) | **Post** /api/1.0/projects/{projectKey}/settings/pull-requests/{scmId} | 
[**UpdateRepository**](DefaultApi.md#UpdateRepository) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug} | 
[**UpdateSettings**](DefaultApi.md#UpdateSettings) | **Post** /api/1.0/users/{userSlug}/settings | 
[**UpdateStatus**](DefaultApi.md#UpdateStatus) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug} | 
[**UpdateTask**](DefaultApi.md#UpdateTask) | **Put** /api/1.0/tasks/{taskId} | 
[**UpdateUserDetails**](DefaultApi.md#UpdateUserDetails) | **Put** /api/1.0/admin/users | 
[**UpdateUserDetails_0**](DefaultApi.md#UpdateUserDetails_0) | **Put** /api/1.0/users | 
[**UpdateUserPassword**](DefaultApi.md#UpdateUserPassword) | **Put** /api/1.0/admin/users/credentials | 
[**UpdateUserPassword_0**](DefaultApi.md#UpdateUserPassword_0) | **Put** /api/1.0/users/credentials | 
[**UpdateWebhook**](DefaultApi.md#UpdateWebhook) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | 
[**Update_0**](DefaultApi.md#Update_0) | **Put** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | 
[**UploadAvatar**](DefaultApi.md#UploadAvatar) | **Post** /api/1.0/projects/{projectKey}/avatar.png | 
[**UploadAvatar_0**](DefaultApi.md#UploadAvatar_0) | **Post** /api/1.0/users/{userSlug}/avatar.png | 
[**Watch**](DefaultApi.md#Watch) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/watch | 
[**Watch_0**](DefaultApi.md#Watch_0) | **Post** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/watch | 
[**WithdrawApproval**](DefaultApi.md#WithdrawApproval) | **Delete** /api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve | 


# **AddGroupToUser**
> AddGroupToUser(ctx, )


<strong>Deprecated since 2.10</strong>. Use /rest/users/add-groups instead.  <p>  Add a user to a group. This is very similar to <code>groups/add-user</code>, but with the <em>context</em> and  <em>itemName</em> attributes of the supplied request entity reversed. On the face of it this may appear  redundant, but it facilitates a specific UI component in Stash.  <p>  In the request entity, the <em>context</em> attribute is the user and the <em>itemName</em> is the group.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUserToGroup**
> AddUserToGroup(ctx, )


<strong>Deprecated since 2.10</strong>. Use /rest/users/add-groups instead.  <p>  Add a user to a group.  <p>  In the request entity, the <em>context</em> attribute is the group and the <em>itemName</em> is the user.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUserToGroups**
> AddUserToGroups(ctx, )


Add a user to one or more groups.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUsersToGroup**
> AddUsersToGroup(ctx, )


Add multiple users to a group.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Approve**
> Approve(ctx, pullRequestId)


Approve a pull request as the current user. Implicitly adds the user as a participant if they are not already.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.  <p>  <strong>Deprecated since 4.2</strong>. Use  /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug} instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignParticipantRole**
> AssignParticipantRole(ctx, pullRequestId)


Assigns a participant to an explicit role in pull request. Currently only the REVIEWER role may be assigned.  <p>  If the user is not yet a participant in the pull request, they are made one and assigned the supplied role.  <p>  If the user is already a participant in the pull request, their previous role is replaced with the supplied role  unless they are already assigned the AUTHOR role which cannot be changed and will result in a Bad Request (400)  response code.  <p>  The authenticated user must have <strong>REPO_WRITE</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CanMerge**
> CanMerge(ctx, pullRequestId)


Test whether a pull request can be merged.  <p>  A pull request may not be merged if:  <ul>      <li>there are conflicts that need to be manually resolved before merging; and/or</li>      <li>one or more merge checks have vetoed the merge.</li>  </ul>  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the ID of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearSenderAddress**
> ClearSenderAddress(ctx, )


Clears the server email address.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearUserCaptchaChallenge**
> ClearUserCaptchaChallenge(ctx, optional)


Clears any CAPTCHA challenge that may constrain the user with the supplied username when they authenticate.  Additionally any counter or metric that contributed towards the user being issued the CAPTCHA challenge  (for instance too many consecutive failed logins) will also be reset.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource, and may not clear  the CAPTCHA of a user with greater permissions than themselves.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountPullRequestTasks**
> CountPullRequestTasks(ctx, )


Retrieve the total number of {@link TaskState#OPEN open} and  {@link TaskState#RESOLVED resolved} tasks associated with a pull request.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create**
> Create(ctx, )


Create a new pull request between two branches. The branches may be in the same repository, or different ones.  When using different repositories, they must still be in the same {@link Repository#getHierarchyId() hierarchy}.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the \"from\" and \"to\"repositories to  call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBranch**
> CreateBranch(ctx, )


Creates a branch using the information provided in the {@link RestCreateBranchRequest request}  <p>  The authenticated user must have <strong>REPO_WRITE</strong> permission for the context repository to call  this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateComment**
> CreateComment(ctx, commitId, optional)


Add a new comment.  <p>  Comments can be added in a few places by setting different attributes:  <p>  General commit comment:   <pre>      {          \"text\": \"An insightful general comment on a commit.\"      }      </pre>   Reply to a comment:   <pre>      {          \"text\": \"A measured reply.\",          \"parent\": {              \"id\": 1          }      }      </pre>   General file comment:   <pre>      {          \"text\": \"An insightful general comment on a file.\",          \"anchor\": {              \"diffType\": \"COMMIT\",              \"fromHash\": \"6df3858eeb9a53a911cd17e66a9174d44ffb02cd\",              \"path\": \"path/to/file\",              \"srcPath\": \"path/to/file\",              \"toHash\": \"04c7c5c931b9418ca7b66f51fe934d0bd9b2ba4b\"          }      }      </pre>   File line comment:   <pre>      {          \"text\": \"A pithy comment on a particular line within a file.\",          \"anchor\": {              \"diffType\": \"COMMIT\",              \"line\": 1,              \"lineType\": \"CONTEXT\",              \"fileType\": \"FROM\",              \"fromHash\": \"6df3858eeb9a53a911cd17e66a9174d44ffb02cd\",              \"path\": \"path/to/file\",              \"srcPath\": \"path/to/file\",              \"toHash\": \"04c7c5c931b9418ca7b66f51fe934d0bd9b2ba4b\"      }      }      </pre>  <strong>Note: general file comments are an experimental feature and may change in the near future!</strong>  <p>  For file and line comments, 'path' refers to the path of the file to which the comment should be applied and  'srcPath' refers to the path the that file used to have (only required for copies and moves). Also,  fromHash and toHash refer to the sinceId / untilId (respectively) used to produce the diff on which the comment  was added. Finally diffType refers to the type of diff the comment was added on.  <p>  For line comments, 'line' refers to the line in the diff that the comment should apply to. 'lineType' refers to  the type of diff hunk, which can be:  <ul>      <li>'ADDED' - for an added line;</li>      <li>'REMOVED' - for a removed line; or</li>      <li>'CONTEXT' - for a line that was unmodified but is in the vicinity of the diff.</li>  </ul>  'fileType' refers to the file of the diff to which the anchor should be attached - which is of relevance when  displaying the diff in a side-by-side way. Currently the supported values are:  <ul>      <li>'FROM' - the source file of the diff</li>      <li>'TO' - the destination file of the diff</li>  </ul>  If the current user is not a participant the user is added as one and updated to watch the commit.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that the commit  is in to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the commit to which the comments must be anchored | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitId** | **string**| the commit to which the comments must be anchored | 
 **since** | **string**| For a merge commit, a parent can be provided to specify which diff the comments should be on. For                  a commit range, a {@code sinceId} can be provided to specify where the comments should be                  anchored from. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateComment_0**
> CreateComment_0(ctx, )


Add a new comment.  <p>  Comments can be added in a few places by setting different attributes:  <p>  General pull request comment:   <pre>      {          \"text\": \"An insightful general comment on a pull request.\"      }      </pre>   Reply to a comment:   <pre>      {          \"text\": \"A measured reply.\",          \"parent\": {              \"id\": 1          }      }      </pre>   General file comment:   <pre>      {          \"text\": \"An insightful general comment on a file.\",          \"anchor\": {              \"diffType\": \"RANGE\",              \"fromHash\": \"6df3858eeb9a53a911cd17e66a9174d44ffb02cd\",              \"path\": \"path/to/file\",              \"srcPath\": \"path/to/file\",              \"toHash\": \"04c7c5c931b9418ca7b66f51fe934d0bd9b2ba4b\"          }      }      </pre>   File line comment:   <pre>      {          \"text\": \"A pithy comment on a particular line within a file.\",          \"anchor\": {              \"diffType\": \"COMMIT\",              \"line\": 1,              \"lineType\": \"CONTEXT\",              \"fileType\": \"FROM\",              \"fromHash\": \"6df3858eeb9a53a911cd17e66a9174d44ffb02cd\",              \"path\": \"path/to/file\",              \"srcPath\": \"path/to/file\",              \"toHash\": \"04c7c5c931b9418ca7b66f51fe934d0bd9b2ba4b\"          }      }      </pre>  <p>  For file and line comments, 'path' refers to the path of the file to which the comment should be applied and  'srcPath' refers to the path the that file used to have (only required for copies and moves). Also,  fromHash and toHash refer to the sinceId / untilId (respectively) used to produce the diff on which the comment  was added. Finally diffType refers to the type of diff the comment was added on. For backwards compatibility  purposes if no diffType is provided and no fromHash/toHash pair is provided the diffType will be resolved to  'EFFECTIVE'. In any other cases the diffType is REQUIRED.  <p>  For line comments, 'line' refers to the line in the diff that the comment should apply to. 'lineType' refers to  the type of diff hunk, which can be:  <ul>      <li>'ADDED' - for an added line;</li>      <li>'REMOVED' - for a removed line; or</li>      <li>'CONTEXT' - for a line that was unmodified but is in the vicinity of the diff.</li>  </ul>  'fileType' refers to the file of the diff to which the anchor should be attached - which is of relevance when  displaying the diff in a side-by-side way. Currently the supported values are:  <ul>      <li>'FROM' - the source file of the diff</li>      <li>'TO' - the destination file of the diff</li>  </ul>  If the current user is not a participant the user is added as a watcher of the pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> CreateGroup(ctx, optional)


Create a new group.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of the group. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> CreateProject(ctx, )


Create a new project.  <p>  To include a custom avatar for the project, the project definition should contain an additional attribute with  the key <code>avatar</code> and the value a data URI containing Base64-encoded image data. The URI should be in  the following format:  <pre>      data:(content type, e.g. image/png);base64,(data)  </pre>  If the data is not Base64-encoded, or if a character set is defined in the URI, or the URI is otherwise invalid,  <em>project creation will fail</em>.  <p>  The authenticated user must have <strong>PROJECT_CREATE</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository**
> CreateRepository(ctx, projectKey)


Create a new repository. Requires an existing project in which this repository will be created.  The only parameters which will be used are name and scmId.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the context project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTag**
> CreateTag(ctx, )


Creates a tag using the information provided in the {@link RestCreateTagRequest request}  <p>  The authenticated user must have <strong>REPO_WRITE</strong> permission for the context repository to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTask**
> CreateTask(ctx, )


Create a new task.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> CreateUser(ctx, optional)


Creates a new user from the assembled query parameters.  <p>  The default group can be used to control initial permissions for new users, such as granting users the ability  to login or providing read access to certain projects or repositories. If the user is not added to the default  group, they may not be able to login after their account is created until explicit permissions are configured.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the username for the new user | 
 **password** | **string**| the password for the new user | 
 **displayName** | **string**| the display name for the new user | 
 **emailAddress** | **string**| the e-mail address for the new user | 
 **addToDefaultGroup** | **bool**| &lt;code&gt;true&lt;/code&gt; to add the user to the default group, which can be used to grant them                           a set of initial permissions; otherwise, &lt;code&gt;false&lt;/code&gt; to not add them to a group | [default to true]
 **notify** | **string**| if present and not &lt;code&gt;false&lt;/code&gt; instead of requiring a password,                           the create user will be notified via email their account has been created and requires                           a password to be reset. This option can only be used if a mail server has been configured | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWebhook**
> CreateWebhook(ctx, )


Create a webhook for the repository specified via the URL.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Decline**
> Decline(ctx, pullRequestId, optional)


Decline a pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pullRequestId** | **int64**|  | 
 **version** | **int32**| the current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified                 version the operation will fail. To determine the current version of the pull request it should be                 fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned                 JSON structure. | [default to -1]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, pullRequestId)


Deletes a pull request.  <p>  To call this resource, users must be authenticated and have permission to view the pull request.  Additionally, they must:  <ul>      <li>          be the pull request author, if the system is configured to allow authors to delete their own          pull requests (this is the default) OR      </li>      <li>have repository administrator permission for the repository the pull request is targeting</li>  </ul>   A body containing the version of the pull request must be provided with this request.   <pre><code>{ \"version\": 1 }</code></pre>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the ID of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAvatar**
> DeleteAvatar(ctx, )


Delete the avatar associated to a user.  <p>  Users are always allowed to delete their own avatar. To delete someone else's avatar the authenticated user must  have global <strong>ADMIN</strong> permission, or global <strong>SYS_ADMIN</strong> permission to update a  <strong>SYS_ADMIN</strong> user's avatar.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComment**
> DeleteComment(ctx, commitId, commentId, commitId2, optional)


Delete a commit comment. Anyone can delete their own comment. Only users with <strong>REPO_ADMIN</strong>  and above may delete comments created by other users. Comments which have replies <i>may not be deleted</i>,  regardless of the user's granted permissions.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that the commit  is in to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the commit to which the comments must be anchored | 
  **commentId** | **int64**| the ID of the comment to retrieve | 
  **commitId2** | **string**| the &lt;i&gt;full {@link Commit#getId() ID}&lt;/i&gt; of the commit within the repository | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitId** | **string**| the commit to which the comments must be anchored | 
 **commentId** | **int64**| the ID of the comment to retrieve | 
 **commitId2** | **string**| the &lt;i&gt;full {@link Commit#getId() ID}&lt;/i&gt; of the commit within the repository | 
 **version** | **int32**| The expected version of the comment. This must match the server&#39;s version of the comment or                   the delete will fail. To determine the current version of the comment, the comment should be                   fetched from the server prior to the delete. Look for the &#39;version&#39; attribute in the returned                   JSON structure. | [default to -1]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComment_0**
> DeleteComment_0(ctx, commentId, optional)


Delete a pull request comment. Anyone can delete their own comment. Only users with <strong>REPO_ADMIN</strong>  and above may delete comments created by other users.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commentId** | **int64**| the id of the comment to retrieve | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentId** | **int64**| the id of the comment to retrieve | 
 **version** | **int32**| The expected version of the comment. This must match the server&#39;s version of the comment or                     the delete will fail. To determine the current version of the comment, the comment should be                     fetched from the server prior to the delete. Look for the &#39;version&#39; attribute in the                     returned JSON structure. | [default to -1]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, optional)


Deletes the specified group, removing them from the system. This also removes any permissions that may have been  granted to the group.  <p>  A user may not delete the last group that is granting them administrative permissions, or a group with greater  permissions than themselves.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the name identifying the group to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMailConfig**
> DeleteMailConfig(ctx, )


Deletes the current mail configuration.  <p>  The authenticated user must have the <strong>SYS_ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, )


Delete the project matching the supplied <strong>projectKey</strong>.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepository**
> DeleteRepository(ctx, projectKey, projectKey2, repositorySlug)


Schedule the repository matching the supplied <strong>projectKey</strong> and <strong>repositorySlug</strong> to  be deleted. If the request repository is not present  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 
  **projectKey2** | **string**| the parent project key | 
  **repositorySlug** | **string**| the repository slug | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryHook**
> DeleteRepositoryHook(ctx, hookKey)


Delete repository hook configuration for the supplied <strong>hookKey</strong> and <strong>repositorySlug</strong>  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTask**
> DeleteTask(ctx, taskId)


Delete a task.  <p>  Note that only the task's creator, the context's author or an admin of the context's repository can delete a  task. (For a pull request task, those are the task's creator, the pull request's author or an admin on the  repository containing the pull request). Additionally a task cannot be deleted if it has already been resolved.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **taskId** | **int64**| the id identifying the task to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, optional)


Deletes the specified user, removing them from the system. This also removes any permissions that may have been  granted to the user.  <p>  A user may not delete themselves, and a user with <strong>ADMIN</strong> permissions may not delete a user with  <strong>SYS_ADMIN</strong>permissions.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the username identifying the user to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook(ctx, webhookId)


Delete a webhook for the repository specified via the URL.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webhookId** | **int32**| the existing webhook id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableHook**
> DisableHook(ctx, hookKey)


Disable a repository hook for this repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableHook_0**
> DisableHook_0(ctx, hookKey)


Disable a repository hook for this project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditFile**
> EditFile(ctx, path)


Update the content of <code>path</code>, on the given <code>repository</code> and <code>branch</code>.  <p>  This resource accepts PUT multipart form data, containing the file in a form-field named <code>content</code>.  <p>  An example <a href=\"http://curl.haxx.se/\">curl</a> request to update 'README.md' would be:  <pre>  curl -X PUT -u username:password -F content=@README.md  -F 'message=Updated using file-edit REST API'  -F branch=master -F  sourceCommitId=5636641a50b   http://example.com/rest/api/latest/projects/PROJECT_1/repos/repo_1/browse/README.md  </pre>  <p>  <ui>  <li>branch:  the branch on which the <code>path</code> should be modified or created</li>  <li>content: the full content of the file at <code>path</code> </li>  <li>message: the message associated with this change, to be used as the commit message.  Or null if the default message should be used.</li>  <li>sourceCommitId: the commit ID of the file before it was edited, used to identify if  content has changed. Or null if this is a new file</li>  </ui>  <p>  The file can be updated or created on a new branch. In this case, the <code>sourceBranch</code> parameter should  be provided to identify the starting point for the new branch and the <code>branch</code> parameter identifies  the branch to create the new commit on.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the file path to retrieve content from | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableHook**
> EnableHook(ctx, hookKey, optional)


Enable a repository hook for this repository and optionally apply new configuration.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.  <p>  A JSON document may be provided to use as the settings for the hook. These structure and validity of  the document is decided by the plugin providing the hook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hookKey** | **string**|  | 
 **contentLength** | **int32**|  | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableHook_0**
> EnableHook_0(ctx, hookKey, optional)


Enable a repository hook for this project and optionally apply new configuration.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project to call this  resource.  <p>  A JSON document may be provided to use as the settings for the hook. These structure and validity of  the document is decided by the plugin providing the hook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hookKey** | **string**|  | 
 **contentLength** | **int32**|  | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindGroupsForUser**
> FindGroupsForUser(ctx, optional)


Retrieves a list of groups the specified user is a member of.  <p>  The authenticated user must have the <strong>LICENSED_USER</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string**| the user which should be used to locate groups | 
 **filter** | **string**| if specified only groups with names containing the supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOtherGroupsForUser**
> FindOtherGroupsForUser(ctx, optional)


Retrieves a list of groups the specified user is <em>not</em> a member of.  <p>  The authenticated user must have the <strong>LICENSED_USER</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string**| the user which should be used to locate groups | 
 **filter** | **string**| if specified only groups with names containing the supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsersInGroup**
> FindUsersInGroup(ctx, optional)


Retrieves a list of users that are members of a specified group.  <p>  The authenticated user must have the <strong>LICENSED_USER</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string**| the group which should be used to locate members | 
 **filter** | **string**| if specified only users with usernames, display names or email addresses containing the                   supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsersNotInGroup**
> FindUsersNotInGroup(ctx, optional)


Retrieves a list of users that are <em>not</em> members of a specified group.  <p>  The authenticated user must have the <strong>LICENSED_USER</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string**| the group which should be used to locate non-members | 
 **filter** | **string**| if specified only users with usernames, display names or email addresses containing the                   supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindWebhooks**
> FindWebhooks(ctx, optional)


Find webhooks in this repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **event** | **string**| list of {@link com.atlassian.webhooks.WebhookEvent} ids to filter for | 
 **statistics** | **bool**| {@code true} if statistics should be provided for all found webhooks | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForkRepository**
> ForkRepository(ctx, projectKey, projectKey2, repositorySlug)


Create a new repository forked from an existing repository.  <p>  The JSON body for this {@code POST} is not required to contain <i>any</i> properties. Even the name may be  omitted. The following properties will be used, if provided:  <ul>      <li>{@code \"name\":\"Fork name\"} - Specifies the forked repository's name      <ul>          <li>Defaults to the name of the origin repository if not specified</li>      </ul>      </li>      <li>{@code \"project\":{\"key\":\"TARGET_KEY\"}} - Specifies the forked repository's target project by key      <ul>          <li>Defaults to the current user's personal project if not specified</li>      </ul>      </li>  </ul>  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository and  <strong>PROJECT_ADMIN</strong> on the target project to call this resource. Note that users <i>always</i>  have <b>PROJECT_ADMIN</b> permission on their personal projects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 
  **projectKey2** | **string**| the parent project key | 
  **repositorySlug** | **string**| the repository slug | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> Get(ctx, )


Retrieves details about the current license, as well as the current status of the system with regards to the  installed license. The status includes the current number of users applied toward the license limit, as well  as any status messages about the license (warnings about expiry or user counts exceeding license limits).  <p>  The authenticated user must have <b>ADMIN</b> permission. Unauthenticated users, and non-administrators, are  not permitted to access license details.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivities**
> GetActivities(ctx, pullRequestId, optional)


Retrieve a page of activity associated with a pull request.  <p>  Activity items include comments, approvals, rescopes (i.e. adding and removing of commits), merges and more.  <p>  Different types of activity items may be introduced in newer versions of Stash or by user installed plugins, so  clients should be flexible enough to handle unexpected entity shapes in the returned page.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pullRequestId** | **int64**| the id of the pull request within the repository | 
 **fromId** | **int64**| (optional) the id of the activity item to use as the first item in the returned page | 
 **fromType** | **string**| (required if &lt;strong&gt;fromId&lt;/strong&gt; is present) the type of the activity item specified by                  &lt;strong&gt;fromId&lt;/strong&gt; (either &lt;strong&gt;COMMENT&lt;/strong&gt; or &lt;strong&gt;ACTIVITY&lt;/strong&gt;) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationProperties**
> GetApplicationProperties(ctx, )


Retrieve version information and other application properties.  <p>  No authentication is required to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArchive**
> GetArchive(ctx, optional)


Streams an archive of the repository's contents at the requested commit. If no <code>at=</code> commit is  requested, an archive of the default branch is streamed.  <p>  The <code>filename=</code> query parameter may be used to specify the exact filename to include in the  <code>\"Content-Disposition\"</code> header. If an explicit filename is not provided, one will be automatically  generated based on what is being archived. Its format depends on the <code>at=</code> value:  <ul>      <li>No <code>at=</code> commit:      <code>&lt;slug&gt;-&lt;default-branch-name&gt;@&lt;commit&gt;.&lt;format&gt;</code>;      e.g. example-master@43c2f8a0fe8.zip</li>      <li><code>at=sha</code>: <code>&lt;slug&gt;-&lt;at&gt;.&lt;format&gt;</code>; e.g.      example-09bcbb00100cfbb5310fb6834a1d5ce6cac253e9.tar.gz</li>      <li><code>at=branchOrTag</code>: <code>&lt;slug&gt;-&lt;branchOrTag&gt;@&lt;commit&gt;.&lt;format&gt;</code>;      e.g. example-feature@bbb225f16e1.tar      <ul>          <li>If the branch or tag is qualified (e.g. <code>refs/heads/master</code>, the short name          (<code>master</code>) will be included in the filename</li>          <li>If the branch or tag's <i>short name</i> includes slashes (e.g. <code>release/4.6</code>),          they will be converted to hyphens in the filename (<code>release-4.5</code>)</li>      </ul>      </li>  </ul>  <p>  Archives may be requested in the following formats by adding the <code>format=</code> query parameter:  <ul>      <li><code>zip</code>: A zip file using standard compression (Default)</li>      <li><code>tar</code>: An uncompressed tarball</li>      <li><code>tar.gz</code> or <code>tgz</code>: A GZip-compressed tarball</li>  </ul>  The contents of the archive may be filtered by using the <code>path=</code> query parameter to specify paths to  include. <code>path=</code> may be specified multiple times to include multiple paths.  <p>  The <code>prefix=</code> query parameter may be used to define a directory (or multiple directories) where  the archive's contents should be placed. If the prefix does not end with <code>/</code>, one will be added  automatically. The prefix is <i>always</i> treated as a directory; it is not possible to use it to prepend  characters to the entries in the archive.  <p>  Archives of public repositories may be streamed by any authenticated or anonymous user. Streaming archives for  non-public repositories requires an <i>authenticated user</i> with at least <b>REPO_READ</b> permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **string**| the commit to stream an archive of; if not supplied, an archive of the default branch is streamed | 
 **filename** | **string**| a filename to include the \&quot;Content-Disposition\&quot; header | 
 **format** | **string**| the format to stream the archive in; must be one of: zip, tar, tar.gz or tgz | 
 **path** | **string**| paths to include in the streamed archive; may be repeated to include multiple paths | 
 **prefix** | **string**| a prefix to apply to all entries in the streamed archive; if the supplied prefix does not end                  with a trailing &lt;code&gt;/&lt;/code&gt;, one will be added automatically | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvatar**
> GetAvatar(ctx, hookKey, optional)


Retrieve the avatar for the project matching the supplied <strong>moduleKey</strong>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**| the complete module key of the hook module | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hookKey** | **string**| the complete module key of the hook module | 
 **version** | **string**| optional version used for HTTP caching only - any non-blank version will result in a large max-age Cache-Control header.                 Note that this does not affect the Last-Modified header. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranches**
> GetBranches(ctx, optional)


Retrieve the branches matching the supplied <strong>filterText</strong> param.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **base** | **string**| base branch or tag to compare each branch to (for the metadata providers that uses that information) | 
 **details** | **bool**| whether to retrieve plugin-provided metadata about each branch | 
 **filterText** | **string**| the text to match on | 
 **orderBy** | **string**| ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChanges**
> GetChanges(ctx, optional)


Retrieve a page of changes made in a specified commit.  <p>  <strong>Note:</strong> The implementation will apply a hard cap ({@code page.max.changes}) and it is not  possible to request subsequent content when that cap is exceeded.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string**| the commit to which &lt;code&gt;until&lt;/code&gt; should be compared to produce a page of changes.                      If not specified the commit&#39;s first parent is assumed (if one exists) | 
 **until** | **string**| the commit to retrieve changes for | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChanges_0**
> GetChanges_0(ctx, commitId, optional)


Retrieve a page of changes made in a specified commit.  <p>  <strong>Note:</strong> The implementation will apply a hard cap ({@code page.max.changes}) and it is not  possible to request subsequent content when that cap is exceeded.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the commit to retrieve changes for | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitId** | **string**| the commit to retrieve changes for | 
 **since** | **string**| the commit to which &lt;code&gt;until&lt;/code&gt; should be compared to produce a page of changes.                       If not specified the commit&#39;s first parent is assumed (if one exists) | 
 **withComments** | **bool**| {@code true} to apply comment counts in the changes (the default); otherwise, {@code false}                       to stream changes without comment counts | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComment**
> GetComment(ctx, commitId, commentId, commitId2)


Retrieves a commit discussion comment.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that the commit  is in to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the commit to which the comments must be anchored | 
  **commentId** | **int64**| the ID of the comment to retrieve | 
  **commitId2** | **string**| the &lt;i&gt;full {@link Commit#getId() ID}&lt;/i&gt; of the commit within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComment_0**
> GetComment_0(ctx, commentId)


Retrieves a pull request comment.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commentId** | **int64**| the id of the comment to retrieve | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComments**
> GetComments(ctx, commitId, optional)


Retrieves the commit discussion comments that match the specified search criteria.  <p>  It is possible to retrieve commit discussion comments that are anchored to a range of commits by providing the  {@code sinceId} that the comments anchored from.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that the commit  is in to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the commit to which the comments must be anchored | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitId** | **string**| the commit to which the comments must be anchored | 
 **path** | **string**| the path to the file on which comments were made | 
 **since** | **string**| For a merge commit, a parent can be provided to specify which diff the comments are on. For                       a commit range, a {@code sinceId} can be provided to specify where the comments are anchored                       from. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComments_0**
> GetComments_0(ctx, optional)


Gets comments for the specified PullRequest.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anchorState** | **string**| {@code ACTIVE} to stream the active comments;                     {@code ORPHANED} to stream the orphaned comments;                     {@code ALL} to stream both the active and the orphaned comments; | [default to ACTIVE]
 **diffType** | **string**| {@code EFFECTIVE} to stream the comments related to the effective diff of the pull request;                     {@code RANGE} to stream comments related to a commit range between two arbitrary commits                                   (requires {@code fromHash} and {@code toHash});                     {@code COMMIT} to stream comments related to a commit between two arbitrary commits (requires                         {@code fromHash} and {@code toHash}) | 
 **fromHash** | **string**| the from commit hash to stream comments for a {@code RANGE} or {@code COMMIT} arbitrary change scope | 
 **path** | **string**| the path to stream comments for a given path | 
 **toHash** | **string**| the to commit hash to stream comments for a {@code RANGE} or {@code COMMIT} arbitrary change scope | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommit**
> GetCommit(ctx, commitId, optional)


Retrieve a single commit <i>identified by its ID</i>>. In general, that ID is a SHA1. <u>From 2.11, ref names  like \"refs/heads/master\" are no longer accepted by this resource.</u>  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the commit ID to retrieve | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitId** | **string**| the commit ID to retrieve | 
 **path** | **string**| an optional path to filter the commit by. If supplied the details returned &lt;i&gt;may not&lt;/i&gt;              be for the specified commit. Instead, starting from the specified commit, they will be the              details for the first commit affecting the specified path. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommits**
> GetCommits(ctx, optional)


Retrieve a page of commits from a given starting commit or \"between\" two commits. If no explicit commit is  specified, the tip of the repository's default branch is assumed. commits may be identified by branch or tag  name or by ID. A path may be supplied to restrict the returned commits to only those which affect that path.  <p>  The authenticated user must have <b>REPO_READ</b> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **followRenames** | **bool**| if &lt;code&gt;true&lt;/code&gt;, the commit history of the specified file will be followed past renames.                       Only valid for a path to a single file. | [default to false]
 **ignoreMissing** | **bool**| &lt;code&gt;true&lt;/code&gt; to ignore missing commits, &lt;code&gt;false&lt;/code&gt; otherwise | [default to false]
 **merges** | **string**| if present, controls how merge commits should be filtered. Can be either &lt;code&gt;exclude&lt;/code&gt;,                to exclude merge commits, &lt;code&gt;include&lt;/code&gt;, to include both merge commits and non-merge                commits or &lt;code&gt;only&lt;/code&gt;, to only return merge commits. | 
 **path** | **string**| an optional path to filter commits by | 
 **since** | **string**| the commit ID or ref (exclusively) to retrieve commits after | 
 **until** | **string**| the commit ID (SHA1) or ref (inclusively) to retrieve commits before | 
 **withCounts** | **bool**| optionally include the total number of commits and total number of unique authors | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommits_0**
> GetCommits_0(ctx, pullRequestId, optional)


Retrieve commits for the specified pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pullRequestId** | **int64**|  | 
 **withCounts** | **bool**| if set to true, the service will add \&quot;authorCount\&quot; and \&quot;totalCount\&quot; at the end of the page.                      \&quot;authorCount\&quot; is the number of different authors and \&quot;totalCount\&quot; is the total number of commits. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent**
> GetContent(ctx, optional)


Retrieve a page of content for a file path at a specified revision.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **string**| the commit ID or ref to retrieve the content for. | [default to ]
 **type_** | **bool**| if true only the type will be returned for the file path instead of the contents. | [default to false]
 **blame** | **string**| if present the blame will be returned for the file as well. | 
 **noContent** | **string**| if present and used with blame only the blame is retrieved instead of the contents. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent_0**
> GetContent_0(ctx, path, optional)


Retrieve a page of content for a file path at a specified revision.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the file path to retrieve content from | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| the file path to retrieve content from | 
 **at** | **string**| the commit ID or ref to retrieve the content for. | [default to ]
 **type_** | **bool**| if true only the type will be returned for the file path instead of the contents. | [default to false]
 **blame** | **string**| if present the blame will be returned for the file as well. | 
 **noContent** | **string**| if present and used with blame only the blame is retrieved instead of the contents. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent_1**
> GetContent_1(ctx, optional)


Retrieve the raw content for a file path at a specified revision.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **string**| the commit ID or ref to retrieve the content for. | 
 **markup** | **string**| if present or &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, triggers the raw content to be markup-rendered and returned                    as HTML; otherwise, if not specified, or any value other than &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, the content                    is streamed without markup | 
 **hardwrap** | **bool**| (Optional) Whether the markup implementation should convert newlines to breaks.                    If not specified, {@link MarkupService} will use the value of the                    &lt;code&gt;markup.render.hardwrap&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default | 
 **htmlEscape** | **bool**| (Optional) true if HTML should be escaped in the input markup, false otherwise.                    If not specified, {@link MarkupService} will use the value of the                    &lt;code&gt;markup.render.html.escape&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent_2**
> GetContent_2(ctx, path, optional)


Retrieve the raw content for a file path at a specified revision.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the file path to retrieve content from | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| the file path to retrieve content from | 
 **at** | **string**| the commit ID or ref to retrieve the content for. | 
 **markup** | **string**| if present or &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, triggers the raw content to be markup-rendered and returned                    as HTML; otherwise, if not specified, or any value other than &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, the content                    is streamed without markup | 
 **hardwrap** | **bool**| (Optional) Whether the markup implementation should convert newlines to breaks.                    If not specified, {@link MarkupService} will use the value of the                    &lt;code&gt;markup.render.hardwrap&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default | 
 **htmlEscape** | **bool**| (Optional) true if HTML should be escaped in the input markup, false otherwise.                    If not specified, {@link MarkupService} will use the value of the                    &lt;code&gt;markup.render.html.escape&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultBranch**
> GetDefaultBranch(ctx, )


Get the default branch of the repository.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForkedRepositories**
> GetForkedRepositories(ctx, projectKey)


Retrieve repositories which have been forked from this one. Unlike {@link #getRelatedRepositories(Repository,  PageRequest) related repositories}, this only looks at a given repository's direct forks. If those forks have  themselves been the origin of more forks, such \"grandchildren\" repositories will not be retrieved.  <p>  Only repositories to which the authenticated user has <b>REPO_READ</b> permission will be included, even  if other repositories have been forked from this one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups**
> GetGroups(ctx, optional)


Retrieve a page of groups.  <p>  The authenticated user must have <strong>LICENSED_USER</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsWithAnyPermission**
> GetGroupsWithAnyPermission(ctx, optional)


Retrieve a page of groups that have been granted at least one global permission.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsWithAnyPermission_0**
> GetGroupsWithAnyPermission_0(ctx, optional)


Retrieve a page of groups that have been granted at least one permission for the specified project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsWithAnyPermission_1**
> GetGroupsWithAnyPermission_1(ctx, optional)


Retrieve a page of groups that have been granted at least one permission for the specified repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsWithoutAnyPermission**
> GetGroupsWithoutAnyPermission(ctx, optional)


Retrieve a page of groups that have no granted global permissions.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsWithoutAnyPermission_0**
> GetGroupsWithoutAnyPermission_0(ctx, optional)


Retrieve a page of groups that have no granted permissions for the specified project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsWithoutAnyPermission_1**
> GetGroupsWithoutAnyPermission_1(ctx, optional)


Retrieve a page of groups that have no granted permissions for the specified repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups_0**
> GetGroups_0(ctx, optional)


Retrieve a page of group names.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInformation**
> GetInformation(ctx, )


Gets information about the nodes that currently make up the stash cluster.  <p>  The authenticated user must have the <strong>SYS_ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestInvocation**
> GetLatestInvocation(ctx, webhookId, optional)


Get the latest invocations for a specific webhook.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webhookId** | **int32**| id of the webhook | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **int32**| id of the webhook | 
 **event** | **string**| the string id of a specific event to retrieve the last invocation for. | 
 **outcome** | **string**| the outcome to filter for. Can be SUCCESS, FAILURE, ERROR. None specified means that the all                   will be considered | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLevel**
> GetLevel(ctx, loggerName)


Retrieve the current log level for a given logger.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **loggerName** | **string**| the name of the logger. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMailConfig**
> GetMailConfig(ctx, )


Retrieves the current mail configuration.   The authenticated user must have the <strong>SYS_ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMergeConfig**
> GetMergeConfig(ctx, scmId)


Retrieve the merge strategies available for this instance.  <p>  The user must be authenticated to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scmId** | **string**| the id of the scm to get strategies for | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPage**
> GetPage(ctx, optional)


Retrieve a page of pull requests to or from the specified repository.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call  this resource.   Optionally clients can specify PR participant filters. Each filter has a mandatory {@code username.N}  parameter, and the optional {@code role.N} and {@code approved.N} parameters.  <ul>      <li>          {@code username.N} - the \"root\" of a single participant filter, where \"N\" is a natural number          starting from 1. This allows clients to specify multiple participant filters, by providing consecutive          filters as {@code username.1}, {@code username.2} etc. Note that the filters numbering has to start          with 1 and be continuous for all filters to be processed. The total allowed number of participant          filters is 10 and all filters exceeding that limit will be dropped.      </li>      <li>          {@code role.N}(optional) the role associated with {@code username.N}.          This must be one of {@code AUTHOR}, {@code REVIEWER}, or{@code PARTICIPANT}      </li>      <li>          {@code approved.N}(optional) the approved status associated with {@code username.N}.          That is whether {@code username.N} has approved the PR. Either {@code true}, or {@code false}      </li>  </ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **direction** | **string**| (optional, defaults to &lt;strong&gt;INCOMING&lt;/strong&gt;) the direction relative to the specified                   repository. Either &lt;strong&gt;INCOMING&lt;/strong&gt; or &lt;strong&gt;OUTGOING&lt;/strong&gt;. | [default to incoming]
 **at** | **string**| (optional) a &lt;i&gt;fully-qualified&lt;/i&gt; branch ID to find pull requests to or from,            such as {@code refs/heads/master} | 
 **state** | **string**| (optional, defaults to &lt;strong&gt;OPEN&lt;/strong&gt;). Supply &lt;strong&gt;ALL&lt;/strong&gt; to return pull request                in any state. If a state is supplied only pull requests in the specified state will be returned.                Either &lt;strong&gt;OPEN&lt;/strong&gt;, &lt;strong&gt;DECLINED&lt;/strong&gt; or &lt;strong&gt;MERGED&lt;/strong&gt;. | 
 **order** | **string**| (optional, defaults to &lt;strong&gt;NEWEST&lt;/strong&gt;) the order to return pull requests in, either               &lt;strong&gt;OLDEST&lt;/strong&gt; (as in: \&quot;oldest first\&quot;) or &lt;strong&gt;NEWEST&lt;/strong&gt;. | 
 **withAttributes** | **bool**| (optional) defaults to true, whether to return additional pull request attributes | [default to true]
 **withProperties** | **bool**| (optional) defaults to true, whether to return additional pull request properties | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> GetProject(ctx, )


Retrieve the project matching the supplied <strong>projectKey</strong>.  <p>  The authenticated user must have <strong>PROJECT_VIEW</strong> permission for the specified project to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectAvatar**
> GetProjectAvatar(ctx, optional)


Retrieve the avatar for the project matching the supplied <strong>projectKey</strong>.  <p>  The authenticated user must have <strong>PROJECT_VIEW</strong> permission for the specified project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **int32**| The desired size of the image. The server will return an image as close as possible to the specified              size. | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**
> GetProjects(ctx, optional)


Retrieve a page of projects.  <p>  Only projects for which the authenticated user has the <strong>PROJECT_VIEW</strong> permission will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **permission** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestCount**
> GetPullRequestCount(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestSettings**
> GetPullRequestSettings(ctx, )


Retrieve the pull request settings for the context repository.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the context repository to call this  resource.  <p>  This resource will call all RestFragments that are registered with the key  <strong>bitbucket.repository.settings.pullRequests</strong>. If any fragment fails validations by returning a  non-empty Map of errors, then no fragments will execute.  <p>  The property keys for the settings that are bundled with the application are  <ul>      <li>mergeConfig - the merge strategy configuration for pull requests</li>      <li>requiredApprovers - (Deprecated, please use com.atlassian.bitbucket.server.bundled-hooks.requiredApproversMergeHook instead) the number of approvals required on a pull request for it to be mergeable, or 0 if the merge check is disabled</li>      <li>com.atlassian.bitbucket.server.bundled-hooks.requiredApproversMergeHook - the merge check configuration for required approvers</li>      <li>requiredAllApprovers - whether or not all approvers must approve a pull request for it to be mergeable</li>      <li>requiredAllTasksComplete - whether or not all tasks on a pull request need to be completed for it to be mergeable</li>      <li>requiredSuccessfulBuilds - (Deprecated, please use com.atlassian.bitbucket.server.bitbucket-build.requiredBuildsMergeCheck instead) the number of successful builds on a pull request for it to be mergeable, or 0 if the merge check is disabled</li>      <li>com.atlassian.bitbucket.server.bitbucket-build.requiredBuildsMergeCheck - the merge check configuration for required builds</li>  </ul>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestSettings_0**
> GetPullRequestSettings_0(ctx, scmId)


Retrieve the merge strategy configuration for this project and SCM.  <p>  The authenticated user must have <strong>PROJECT_READ</strong> permission for the context repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scmId** | **string**| the SCM to get strategies for | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestSuggestions**
> GetPullRequestSuggestions(ctx, optional)


Retrieves a page of suggestions for pull requests that the currently authenticated user may wish to  raise. Such suggestions are based on ref changes occurring and so contain the ref change that  prompted the suggestion plus the time the change event occurred. Changes will be returned in  descending order based on the time the change that prompted the suggestion occurred.  <p>  Note that although the response is a page object, the interface does not support paging, however  a limit can be applied to the size of the returned page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changesSince** | **string**| restrict pull request suggestions to be based on events that occurred since some time                      in the past. This is expressed in seconds since \&quot;now\&quot;. So to return suggestions                      based only on activity within the past 48 hours, pass a value of 172800. | [default to 172800]
 **limit** | **int32**| restricts the result set to return at most this many suggestions. | [default to 3]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestTasks**
> GetPullRequestTasks(ctx, )


Retrieve the tasks associated with a pull request.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequests**
> GetPullRequests(ctx, optional)


Retrieve a page of pull requests where the current authenticated user is involved as either a reviewer, author  or a participant. The request may be filtered by pull request state, role or participant status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string**| (optional, defaults to returning pull requests in any state). If a state is supplied only pull               requests in the specified state will be returned. Either &lt;strong&gt;OPEN&lt;/strong&gt;,               &lt;strong&gt;DECLINED&lt;/strong&gt; or &lt;strong&gt;MERGED&lt;/strong&gt;.               Omit this parameter to return pull request in any state. | 
 **role** | **string**| (optional, defaults to returning pull requests for any role). If a role is supplied only pull               requests where the authenticated user is a participant in the given role will be returned.               Either &lt;strong&gt;REVIEWER&lt;/strong&gt;, &lt;strong&gt;AUTHOR&lt;/strong&gt; or &lt;strong&gt;PARTICIPANT&lt;/strong&gt;. | 
 **participantStatus** | **string**| (optional, defaults to returning pull requests with any participant status). A comma                           separated list of participant status. That is, one or more of                           &lt;strong&gt;UNAPPROVED&lt;/strong&gt;, &lt;strong&gt;NEEDS_WORK&lt;/strong&gt;, or &lt;strong&gt;APPROVED&lt;/strong&gt;. | 
 **order** | **string**| (optional, defaults to &lt;strong&gt;NEWEST&lt;/strong&gt;) the order to return pull requests in, either               &lt;strong&gt;OLDEST&lt;/strong&gt; (as in: \&quot;oldest first\&quot;), &lt;strong&gt;NEWEST&lt;/strong&gt;,               &lt;strong&gt;PARTICIPANT_STATUS&lt;/strong&gt;, or &lt;strong&gt;CLOSED_DATE&lt;/strong&gt;. Where               &lt;strong&gt;CLOSED_DATE&lt;/strong&gt; is specified and the result set includes pull requests that are not in               the closed state, these pull requests will appear first in the result set, followed by most recently               closed pull requests. | 
 **closedSince** | **string**| (optional, defaults to returning pull requests regardless of closed since date). Permits                       returning only pull requests with a closed timestamp set more recently that                       (now - closedSince). Units are in seconds. So for example if closed since 86400 is set only                       pull requests closed in the previous 24 hours will be returned. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequests_0**
> GetPullRequests_0(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32**|  | [default to 0]
 **limit** | **int32**|  | [default to 25]
 **role** | **string**|  | [default to reviewer]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelatedRepositories**
> GetRelatedRepositories(ctx, projectKey)


Retrieve repositories which are related to this one. Related repositories are from the same  {@link Repository#getHierarchyId() hierarchy} as this repository.  <p>  Only repositories to which the authenticated user has <b>REPO_READ</b> permission will be included, even  if more repositories are part of this repository's hierarchy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositories**
> GetRepositories(ctx, projectKey)


Retrieve repositories from the project corresponding to the supplied <strong>projectKey</strong>.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoriesRecentlyAccessed**
> GetRepositoriesRecentlyAccessed(ctx, optional)


Retrieve a page of recently accessed repositories for the currently authenticated user.  <p>  Repositories are ordered from most recently to least recently accessed.  <p>  Only authenticated users may call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| (optional) if specified, it must be a valid repository permission level name and will limit                    the resulting repository list to ones that the requesting user has the specified permission                    level to. If not specified, the default &lt;code&gt;REPO_READ&lt;/code&gt; permission level will be assumed. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositories_0**
> GetRepositories_0(ctx, optional)


Retrieve a page of repositories based on query parameters that control the search. See the documentation of the  parameters for more details.  <p>  This resource is anonymously accessible.  <p>  <b>Note on permissions.</b> In absence of the {@code permission} query parameter the implicit 'read' permission  is assumed. Please note that this permission is lower than the REPO_READ permission rather than being equal to  it. The implicit 'read' permission for a given repository is assigned to any user that has any of the higher  permissions, such as <tt>REPO_READ</tt>, as well as to anonymous users if the repository is marked as public.  The important implication of the above is that an anonymous request to this resource with a permission level  <tt>REPO_READ</tt> is guaranteed to receive an empty list of repositories as a result. For anonymous requests  it is therefore recommended to not specify the <tt>permission</tt> parameter at all.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| (optional) if specified, this will limit the resulting repository list to ones whose name                     matches this parameter&#39;s value. The match will be done case-insensitive and any leading                     and/or trailing whitespace characters on the &lt;code&gt;name&lt;/code&gt; parameter will be stripped. | 
 **projectname** | **string**| (optional) if specified, this will limit the resulting repository list to ones whose project&#39;s                     name matches this parameter&#39;s value. The match will be done case-insensitive and any leading                     and/or trailing whitespace characters on the &lt;code&gt;projectname&lt;/code&gt; parameter will                     be stripped. | 
 **permission** | **string**| (optional) if specified, it must be a valid repository permission level name and will limit                     the resulting repository list to ones that the requesting user has the specified permission                     level to. If not specified, the default implicit &#39;read&#39; permission level will be assumed. The                     currently supported explicit permission values are &lt;tt&gt;REPO_READ&lt;/tt&gt;, &lt;tt&gt;REPO_WRITE&lt;/tt&gt;                     and &lt;tt&gt;REPO_ADMIN&lt;/tt&gt;. | 
 **visibility** | **string**| (optional) if specified, this will limit the resulting repository list based on the                     repositories visibility. Valid values are &lt;em&gt;public&lt;/em&gt; or &lt;em&gt;private&lt;/em&gt;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository**
> GetRepository(ctx, projectKey, projectKey2, repositorySlug)


Retrieve the repository matching the supplied <strong>projectKey</strong> and <strong>repositorySlug</strong>.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 
  **projectKey2** | **string**| the parent project key | 
  **repositorySlug** | **string**| the repository slug | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryHook**
> GetRepositoryHook(ctx, hookKey)


Retrieve a repository hook for this repository.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryHook_0**
> GetRepositoryHook_0(ctx, hookKey)


Retrieve a repository hook for this project.  <p>  The authenticated user must have <strong>PROJECT_READ</strong> permission for the specified project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryHooks**
> GetRepositoryHooks(ctx, optional)


Retrieve a page of repository hooks for this repository.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| the optional type to filter by. Valid values are &lt;code&gt;PRE_RECEIVE&lt;/code&gt; or &lt;code&gt;POST_RECEIVE&lt;/code&gt; | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryHooks_0**
> GetRepositoryHooks_0(ctx, optional)


Retrieve a page of repository hooks for this project.  <p>  The authenticated user must have <strong>PROJECT_READ</strong> permission for the specified project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| the optional type to filter by. Valid values are &lt;code&gt;PRE_RECEIVE&lt;/code&gt; or &lt;code&gt;POST_RECEIVE&lt;/code&gt; | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRootLevel**
> GetRootLevel(ctx, )


Retrieve the current log level for the root logger.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderAddress**
> GetSenderAddress(ctx, )


Retrieves the server email address

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettings**
> GetSettings(ctx, hookKey)


Retrieve the settings for a repository hook for this repository.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettings_0**
> GetSettings_0(ctx, hookKey)


Retrieve the settings for a repository hook for this project.  <p>  The authenticated user must have <strong>PROJECT_READ</strong> permission for the specified project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatistics**
> GetStatistics(ctx, webhookId, optional)


Get the statistics for a specific webhook.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webhookId** | **int32**| id of the webhook | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **int32**| id of the webhook | 
 **event** | **string**| the string id of a specific event to retrieve the last invocation for. May be empty, in which                   case all events are considered | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatisticsSummary**
> GetStatisticsSummary(ctx, webhookId)


Get the statistics summary for a specific webhook.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webhookId** | **int32**| id of the webhook | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTag**
> GetTag(ctx, name)


Retrieve a tag in the specified repository.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the context repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| the name of the tag to be retrieved | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTags**
> GetTags(ctx, optional)


Retrieve the tags matching the supplied <strong>filterText</strong> param.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the context repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterText** | **string**| the text to match on | 
 **orderBy** | **string**| ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTask**
> GetTask(ctx, taskId)


Retrieve a existing task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **taskId** | **int64**| the id identifying the task to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> GetUser(ctx, )


Retrieve the user matching the supplied <strong>userSlug</strong>.  <p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSettings**
> GetUserSettings(ctx, )


Retrieve a map of user setting key values for a specific user identified by the user slug.  <p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> GetUsers(ctx, optional)


Retrieve a page of users.  <p>  The authenticated user must have the <strong>LICENSED_USER</strong> permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only users with usernames, display name or email addresses containing the supplied                string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWithAnyPermission**
> GetUsersWithAnyPermission(ctx, optional)


Retrieve a page of users that have been granted at least one global permission.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only user names containing the supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWithAnyPermission_0**
> GetUsersWithAnyPermission_0(ctx, optional)


Retrieve a page of users that have been granted at least one permission for the specified project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWithAnyPermission_1**
> GetUsersWithAnyPermission_1(ctx, optional)


Retrieve a page of users that have been granted at least one permission for the specified repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWithoutAnyPermission**
> GetUsersWithoutAnyPermission(ctx, optional)


Retrieve a page of users that have no granted global permissions.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only user names containing the supplied string will be returned | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWithoutPermission**
> GetUsersWithoutPermission(ctx, optional)


Retrieve a page of <i>licensed</i> users that have no granted permissions for the specified project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersWithoutPermission_0**
> GetUsersWithoutPermission_0(ctx, optional)


Retrieve a page of <i>licensed</i> users that have no granted permissions for the specified repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| if specified only group names containing the supplied string will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers_0**
> GetUsers_0(ctx, )


Retrieve a page of users, optionally run through provided filters.  <p>  Only authenticated users may call this resource.   <h3>Supported Filters</h3>  <p>  Filters are provided in query parameters in a standard <code>name=value</code> fashion. The following filters are  currently supported:  <ul>      <li>          {@code filter} - return only users, whose username, name or email address <i>contain</i> the          {@code filter} value      </li>      <li>          {@code group} - return only users who are members of the given group      </li>      <li>          {@code permission} - the \"root\" of a permission filter, whose value must be a valid global,          project, or repository permission. Additional filter parameters referring to this filter that specify the          resource (project or repository) to apply the filter to must be prefixed with <code>permission.</code>. See the          section \"Permission Filters\" below for more details.      </li>      <li>          {@code permission.N} - the \"root\" of a single permission filter, similar to the {@code permission}          parameter, where \"N\" is a natural number starting from 1. This allows clients to specify multiple permission          filters, by providing consecutive filters as {@code permission.1}, {@code permission.2} etc. Note that          the filters numbering has to start with 1 and be continuous for all filters to be processed. The total allowed          number of permission filters is 50 and all filters exceeding that limit will be dropped. See the section          \"Permission Filters\" below for more details on how the permission filters are processed.      </li>  </ul>     <h3>Permission Filters</h3>  <p>  The following three sub-sections list parameters supported for permission filters (where <code>[root]</code> is  the root permission filter name, e.g. {@code permission}, {@code permission.1} etc.) depending on the  permission resource. The system determines which filter to apply (Global, Project or Repository permission)  based on the <code>[root]</code> permission value. E.g. {@code ADMIN} is a global permission,  {@code PROJECT_ADMIN} is a project permission and {@code REPO_ADMIN} is a repository permission. Note  that the parameters for a given resource will be looked up in the order as they are listed below, that is e.g.  for a project resource, if both {@code projectId} and {@code projectKey} are provided, the system will  use {@code projectId} for the lookup.  <h4>Global permissions</h4>  <p>  The permission value under <code>[root]</code> is the only required and recognized parameter, as global  permissions do not apply to a specific resource.  <p>  Example valid filter: <code>permission=ADMIN</code>.  <h4>Project permissions</h4>  <ul>      <li><code>[root]</code>- specifies the project permission</li>      <li><code>[root].projectId</code> - specifies the project ID to lookup the project by</li>      <li><code>[root].projectKey</code> - specifies the project key to lookup the project by</li>  </ul>  <p>  Example valid filter: <code>permission.1=PROJECT_ADMIN&permission.1.projectKey=TEST_PROJECT</code>.  <h4>Repository permissions</h4>  <ul>      <li><code>[root]</code>- specifies the repository permission</li>      <li><code>[root].projectId</code> - specifies the repository ID to lookup the repository by</li>      <li><code>[root].projectKey</code> and <code>[root].repositorySlug</code>- specifies the project key and      repository slug to lookup the repository by; both values <i>need to</i> be provided for this look up to be      triggered</li>  </ul>  Example valid filter: <code>permission.2=REPO_ADMIN&permission.2.projectKey=TEST_PROJECT&permission.2.repositorySlug=test_repo</code>.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> GetWebhook(ctx, webhookId, optional)


Get a webhook by id.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webhookId** | **int32**| the existing webhook id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **int32**| the existing webhook id | 
 **statistics** | **bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> Get_0(ctx, pullRequestId)


Retrieve a pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the ID of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HasAllUserPermission**
> HasAllUserPermission(ctx, permission)


Check whether the specified permission is the default permission (granted to all users) for a project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **permission** | **string**| the permission to grant                        Available project permissions are:                        &lt;ul&gt;                            &lt;li&gt;PROJECT_READ&lt;/li&gt;                            &lt;li&gt;PROJECT_WRITE&lt;/li&gt;                            &lt;li&gt;PROJECT_ADMIN&lt;/li&gt;                        &lt;/ul&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListParticipants**
> ListParticipants(ctx, pullRequestId)


Retrieves a page of the participants for a given pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Merge**
> Merge(ctx, pullRequestId, optional)


Merge the specified pull request.  <p>  The authenticated user must have <strong>REPO_WRITE</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the ID of the pull request within the repository | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pullRequestId** | **int64**| the ID of the pull request within the repository | 
 **version** | **int32**| the current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified                 version the operation will fail. To determine the current version of the pull request it should be                 fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned                 JSON structure. | [default to -1]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyAllUserPermission**
> ModifyAllUserPermission(ctx, permission, optional)


Grant or revoke a project permission to all users, i.e. set the default permission.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **permission** | **string**| the permission to grant                        Available project permissions are:                        &lt;ul&gt;                            &lt;li&gt;PROJECT_READ&lt;/li&gt;                            &lt;li&gt;PROJECT_WRITE&lt;/li&gt;                            &lt;li&gt;PROJECT_ADMIN&lt;/li&gt;                        &lt;/ul&gt; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| the permission to grant                        Available project permissions are:                        &lt;ul&gt;                            &lt;li&gt;PROJECT_READ&lt;/li&gt;                            &lt;li&gt;PROJECT_WRITE&lt;/li&gt;                            &lt;li&gt;PROJECT_ADMIN&lt;/li&gt;                        &lt;/ul&gt; | 
 **allow** | **bool**| &lt;em&gt;true&lt;/em&gt; to grant the specified permission to all users, or &lt;em&gt;false&lt;/em&gt; to revoke it | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Preview**
> Preview(ctx, optional)


Preview the generated html for given markdown contents.  <p>  Only authenticated users may call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **urlMode** | **string**| (Optional) The UrlMode used when building the url. One of: ABSOLUTE, RELATIVE and CONFIGURED                       By default this is RELATIVE. | 
 **hardwrap** | **bool**| (Optional) Whether the markup implementation should convert newlines to breaks.                       By default this is false which reflects the standard markdown specification. | 
 **htmlEscape** | **bool**| (Optional) true if HTML should be escaped in the input markup, false otherwise. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGroupFromUser**
> RemoveGroupFromUser(ctx, )


Remove a user from a group. This is very similar to <code>groups/remove-user</code>, but with the <em>context</em>  and <em>itemName</em> attributes of the supplied request entity reversed. On the face of it this may appear  redundant, but it facilitates a specific UI component in Stash.  <p>  In the request entity, the <em>context</em> attribute is the user and the <em>itemName</em> is the group.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserFromGroup**
> RemoveUserFromGroup(ctx, )


<strong>Deprecated since 2.10</strong>. Use /rest/users/remove-groups instead.  <p>  Remove a user from a group.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.  <p>  In the request entity, the <em>context</em> attribute is the group and the <em>itemName</em> is the user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameUser**
> RenameUser(ctx, )


Rename a user.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Reopen**
> Reopen(ctx, pullRequestId, optional)


Re-open a declined pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pullRequestId** | **int64**| the id of the pull request within the repository | 
 **version** | **int32**| the current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified                 version the operation will fail. To determine the current version of the pull request it should be                 fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned                 JSON structure. | [default to -1]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryCreateRepository**
> RetryCreateRepository(ctx, projectKey)


If a create or fork operation fails, calling this method will clean up the broken repository and try again. The  repository must be in an INITIALISATION_FAILED state.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokePermissionsForGroup**
> RevokePermissionsForGroup(ctx, optional)


Revoke all global permissions for a group.   <p>  The authenticated user must have:  <ul>      <li><strong>ADMIN</strong> permission or higher; and</li>      <li>greater or equal permissions than the current permission level of the group (a user may not demote the      permission level of a group with higher permissions than them)</li>  </ul>  to call this resource. In addition, a user may not revoke a group's permissions if their own permission level  would be reduced as a result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the name of the group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokePermissionsForGroup_0**
> RevokePermissionsForGroup_0(ctx, optional)


Revoke all permissions for the specified project for a group.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.  <p>  In addition, a user may not revoke a group's permissions if it will reduce their own permission level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the name of the group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokePermissionsForGroup_1**
> RevokePermissionsForGroup_1(ctx, optional)


Revoke all permissions for the specified repository for a group.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource.  <p>  In addition, a user may not revoke a group's permissions if it will reduce their own permission level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the name of the group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokePermissionsForUser**
> RevokePermissionsForUser(ctx, optional)


Revoke all global permissions for a user.  <p>  The authenticated user must have:  <ul>      <li><strong>ADMIN</strong> permission or higher; and</li>      <li>greater or equal permissions than the current permission level of the user (a user may not demote the      permission level of a user with higher permissions than them)</li>  </ul>  to call this resource. In addition, a user may not demote their own permission level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the name of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokePermissionsForUser_0**
> RevokePermissionsForUser_0(ctx, optional)


Revoke all permissions for the specified project for a user.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource.  <p>  In addition, a user may not revoke their own project permissions if they do not have a higher global permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the name of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokePermissionsForUser_1**
> RevokePermissionsForUser_1(ctx, optional)


Revoke all permissions for the specified repository for a user.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource.  <p>  In addition, a user may not revoke their own repository permissions if they do not have a higher  project or global permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the name of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search**
> Search(ctx, optional)


Retrieve a page of participant users for all the pull requests to or from the specified repository.  <p>  <p>  Optionally clients can specify following filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **direction** | **string**| (optional, defaults to &lt;strong&gt;INCOMING&lt;/strong&gt;) the direction relative to the specified                   repository. Either &lt;strong&gt;INCOMING&lt;/strong&gt; or &lt;strong&gt;OUTGOING&lt;/strong&gt;. | [default to incoming]
 **filter** | **string**| (optional) return only users, whose username, name or email address &lt;i&gt;contain&lt;/i&gt;                   the {@code filter} value | 
 **role** | **string**| (optional) The role associated with the pull request participant.                   This must be one of {@code AUTHOR}, {@code REVIEWER}, or{@code PARTICIPANT} | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDefaultBranch**
> SetDefaultBranch(ctx, )


Update the default branch of a repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetLevel**
> SetLevel(ctx, levelName, loggerName)


Set the current log level for a given logger.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **levelName** | **string**| the level to set the logger to. Either TRACE, DEBUG, INFO, WARN or ERROR | 
  **loggerName** | **string**| the name of the logger. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMailConfig**
> SetMailConfig(ctx, )


Updates the mail configuration   The authenticated user must have the <strong>SYS_ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMergeConfig**
> SetMergeConfig(ctx, scmId)


Update the pull request merge strategies for the context repository.  <p>  The authenticated user must have <strong>ADMIN</strong> permission for the context repository to call this  resource.  <p>  Only the strategies provided will be enabled, only one may be set to default  <p>  An explicitly set pull request merge strategy configuration can be deleted by POSTing a document with  an empty \"mergeConfig\" attribute. i.e:  <pre>  {      \"mergeConfig\": {      }  }  </pre>  Upon completion of this request, the effective configuration will be the default configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scmId** | **string**| the id of the scm to get strategies for | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPermissionForGroup**
> SetPermissionForGroup(ctx, optional)


Promote or demote a group's permission level for the specified repository. Available repository permissions are:  <ul>      <li>REPO_READ</li>      <li>REPO_WRITE</li>      <li>REPO_ADMIN</li>  </ul>  See the <a href=\"https://confluence.atlassian.com/display/BitbucketServer/Using+repository+permissions\">Bitbucket  Server documentation</a> for a detailed explanation of what each permission entails.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource. In addition, a user may not demote a group's permission level if their  own permission level would be reduced as a result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| the permission to grant | 
 **name** | **string**| the names of the groups | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPermissionForGroups**
> SetPermissionForGroups(ctx, optional)


Promote or demote a user's global permission level. Available global permissions are:  <ul>      <li>LICENSED_USER</li>      <li>PROJECT_CREATE</li>      <li>ADMIN</li>      <li>SYS_ADMIN</li>  </ul>  See the <a href=\"https://confluence.atlassian.com/display/BitbucketServer/Global+permissions\">Bitbucket Server  documentation</a> for a detailed explanation of what each permission entails.  <p>  The authenticated user must have:  <ul>      <li><strong>ADMIN</strong> permission or higher; and</li>      <li>the permission they are attempting to grant or higher; and</li>      <li>greater or equal permissions than the current permission level of the group (a user may not demote the      permission level of a group with higher permissions than them)</li>  </ul>  to call this resource. In addition, a user may not demote a group's permission level if their own permission  level would be reduced as a result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| the permission to grant | 
 **name** | **string**| the names of the groups | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPermissionForGroups_0**
> SetPermissionForGroups_0(ctx, optional)


Promote or demote a group's permission level for the specified project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource. In addition, a user may not demote a group's permission level if their  own permission level would be reduced as a result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string**| The permission to grant.                        See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)                        for a detailed explanation of what each permission entails.                        Available project permissions are:                        &lt;ul&gt;                            &lt;li&gt;PROJECT_READ&lt;/li&gt;                            &lt;li&gt;PROJECT_WRITE&lt;/li&gt;                            &lt;li&gt;PROJECT_ADMIN&lt;/li&gt;                        &lt;/ul&gt; | 
 **name** | **string**| the names of the groups | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPermissionForUser**
> SetPermissionForUser(ctx, optional)


Promote or demote a user's permission level for the specified repository. Available repository permissions are:  <ul>      <li>REPO_READ</li>      <li>REPO_WRITE</li>      <li>REPO_ADMIN</li>  </ul>  See the <a href=\"https://confluence.atlassian.com/display/BitbucketServer/Using+repository+permissions\">Bitbucket  Server documentation</a> for a detailed explanation of what each permission entails.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository or a higher  project or global permission to call this resource. In addition, a user may not reduce their own permission level unless  they have a project or global permission that already implies that permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the names of the users | 
 **permission** | **string**| the permission to grant | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPermissionForUsers**
> SetPermissionForUsers(ctx, optional)


Promote or demote the global permission level of a user. Available global permissions are:  <ul>      <li>LICENSED_USER</li>      <li>PROJECT_CREATE</li>      <li>ADMIN</li>      <li>SYS_ADMIN</li>  </ul>  See the <a href=\"https://confluence.atlassian.com/display/BitbucketServer/Global+permissions\">Bitbucket Server  documentation</a> for a detailed explanation of what each permission entails.  <p>  The authenticated user must have:  <ul>      <li><strong>ADMIN</strong> permission or higher; and</li>      <li>the permission they are attempting to grant; and</li>      <li>greater or equal permissions than the current permission level of the user (a user may not demote the      permission level of a user with higher permissions than them)</li>  </ul>  to call this resource. In addition, a user may not demote their own permission level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the names of the users | 
 **permission** | **string**| the permission to grant | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPermissionForUsers_0**
> SetPermissionForUsers_0(ctx, optional)


Promote or demote a user's permission level for the specified project.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project or a higher  global permission to call this resource. In addition, a user may not reduce their own permission level unless  they have a global permission that already implies that permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| the names of the users | 
 **permission** | **string**| the permission to grant. See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)                        for a detailed explanation of what each permission entails.                        Available project permissions are:                        &lt;ul&gt;                            &lt;li&gt;PROJECT_READ&lt;/li&gt;                            &lt;li&gt;PROJECT_WRITE&lt;/li&gt;                            &lt;li&gt;PROJECT_ADMIN&lt;/li&gt;                        &lt;/ul&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRootLevel**
> SetRootLevel(ctx, levelName)


Set the current log level for the root logger.  <p>  The authenticated user must have <strong>ADMIN</strong> permission or higher to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **levelName** | **string**| the level to set the logger to. Either TRACE, DEBUG, INFO, WARN or ERROR | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSenderAddress**
> SetSenderAddress(ctx, )


Updates the server email address   The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSettings**
> SetSettings(ctx, hookKey)


Modify the settings for a repository hook for this repository.  <p>  The service will reject any settings which are too large, the current limit is 32KB once serialized.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.  <p>  A JSON document can be provided to use as the settings for the hook. These structure and validity of the document  is decided by the plugin providing the hook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSettings_0**
> SetSettings_0(ctx, hookKey)


Modify the settings for a repository hook for this project.  <p>  The service will reject any settings which are too large, the current limit is 32KB once serialized.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project to call this  resource.  <p>  A JSON document can be provided to use as the settings for the hook. These structure and validity of the document  is decided by the plugin providing the hook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Stream**
> Stream(ctx, optional)


Streams files in the requested <code>path</code> with the last commit to modify each file. Commit modifications  are traversed starting from the <code>at</code> commit or, if not specified, from the tip of the default branch.  <p>  Unless the repository is public, the authenticated user must have <b>REPO_READ</b> access to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **string**| the commit to use as the starting point when listing files and calculating modifications | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamChanges**
> StreamChanges(ctx, optional)


Gets the file changes available in the {@code from} commit but not in the {@code to} commit.  <p>  If either the {@code from} or {@code to} commit are not specified, they will be replaced by the  default branch of their containing repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string**| the source commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **to** | **string**| the target commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **fromRepo** | **string**| an optional parameter specifying the source repository containing the source commit                  if that commit is not present in the current repository; the repository can be specified                  by either its ID &lt;em&gt;fromRepo&#x3D;42&lt;/em&gt; or by its project key plus its repo slug separated by                  a slash: &lt;em&gt;fromRepo&#x3D;projectKey/repoSlug&lt;/em&gt; | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamChanges_0**
> StreamChanges_0(ctx, optional)


Gets changes for the specified PullRequest.  <p>  If the {@code changeScope} query parameter is set to {@code unreviewed}, the application will attempt to stream  unreviewed changes based on the {@code lastReviewedCommit} of the current user, which are the changes between the  {@code lastReviewedCommit} and the latest commit of the source branch. The current user is considered to  <i>not</i> have any unreviewed changes for the pull request when the {@code lastReviewedCommit} is either  {@code null} (everything is unreviewed, so all changes are streamed), equal to the latest commit of the source  branch (everything is reviewed), or no longer on the source branch (the source branch has been rebased). In these  cases, the application will fall back to streaming all changes (the default), which is the effective diff for the  pull request. The type of changes streamed can be determined by the {@code changeScope} parameter included in the  properties map of the response.  <p>  Note: This resource is currently <i>not paged</i>. The server will return at most one page. The server will  truncate the number of changes to either the request's page limit or an internal maximum, whichever is smaller.  The start parameter of the page request is also ignored.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeScope** | **string**| {@code UNREVIEWED} to stream the unreviewed changes for the current user (if they exist);                     {@code RANGE} to stream changes between two arbitrary commits (requires {@code sinceId} and                     {@code untilId}); otherwise {@code ALL} to stream all changes (the default) | [default to ALL]
 **sinceId** | **string**| the since commit hash to stream changes for a {@code RANGE} arbitrary change scope | 
 **untilId** | **string**| the until commit hash to stream changes for a {@code RANGE} arbitrary change scope | 
 **withComments** | **bool**| {@code true} to apply comment counts in the changes (the default); otherwise, {@code false}                      to stream changes without comment counts | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamCommits**
> StreamCommits(ctx, optional)


Gets the commits accessible from the {@code from} commit but not in the {@code to} commit.  <p>  If either the {@code from} or {@code to} commit are not specified, they will be replaced by the  default branch of their containing repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string**| the source commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **to** | **string**| the target commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **fromRepo** | **string**| an optional parameter specifying the source repository containing the source commit                  if that commit is not present in the current repository; the repository can be specified                  by either its ID &lt;em&gt;fromRepo&#x3D;42&lt;/em&gt; or by its project key plus its repo slug separated by                  a slash: &lt;em&gt;fromRepo&#x3D;projectKey/repoSlug&lt;/em&gt; | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamDiff**
> StreamDiff(ctx, commitId, optional)


Retrieve the diff between two provided revisions.  <p>  <strong>Note:</strong> This resource is currently <i>not paged</i>. The server will internally apply a hard cap  to the streamed lines, and it is not possible to request subsequent pages if that cap is exceeded. In the event  that the cap is reached, the diff will be cut short and one or more {@code truncated} flags will be set to  {@code true} on the {@code \"segments\"}, {@code \"hunks\"} and {@code \"diffs\"} properties, as well as the top-level  object, in the returned JSON response.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitId** | **string**|  | 
 **autoSrcPath** | **bool**| {@code true} to automatically try to find the source path when it&#39;s not provided,                       {@code false} otherwise. Requires the {@code path} to be provided. | [default to false]
 **contextLines** | **int32**| the number of context lines to include around added/removed lines in the diff | [default to -1]
 **since** | **string**| the base revision to diff from. If omitted the parent revision of the until revision is used | 
 **srcPath** | **string**| the source path for the file, if it was copied, moved or renamed | 
 **whitespace** | **string**| optional whitespace flag which can be set to {@code ignore-all} | 
 **withComments** | **bool**| {@code true} to embed comments in the diff (the default); otherwise {@code false}                       to stream the diff without comments | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamDiff_0**
> StreamDiff_0(ctx, commitId, path, commitId2, optional)


Retrieve the diff between two provided revisions.  <p>  <strong>Note:</strong> This resource is currently <i>not paged</i>. The server will internally apply a hard cap  to the streamed lines, and it is not possible to request subsequent pages if that cap is exceeded. In the event  that the cap is reached, the diff will be cut short and one or more {@code truncated} flags will be set to  {@code true} on the {@code \"segments\"}, {@code \"hunks\"} and {@code \"diffs\"} properties, as well as the top-level  object, in the returned JSON response.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**|  | 
  **path** | **string**| the path to the file which should be diffed (optional) | 
  **commitId2** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitId** | **string**|  | 
 **path** | **string**| the path to the file which should be diffed (optional) | 
 **commitId2** | **string**|  | 
 **autoSrcPath** | **bool**| {@code true} to automatically try to find the source path when it&#39;s not provided,                       {@code false} otherwise. Requires the {@code path} to be provided. | [default to false]
 **contextLines** | **int32**| the number of context lines to include around added/removed lines in the diff | [default to -1]
 **since** | **string**| the base revision to diff from. If omitted the parent revision of the until revision is used | 
 **srcPath** | **string**| the source path for the file, if it was copied, moved or renamed | 
 **whitespace** | **string**| optional whitespace flag which can be set to {@code ignore-all} | 
 **withComments** | **bool**| {@code true} to embed comments in the diff (the default); otherwise {@code false}                       to stream the diff without comments | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamDiff_1**
> StreamDiff_1(ctx, path, optional)


Gets a diff of the changes available in the {@code from} commit but not in the {@code to} commit.  <p>  If either the {@code from} or {@code to} commit are not specified, they will be replaced by the  default branch of their containing repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the path to the file to diff (optional) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| the path to the file to diff (optional) | 
 **from** | **string**| the source commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **to** | **string**| the target commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **fromRepo** | **string**| an optional parameter specifying the source repository containing the source commit                  if that commit is not present in the current repository; the repository can be specified                  by either its ID &lt;em&gt;fromRepo&#x3D;42&lt;/em&gt; or by its project key plus its repo slug separated by                  a slash: &lt;em&gt;fromRepo&#x3D;projectKey/repoSlug&lt;/em&gt; | [default to ]
 **srcPath** | **string**|  | 
 **contextLines** | **int32**| an optional number of context lines to include around each added or removed lines in the diff | [default to -1]
 **whitespace** | **string**| an optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamDiff_2**
> StreamDiff_2(ctx, optional)


Retrieve the diff for a specified file path between two provided revisions.  <p>  <strong>Note:</strong> This resource is currently <i>not paged</i>. The server will internally apply a hard cap  to the streamed lines, and it is not possible to request subsequent pages if that cap is exceeded. In the event  that the cap is reached, the diff will be cut short and one or more <code>truncated</code> flags will be set to  <code>true</code> on the segments, hunks and diffs substructures in the returned JSON response.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextLines** | **int32**| the number of context lines to include around added/removed lines in the diff | [default to -1]
 **since** | **string**| the base revision to diff from. If omitted the parent revision of the until revision is used | 
 **srcPath** | **string**| the source path for the file, if it was copied, moved or renamed | 
 **until** | **string**| the target revision to diff to (required) | 
 **whitespace** | **string**| optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamDiff_3**
> StreamDiff_3(ctx, path, optional)


Retrieve the diff for a specified file path between two provided revisions.  <p>  <strong>Note:</strong> This resource is currently <i>not paged</i>. The server will internally apply a hard cap  to the streamed lines, and it is not possible to request subsequent pages if that cap is exceeded. In the event  that the cap is reached, the diff will be cut short and one or more <code>truncated</code> flags will be set to  <code>true</code> on the segments, hunks and diffs substructures in the returned JSON response.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the path to the file which should be diffed (required) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| the path to the file which should be diffed (required) | 
 **contextLines** | **int32**| the number of context lines to include around added/removed lines in the diff | [default to -1]
 **since** | **string**| the base revision to diff from. If omitted the parent revision of the until revision is used | 
 **srcPath** | **string**| the source path for the file, if it was copied, moved or renamed | 
 **until** | **string**| the target revision to diff to (required) | 
 **whitespace** | **string**| optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamDiff_4**
> StreamDiff_4(ctx, optional)


Streams a diff within a pull request.  <p>  If the specified file has been copied, moved or renamed, the <code>srcPath</code> must also be specified to  produce the correct diff.  <p>  Note: This RESTful endpoint is currently <i>not paged</i>. The server will internally apply a hard cap to the  streamed lines, and it is not possible to request subsequent pages if that cap is exceeded.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextLines** | **int32**| the number of context lines to include around added/removed lines in the diff | [default to -1]
 **diffType** | **string**| the type of diff being requested. When {@code withComments} is {@code true}                      this works as a hint to the system to attach the correct set of comments to the diff | [default to EFFECTIVE]
 **sinceId** | **string**| the since commit hash to stream a diff between two arbitrary hashes | 
 **srcPath** | **string**| the previous path to the file, if the file has been copied, moved or renamed | 
 **untilId** | **string**| the until commit hash to stream a diff between two arbitrary hashes | 
 **whitespace** | **string**| optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 
 **withComments** | **bool**| &lt;code&gt;true&lt;/code&gt; to embed comments in the diff (the default); otherwise, &lt;code&gt;false&lt;/code&gt;                      to stream the diff without comments | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamDiff_5**
> StreamDiff_5(ctx, path, optional)


Streams a diff within a pull request.  <p>  If the specified file has been copied, moved or renamed, the <code>srcPath</code> must also be specified to  produce the correct diff.  <p>  Note: This RESTful endpoint is currently <i>not paged</i>. The server will internally apply a hard cap to the  streamed lines, and it is not possible to request subsequent pages if that cap is exceeded.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the path to the file which should be diffed (optional) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| the path to the file which should be diffed (optional) | 
 **contextLines** | **int32**| the number of context lines to include around added/removed lines in the diff | [default to -1]
 **diffType** | **string**| the type of diff being requested. When {@code withComments} is {@code true}                      this works as a hint to the system to attach the correct set of comments to the diff | [default to EFFECTIVE]
 **sinceId** | **string**| the since commit hash to stream a diff between two arbitrary hashes | 
 **srcPath** | **string**| the previous path to the file, if the file has been copied, moved or renamed | 
 **untilId** | **string**| the until commit hash to stream a diff between two arbitrary hashes | 
 **whitespace** | **string**| optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 
 **withComments** | **bool**| &lt;code&gt;true&lt;/code&gt; to embed comments in the diff (the default); otherwise, &lt;code&gt;false&lt;/code&gt;                      to stream the diff without comments | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamFiles**
> StreamFiles(ctx, optional)


Retrieve a page of files from particular directory of a repository. The search is done recursively, so all files  from any sub-directory of the specified directory will be returned.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **string**| the commit ID or ref (e.g. a branch or tag) to list the files at.              If not specified the default branch will be used instead. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamFiles_0**
> StreamFiles_0(ctx, path, optional)


Retrieve a page of files from particular directory of a repository. The search is done recursively, so all files  from any sub-directory of the specified directory will be returned.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the directory to list files for. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| the directory to list files for. | 
 **at** | **string**| the commit ID or ref (e.g. a branch or tag) to list the files at.              If not specified the default branch will be used instead. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Stream_0**
> Stream_0(ctx, path, optional)


Streams files in the requested <code>path</code> with the last commit to modify each file. Commit modifications  are traversed starting from the <code>at</code> commit or, if not specified, from the tip of the default branch.  <p>  Unless the repository is public, the authenticated user must have <b>REPO_READ</b> access to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| the path within the repository whose files should be streamed | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| the path within the repository whose files should be streamed | 
 **at** | **string**| the commit to use as the starting point when listing files and calculating modifications | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestWebhook**
> TestWebhook(ctx, optional)


Test connectivity to a specific endpoint.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string**| the url in which to connect to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignParticipantRole**
> UnassignParticipantRole(ctx, pullRequestId, optional)


Unassigns a participant from the REVIEWER role they may have been given in a pull request.  <p>  If the participant has no explicit role this method has no effect.  <p>  Afterwards, the user will still remain a participant in the pull request but their role will be reduced to  PARTICIPANT. This is because once made a participant of a pull request,  a user will forever remain a participant. Only their role may be altered.  <p>  The authenticated user must have <strong>REPO_WRITE</strong> permission for the repository that this pull request  targets to call this resource.  <p>  <strong>Deprecated since 4.2</strong>. Use  /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug}  instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pullRequestId** | **int64**| the id of the pull request within the repository | 
 **username** | **string**| the participant&#39;s user name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignParticipantRole_0**
> UnassignParticipantRole_0(ctx, pullRequestId, userSlug, pullRequestId2)


Unassigns a participant from the REVIEWER role they may have been given in a pull request.  <p>  If the participant has no explicit role this method has no effect.  <p>  Afterwards, the user will still remain a participant in the pull request but their role will be reduced to  PARTICIPANT. This is because once made a participant of a pull request,  a user will forever remain a participant. Only their role may be altered.  <p>  The authenticated user must have <strong>REPO_WRITE</strong> permission for the repository that this pull request  targets to call this resource.  <p>  <strong>Deprecated since 4.2</strong>. Use  /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug}  instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 
  **userSlug** | **string**| the slug for the user changing their status | 
  **pullRequestId2** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unwatch**
> Unwatch(ctx, commitId)


Removes the authenticated user as a watcher for the specified commit.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository containing the commit  to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the &lt;i&gt;full {@link Commit#getId() ID}&lt;/i&gt; of the commit within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unwatch_0**
> Unwatch_0(ctx, pullRequestId)


Make the authenticated user stop watching the specified pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Update(ctx, )


Decodes the provided encoded license and sets it as the active license. If no license was provided, a 400 is  returned. If the license cannot be decoded, or cannot be applied, a 409 is returned. Some possible reasons a  license may not be applied include:  <ul>      <li>It is for a different product</li>      <li>It is already expired</li>  </ul>  Otherwise, if the license is updated successfully, details for the new license are returned with a 200 response.  <p>  <b>Warning</b>: It is possible to downgrade the license during update, applying a license with a lower number  of permitted users. If the number of currently-licensed users exceeds the limits of the new license, pushing  will be disabled until the licensed user count is brought into compliance with the new license.  <p>  The authenticated user must have <b>SYS_ADMIN</b> permission. <b>ADMIN</b> users may <i>view</i> the current  license details, but they may not <i>update</i> the license.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateComment**
> UpdateComment(ctx, commitId, commentId, commitId2)


Update the text of a comment. Only the user who created a comment may update it.  <p>  <strong>Note:</strong> the supplied supplied JSON object must contain a <code>version</code> that must match  the server's version of the comment or the update will fail. To determine the current version of the comment,  the comment should be fetched from the server prior to the update. Look for the 'version' attribute in the  returned JSON structure.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that the commit  is in to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the commit to which the comments must be anchored | 
  **commentId** | **int64**| the ID of the comment to retrieve | 
  **commitId2** | **string**| the &lt;i&gt;full {@link Commit#getId() ID}&lt;/i&gt; of the commit within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateComment_0**
> UpdateComment_0(ctx, commentId)


Update the text of a comment. Only the user who created a comment may update it.  <p>  <strong>Note:</strong> the supplied supplied JSON object must contain a <code>version</code> that must match the  server's version of the comment or the update will fail. To determine the current version of  the comment, the comment should be fetched from the server prior to the update. Look for the  'version' attribute in the returned JSON structure.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commentId** | **int64**| the id of the comment to retrieve | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> UpdateProject(ctx, )


Update the project matching the <strong>projectKey</strong> supplied in the resource path.  <p>  To include a custom avatar for the updated project, the project definition should contain an additional attribute  with the key <code>avatar</code> and the value a data URI containing Base64-encoded image data. The URI should be  in the following format:  <code>      data:(content type, e.g. image/png);base64,(data)  </code>  If the data is not Base64-encoded, or if a character set is defined in the URI, or the URI is otherwise invalid,  <em>project creation will fail</em>.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePullRequestSettings**
> UpdatePullRequestSettings(ctx, )


Update the pull request settings for the context repository.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the context repository to call this  resource.  <p>  This resource will call all RestFragments that are registered with the key  <strong>bitbucket.repository.settings.pullRequests</strong>. If any fragment fails validations by returning a  non-empty Map of errors, then no fragments will execute.  <p>  Only the settings that should be updated need to be included in the request.  <p>  The property keys for the settings that are bundled with the application are  <ul>      <li>mergeConfig - the merge strategy configuration for pull requests</li>      <li>requiredApprovers - (Deprecated, please use com.atlassian.bitbucket.server.bundled-hooks.requiredApproversMergeHook instead) the number of approvals required on a pull request for it to be mergeable, or 0 to disable the merge check</li>      <li>com.atlassian.bitbucket.server.bundled-hooks.requiredApproversMergeHook - a json map containing the keys 'enabled' (a boolean to enable or disable this merge check) and 'count' (an integer to set the number of required approvals)</li>      <li>requiredAllApprovers - whether or not all approvers must approve a pull request for it to be mergeable</li>      <li>requiredAllTasksComplete - whether or not all tasks on a pull request need to be completed for it to be mergeable</li>      <li>requiredSuccessfulBuilds - (Deprecated, please use com.atlassian.bitbucket.server.bitbucket-build.requiredBuildsMergeCheck instead) the number of successful builds on a pull request for it to be mergeable, or 0 to disable the merge check</li>      <li>com.atlassian.bitbucket.server.bitbucket-build.requiredBuildsMergeCheck - a json map containing the keys 'enabled' (a boolean to enable or disable this merge check) and 'count' (an integer to set the number of required builds)</li>  </ul>  <strong>Merge strategy configuration deletion:</strong>  <p>  An explicitly set pull request merge strategy configuration can be deleted by POSTing a document with an empty  \"mergeConfig\" attribute. i.e:  <pre>  {      \"mergeConfig\": {      }  }  </pre>  Upon completion of this request, the effective configuration will be:  <ul>      <li>The configuration set for this repository's SCM type as set at the project level, if present, otherwise</li>      <li>the configuration set for this repository's SCM type as set at the instance level, if present, otherwise</li>      <li>the default configuration for this repository's SCM type</li>  <ul>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePullRequestSettings_0**
> UpdatePullRequestSettings_0(ctx, scmId)


Update the pull request merge strategy configuration for this project and SCM.  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the context repository to call this  resource.  <p>  Only the strategies provided will be enabled, the default must be set and included in the set of strategies.  <p>  An explicitly set pull request merge strategy configuration can be deleted by POSTing a document with  an empty \"mergeConfig\" attribute. i.e:  <pre>  {      \"mergeConfig\": {      }  }  </pre>  Upon completion of this request, the effective configuration will be the configuration explicitly set for  the SCM, or if no such explicit configuration is set then the default configuration will be used.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scmId** | **string**| the SCM to get strategies for | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository**
> UpdateRepository(ctx, projectKey, projectKey2, repositorySlug)


Update the repository matching the <strong>repositorySlug</strong> supplied in the resource path.  <p>  The repository's slug is derived from its name. If the name changes the slug may also change.  <p>  This API can be used to move the repository to a different project by setting the new project in the request,  example: {@code {\"project\":{\"key\":\"NEW_KEY\"}}} .  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectKey** | **string**| the parent project key | 
  **projectKey2** | **string**| the parent project key | 
  **repositorySlug** | **string**| the repository slug | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettings**
> UpdateSettings(ctx, )


Update the entries of a map of user setting key/values for a specific user identified by the user slug.  <p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStatus**
> UpdateStatus(ctx, pullRequestId, userSlug, pullRequestId2)


Change the current user's status for a pull request. Implicitly adds the user as a participant if they are not  already. If the current user is the author, this method will fail.  <p>  The possible values for {@code status} are <strong>UNAPPROVED</strong>, <strong>NEEDS_WORK</strong>, or  <strong>APPROVED</strong>.  <p>  If the new {@code status} is <strong>NEEDS_WORK</strong> or <strong>APPROVED</strong> then the  {@code lastReviewedCommit} for the participant will be updated to the latest commit of the source branch of the  pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 
  **userSlug** | **string**| the slug for the user changing their status | 
  **pullRequestId2** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTask**
> UpdateTask(ctx, taskId)


Update a existing task.  <p>  As of Stash 3.3, only the state and text of a task can be updated.  <p>  Updating the state of a task is allowed for any user having <em>READ</em> access to the repository.  However only the task's creator, the context's author or an admin of the context's repository can update the  task's text. (For a pull request task, those are the task's creator, the pull request's author or an admin on the  repository containing the pull request). Additionally the task's text cannot be updated if it has been resolved.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **taskId** | **int64**| the id identifying the task to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserDetails**
> UpdateUserDetails(ctx, )


Update a user's details.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserDetails_0**
> UpdateUserDetails_0(ctx, )


Update the currently authenticated user's details. The update will always be applied to the currently  authenticated user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserPassword**
> UpdateUserPassword(ctx, )


Update a user's password.  <p>  The authenticated user must have the <strong>ADMIN</strong> permission to call this resource, and may not update  the password of a user with greater permissions than themselves.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserPassword_0**
> UpdateUserPassword_0(ctx, )


Update the currently authenticated user's password.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> UpdateWebhook(ctx, webhookId)


Update an existing webhook.  <p>  The authenticated user must have <strong>REPO_ADMIN</strong> permission for the specified repository to call this  resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webhookId** | **int32**| the existing webhook id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update_0**
> Update_0(ctx, pullRequestId)


Update the title, description, reviewers or destination branch of an existing pull request.  <p>  <strong>Note:</strong> the <em>reviewers</em> list may be updated using this resource. However the  <em>author</em> and <em>participants</em> list may not.  <p>  The authenticated user must either:  <ul>      <li>be the author of the pull request and have the <strong>REPO_READ</strong> permission for the repository      that this pull request targets; or</li>      <li>have the <strong>REPO_WRITE</strong> permission for the repository that this pull request targets</li>  </ul>  to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the ID of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadAvatar**
> UploadAvatar(ctx, )


Update the avatar for the project matching the supplied <strong>projectKey</strong>.  <p>  This resource accepts POST multipart form data, containing a single image in a form-field named 'avatar'.  <p>  There are configurable server limits on both the dimensions (1024x1024 pixels by default) and uploaded file size  (1MB by default). Several different image formats are supported, but <strong>PNG</strong> and  <strong>JPEG</strong> are preferred due to the file size limit.  <p>  This resource has Cross-Site Request Forgery (XSRF) protection. To allow the request to  pass the XSRF check the caller needs to send an <code>X-Atlassian-Token</code> HTTP header with the  value <code>no-check</code>.  <p>  An example <a href=\"http://curl.haxx.se/\">curl</a> request to upload an image name 'avatar.png' would be:  <pre>  curl -X POST -u username:password -H \"X-Atlassian-Token: no-check\" http://example.com/rest/api/1.0/projects/STASH/avatar.png -F avatar=@avatar.png  </pre>  <p>  The authenticated user must have <strong>PROJECT_ADMIN</strong> permission for the specified project to call this  resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadAvatar_0**
> UploadAvatar_0(ctx, )


Update the avatar for the user with the supplied <strong>slug</strong>.  <p>  This resource accepts POST multipart form data, containing a single image in a form-field named 'avatar'.  <p>  There are configurable server limits on both the dimensions (1024x1024 pixels by default) and uploaded  file size (1MB by default). Several different image formats are supported, but <strong>PNG</strong> and  <strong>JPEG</strong> are preferred due to the file size limit.  <p>  This resource has Cross-Site Request Forgery (XSRF) protection. To allow the request to  pass the XSRF check the caller needs to send an <code>X-Atlassian-Token</code> HTTP header with the  value <code>no-check</code>.  <p>  An example <a href=\"http://curl.haxx.se/\">curl</a> request to upload an image name 'avatar.png' would be:  <pre>  curl -X POST -u username:password -H \"X-Atlassian-Token: no-check\" http://example.com/rest/api/latest/users/jdoe/avatar.png -F avatar=@avatar.png  </pre>  <p>  Users are always allowed to update their own avatar. To update someone else's avatar the authenticated user must  have global <strong>ADMIN</strong> permission, or global <strong>SYS_ADMIN</strong> permission to update a  <strong>SYS_ADMIN</strong> user's avatar.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Watch**
> Watch(ctx, commitId)


Adds the authenticated user as a watcher for the specified commit.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository containing the commit  to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **commitId** | **string**| the &lt;i&gt;full {@link Commit#getId() ID}&lt;/i&gt; of the commit within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Watch_0**
> Watch_0(ctx, pullRequestId)


Make the authenticated user watch the specified pull request.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawApproval**
> WithdrawApproval(ctx, pullRequestId)


Remove approval from a pull request as the current user. This does not remove the user as a participant.  <p>  The authenticated user must have <strong>REPO_READ</strong> permission for the repository that this pull request  targets to call this resource.  <p>  <strong>Deprecated since 4.2</strong>. Use  /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug} instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pullRequestId** | **int64**| the id of the pull request within the repository | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

