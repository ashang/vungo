package vast

import "encoding/xml"

// Vast type represents the root <VAST> tag.
type Vast struct {
	XMLName xml.Name `xml:"VAST"`
	Version Version  `xml:"version,attr"`

	Ads    []*Ad    `xml:"Ad"`
	Errors []string `xml:"Error,omitempty"` // One or more URI's, likely tracking pixels, to request in case of no ad.
}

// FindFirstInlineLinearCreative method inspects through all of its inline ads and finds the first
// linear creative within, or returns nil when found nothing.
func (v Vast) FindFirstInlineLinearCreative() *Linear {
	for _, ad := range v.Ads {
		if ad.InLine == nil {
			continue
		}

		for _, c := range ad.InLine.Creatives {
			if c.Linear != nil {
				return c.Linear
			}
		}
	}
	return nil
}

// Validate method validates the Vast element according to the VAST.
// Version, and Ads are required.
func (vast *Vast) Validate() error {

	if err := vast.Version.Validate(); err != nil {
		return err
	}

	if len(vast.Ads) == 0 {
		return ErrVastMissAd
	}

	for _, ad := range vast.Ads {
		if err := ad.Validate(); err != nil {
			return err
		}
	}

	return nil
}
