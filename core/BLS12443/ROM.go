/*
 * Copyright (c) 2012-2020 MIRACL UK Ltd.
 *
 * This file is part of MIRACL Core
 * (see https://github.com/miracl/core).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* Fixed Data in ROM - Field and Curve parameters */

package BLS12443

// Base  bits= 60
var Modulus = [...]Chunk{0xFFAAAAAB0AAAAAB, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}
var R2modp = [...]Chunk{0xF5CDA0EB0AD64E2, 0xDE66AA74FAE046B, 0x8B9229B598075AE, 0xDFAFBC5DB0E321E, 0x9AF61017ADA96A8, 0xFFAEA657DE81FE0, 0x1800170F84B9395, 0x8E3DD}
var ROI = [...]Chunk{0xFFAAAAAB0AAAAAA, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}
var SQRTm3 = [...]Chunk{0xFDAAAAAD4AAAAA8, 0x6035B554DBDEB3F, 0x76827C02FE4D504, 0x58F2FA9D667766E, 0xE1E6D0C29664B91, 0xBAD4B70A8CFD628, 0xDCB20A9C8688A07, 0x575A53}
var CRu = [...]Chunk{0xFEAAAAAC2AAAAA9, 0xEAB145552032AFF, 0xE6D162819938FE1, 0xA51F36592AC40E7, 0xE71AEDBDFA5AE98, 0xE2AB89302F9EF2A, 0xDCB20A9C8688E1B, 0x575A53}

const MConst Chunk = 0xC04000035FFFFFD

var Fra = [...]Chunk{0xD3B144FC49551C8, 0xB811B6AF7EB3463, 0xC145F9D97D38DC4, 0x75E17CAAF11AFC7, 0x300FB086DD029E2, 0x8C6E48295004F76, 0x31DFF2E447F979B, 0x440309}
var Frb = [...]Chunk{0x2BF965AEC1558E3, 0xBD1B1EA5E5D365C, 0x95DA4F26B6EBCFA, 0x7B69F569FDF5B99, 0xBC3F5A32814E7BD, 0x7E14132C823B8B6, 0xAAD217B83E8FA94, 0x13574A}

// *** rom curve parameters *****
// Ate Bits= 75
// G2 Table size= 78
const CURVE_Cof_I int = 0

