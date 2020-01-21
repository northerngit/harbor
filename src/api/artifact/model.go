// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package artifact

import (
	"github.com/goharbor/harbor/src/pkg/artifact"
	"github.com/goharbor/harbor/src/pkg/tag/model/tag"
)

// Artifact is the overall view of artifact
type Artifact struct {
	artifact.Artifact
	Tags             []*Tag                     // the list of tags that attached to the artifact
	SubResourceLinks map[string][]*ResourceLink // the resource link for build history(image), values.yaml(chart), dependency(chart), etc
	// TODO add other attrs: signature, scan result, etc
}

// Tag is the overall view of tag
type Tag struct {
	tag.Tag
	// TODO add other attrs: signature, label, immutable status, etc
}

// Resource defines the specific resource of different artifacts: build history for image, values.yaml for chart, etc
type Resource struct {
	Content     []byte // the content of the resource
	ContentType string // the content type of the resource, returned as "Content-Type" header in API
}

// ResourceLink is a link via that a resource can be fetched
type ResourceLink struct {
	HREF     string
	Absolute bool // specify the href is an absolute URL or not
}

// Option is used to specify the properties returned when listing/getting artifacts
type Option struct {
	WithTag        bool
	TagOption      *TagOption // only works when WithTag is set to true
	WithScanResult bool
	WithSignature  bool // TODO move it to TagOption?
}

// TagOption is used to specify the properties returned when listing/getting tags
type TagOption struct {
	WithLabel           bool
	WithImmutableStatus bool
}

// TODO move this to GC controller?
// Option for pruning artifact records
// type Option struct {
//	 KeepUntagged bool // keep the untagged artifacts or not
// }
