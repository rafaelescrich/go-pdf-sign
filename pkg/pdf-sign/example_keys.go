package pdf_sign

import (
	"crypto"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

var test1024Key, test2048Key, test3072Key, test4096Key *rsa.PrivateKey

func init() {
	test1024Key = &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("123024078101403810516614073341068864574068590522569345017786163424062310013967742924377390210586226651760719671658568413826602264886073432535341149584680111145880576802262550990305759285883150470245429547886689754596541046564560506544976611114898883158121012232676781340602508151730773214407220733898059285561"),
			E: 65537,
		},
		D: fromBase10("118892427340746627750435157989073921703209000249285930635312944544706203626114423392257295670807166199489096863209592887347935991101581502404113203993092422730000157893515953622392722273095289787303943046491132467130346663160540744582438810535626328230098940583296878135092036661410664695896115177534496784545"),
		Primes: []*big.Int{
			fromBase10("12172745919282672373981903347443034348576729562395784527365032103134165674508405592530417723266847908118361582847315228810176708212888860333051929276459099"),
			fromBase10("10106518193772789699356660087736308350857919389391620140340519320928952625438936098550728858345355053201610649202713962702543058578827268756755006576249339"),
		},
	}
	test1024Key.Precompute()
	test2048Key = &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("14314132931241006650998084889274020608918049032671858325988396851334124245188214251956198731333464217832226406088020736932173064754214329009979944037640912127943488972644697423190955557435910767690712778463524983667852819010259499695177313115447116110358524558307947613422897787329221478860907963827160223559690523660574329011927531289655711860504630573766609239332569210831325633840174683944553667352219670930408593321661375473885147973879086994006440025257225431977751512374815915392249179976902953721486040787792801849818254465486633791826766873076617116727073077821584676715609985777563958286637185868165868520557"),
			E: 3,
		},
		D: fromBase10("9542755287494004433998723259516013739278699355114572217325597900889416163458809501304132487555642811888150937392013824621448709836142886006653296025093941418628992648429798282127303704957273845127141852309016655778568546006839666463451542076964744073572349705538631742281931858219480985907271975884773482372966847639853897890615456605598071088189838676728836833012254065983259638538107719766738032720239892094196108713378822882383694456030043492571063441943847195939549773271694647657549658603365629458610273821292232646334717612674519997533901052790334279661754176490593041941863932308687197618671528035670452762731"),
		Primes: []*big.Int{
			fromBase10("130903255182996722426771613606077755295583329135067340152947172868415809027537376306193179624298874215608270802054347609836776473930072411958753044562214537013874103802006369634761074377213995983876788718033850153719421695468704276694983032644416930879093914927146648402139231293035971427838068945045019075433"),
			fromBase10("109348945610485453577574767652527472924289229538286649661240938988020367005475727988253438647560958573506159449538793540472829815903949343191091817779240101054552748665267574271163617694640513549693841337820602726596756351006149518830932261246698766355347898158548465400674856021497190430791824869615170301029"),
		},
	}
	test2048Key.Precompute()
	test3072Key = &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("4799422180968749215324244710281712119910779465109490663934897082847293004098645365195947978124390029272750644394844443980065532911010718425428791498896288210928474905407341584968381379157418577471272697781778686372450913810019702928839200328075568223462554606149618941566459398862673532997592879359280754226882565483298027678735544377401276021471356093819491755877827249763065753555051973844057308627201762456191918852016986546071426986328720794061622370410645440235373576002278045257207695462423797272017386006110722769072206022723167102083033531426777518054025826800254337147514768377949097720074878744769255210076910190151785807232805749219196645305822228090875616900385866236956058984170647782567907618713309775105943700661530312800231153745705977436176908325539234432407050398510090070342851489496464612052853185583222422124535243967989533830816012180864309784486694786581956050902756173889941244024888811572094961378021"),
			E: 65537,
		},
		D: fromBase10("4068124900056380177006532461065648259352178312499768312132802353620854992915205894105621345694615110794369150964768050224096623567443679436821868510233726084582567244003894477723706516831312989564775159596496449435830457803384416702014837685962523313266832032687145914871879794104404800823188153886925022171560391765913739346955738372354826804228989767120353182641396181570533678315099748218734875742705419933837638038793286534641711407564379950728858267828581787483317040753987167237461567332386718574803231955771633274184646232632371006762852623964054645811527580417392163873708539175349637050049959954373319861427407953413018816604365474462455009323937599275324390953644555294418021286807661559165324810415569396577697316798600308544755741549699523972971375304826663847015905713096287495342701286542193782001358775773848824496321550110946106870685499577993864871847542645561943034990484973293461948058147956373115641615329"),
		Primes: []*big.Int{
			fromBase10("2378529069722721185825622840841310902793949682948530343491428052737890236476884657507685118578733560141370511507721598189068683665232991988491561624429938984370132428230072355214627085652359350722926394699707232921674771664421591347888367477300909202851476404132163673865768760147403525700174918450753162242834161458300343282159799476695001920226357456953682236859505243928716782707623075239350380352265954107362618991716602898266999700316937680986690964564264877"),
			fromBase10("2017811025336026464312837780072272578817919741496395062543647660689775637351085991504709917848745137013798005682591633910555599626950744674459976829106750083386168859581016361317479081273480343110649405858059581933773354781034946787147300862495438979895430001323443224335618577322449133208754541656374335100929456885995320929464029817626916719434010943205170760536768893924932021302887114400922813817969176636993508191950649313115712159241971065134077636674146073"),
		},
	}
	test3072Key.Precompute()
	test4096Key = &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("633335480064287130853997429184971616419051348693342219741748040433588285601270210251206421401040394238592139790962887290698043839174341843721930134010306454716566698330215646704263665452264344664385995704186692432827662862845900348526672531755932642433662686500295989783595767573119607065791980381547677840410600100715146047382485989885183858757974681241303484641390718944520330953604501686666386926996348457928415093305041429178744778762826377713889019740060910363468343855830206640274442887621960581569183233822878661711798998132931623726434336448716605363514220760343097572198620479297583609779817750646169845195672483600293522186340560792255595411601450766002877850696008003794520089358819042318331840490155176019070646738739580486357084733208876620846449161909966690602374519398451042362690200166144326179405976024265116931974936425064291406950542193873313447617169603706868220189295654943247311295475722243471700112334609817776430552541319671117235957754556272646031356496763094955985615723596562217985372503002989591679252640940571608314743271809251568670314461039035793703429977801961867815257832671786542212589906513979094156334941265621017752516999186481477500481433634914622735206243841674973785078408289183000133399026553"),
			E: 65537,
		},
		D: fromBase10("439373650557744155078930178606343279553665694488479749802070836418412881168612407941793966086633543867614175621952769177088930851151267623886678906158545451731745754402575409204816390946376103491325109185445659065122640946673660760274557781540431107937331701243915001777636528502669576801704352961341634812275635811512806966908648671988644114352046582195051714797831307925775689566757438907578527366568747104508496278929566712224252103563340770696548181508180254674236716995730292431858611476396845443056967589437890065663497768422598977743046882539288481002449571403783500529740184608873520856954837631427724158592309018382711485601884461168736465751756282510065053161144027097169985941910909130083273691945578478173708396726266170473745329617793866669307716920992380350270584929908460462802627239204245339385636926433446418108504614031393494119344916828744888432279343816084433424594432427362258172264834429525166677273382617457205387388293888430391895615438030066428745187333897518037597413369705720436392869403948934993623418405908467147848576977008003556716087129242155836114780890054057743164411952731290520995017097151300091841286806603044227906213832083363876549637037625314539090155417589796428888619937329669464810549362433"),
		Primes: []*big.Int{
			fromBase10("25745433817240673759910623230144796182285844101796353869339294232644316274580053211056707671663014355388701931204078502829809738396303142990312095225333440050808647355535878394534263839500592870406002873182360027755750148248672968563366185348499498613479490545488025779331426515670185366021612402246813511722553210128074701620113404560399242413747318161403908617342170447610792422053460359960010544593668037305465806912471260799852789913123044326555978680190904164976511331681163576833618899773550873682147782263100803907156362439021929408298804955194748640633152519828940133338948391986823456836070708197320166146761"),
			fromBase10("24599914864909676687852658457515103765368967514652318497893275892114442089314173678877914038802355565271545910572804267918959612739009937926962653912943833939518967731764560204997062096919833970670512726396663920955497151415639902788974842698619579886297871162402643104696160155894685518587660015182381685605752989716946154299190561137541792784125356553411300817844325739404126956793095254412123887617931225840421856505925283322918693259047428656823141903489964287619982295891439430302405252447010728112098326033634688757933930065610737780413018498561434074501822951716586796047404555397992425143397497639322075233073"),
		},
	}
	test4096Key.Precompute()
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

