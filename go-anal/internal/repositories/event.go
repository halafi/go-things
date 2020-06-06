package repositories

import (
	"github.com/halafi/go-things/go-anal/internal/models"
	"github.com/halafi/go-things/go-anal/pkg/utils"
)

// Stats are constants for which event stats can be retrieved
type Stats string

const (
	// StatsLocationCountry is stats for Country
	StatsLocationCountry Stats = "country"
	// StatsLocationCity is stats for City
	StatsLocationCity = "city"
	// StatsDeviceType is stats for Device Type
	StatsDeviceType = "deviceType"
	// StatsDevicePlatform is stats for Device Platform
	StatsDevicePlatform = "devicePlatform"
	// StatsDeviceOS is stats for OS
	StatsDeviceOS = "os"
	// StatsDeviceBrowser is stats for Browser
	StatsDeviceBrowser = "browser"
	// StatsDeviceLanguage is stats for Language
	StatsDeviceLanguage = "language"
	// StatsReferral is stats for Referral
	StatsReferral = "referral"
)

// EventRepository is storage repository for Events.
// We are not using a persistent datastore like SQL.
// Once the application is exited all stats will be gone.
type EventRepository struct {
	locationCountry *utils.Counter
	locationCity    *utils.Counter
	deviceType      *utils.Counter
	devicePlatform  *utils.Counter
	deviceOS        *utils.Counter
	deviceBrowser   *utils.Counter
	deviceLanguage  *utils.Counter
	referral        *utils.Counter
}

// NewEventRepository creates and returns new EventRepository
func NewEventRepository() *EventRepository {
	return &EventRepository{
		locationCountry: utils.NewCounter(),
		locationCity:    utils.NewCounter(),
		deviceType:      utils.NewCounter(),
		devicePlatform:  utils.NewCounter(),
		deviceOS:        utils.NewCounter(),
		deviceBrowser:   utils.NewCounter(),
		deviceLanguage:  utils.NewCounter(),
		referral:        utils.NewCounter(),
	}
}

// AddEvent adds an event to the repository
func (tr *EventRepository) AddEvent(ev *models.Event) {
	tr.locationCountry.Incr(ev.Location.Country)
	tr.locationCity.Incr(ev.Location.City)
	tr.deviceType.Incr(ev.Device.Type)
	tr.devicePlatform.Incr(ev.Device.Platform)
	tr.deviceOS.Incr(ev.Device.OS)
	tr.deviceBrowser.Incr(ev.Device.Browser)
	tr.deviceLanguage.Incr(ev.Device.Language)
	tr.referral.Incr(ev.Referral)
}

// Events returns stats for the specified event query
func (tr *EventRepository) Events(d Stats) map[string]uint64 {
	m := make(map[string]uint64)
	switch d {
	case StatsLocationCountry:
		m = tr.locationCountry.Items()
	case StatsLocationCity:
		m = tr.locationCity.Items()
	case StatsDeviceType:
		m = tr.deviceType.Items()
	case StatsDevicePlatform:
		m = tr.devicePlatform.Items()
	case StatsDeviceOS:
		m = tr.deviceOS.Items()
	case StatsDeviceBrowser:
		m = tr.deviceBrowser.Items()
	case StatsDeviceLanguage:
		m = tr.deviceLanguage.Items()
	case StatsReferral:
		m = tr.referral.Items()
	}
	return m
}
