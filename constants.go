package mystem

// Lemma quality
const (
	Dictionary  = 0        // слово из словаря
	Bastard     = 1        // не словарное
	Sob         = 2        // из "быстрого словаря"
	Prefixoid   = 4        // словарное + стандартный префикс (авто- мото- кино- фото-) всегда в компании с Bastard или Sob
	Foundling   = 8        // непонятный набор букв, но проходящий в алфавит
	BadRequest  = 16       // доп. флаг.: "плохая лемма" при наличии "хорошей" альтернативы ("махать" по форме "маша")
	FromEnglish = 65536    // переведено с английского
	ToEnglish   = 131072   // переведено на английский
	Untranslit  = 262144   // "переведено" с транслита
	Overrode    = 1048576  // текст леммы был перезаписан
	Fix         = 16777216 // слово из фикс-листа
)

// Lemma grammemes
const (
	Invalid          = 0
	Before           = 126
	Postposition     = 127          // POSTP
	First            = Postposition // same as Postposition
	Adjective        = 128          // A      // Nomenus
	Adverb           = 129          // ADV
	Composite        = 130          // COM(P)
	Conjunction      = 131          // CONJ
	Interjunction    = 132          // INTJ
	Numeral          = 133          // NUM
	Particle         = 134          // PCL
	Preposition      = 135          // PRE(P)
	Substantive      = 136          // S
	Verb             = 137          // V
	AdjNumeral       = 138          // ANUM
	AdjPronoun       = 139          // APRO
	AdvPronoun       = 140          // ADVPRO
	SubstPronoun     = 141          // SPRO
	Article          = 142          // артикли
	PartOfIdiom      = 143          // части идиом (прежде всего иностр. слов)
	LastPartOfSpeech = PartOfIdiom  // same as PartOfIdiom
	Reserved         = 144          // зарезервировано    // особые пометы
	Abbreviation     = 145          // сокращения
	IrregularStem    = 146          // чередование в корне (или супплетивизм)
	Informal         = 147          // разговорная форма
	Distort          = 148          // искаженная форма
	Contracted       = 149          // стяжённая форма (фр. q' и т.п.)
	Obscene          = 150          // обсц
	Rare             = 151          // редк
	Awkward          = 152          // затр
	Obsolete         = 153          // устар
	SubstAdjective   = 154          // адъект
	FirstName        = 155          // имя
	Surname          = 156          // фам
	Patr             = 157          // отч
	Geo              = 158          // гео
	Proper           = 159          // собств
	Present          = 160          // наст  // Tempus
	Notpast          = 161          // непрош
	Past             = 162          // прош
	Future           = 163          // буд. время (фр., ит.)
	Past2            = 164          // фр. passe simple, ит. passato remoto
	Nominative       = 165          // им    // Casus
	Genitive         = 166          // род
	Dative           = 167          // дат
	Accusative       = 168          // вин
	Instrumental     = 169          // твор
	Ablative         = 170          // пр
	Partitive        = 171          // парт(вин2)
	Locative         = 172          // местн(пр2)
	Vocative         = 173          // звательный
	Singular         = 174          // ед    // Numerus
	Plural           = 175          // мн
	Gerund           = 176          // деепр // Modus
	Infinitive       = 177          // инф
	Participle       = 178          // прич
	Indicative       = 179          // изъяв
	Imperative       = 180          // пов
	Conditional      = 181          // усл. наклонение (фр. =ит.)
	Subjunctive      = 182          // сослаг. накл. (фр. =ит.)
	Short            = 183          // кр    // Gradus
	Full             = 184          // полн
	Superlative      = 185          // прев
	Comparative      = 186          // срав
	Possessive       = 187          // притяж
	Person1          = 188          // 1-л   // Personae
	Person2          = 189          // 2-л
	Person3          = 190          // 3-л
	Feminine         = 191          // жен   // Gender (genus)
	Masculine        = 192          // муж
	Neuter           = 193          // сред
	MasFem           = 194          // мж
	Perfect          = 195          // сов   // Perfectum-imperfectum (Accept)
	Imperfect        = 196          // несов
	Passive          = 197          // страд // Voice (Genus)
	Active           = 198          // действ
	Reflexive        = 199          // возвратные
	Impersonal       = 200          // безличные
	Animated         = 201          // од    // Animated
	Inanimated       = 202          // неод
	Praedic          = 203          // прдк
	Parenth          = 204          // вводн
	Transitive       = 205          // пе     //transitivity
	Intransitive     = 206          // нп
	Definite         = 207          // опред. члены   //definiteness
	Indefinite       = 208          // неопред. члены
	SimConj          = 209          // сочинит. (для союзов)
	SubConj          = 210          // подчинит. (для союзов)
	PronounConj      = 211          // местоимение-союз ("я знаю, _что_ вы сделали прошлым летом")
	CorrelateConj    = 212          // вторая зависимая часть парных союзов - "если ... _то_ ... ", "как ... _так_ и ..."
	AuxVerb          = 213          //вспомогательный глагол в аналитической форме ("_будем_ думать")
)