type certKeyPair struct {
	Certificate *x509.Certificate
	PrivateKey  *crypto.PrivateKey
}

func createTestCertificate(sigAlg x509.SignatureAlgorithm) (certKeyPair, error) {
	signer, err := createTestCertificateByIssuer("Eddard Stark", nil, sigAlg, true)
	if err != nil {
		return certKeyPair{}, err
	}
	pair, err := createTestCertificateByIssuer("Jon Snow", signer, sigAlg, false)
	if err != nil {
		return certKeyPair{}, err
	}
	return *pair, nil
}

func createTestCertificateByIssuer(name string, issuer *certKeyPair, sigAlg x509.SignatureAlgorithm, isCA bool) (*certKeyPair, error) {
	var (
		err        error
		priv       crypto.PrivateKey
		derCert    []byte
		issuerCert *x509.Certificate
		issuerKey  crypto.PrivateKey
	)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 32)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   name,
			Organization: []string{"Acme Co"},
		},
		NotBefore:   time.Now().Add(-1 * time.Second),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageEmailProtection},
	}
	if issuer != nil {
		issuerCert = issuer.Certificate
		issuerKey = *issuer.PrivateKey
	}
	switch sigAlg {
	case x509.SHA1WithRSA:
		priv = test1024Key
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA1WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA1
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA1
		}
	case x509.SHA256WithRSA:
		priv = test2048Key
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA256WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA256
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA256
		}
	case x509.SHA384WithRSA:
		priv = test3072Key
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA384WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA384
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA256
		}
	case x509.SHA512WithRSA:
		priv = test4096Key
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA512WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA512
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA256
		}
	case x509.ECDSAWithSHA1:
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, err
		}
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA1WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA1
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA1
		}
	case x509.ECDSAWithSHA256:
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, err
		}
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA256WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA256
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA256
		}
	case x509.ECDSAWithSHA384:
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			return nil, err
		}
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA384WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA384
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA256
		}
	case x509.ECDSAWithSHA512:
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		if err != nil {
			return nil, err
		}
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA512WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA512
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA256
		}
	case x509.DSAWithSHA1:
		var dsaPriv dsa.PrivateKey
		params := &dsaPriv.Parameters
		err = dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160)
		if err != nil {
			return nil, err
		}
		err = dsa.GenerateKey(&dsaPriv, rand.Reader)
		if err != nil {
			return nil, err
		}
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			template.SignatureAlgorithm = x509.SHA1WithRSA
		case *ecdsa.PrivateKey:
			template.SignatureAlgorithm = x509.ECDSAWithSHA1
		case *dsa.PrivateKey:
			template.SignatureAlgorithm = x509.DSAWithSHA1
		}
		priv = &dsaPriv
	}
	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
		template.BasicConstraintsValid = true
	}
	if issuer == nil {
		// no issuer given,make this a self-signed root cert
		issuerCert = &template
		issuerKey = priv
	}

	log.Println("creating cert", name, "issued by", issuerCert.Subject.CommonName, "with sigalg", sigAlg)
	switch priv.(type) {
	case *rsa.PrivateKey:
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*rsa.PrivateKey).Public(), issuerKey.(*rsa.PrivateKey))
		case *ecdsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*rsa.PrivateKey).Public(), issuerKey.(*ecdsa.PrivateKey))
		case *dsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*rsa.PrivateKey).Public(), issuerKey.(*dsa.PrivateKey))
		}
	case *ecdsa.PrivateKey:
		switch issuerKey.(type) {
		case *rsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*ecdsa.PrivateKey).Public(), issuerKey.(*rsa.PrivateKey))
		case *ecdsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*ecdsa.PrivateKey).Public(), issuerKey.(*ecdsa.PrivateKey))
		case *dsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*ecdsa.PrivateKey).Public(), issuerKey.(*dsa.PrivateKey))
		}
	case *dsa.PrivateKey:
		pub := &priv.(*dsa.PrivateKey).PublicKey
		switch issuerKey := issuerKey.(type) {
		case *rsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, pub, issuerKey)
		case *ecdsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*dsa.PublicKey), issuerKey)
		case *dsa.PrivateKey:
			derCert, err = x509.CreateCertificate(rand.Reader, &template, issuerCert, priv.(*dsa.PublicKey), issuerKey)
		}
	}
	if err != nil {
		return nil, err
	}
	if len(derCert) == 0 {
		return nil, fmt.Errorf("no certificate created, probably due to wrong keys. types were %T and %T", priv, issuerKey)
	}
	cert, err := x509.ParseCertificate(derCert)
	if err != nil {
		return nil, err
	}
	pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	return &certKeyPair{
		Certificate: cert,
		PrivateKey:  &priv,
	}, nil
}

type TestFixture struct {
	Input       []byte
	Certificate *x509.Certificate
	PrivateKey  *rsa.PrivateKey
}
