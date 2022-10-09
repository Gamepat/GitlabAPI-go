package package_api

import "time"

/* ---------- Attributes ----------*/

// orderBy The field to use as order
type orderBy struct {
	slug string
}

func (o orderBy) String() string {
	return o.slug
}

// The field to use as order
var (
	OrderByCreatedAt = orderBy{"created_at"}
	OrderByName      = orderBy{"name"}
	OrderByVersion   = orderBy{"version"}
	OrderByType      = orderBy{"type"}
)

// sort The direction of the order
type sort struct {
	slug string
}

func (s sort) String() string {
	return s.slug
}

// The direction of the order
var (
	SortAscending  = sort{"asc"}
	SortDescending = sort{"desc"}
)

// packageType Filter the returned packages by type
type packageType struct {
	slug string
}

func (p packageType) String() string {
	return p.slug
}

// Filter the returned packages by type
var (
	PackageTypeAll       = packageType{""}
	PackageTypeConan     = packageType{"conan"}
	PackageTypeMaven     = packageType{"maven"}
	PackageTypeNpm       = packageType{"npm"}
	PackageTypePypi      = packageType{"pypi"}
	PackageTypeComposer  = packageType{"composer"}
	PackageTypeNuget     = packageType{"nuget"}
	PackageTypeHelm      = packageType{"helm"}
	PackageTypeTerraform = packageType{"terraform_module"}
	PackageTypeGolang    = packageType{"golang"}
	PackageTypeGeneric   = packageType{"generic"}
)

// packageStatus Filter the returned packages by status
type packageStatus struct {
	slug string
}

func (p packageStatus) String() string {
	return p.slug
}

// Filter the returned packages by status
var (
	PackageStatusDefault            = packageStatus{"default"}
	PackageStatusHidden             = packageStatus{"hidden"}
	PackageStatusProcessing         = packageStatus{"processing"}
	PackageStatusError              = packageStatus{"error"}
	PackageStatusPendingDestruction = packageStatus{"pending_destruction"}
)

/* ------- Api Response Types ------- */

type PackageApiResponse []struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	PackageType string `json:"package_type"`
	Links       struct {
		WebPath       string `json:"web_path"`
		DeleteAPIPath string `json:"delete_api_path"`
	} `json:"_links"`
	CreatedAt        time.Time `json:"created_at"`
	LastDownloadedAt time.Time `json:"last_downloaded_at"`
	Pipelines        []struct {
		ID        int       `json:"id"`
		Status    string    `json:"status"`
		Ref       string    `json:"ref"`
		Sha       string    `json:"sha"`
		WebURL    string    `json:"web_url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      struct {
			Name      string `json:"name"`
			AvatarURL string `json:"avatar_url"`
		} `json:"user"`
	} `json:"pipelines"`
	Versions []struct {
		ID        int       `json:"id"`
		Version   string    `json:"version"`
		CreatedAt time.Time `json:"created_at"`
		Pipelines []struct {
			ID        int       `json:"id"`
			Status    string    `json:"status"`
			Ref       string    `json:"ref"`
			Sha       string    `json:"sha"`
			WebURL    string    `json:"web_url"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			User      struct {
				Name      string `json:"name"`
				AvatarURL string `json:"avatar_url"`
			} `json:"user"`
		} `json:"pipelines"`
	} `json:"versions"`
}
