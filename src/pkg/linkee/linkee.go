package linkee

import (
	"encoding/json"
)

type LinkyLixee struct {
	ActiveEnerfyOutD01          int         `json:"active_enerfy_out_d01"`
	ActiveEnergyOutD01          int         `json:"active_energy_out_d01"`
	ActivePower                 int         `json:"active_power"`
	ActivePowerMax              int         `json:"active_power_max"`
	ActivePowerPhB              int         `json:"active_power_ph_b"`
	ActiveRegisterTierDelivered string      `json:"active_register_tier_delivered"`
	ApparentPower               int         `json:"apparent_power"`
	AvailablePower              int         `json:"available_power"`
	AverageRmsVoltageMeasPeriod int         `json:"average_rms_voltage_meas_period"`
	CurrentDate                 string      `json:"current_date"`
	CurrentIndexTarif           int         `json:"current_index_tarif"`
	CurrentPrice                string      `json:"current_price"`
	CurrentSummDelivered        float64     `json:"current_summ_delivered"`
	CurrentTarif                string      `json:"current_tarif"`
	CurrentTier1SummDelivered   float64     `json:"current_tier1_summ_delivered"`
	CurrentTier2SummDelivered   int         `json:"current_tier2_summ_delivered"`
	DrawnVAMaxN1                int         `json:"drawn_v_a_max_n1"`
	Linkquality                 int         `json:"linkquality"`
	Message1                    string      `json:"message1"`
	Message2                    string      `json:"message2"`
	MeterSerialNumber           string      `json:"meter_serial_number"`
	PowerThreshold              interface{} `json:"power_threshold"`
	Relais                      int         `json:"relais"`
	RmsCurrent                  int         `json:"rms_current"`
	RmsCurrentMax               int         `json:"rms_current_max"`
	RmsVoltage                  int         `json:"rms_voltage"`
	SiteID                      string      `json:"site_id"`
	SoftwareRevision            interface{} `json:"software_revision"`
	StatusRegister              string      `json:"status_register"`
	StatusRegisterBreakout      struct {
		CacheBorneDist      string `json:"cache_borne_dist"`
		CommEuridis         string `json:"comm_euridis"`
		ContactSec          string `json:"contact_sec"`
		DepassementRefPow   int    `json:"depassement_ref_pow"`
		EtatCpl             string `json:"etat_cpl"`
		Horloge             string `json:"horloge"`
		OrganeCoupure       string `json:"organe_coupure"`
		PointeMobile        string `json:"pointe_mobile"`
		PreavisPointeMobile string `json:"preavis_pointe_mobile"`
		Producteur          int    `json:"producteur"`
		SensEnergieActive   string `json:"sens_energie_active"`
		SurtensionPhase     int    `json:"surtension_phase"`
		SyncCpl             string `json:"sync_cpl"`
		TarifDist           string `json:"tarif_dist"`
		TarifFour           string `json:"tarif_four"`
		TempoDemain         string `json:"tempo_demain"`
		TempoJour           string `json:"tempo_jour"`
		TypeTic             string `json:"type_tic"`
	} `json:"status_register_breakout"`
	StatusRegisterCacheBorneDist      string `json:"status_register_cache_borne_dist"`
	StatusRegisterCommEuridis         string `json:"status_register_comm_euridis"`
	StatusRegisterContactSec          string `json:"status_register_contact_sec"`
	StatusRegisterDepassementRefPow   int    `json:"status_register_depassement_ref_pow"`
	StatusRegisterEtatCpl             string `json:"status_register_etat_cpl"`
	StatusRegisterHorloge             string `json:"status_register_horloge"`
	StatusRegisterOrganeCoupure       string `json:"status_register_organe_coupure"`
	StatusRegisterPointeMobile        string `json:"status_register_pointe_mobile"`
	StatusRegisterPreavisPointeMobile string `json:"status_register_preavis_pointe_mobile"`
	StatusRegisterProducteur          int    `json:"status_register_producteur"`
	StatusRegisterSensEnergieActive   string `json:"status_register_sens_energie_active"`
	StatusRegisterSurtensionPhase     int    `json:"status_register_surtension_phase"`
	StatusRegisterSyncCpl             string `json:"status_register_sync_cpl"`
	StatusRegisterTarifDist           string `json:"status_register_tarif_dist"`
	StatusRegisterTarifFour           string `json:"status_register_tarif_four"`
	StatusRegisterTempoDemain         string `json:"status_register_tempo_demain"`
	StatusRegisterTempoJour           string `json:"status_register_tempo_jour"`
	StatusRegisterTypeTic             string `json:"status_register_type_tic"`
	Update                            struct {
		InstalledVersion int    `json:"installed_version"`
		LatestVersion    int    `json:"latest_version"`
		State            string `json:"state"`
	} `json:"update"`
	UpdateAvailable bool `json:"update_available"`
	WarnDPS         int  `json:"warn_d_p_s"`
}

func New(data []byte) (*LinkyLixee, error) {
	var newValue LinkyLixee
	err := json.Unmarshal(data, &newValue)
	if err != nil {
		return nil, err
	}
	return &newValue, nil
}
