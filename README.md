# auth-acl-service
Authentication and Authorization Service built with Go

## Database Design

### Entities

#### users
- id
- email
- password_hash
- first_name
- last_name
- is_active
- created_at
- updated_at

#### roles
- id
- name
- description
- created_at
- updated_at

#### permissions
- id
- name
- description
- created_at
- updated_at

#### user_roles
- user_id
- role_id

#### role_permissions
- role_id
- permission_id

#### refresh_tokens
- id
- user_id
- token
- expires_at
- revoked_at
- created_at

### Relationships

- User has many Roles through user_roles
- Role has many Users through user_roles
- Role has many Permissions through role_permissions
- Permission has many Roles through role_permissions
- User has many RefreshTokens

### Design Decisions

- Every user must have at least one role.
- Permissions are assigned to roles, not directly to users.
- Role names and permission names should be lowercase.
- Refresh tokens support multiple sessions/devices.
- Device limits are business rules handled by consuming apps, not by the ACL service.
