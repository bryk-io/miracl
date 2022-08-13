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

package BLS12381

// Base Bits= 58

var Modulus = [...]Chunk{0x1FEFFFFFFFFAAAB, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}
var ROI = [...]Chunk{0x1FEFFFFFFFFAAAA, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}
var R2modp = [...]Chunk{0x20639A1D5BEF7AE, 0x1244C6462DD93E8, 0x22D09B54E6E2CD2, 0x111C4B63170E5DB, 0x38A6DE8FB366399, 0x4F16CFED1F9CBC, 0x19EA66A2B}
var SQRTm3 = [...]Chunk{0x1FB00000001AAAE, 0x313F5FB4FFFFED7, 0x2928BFC912627, 0x1D87D988BA6AF26, 0x2845E1033EFA3BF, 0x25FF9A6633A3655, 0x1A0111EA3}

const MConst Chunk = 0x1F3FFFCFFFCFFFD

const CURVE_Cof_I int = 0
const CURVE_B_I int = 4

var CURVE_B = [...]Chunk{0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Order = [...]Chunk{0x3FFFFFF00000001, 0x36900BFFF96FFBF, 0x180809A1D80553B, 0x14CA675F520CCE7, 0x73EDA7, 0x0, 0x0}
var CURVE_Gx = [...]Chunk{0x33AF00ADB22C6BB, 0x17A0FFE5E86BBFE, 0x3A3F171BAC586C5, 0x13E5DD2E4168538, 0x4FA9AC0FC3688C, 0x65F5E509A558E3, 0x17F1D3A73}
var CURVE_Gy = [...]Chunk{0xAA232946C5E7E1, 0x331D128A222B903, 0x18CB2C04B3EDD03, 0x25757402BD8036C, 0x1741D8AE4FCF5E0, 0xEAA83C68278C3B, 0x8B3F481E}
var CURVE_HTPC = [...]Chunk{0xC51062BDE821B8, 0x1A5483B9715FEDF, 0x1BDD403FC31088B, 0x3D2523427FC11BB, 0x1A3D71BDA12F01D, 0x2DB2FDD36CE3D2A, 0x1F7462C8}

var Fra = [...]Chunk{0x10775ED92235FB8, 0x3A94F58F9E04F63, 0x3D784BAB9C4F67, 0x3F4F2F57D3DEC91, 0x202C0D1F0FD603, 0xAEC199F08C6FAD, 0x1904D3BF0}
var Frb = [...]Chunk{0xF78A126DDC4AF3, 0x356B0535B1FB08B, 0xEC971F63C5F282, 0x21EDB1ECDBFB032, 0x2231F9FB854A147, 0x1B1380CA23A7A40, 0xFC3E2B3}
var CURVE_Bnx = [...]Chunk{0x201000000010000, 0x34, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Cof = [...]Chunk{0x201000000010001, 0x34, 0x0, 0x0, 0x0, 0x0, 0x0}

// var CURVE_Cof = [...]Chunk{0xAAAB0000AAAB, 0x3230015557855A3, 0x396, 0x0, 0x0, 0x0, 0x0}
var CRu = [...]Chunk{0x201FFFFFFFEFFFE, 0x1F604D88280008B, 0x293BE6F89688DE1, 0x1DA83DDFAB76CE, 0x3DF76CE51BA69C6, 0x17C659CB, 0x0}

var CURVE_Pxa = [...]Chunk{0x8056C8C121BDB8, 0x300C9AA016EFBF5, 0xB647AE3D1770BA, 0x353E900EC0AD144, 0x32DC51051C6E47A, 0x23C2A449820149, 0x24AA2B2F}
var CURVE_Pxb = [...]Chunk{0x1AC7D055D042B7E, 0x33C4484E51755F9, 0x21BBDC7F5049334, 0x3426482D86AD769, 0x88274F65596BD0, 0x9C67D81F6B34E8, 0x13E02B605}
var CURVE_Pya = [...]Chunk{0x193548608B82801, 0x2B2730EEB28A278, 0x1A695160D12C923, 0x2AA32F74E9DB50A, 0x2DA2E351AADFD9B, 0x9F5B8463327371, 0xCE5D5277}
var CURVE_Pyb = [...]Chunk{0x2A9075FF05F79BE, 0x1C349D73B07686A, 0x12AB572E99AB3F3, 0x1FA169D8EBC99D2, 0x2BC28B99CB3E28, 0x3A9CD330CAB34AC, 0x606C4A02}

var CURVE_Ad = [...]Chunk{0xF428082D584C1D, 0xDBE368383E5FD7, 0x181AEFD881AC989, 0x14E0FB99AA363A2, 0x2C96D4982B0EA98, 0xEE3A50CF5A4E80, 0x144698A}
var CURVE_Bd = [...]Chunk{0x1CC48E98E172BE0, 0xC8568C5B3AA974, 0x14FCEF35EF55A2, 0x3C3C93D01C282E7, 0x753EEE3B2016C1, 0x5A200C0062C4BA, 0x12E2908D1}
var PC = [53][7]Chunk{{0x1C8BA2E8BA2D229, 0x2C6E02D934E47EA, 0x3F1BC24C6B68C24, 0x1F88B20DEF08F02, 0x381EDEE3D31D79D, 0x389839C2F47A588, 0x6E08C248}, {0x267DF3F1605FB7B, 0x2DDC7E30A177B32, 0x336003B14866F69, 0x37799E1FE5B542B, 0x1D2565B0DFA7DCC, 0x27381F89CB63B02, 0x10321DA07}, {0x3241067BE390C9E, 0x242CBB700C9DE5F, 0x14BAF4BB1B7FA31, 0x200E83172659D8C, 0x15D138F22DD2ECB, 0x2F3E9F10B830DD4, 0x169B1F8E1}, {0x171986A8497E317, 0xA57CA5ADD3A55B, 0x16C928C5D1DE4FA, 0x1B39E7D55D28B16, 0x163BE990DC43B75, 0x269E3F11EE42CCD, 0x80D3CF1F}, {0xCB5618E3F0C88E, 0x1F23E323D1D6BE7, 0x62EF0F2753339B, 0x2AC9D6D36C69A0B, 0xD1117E53356DE5, 0x6AF6F8BA1D0E21, 0x17B81E770}, {0x1D7F225A139ED84, 0x944A30414BB2B7, 0x2218F9C86B2A8DA, 0x993C3E33864023, 0x38AE652BFB11586, 0x3F9134A5A8DC9B0, 0xD6ED6553}, {0x113C1C66F652983, 0x1C34B72B9CF4673, 0x2B9097E68F90A08, 0x1F76549E66E7B4E, 0x3F7A74AB5DB3CB1, 0x35CC4FFC0744806, 0x1630C3250}, {0x1154CE9AC8895D9, 0x28A1BCC079DF114, 0x2B65982FAC18985, 0x168495FECFC21BB, 0x3E4118E5499DB99, 0x667D10D990AD2C, 0xE99726A3}, {0x1B388641D9B6861, 0x1B89738C41C64F1, 0x3289F1B33083533, 0x195AA36FC97C6CC, 0x307E55412D7F5E4, 0x3F31B6DD3818274, 0x1778E7166}, {0x179F9DAC9EDCB0, 0x30F8F4A825CA7F8, 0x2501EC68E25C958, 0x1CCA5660F95A1E3, 0x1D10A9A1BCE0324, 0x25D9E3B07441231, 0xD54005DB}, {0x34EEF1B3CB83BB, 0x23CA9BCC630D5BA, 0x233C70D1E86B483, 0x16CBDAA105FD597, 0x22147A81C7C17E7, 0x250EACBC1622EAC, 0x17294ED3E}, {0x2AC1662734649B7, 0x30B57CB98B5BAB, 0x3B56CDB4E2C8561, 0x2228B5C017FC989, 0x1D99815856B303E, 0x3A0CCD02E024407, 0x11A05F2B1}, {0x16384D168ECDD0A, 0x1D392D2DE19400B, 0x133978F31C15931, 0x3BA5BDF40DDDB7D, 0x2B3A56680F682B4, 0x27A4AB511DB5B8F, 0x95FC13AB}, {0x376EC3A79A1D641, 0x99A4AAEE90DC11, 0xDA67F398835038, 0x75C584D9ADD040, 0x1AFC7A3CCE07F8D, 0x36953E097A482CF, 0xA10ECF6A}, {0x1F7D99BBDCC5A5E, 0x16E52274478B4C4, 0x21CDF9822C580FA, 0x3086F29A2A0665B, 0x74CF01996E7F63, 0x3592A2C8C2CFD6C, 0x14A7AC2A9}, {0x2574496EE84A3A, 0xECD4E3C3781B3B, 0x73062AEDE9CEA7, 0x266BD4E862538B8, 0x3E0596721570F57, 0x5A4D8643CF8318, 0x772CAACF}, {0x2DF9A29F6304A5, 0x3492F108A3C470, 0x3CEF24B8982F740, 0x3A73A72B534290E, 0x30506C6E9395735, 0x13999EE554E43DF, 0xE7355F8E}, {0x39D395B3532A21E, 0xA6EA07CD5E0754, 0x4E833B306DA9BD, 0x16684818AEE35AD, 0x343E7A07DFFDFC7, 0x8A452A029BC757, 0x13A8E1620}, {0x30DE8938DC62CD8, 0x1B5490FBB3D7104, 0x28ABC28D6FD0497, 0xFC5AC595455332, 0x37C40EB545B0824, 0x162B8BFB20EABFB, 0x3425581A}, {0xC239BA5CB83E19, 0xF4259F253FB73F, 0xE00B11ACEACD6A, 0x1BD69C63347F299, 0x1BFF2991F6F8941, 0x1E8C897A04DF98A, 0xB2962FE5}, {0x1C8276EC82B3BFF, 0x2AA211B2C09BA79, 0x2588C48BF5713D, 0x32833C20030049B, 0x298E536367041E8, 0x2D56710D22D1C44, 0x12561A5DE}, {0x13CF9FA40D21B1C, 0x235A06F8D0F7E26, 0x8617FC8AC62B55, 0x12E8D6D22EA7256, 0x34BD3FA6F01D5EF, 0x33FC66B862CB98B, 0x8CA8D548}, {0xB456BE69C8B604, 0x1409FBFB0071DC1, 0x14FA95AF01B2B66, 0x23E125968E55EB7, 0x342DF2EB5CB181D, 0x243C0F393A942CE, 0x15E6BE4E9}, {0x26B1E715475224B, 0x4126D95E6BEDE1, 0xF5D396A7CE46BA, 0x2075FA195A366AC, 0x348C4A3FC5E673D, 0x39133C440A8567D, 0x5C129645}, {0x2D9D3F5DB980133, 0x3E42B4708CA9910, 0x232D3C40659CC6C, 0x20353056004F99, 0x27BE315DC757B3B, 0x347B2A6DCBF002B, 0x245A394A}, {0x14C04F00B971EF8, 0x214706464847C83, 0x10E807B4633F06C, 0xA8D09AC23B009C, 0x4F53F447AA7B1, 0x6E4E674554258, 0xB182CAC1}, {0x207C8A4D0074D8E, 0x2737D06D13581B3, 0x3E7F911F643249D, 0x2E2ABC30918B9AF, 0x3FED2EDCC523559, 0x3CDBDB7AE463050, 0x18B46A908}, {0x13711AD011C132, 0x3CE97338FEEBF3A, 0x3E416389E61031B, 0x32DB2BD24FF4460, 0x31D43FB93CD2FCB, 0xDF346F837F42E3, 0x19713E479}, {0x3AFAAEBCA731C30, 0x3DC157753AE9BCA, 0x1E7ED1E4D43B9B3, 0x29E456BDBF81A61, 0x3ADA14A23C42A0C, 0x61AF6D488EAF79, 0xE1BBA7A1}, {0x370E577BDBA587, 0x1948071E181E8D8, 0x2E6A1F20CABE69D, 0x599E7709B07A2D, 0x21E4DA1BB8F3ABD, 0x3659A12FA232788, 0x9FC4018B}, {0x15E4CA31870FB29, 0x191543FB7FA4D68, 0xDA6C26C842642F, 0x2FF8EF7607FF40E, 0x12CA6C674170A05, 0xCEAE1BF7A649AF, 0x987C8D53}, {0x161F8855FE9D6F2, 0x21EB09183D057B2, 0x13C4D634F3747A, 0x328AF86132D48C5, 0x27796B3CE75BB8, 0x3EB06EF2CB25DF4, 0x4AB0B9BC}, {0x1B23AB13633A5F0, 0x3D8C9B256A01CA6, 0x1C3D3AD5544E203, 0x352BEB6DEF5D941, 0x1B8F0A6A074A7D0, 0x18D2DA88847847, 0x16603FCA4}, {0x1B6DAECF2E8FEDB, 0x1FE370264102A10, 0x3FD221351ADC2EE, 0x3EF8F3942E1E60C, 0x2A21529C4195536, 0x3F83FC4D72BD3F8, 0x8CC03FDE}, {0x2355C77B0E5F4CB, 0x16AEA7B1877B29, 0x23EC03251CF9DE4, 0x2E43BADE4702792, 0x2D8746757D42AA7, 0x22607085E261D46, 0x1F86376E}, {0xDFE240C72DE1F6, 0x354858A2C0148EE, 0x3E4B91400DA7D26, 0x359628C738B0D12, 0x6A3B49942552E2, 0x2A59B99BD28E132, 0xCC786BA}, {0x97E75A2E41C696, 0x159C4658BEA2FF8, 0x2343EB67AD34D6C, 0x1B0953CE0F43E41, 0x376FB46831223E9, 0x13B960475440DB5, 0x134996A10}, {0x29845719707BB33, 0x31EBBA6CEE8F0AF, 0x2F6C956543D3CD0, 0x23922A1A548AD4A, 0x14980DCFA11AD13, 0x2E893B8096747C2, 0x90D97C81}, {0x15473A1D634B8F, 0xBD5C3C4D25E011, 0x3CD6356CAA205CA, 0x19789CEE14CC93B, 0x20D7819C171C40F, 0x1B7700F9AC90957, 0xE0FA1D81}, {0x1FADC1326ED06F7, 0x145EF61C5332034, 0xDF27942480E420, 0x2539CA49F072DD2, 0x153CD76F2BF565B, 0x2CB93CED8A2F743, 0x2660400E}, {0x299B138573345CC, 0x1D8F8EE42B047, 0x2EF9A00D9B86930, 0x3662B7C0899F573, 0xB45F1496543346, 0x31D9FF8F0D84C51, 0xAD6B9514}, {0x284B529E2561092, 0x25A261BDFAEFAA5, 0x1A88CEA7913516F, 0x22BBF390B4A303E, 0x248C50C477F94FF, 0x20740CFFD614B07, 0xACCBB674}, {0xF8B49CBA8F6AA8, 0x170A7D3E0C18100, 0x1B36E636A5C871A, 0xE6ED8698A43964, 0x1AD2911D9C6DD0, 0x3A9016F523C0428, 0x4D2F259E}, {0x6EF48BB8913F55, 0x217A8F54A6CD78D, 0x192E7EA7D4FBC73, 0x18F84F61EED4C21, 0x3D94A84903216F7, 0x1C29B873AA08165, 0x167A55CDA}, {0x3233D9D55535D4A, 0x3F8BDEEE49220DA, 0x350C4BF39B4852C, 0x3931ABD6482AF15, 0x3D1D74CC4F9FB0C, 0xDB1848C686F953, 0x1866C8ED3}, {0xEE415A15812ED9, 0x3D6C020077B918, 0xFD206357132B92, 0x17BE87D3F5FFACD, 0x2BBA6FF6EE5A437, 0x38FA9FA80EF377E, 0x16A3EF08B}, {0x11A1399126A775C, 0x2A7006962C7EE4F, 0x25BC400A0051D5F, 0x3EA3433E3BD774D, 0xACE9824B5EECFD, 0x2A676CBF0EEA1CD, 0x166007C08}, {0x2C6477FAAF9B7AC, 0xE36E77EA733880, 0x187B6F0F5A6449F, 0x3195543620717B3, 0x2AC783182B70152, 0x61B6CB67EC99BA, 0x8D9E5297}, {0x239142311A5001D, 0x2C57703F4BB7B76, 0x1A0FC9DEC916A20, 0x27C3DA6EEC150BB, 0x2F8228DDCC6D19C, 0x117D0F92C033244, 0xBE0E0795}, {0xD26D98445F5416, 0xD93CB0A0A5EB6A, 0x2489E726AF41727, 0x36F76F34C3848F6, 0x389EDB4D1D115C5, 0x26394E57C8348EF, 0x16B7D2887}, {0x22538B53DBF67F2, 0x15F358DBE5BE247, 0x25DD279CD2ECA67, 0x15546B9FCC430D6, 0x16E8EB15778C485, 0x1903689DBEAAB9F, 0x58DF3306}, {0x2F6102C2E49A03D, 0x10981D8D4A78D4C, 0x356F453E01F78A, 0x3DCC71356729284, 0x43C348B885C84F, 0xE0480786832F5B, 0x1962D75C2}, {0x1479253B03663C1, 0xDA23BD83081B40, 0x232B5BE72E7A07F, 0x395E2602F9BBB0C, 0xFAD0EAE9601A6D, 0x2A7262C94860450, 0x16112C4C3}}

var CURVE_HTPC2 = [...]Chunk{0x27713A80F8492B, 0x211421FBAA68D1F, 0x361DD4CB6D9723B, 0x1B89D475CD7D27C, 0x21ECE6B49FAD53, 0x301E011E4075923, 0x52988B9}
var CURVE_Adr = [...]Chunk{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Adi = [...]Chunk{0xF0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Bdr = [...]Chunk{0x3F4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Bdi = [...]Chunk{0x3F4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var PCR = [13][7]Chunk{{0xE2AAAAAAAA5ED1, 0x238E343D9C71C62, 0x108F142B8575709, 0x39FD3A042A88B58, 0x11F5FB614CB14B4, 0x28E333EBB5B7A9A, 0x171D6541F}, {0x2A9FFFFFFFFC71E, 0xAAAA72E3555549, 0xC6B4F20A418147, 0x2B7DEB831FE6882, 0x2D787C88F984F87, 0x2EAA66F0C849BF3, 0x11560BF17}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x238AAAAAAAA97D6, 0x18E38D0F671C718, 0x423C50AE15D5C2, 0xE7F4E810AA22D6, 0x247D7ED8532C52D, 0x3A38CCFAED6DEA6, 0x5C759507}, {0xC, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x1B371C71C718B10, 0x2425E95B712F678, 0x37C69AA274524E7, 0xDE87898A1AC3A5, 0x1E3811AD0761B0F, 0x2DB3DE6FEFDC10F, 0x124C9AD43}, {0x2A9FFFFFFFFC71C, 0xAAAA72E3555549, 0xC6B4F20A418147, 0x2B7DEB831FE6882, 0x2D787C88F984F87, 0x2EAA66F0C849BF3, 0x11560BF17}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x2CFC71C71C6D706, 0x3097AFE324BDA04, 0x39D87D27E500FC8, 0x35281FD926FD510, 0x3076D11930F7DA5, 0x2AD044ED6693062, 0x1530477C7}, {0x12, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x1FEFFFFFFFFA8FB, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}}
var PCI = [13][7]Chunk{{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x354FFFFFFFFE38D, 0x255553971AAAAA4, 0x635A790520C0A3, 0x35BEF5C18FF3441, 0x36BC3E447CC27C3, 0x375533786424DF9, 0x8AB05F8B}, {0x2A9FFFFFFFFC71A, 0xAAAA72E3555549, 0xC6B4F20A418147, 0x2B7DEB831FE6882, 0x2D787C88F984F87, 0x2EAA66F0C849BF3, 0x11560BF17}, {0x238AAAAAAAA97D6, 0x18E38D0F671C718, 0x423C50AE15D5C2, 0xE7F4E810AA22D6, 0x247D7ED8532C52D, 0x3A38CCFAED6DEA6, 0x5C759507}, {0x1FEFFFFFFFFAA9F, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}, {0x1FEFFFFFFFFAA63, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x354FFFFFFFFE38F, 0x255553971AAAAA4, 0x635A790520C0A3, 0x35BEF5C18FF3441, 0x36BC3E447CC27C3, 0x375533786424DF9, 0x8AB05F8B}, {0x238AAAAAAAA97BE, 0x18E38D0F671C718, 0x423C50AE15D5C2, 0xE7F4E810AA22D6, 0x247D7ED8532C52D, 0x3A38CCFAED6DEA6, 0x5C759507}, {0x2CFC71C71C6D706, 0x3097AFE324BDA04, 0x39D87D27E500FC8, 0x35281FD926FD510, 0x3076D11930F7DA5, 0x2AD044ED6693062, 0x1530477C7}, {0x1FEFFFFFFFFAA99, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}, {0x1FEFFFFFFFFA9D3, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}, {0x1FEFFFFFFFFA8FB, 0x2FFFFAC54FFFFEE, 0x12A0F6B0F6241EA, 0x213CE144AFD9CC3, 0x2434BACD764774B, 0x25FF9A692C6E9ED, 0x1A0111EA3}}
