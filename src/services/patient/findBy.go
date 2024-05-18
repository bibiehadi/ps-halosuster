package patientservice

func (s *patientService) IDisExist(patientId int64) bool {
	return s.patientRepository.IDisExist(patientId)
}
