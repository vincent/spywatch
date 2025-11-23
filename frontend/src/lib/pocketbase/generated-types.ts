/**
* This file was @generated using pocketbase-typegen
*/

import type PocketBase from 'pocketbase'
import type { RecordService } from 'pocketbase'

export enum Collections {
	Authorigins = "_authOrigins",
	Externalauths = "_externalAuths",
	Mfas = "_mfas",
	Otps = "_otps",
	Superusers = "_superusers",
	Competitors = "competitors",
	Resources = "resources",
	Snapshots = "snapshots",
	Users = "users",
	Workspaces = "workspaces",
}

// Alias types for improved usability
export type IsoDateString = string
export type RecordIdString = string
export type HTMLString = string

type ExpandType<T> = unknown extends T
	? T extends unknown
		? { expand?: unknown }
		: { expand: T }
	: { expand: T }

// System fields
export type BaseSystemFields<T = unknown> = {
	id: RecordIdString
	collectionId: string
	collectionName: Collections
} & ExpandType<T>

export type AuthSystemFields<T = unknown> = {
	email: string
	emailVisibility: boolean
	username: string
	verified: boolean
} & BaseSystemFields<T>

// Record types for each collection

export type AuthoriginsRecord = {
	collectionRef: string
	created?: IsoDateString
	fingerprint: string
	id: string
	recordRef: string
	updated?: IsoDateString
}

export type ExternalauthsRecord = {
	collectionRef: string
	created?: IsoDateString
	id: string
	provider: string
	providerId: string
	recordRef: string
	updated?: IsoDateString
}

export type MfasRecord = {
	collectionRef: string
	created?: IsoDateString
	id: string
	method: string
	recordRef: string
	updated?: IsoDateString
}

export type OtpsRecord = {
	collectionRef: string
	created?: IsoDateString
	id: string
	password: string
	recordRef: string
	sentTo?: string
	updated?: IsoDateString
}

export type SuperusersRecord = {
	created?: IsoDateString
	email: string
	emailVisibility?: boolean
	id: string
	password: string
	tokenKey: string
	updated?: IsoDateString
	verified?: boolean
}

export type CompetitorsRecord = {
	created?: IsoDateString
	external_id?: string
	found_resources?: string
	id: string
	logo?: string
	name: string
	owner?: RecordIdString
	updated?: IsoDateString
	url: string
	workspace?: RecordIdString
}

export type ResourcesRecord = {
	checked?: IsoDateString
	competitor?: RecordIdString
	created?: IsoDateString
	id: string
	updated?: IsoDateString
	url?: string
}

export type SnapshotsRecord = {
	content?: string
	created?: IsoDateString
	diff?: string
	id: string
	narrative?: string
	resource?: RecordIdString
	updated?: IsoDateString
}

export type UsersRecord = {
	avatar?: string
	created?: IsoDateString
	email: string
	emailVisibility?: boolean
	id: string
	name?: string
	password: string
	tokenKey: string
	updated?: IsoDateString
	verified?: boolean
}

export type WorkspacesRecord = {
	created?: IsoDateString
	id: string
	name?: string
	updated?: IsoDateString
}

// Response types include system fields and match responses from the PocketBase API
export type AuthoriginsResponse<Texpand = unknown> = Required<AuthoriginsRecord> & BaseSystemFields<Texpand>
export type ExternalauthsResponse<Texpand = unknown> = Required<ExternalauthsRecord> & BaseSystemFields<Texpand>
export type MfasResponse<Texpand = unknown> = Required<MfasRecord> & BaseSystemFields<Texpand>
export type OtpsResponse<Texpand = unknown> = Required<OtpsRecord> & BaseSystemFields<Texpand>
export type SuperusersResponse<Texpand = unknown> = Required<SuperusersRecord> & AuthSystemFields<Texpand>
export type CompetitorsResponse<Texpand = unknown> = Required<CompetitorsRecord> & BaseSystemFields<Texpand>
export type ResourcesResponse<Texpand = unknown> = Required<ResourcesRecord> & BaseSystemFields<Texpand>
export type SnapshotsResponse<Texpand = unknown> = Required<SnapshotsRecord> & BaseSystemFields<Texpand>
export type UsersResponse<Texpand = unknown> = Required<UsersRecord> & AuthSystemFields<Texpand>
export type WorkspacesResponse<Texpand = unknown> = Required<WorkspacesRecord> & BaseSystemFields<Texpand>

// Types containing all Records and Responses, useful for creating typing helper functions

export type CollectionRecords = {
	_authOrigins: AuthoriginsRecord
	_externalAuths: ExternalauthsRecord
	_mfas: MfasRecord
	_otps: OtpsRecord
	_superusers: SuperusersRecord
	competitors: CompetitorsRecord
	resources: ResourcesRecord
	snapshots: SnapshotsRecord
	users: UsersRecord
	workspaces: WorkspacesRecord
}

export type CollectionResponses = {
	_authOrigins: AuthoriginsResponse
	_externalAuths: ExternalauthsResponse
	_mfas: MfasResponse
	_otps: OtpsResponse
	_superusers: SuperusersResponse
	competitors: CompetitorsResponse
	resources: ResourcesResponse
	snapshots: SnapshotsResponse
	users: UsersResponse
	workspaces: WorkspacesResponse
}

// Type for usage with type asserted PocketBase instance
// https://github.com/pocketbase/js-sdk#specify-typescript-definitions

export type TypedPocketBase = PocketBase & {
	collection(idOrName: '_authOrigins'): RecordService<AuthoriginsResponse>
	collection(idOrName: '_externalAuths'): RecordService<ExternalauthsResponse>
	collection(idOrName: '_mfas'): RecordService<MfasResponse>
	collection(idOrName: '_otps'): RecordService<OtpsResponse>
	collection(idOrName: '_superusers'): RecordService<SuperusersResponse>
	collection(idOrName: 'competitors'): RecordService<CompetitorsResponse>
	collection(idOrName: 'resources'): RecordService<ResourcesResponse>
	collection(idOrName: 'snapshots'): RecordService<SnapshotsResponse>
	collection(idOrName: 'users'): RecordService<UsersResponse>
	collection(idOrName: 'workspaces'): RecordService<WorkspacesResponse>
}
