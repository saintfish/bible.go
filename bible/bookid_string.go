// generated by stringer --type BookID; DO NOT EDIT

package bible

import "fmt"

const _BookID_name = "InvalidBookGenesisExodusLeviticusNumbersDeuteronomyJoshuaJudgesRuthSamuel1Samuel2Kings1Kings2Chronicles1Chronicles2EzraNehemiahEstherJobPsalmProverbsEcclesiastesSongOfSongsIsaiahJeremiahLamentationsEzekielDanielHoseaJoelAmosObadiahJonahMicahNahumHabakkukZephaniahHaggaiZechariahMalachiMatthewMarkLukeJohnActsRomansCorinthians1Corinthians2GalatiansEphesiansPhilippiansColossiansThessalonians1Thessalonians2Timothy1Timothy2TitusPhilemonHebrewsJamesPeter1Peter2John1John2John3JudeRevelation"

var _BookID_index = [...]uint16{0, 11, 18, 24, 33, 40, 51, 57, 63, 67, 74, 81, 87, 93, 104, 115, 119, 127, 133, 136, 141, 149, 161, 172, 178, 186, 198, 205, 211, 216, 220, 224, 231, 236, 241, 246, 254, 263, 269, 278, 285, 292, 296, 300, 304, 308, 314, 326, 338, 347, 356, 367, 377, 391, 405, 413, 421, 426, 434, 441, 446, 452, 458, 463, 468, 473, 477, 487}

func (i BookID) String() string {
	if i < 0 || i+1 >= BookID(len(_BookID_index)) {
		return fmt.Sprintf("BookID(%d)", i)
	}
	return _BookID_name[_BookID_index[i]:_BookID_index[i+1]]
}