var CURVE_Cof = [...]Chunk{0xFF0000011FFFFFF, 0x403F, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

const CURVE_B_I int = 4

var CURVE_B = [...]Chunk{0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Order = [...]Chunk{0xEBC000000000001, 0x81987F70023FFFF, 0x340905DC5A3E80, 0x94374F2CDDD3713, 0x10405F3D0D062F6, 0x0, 0x0, 0x0}
var CURVE_Gx = [...]Chunk{0xE6CA5B278ABC574, 0xFC6EC8ECBE2A6D2, 0x29BD7C18640BFDD, 0x8814D9BD1F29789, 0xABA282B7CE9049E, 0x43412B4E0DF13E7, 0xB62517BEDDB681E, 0x39E6C4}
var CURVE_Gy = [...]Chunk{0x994157597A9CE3F, 0xF905D6B69C431DC, 0x3476A5568900742, 0xB09E5F3D3F429D6, 0xFAD8CED72F5769D, 0xB19187E10D9FBDD, 0x87504385DD443B2, 0x12621F}
var CURVE_HTPC = [...]Chunk{0x14C40448911B6DF, 0xA05D73B4005575A, 0x19D3FFA0043A6A, 0x9DCBB950453AAB1, 0x8F188A39830074A, 0x886F6C1313A12B, 0x323F40FBC329680, 0x4F6592}

var CURVE_Bnx = [...]Chunk{0xFF0000012000000, 0x403F, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Pxa = [...]Chunk{0x65025557F3AF148, 0xAF4DBD1F46E75B8, 0x1730875544DD695, 0xE37A84E30576317, 0x2760C16CE339FC8, 0x258FB84649C73F2, 0x45DECE4C1502219, 0x2C5EC}
var CURVE_Pxb = [...]Chunk{0xC4F6E6BB36279B6, 0x530E32492D8015B, 0x3F16872759E2951, 0xDD4ACB43BA5202F, 0xFE06E165089BD13, 0xDA296B65CE7FAEA, 0x16EC13B532F64C9, 0xF64B9}
var CURVE_Pya = [...]Chunk{0xBFC0EA271C0B975, 0x971B0987719FB7D, 0x783587349A91EFB, 0x5B26E1E164C5BA2, 0x34AE6E2149122B6, 0x919E2C134F4019F, 0x46D510457BD2F, 0x45EA07}
var CURVE_Pyb = [...]Chunk{0x3EE1FF0082464AD, 0xF8557C8A8F0CF1B, 0x6062A50D0077B77, 0x806D6E17CEEF5EB, 0x8F95D948A372182, 0x1690DECDC940F93, 0xEB1BEA8F586B9E9, 0x34015D}

var CURVE_Ad = [...]Chunk{0x2B7CF1B298E3466, 0x6C9FE584ADB58A, 0xAE618C4441D091D, 0x649C886E6D2CC2B, 0x7BDA1ECDF8CA78F, 0x9FC3FCD80DBBB27, 0x74D4F1F9729760C, 0x19EA0E}
var CURVE_Bd = [...]Chunk{0x1CCB86130F780D9, 0x2D5A7C07F4A51A0, 0x99A51B6FD01A1BB, 0xAE12B32E246480, 0x45C8B048B4D289, 0xAD74B1FED8D288B, 0x276B220D6B08F5B, 0x29D1A2}
var PC = [53][8]Chunk{{0xCA2E8BA3F34E48, 0x9C1F67AB89650A, 0x4FA71DA7BA330, 0xD104A18D275FFDC, 0xE98A8B61C21689C, 0xC03CB4E05961F09, 0xAB977EAB088A0C6, 0xF2910}, {0x8DAB6776604D319, 0x5E333A54FF1F29E, 0xB06B4CD9E4CBF06, 0xF8DAA9493D66B74, 0x26D3AB458B9CC61, 0xBACFA058DFB0295, 0x3F9DF9DF3EB0CB5, 0x3DFC8E}, {0xF5334873EA1F60A, 0x9549B6EC7E13A86, 0xA5D45908C4E2558, 0xE0D733926F79B73, 0xCDE1872F1991243, 0x8E00DEE03D4D2CA, 0xD168F4937C08D2C, 0x53D3D2}, {0x44C150037499741, 0x3BD7556E9D85B54, 0x70B485348521E19, 0x4DCEB44DF6F5242, 0x7F602EC95D38A81, 0xD3B6B4317C39525, 0xAB827397CF26EAB, 0xF40BE}, {0xD780C2EE8A9B7DD, 0x5BFD0A87BA123F6, 0xB704149A656E9B7, 0xFB3D8D8250169A7, 0xDBAFE384E1B5A80, 0xE33E6DB698822E2, 0x1ADFD5731D2D4D2, 0x27E3F5}, {0xEE18F93C41C4477, 0x8822B165E9F3BB5, 0xF81BCEAD751713B, 0xF3FD57505403933, 0x543A69376B6BAE1, 0x2E2C188E07B1B4A, 0xFA47058B97839A4, 0x129214}, {0xC46466ED0815090, 0x82A8477388286F6, 0x55415FA6E5B59C3, 0x4FCDBA32F596982, 0x3BEBA0B8B800647, 0x8067EDEC24FD2DC, 0x23735D491406F8D, 0x27031A}, {0x237171BD500BAD1, 0x54E5DABBDA52179, 0x39EED6D9840047, 0x9F62B3C5BBE70BB, 0x5B682935D460B68, 0xE37CCA944164133, 0xBAA9B754C8769C5, 0x4F71D5}, {0x48821C29E8081C0, 0x8B10F8432B1F73F, 0xF33341BDF35547E, 0x44D8A286F2D64E3, 0x7CF3F353E5AFD5, 0x77CBAA1A38E67B6, 0x474039CD79B26B0, 0x15F7E2}, {0xF92A02857226171, 0x670D84A4DB8412E, 0x4C07F33AC92190E, 0xFE7C239718D531F, 0x16FEB62FC918257, 0x551F5FA232DBF4C, 0xA2262B91DD663A6, 0x513BA6}, {0x139421007A91954, 0xC897089BCA8F8E1, 0xF39670B7F8E69E7, 0xECD5CC0FAC3C084, 0x4A62E46CB8D5D73, 0xD6958651AF783AE, 0x6214D7EAF10E794, 0x4A1CF5}, {0x5035F58B92801A5, 0x9DD21175856DFA0, 0x1C54471AC39F6E8, 0x4A4455B68F864, 0xCEBF3444F479330, 0xA1896ECB026AE8F, 0x13798A8AD8FBB10, 0x40036}, {0x12593C28F9D1C0A, 0x9E53BCD23503B31, 0x74FF15EBD437870, 0x814D22AAA5002D5, 0xE3CD6452AC2C033, 0xCEDA7682EAD8194, 0xCA8D948BF6077F1, 0x4B6163}, {0x695774ADF0A734C, 0x3A4F0FAF18848AE, 0xA329B8FFF150217, 0xCFA4BB81D32CB6C, 0x46753EE2A5BF898, 0x76FF71FFF34CE9C, 0xF985D613FEBD9A1, 0x4AE3B1}, {0xA6F72CFF1F37A8D, 0xC08952B61FB4664, 0x75F3020A8B74C59, 0xF837EF511EC585E, 0xACEE128F587CFCE, 0xC68B982093597FB, 0x11DBD5819E24045, 0x1305B0}, {0x8AD991A3124740B, 0xB1BA6424469E9C6, 0x6D0500D1A8B0336, 0x4697B3B665ECD64, 0x789128B282C2A24, 0x7C5FEA435E1E511, 0xBF4998179A5D89A, 0x370542}, {0x6EB49AAAD27D311, 0xC44B2B709A2C349, 0x97D96FEB98EEF0C, 0xEDFE38DCE8507CB, 0xD7302DB5F6B3FF9, 0xA8E476581D84921, 0xA51CC54C7152F5, 0x47A513}, {0xC2E3ACF1F1FC577, 0x819DC8582B977B, 0x2CAB1522882CAB8, 0xC8900CA9AF3D679, 0x479599396371CE7, 0xCDCE25F02C5B347, 0xA58D8B2A509BE1A, 0x2FCD0A}, {0x657289AE9150A48, 0x38C6B088F4220C0, 0x3483EB38D0F458, 0x589009EAF978350, 0x7EBB824407235B4, 0x56CA72A2CCB14E4, 0xEA66FC272BCDCCF, 0x439DC3}, {0x6623DEF9357D628, 0xFD7BA8E7DE31234, 0x99C040938DB9E8B, 0xB0FE406C9D3CDB6, 0xD38F65BE751A703, 0x7CA37104A247DEB, 0x7ECA5B210A36965, 0x3932D3}, {0xA23FE2FB17D6EEC, 0x7ABF374B72AAE7, 0x62C2BF1860877F7, 0x8C4635085D4226F, 0x295B0F5706768CC, 0x227C909C7AEB12F, 0x521176C4D8AEC4C, 0x2C87AB}, {0x6F88B2A580B7137, 0x6A01C7EC3DCBF40, 0x77A61FE541BC690, 0xC21EED75E09CBBE, 0xD546C64B66EF123, 0xE9BE54149BB92F4, 0x2737E7C402B1E3A, 0x2542AA}, {0x75832C6E05BEFD8, 0x5DFA5C39B3B092F, 0xE8C16C48839C904, 0xCD2EF769EC4E8B6, 0x2C80F565FA5F23C, 0x6E9127B74DF1A18, 0xF8539726D23B186, 0x160D2}, {0x419B4088ACA5AD9, 0x7306B8707C1ADC1, 0xACD61FFAF6099A0, 0xC51DEF3D128F489, 0x81B7085E8A76F4E, 0x2D3B39889E0FB08, 0xD61F1EA307C7E55, 0x4BF3DF}, {0xA8371AA52C72185, 0x8EB49A55D2D9E38, 0xC6C6736817932E9, 0x9DBF37057564284, 0x443A1CBA5B94394, 0xD87AF590E234AA5, 0x7D1FCAC035D89C, 0x2FC680}, {0xFEA149BAEFEE0EA, 0x3CF4640E446D388, 0x1E4CE07FE5EDC2C, 0x5F90A3F48CF058E, 0x9F19E019EED6342, 0xF4E2F3EB5891E98, 0x7D7EECC281EEE08, 0x4B2EE5}, {0x3F500EED6887410, 0x10F5106924F0215, 0x5AF361F0BCBCC51, 0xBA469C2DA94E589, 0xD831D9B3D39CDF1, 0xCB5839282F3E02D, 0x2DE045F7C285BD4, 0x3DFAA0}, {0xFEE121340DD0AA9, 0xC80E715435565F5, 0xC9BDDB418617757, 0x9008A702AC7A865, 0x7B352BC03C0F76A, 0xDB9BAD304E54DE7, 0x20262A2ADBD4C8F, 0x4B32FD}, {0x2DE43411E428BFD, 0xA41124E28EB7D81, 0xEFCAE1597CB7A3F, 0xD7659AAD7AD2D1, 0x44F775441C84B43, 0x53A8AE3DCFF23BD, 0x40645BA7F7D5EB, 0x285EEC}, {0x2727300D2540F1F, 0xFDECB15E09A345, 0x21DEA9022B44831, 0xBBB119161008E89, 0xECA756DE403AF19, 0xEED8A0E6D5E59C, 0x95A363EF5B040D0, 0x522411}, {0xA2C27EEF545BF96, 0xAA5B86F0E4960D7, 0x6AFF6F2F34CA0A0, 0x2D50D7FDB10F76C, 0xC47B4C01E25E74A, 0xCA748839AA13431, 0xB0C0E8C18582A96, 0x4E11DC}, {0xA3678E32C96D5DE, 0xF315422ACB1139C, 0x2AE44725AB3522D, 0xE862D8477DF1321, 0x8E709EED8BC73E4, 0xDAC20BC1126FA15, 0x248E008527C5F1A, 0x365008}, {0xDEAED787A433825, 0x3D4E69DE567CE3B, 0x2F8F7B0845854D8, 0xAECAE67C3C4F7CF, 0x210926FB5CF553C, 0x637AFBF8D84A492, 0x2CD7464BC2B8EBD, 0x19F380}, {0x12C16B5AF92004C, 0xB35CDE4958DDF53, 0xFE46B6BD7B008FF, 0x23AFFA28D733504, 0x4B1CFF42422BBD, 0x149A03F9D6D5E25, 0x30A3BB3DEC0E34E, 0x50A16D}, {0xCF5416E293FF538, 0xE30DE58F0F4B3AC, 0xEF46ECFBCACC622, 0xE9C2E6E58987C09, 0x63DC6C30073E6A1, 0xBFE3144AFF62008, 0x9C5A4805114B2B, 0x9F07F}, {0x56A7C6F9B349B02, 0xE2D5942C7B67A98, 0x3C554E3EEB67832, 0xE7639D982B6E742, 0xAEE771637673FA5, 0xB9542E2111157FD, 0xE709FFAF45FCE8E, 0x3DB8D6}, {0x8B181AFDE8BF419, 0x239F62BDFA32D68, 0x9508D18C21BA687, 0xCAC0B6357E22FA9, 0x2FB95688793946F, 0xF7099FD787682C3, 0x773503632E5DFEA, 0x179F05}, {0x2A7CE8235EA6770, 0x31EF18A9E5281EC, 0x68AE4235B272CD, 0xB1927ED87AEFA12, 0x1E383463BBD9F0A, 0x512707DEA5F069C, 0x5DFFD3541444C2F, 0x4EB5DD}, {0x9BDB2F926C0FF64, 0x7850C5E5EAFEE09, 0xD85E57E18A2E9E9, 0x50A841EB086F8DE, 0x69650BC2A3F0EAD, 0xABC5566E8E03A32, 0xD32254356A821BA, 0x19B7C1}, {0x2B6551A5AD61CD4, 0x6FE4521DBA27B6F, 0x50ADFEABFB234A1, 0xBCD0B26F910D08, 0x9AFEB2E0FE6FD7D, 0x9371E2A9C390092, 0xEB9184881B2B3B4, 0x12D473}, {0x7897C3733196AB7, 0xF21B825648CE28F, 0xD76AD798C5D1A5B, 0x7388E3C88241F04, 0x49D7E3F8163E32A, 0x8B92841F64393C1, 0x1C2C0C4D16C17B2, 0x1C957C}, {0x3E9EB28A64A98B, 0xAE65B5A242437A3, 0xF0322B721BDBE18, 0x5F9D50449E0DC99, 0x35E8044F649E204, 0x688E7AE47885472, 0x1F13C6506336B1A, 0x14DD21}, {0x4415C43C23F15E0, 0xF980455A0894215, 0x1A238B618073F6D, 0x6D5BFF0A9566BE7, 0x9D413901AC79DF9, 0xD04752571822093, 0x584F0D545A4D1AC, 0x1FC227}, {0x1C4F5A4CDEEFD29, 0xF8A5FF8474F1D92, 0xB62F43C17E99583, 0xD84446D082FAC19, 0xC088A58ADE003FC, 0x13EDDE0B794547E, 0x7DBDF380C345BDE, 0x2517C3}, {0xD6E6050AC6AA7A5, 0xA8A1431CD988F9A, 0x4D67940C0E16EF5, 0xE6681FA1B65737, 0xA03CDAA1D6CC34C, 0x981EEF164885152, 0x9ADFE5F7AE38D93, 0x293546}, {0x127C67528B3D97C, 0x5DA0CDE6396FB3B, 0x4AEAC3133994800, 0xEB317904F2AD0C6, 0x24A8084A1872E20, 0x3121DB9650F7831, 0xFA4B5873E9F0C7D, 0x12330B}, {0x95821E917E5B5CD, 0xD9EA97459349281, 0x3935F51383CAC1A, 0x4DCB99C3E56A413, 0xC08D7BD240E4F65, 0x1538AB00D38A250, 0x2758DAC36086F09, 0x17B030}, {0x79A6D86A20CBB48, 0x85251732747B32B, 0xABE02F14C6410B4, 0xAD5E4F4C6D03568, 0xC9BACE95CC25196, 0xC21B86FBE05E5A3, 0xDC7A226D57F4FA6, 0x425BA4}, {0x84103A0D20E7325, 0x6C014D3CC5E31C3, 0x517A9067D1A365, 0x5FFFF5040819E30, 0xE14DB63BC7E52B1, 0x15A0678BED4C4BE, 0x763DC12FF86D424, 0x48D2BE}, {0x64A5FC7E3BCD51F, 0xE432E954F3F184E, 0xE129D39AA351F9E, 0x62E8D1733DA85D5, 0x46AB8EE7998D6D6, 0x88A60C30C8D77F4, 0xD36A9E77F862614, 0x2C4F5E}, {0x5D7EC335F460AFF, 0x6328CF1408444B4, 0xC0245B9996301DF, 0xAC5C9DD5665483, 0x31D5641C2B06612, 0xC8AA9D9ABD4E10, 0xF8BE90010D86599, 0xB1C34}, {0x4C2725FD88381C1, 0x6FFAB752C1B48CE, 0xC2A854562C827C, 0xAA12FEBEC40B3D4, 0x1914AA8AA28F93E, 0x49BC342093BB17F, 0x55E55A9524271BE, 0x88AA9}, {0x5C79D23020F40C4, 0x3340CEDB83BCF90, 0xE82E097E6A3F7AA, 0xD9E9718667F06F8, 0x16681C12DD39A6, 0x3E02E7172BB6F21, 0x13D3F19EFD11728, 0x53A571}}
var CURVE_HTPC2 = [...]Chunk{0x933B490F52C1DFF, 0x832A3579E306168, 0x5C1DE6F2F44EC7F, 0xFBCE0273E65375, 0x60C4B0A24C57CDB, 0xFF3652096A927A1, 0x668298885C86BFA, 0xD802F}
var CURVE_Adr = [...]Chunk{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Adi = [...]Chunk{0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Bdr = [...]Chunk{0x3F4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Bdi = [...]Chunk{0x3F4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var PCR = [13][8]Chunk{{0x55097B42B425ED1, 0x2F444BDA2077B55, 0xDBAAEB8E673D0AA, 0xF5F81D9B80EDAB, 0x9929D0A4C59D6C7, 0x7B1E8A13658EAD2, 0xE09E42523EB2C9C, 0x4DA59F}, {0xFFC71C72071C71E, 0xA37338E39859C7F, 0xE4C030AACD6DC7F, 0x4B87A1634A0B240, 0xF2DF5C7B9436115, 0x5C56E78E8C2B01D, 0xE876B1BDAF06175, 0x3A3C37}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x55425ED0AD097D6, 0x8BD112F6881DED5, 0xF6EABAE399CF42A, 0xC3D7E0766E03B6A, 0xA64A742931675B1, 0x1EC7A284D963AB4, 0xF82790948FACB27, 0x136967}, {0xC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x7174F032EDD3C0D, 0xB334A781A32C59C, 0x834E7F2F9A8220D, 0x50A7775687656CF, 0x25ED4CB280BFE02, 0xDAB66AEA58AFE64, 0xDE00C7D91941AFE, 0x541E17}, {0xFFE38E39038E38F, 0xD1B99C71CC2CE3F, 0x7260185566B6E3F, 0xA5C3D0B1A505920, 0xF96FAE3DCA1B08A, 0xAE2B73C7461580E, 0xF43B58DED7830BA, 0x1D1E1B}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0xC6D3C0CA974F017, 0x6D4C1E065F1D631, 0x7D921BDCD9A7F8, 0x6EBB871B506391A, 0xD2C8129DE80C329, 0x4B5299A7EBFE109, 0xE1ECFF8ED16B56B, 0x4A6963}, {0x12, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0xFFAAAAAB0AAA8FB, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}}
var PCI = [13][8]Chunk{{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0xFFE38E39038E38D, 0xD1B99C71CC2CE3F, 0x7260185566B6E3F, 0xA5C3D0B1A505920, 0xF96FAE3DCA1B08A, 0xAE2B73C7461580E, 0xF43B58DED7830BA, 0x1D1E1B}, {0xFFC71C72071C71A, 0xA37338E39859C7F, 0xE4C030AACD6DC7F, 0x4B87A1634A0B240, 0xF2DF5C7B9436115, 0x5C56E78E8C2B01D, 0xE876B1BDAF06175, 0x3A3C37}, {0x55425ED0AD097D6, 0x8BD112F6881DED5, 0xF6EABAE399CF42A, 0xC3D7E0766E03B6A, 0xA64A742931675B1, 0x1EC7A284D963AB4, 0xF82790948FACB27, 0x136967}, {0xFFAAAAAB0AAAA9F, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}, {0xFFAAAAAB0AAAA63, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0xFFC71C72071C71C, 0xA37338E39859C7F, 0xE4C030AACD6DC7F, 0x4B87A1634A0B240, 0xF2DF5C7B9436115, 0x5C56E78E8C2B01D, 0xE876B1BDAF06175, 0x3A3C37}, {0xAA684BDA5DA12ED, 0xE95BC25EDC68BEA, 0x60358E1C9A55694, 0x2D73919E810CFF6, 0x460496902CE9BEE, 0xEBBAB8D0F8DCD78, 0xE48A7A07F6DC708, 0x43F0EB}, {0xC6D3C0CA974F017, 0x6D4C1E065F1D631, 0x7D921BDCD9A7F8, 0x6EBB871B506391A, 0xD2C8129DE80C329, 0x4B5299A7EBFE109, 0xE1ECFF8ED16B56B, 0x4A6963}, {0xFFAAAAAB0AAAA99, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}, {0xFFAAAAAB0AAA9D3, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}, {0xFFAAAAAB0AAA8FB, 0x752CD5556486ABF, 0x572049003424ABF, 0xF14B7214EF10B61, 0xEC4F0AB95E5119F, 0xA825B55D24082C, 0xDCB20A9C8689230, 0x575A53}}
