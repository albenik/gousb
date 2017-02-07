// Copyright 2013 Google Inc.  All rights reserved.
// Copyright 2016 the gousb Authors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package usbid

import "time"

// LastUpdate stores the latest time that the library was updated.
//
// The baked-in data was last generated:
//   2016-04-03 20:03:19.628811942 -0700 PDT
var LastUpdate = time.Unix(0, 1459738999628811942)

const usbIdListData = `#
#	List of USB ID's
#
#	Maintained by Stephen J. Gowdy <linux.usb.ids@gmail.com>
#	If you have any new entries, please submit them via
#		http://www.linux-usb.org/usb-ids.html
#	or send entries as patches (diff -u old new) in the
#	body of your email (a bot will attempt to deal with it).
#	The latest version can be obtained from
#		http://www.linux-usb.org/usb.ids
#
# Version: 2016.03.03
# Date:    2016-03-03 20:34:05
#

# Vendors, devices and interfaces. Please keep sorted.

# Syntax:
# vendor  vendor_name
#	device  device_name				<-- single tab
#		interface  interface_name		<-- two tabs

0001  Fry's Electronics
	7778  Counterfeit flash drive [Kingston]
0002  Ingram
0003  Club Mac
0004  Nebraska Furniture Mart
0011  Unknown
	7788  counterfeit flash drive
0053  Planex
	5301  GW-US54ZGL 802.11bg
0079  DragonRise Inc.
	0006  PC TWIN SHOCK Gamepad
	0011  Gamepad
0105  Trust International B.V.
	145f  NW-3100 802.11b/g 54Mbps Wireless Network Adapter [zd1211]
0127  IBP
	0002  HDM Interface
0145  Unknown
	0112  Card Reader
017c  MLK
	145f  Trust Deskset
0200  TP-Link
	0201  MA180 UMTS Modem
0204  Chipsbank Microelectronics Co., Ltd
	6025  CBM2080 / CBM2090 Flash drive controller
	6026  CBM1180 Flash drive controller
0218  Hangzhou Worlde
	0301  MIDI Port
02ad  HUMAX Co., Ltd.
	138c  PVR Mass Storage
0300  MM300 eBook Reader
0324  OCZ Technology Inc
	bc06  OCZ ATV USB 2.0 Flash Drive
	bc08  OCZ Rally2/ATV USB 2.0 Flash Drive
0325  OCZ Technology Inc
	ac02  ATV Turbo / Rally2 Dual Channel USB 2.0 Flash Drive
0386  LTS
	0001  PSX for USB Converter
03d9  Shenzhen Sinote Tech-Electron Co., Ltd
	0499  SE340D PC Remote Control
03da  Bernd Walter Computer Technology
	0002  HD44780 LCD interface
03e8  EndPoints, Inc.
	0004  SE401 Webcam
	0008  101 Ethernet [klsi]
	0015  ATAPI Enclosure
	2123  SiPix StyleCam Deluxe
	8004  Aox 99001
03e9  Thesys Microelectronics
03ea  Data Broadcasting Corp.
03eb  Atmel Corp.
	0902  4-Port Hub
	2002  Mass Storage Device
	2015  at90usbkey sample firmware (HID keyboard)
	2018  at90usbkey sample firmware (CDC ACM)
	2019  stk525 sample firmware (microphone)
	201c  at90usbkey sample firmware (HID mouse)
	201d  at90usbkey sample firmware (HID generic)
	2022  at90usbkey sample firmware (composite device)
	2040  LUFA Test PID
	2041  LUFA Mouse Demo Application
	2042  LUFA Keyboard Demo Application
	2043  LUFA Joystick Demo Application
	2044  LUFA CDC Demo Application
	2045  LUFA Mass Storage Demo Application
	2046  LUFA Audio Output Demo Application
	2047  LUFA Audio Input Demo Application
	2048  LUFA MIDI Demo Application
	2049  Stripe Snoop Magnetic Stripe Reader
	204a  LUFA CDC Class Bootloader
	204b  LUFA USB to Serial Adapter Project
	204c  LUFA RNDIS Demo Application
	204d  LUFA Combined Mouse and Keyboard Demo Application
	204e  LUFA Dual CDC Demo Application
	204f  LUFA Generic HID Demo Application
	2060  Benito Programmer Project
	2061  LUFA Combined Mass Storage and Keyboard Demo Application
	2062  LUFA Combined CDC and Mouse Demo Application
	2063  LUFA Datalogger Device
	2064  Interfaceless Control-Only LUFA Devices
	2065  LUFA Test and Measurement Demo Application
	2066  LUFA Multiple Report HID Demo
	2068  LUFA Virtual Serial/Mass Storage Demo
	2069  LUFA Webserver Project
	2103  JTAG ICE mkII
	2104  AVR ISP mkII
	2105  AVRONE!
	2106  STK600 development board
	2107  AVR Dragon
	2109  STK541 ZigBee Development Board
	210d  XPLAIN evaluation kit (CDC ACM)
	2110  AVR JTAGICE3 Debugger and Programmer
	2122  XMEGA-A1 Explained evaluation kit
	2141  ICE debugger
	2310  EVK11xx evaluation board
	2fe4  ATxmega32A4U DFU bootloader
	2fe6  Cactus V6 (DFU)
	2fea  Cactus RF60 (DFU)
	2ff0  atmega32u2 DFU bootloader
	2ff4  atmega32u4 DFU bootloader
	2ffa  at90usb162 DFU bootloader
	2ffb  at90usb AVR DFU bootloader
	2ffd  at89c5130/c5131 DFU bootloader
	2fff  at89c5132/c51snd1c DFU bootloader
	3301  at43301 4-Port Hub
	3312  4-Port Hub
	4102  AirVast W-Buddie WN210
	5601  at76c510 Prism-II 802.11b Access Point
	5603  Cisco 7920 WiFi IP Phone
	6119  AT91SAM CDC Demo Application
	6124  at91sam SAMBA bootloader
	6127  AT91SAM HID Keyboard Demo Application
	6129  AT91SAM Mass Storage Demo Application
	6200  AT91SAM HID Mouse Demo Application
	7603  D-Link DWL-120 802.11b Wireless Adapter [Atmel at76c503a]
	7604  at76c503a 802.11b Adapter
	7605  at76c503a 802.11b Adapter
	7606  at76c505 802.11b Adapter
	7611  at76c510 rfmd2948 802.11b Access Point
	7613  WL-1130 USB
	7614  AT76c505a Wireless Adapter
	7615  AT76C505AMX Wireless Adapter
	7617  AT76C505AS Wireless Adapter
	7800  Mini Album
	ff07  Tux Droid fish dongle
03ec  Iwatsu America, Inc.
03ed  Mitel Corp.
03ee  Mitsumi
	0000  CD-R/RW Drive
	2501  eHome Infrared Receiver
	2502  eHome Infrared Receiver
	5609  Japanese Keyboard
	641f  WIF-0402C Bluetooth Adapter
	6438  Bluetooth Device
	6440  WML-C52APR Bluetooth Adapter
	6901  SmartDisk FDD
	6902  Floppy Disk Drive
	7500  CD-R/RW
	ffff  Dongle with BlueCore in DFU mode
03f0  Hewlett-Packard
	0004  DeskJet 895c
	0011  OfficeJet G55
	0012  DeskJet 1125C Printer Port
	0024  KU-0316 Keyboard
	002a  LaserJet P1102
	0101  ScanJet 4100c
	0102  PhotoSmart S20
	0104  DeskJet 880c/970c
	0105  ScanJet 4200c
	0107  CD-Writer Plus
	010c  Multimedia Keyboard Hub
	0111  G55xi Printer/Scanner/Copier
	0117  LaserJet 3200
	011c  hn210w 802.11b Adapter
	011d  Bluetooth 1.2 Interface [Broadcom BCM2035]
	0121  HP 39g+ [F2224A], 39gs [F2223A], 40gs [F2225A], 48gII [F2226A], 49g+ [F2228A], 50g [F2229A, NW240AA]
	0122  HID Internet Keyboard
	0125  DAT72 Tape
	0139  Barcode Scanner 4430
	0201  ScanJet 6200c
	0202  PhotoSmart S20
	0204  DeskJet 815c
	0205  ScanJet 3300c
	0207  CD-Writer Plus 8200e
	020c  Multimedia Keyboard
	0211  OfficeJet G85
	0212  DeskJet 1220C
	0217  LaserJet 2200
	0218  APOLLO P2500/2600
	0221  StreamSmart 400 [F2235AA]
	022a  Laserjet CP1525nw
	0241  Link-5 micro dongle
	0304  DeskJet 810c/812c
	0305  ScanJet 4300c
	0307  CD-Writer+ CD-4e
	0311  OfficeJet G85xi
	0312  Color Inkjet CP1700
	0314  designjet 30/130 series
	0317  LaserJet 1200
	0324  SK-2885 keyboard
	034a  Elite Keyboard
	0401  ScanJet 5200c
	0404  DeskJet 830c/832c
	0405  ScanJet 3400cse
	0411  OfficeJet G95
	0412  Printing Support
	0417  LaserJet 1200 series
	0423  HS-COMBO Cardreader
	042a  LaserJet M1132 MFP
	0441  Prime [NW280AA, G8X92AA]
	0504  DeskJet 885c
	0505  ScanJet 2100c
	0507  DVD+RW
	050c  5219 Wireless Keyboard
	0511  OfficeJet K60
	0512  DeckJet 450
	0517  LaserJet 1000
	051d  Bluetooth Interface
	0601  ScanJet 6300c
	0604  DeskJet 840c
	0605  ScanJet 2200c
	0611  OfficeJet K60xi
	0612  business inkjet 3000
	0624  Bluetooth Dongle
	0701  ScanJet 5300c/5370c
	0704  DeskJet 825c
	0705  ScanJet 4400c
	070c  Personal Media Drive
	0711  OfficeJet K80
	0712  DeskJet 1180c
	0714  Printing Support
	0741  Prime Wireless Kit [FOK65AA]
	0801  ScanJet 7400c
	0804  DeskJet 816c
	0805  HP4470C
	0811  OfficeJet K80xi
	0817  LaserJet 3300
	0901  ScanJet 2300c
	0904  DeskJet 845c
	0912  Printing Support
	0917  LaserJet 3330
	0924  Modular Smartcard Keyboard
	094a  Optical Mouse [672662-001]
	0a01  ScanJet 2400c
	0a17  color LaserJet 3700
	0b01  ScanJet 82x0C
	0b0c  Wireless Keyboard and Optical Mouse receiver
	0b17  LaserJet 2300d
	0c17  LaserJet 1010
	0c24  Bluetooth Dongle
	0d12  OfficeJet 9100 series
	0d17  LaserJet 1012
	0e17  LaserJet 1015
	0f0c  Wireless Keyboard and Optical Mouse receiver
	0f11  OfficeJet V40
	0f12  Printing Support
	0f17  LaserJet 1150
	0f2a  LaserJet 400 color M451dn
	1001  Photo Scanner 1000
	1002  PhotoSmart 140 series
	1004  DeskJet 970c/970cse
	1005  ScanJet 5400c
	1011  OfficeJet V40xi
	1016  Jornada 548 / iPAQ HW6515 Pocket PC
	1017  LaserJet 1300
	1024  Smart Card Keyboard
	1027  Virtual keyboard and mouse
	1102  PhotoSmart 240 series
	1104  DeskJet 959c
	1105  ScanJet 5470c/5490c
	1111  OfficeJet v60
	1116  Jornada 568 Pocket PC
	1117  LaserJet 1300n
	1151  PSC-750xi Printer/Scanner/Copier
	1198  HID-compliant mouse
	1202  PhotoSmart 320 series
	1204  DeskJet 930c
	1205  ScanJet 4500C/5550C
	1211  OfficeJet v60xi
	1217  LaserJet 2300L
	1227  Virtual CD-ROM
	1302  PhotoSmart 370 series
	1305  ScanJet 4570c
	1311  OfficeJet V30
	1312  DeskJet 460
	1317  LaserJet 1005
	1327  iLO Virtual Hub
	134a  Optical Mouse
	1405  ScanJet 3670
	1411  PSC 750
	1424  f2105 Monitor Hub
	1502  PhotoSmart 420 series
	1504  DeskJet 920c
	150c  Mood Lighting (Microchip Technology Inc.)
	1511  PSC 750xi
	1512  Printing Support
	1517  color LaserJet 3500
	1524  Smart Card Keyboard - KR
	1539  Mini Magnetic Stripe Reader
	1541  Prime [G8X92AA]
	1602  PhotoSmart 330 series
	1604  DeskJet 940c
	1605  ScanJet 5530C PhotoSmart
	1611  psc 780
	1617  LaserJet 3015
	161d  Wireless Rechargeable Optical Mouse (HID)
	1624  Smart Card Keyboard - JP
	1702  PhotoSmart 380 series
	1704  DeskJet 948C
	1705  ScanJet 5590
	1711  psc 780xi
	1712  Printing Support
	1717  LaserJet 3020
	171d  Bluetooth 2.0 Interface [Broadcom BCM2045]
	1801  Inkjet P-2000U
	1802  PhotoSmart 470 series
	1804  DeskJet 916C
	1805  ScanJet 7650
	1811  PSC 720
	1812  OfficeJet Pro K550
	1817  LaserJet 3030
	181d  Bluetooth 2.0 Interface
	1902  PhotoSmart A430 series
	1904  DeskJet 3820
	1911  OfficeJet V45
	1917  LaserJet 3380
	1a02  PhotoSmart A510 series
	1a11  OfficeJet 5100 series
	1a17  color LaserJet 4650
	1b02  PhotoSmart A610 series
	1b04  DeskJet 3810
	1b05  ScanJet 4850C/4890C
	1b07  Premium Starter Webcam
	1c02  PhotoSmart A710 series
	1c17  Color LaserJet 2550l
	1d02  PhotoSmart A310 series
	1d17  LaserJet 1320
	1d24  Barcode scanner
	1e02  PhotoSmart A320 Printer series
	1e11  PSC-950
	1e17  LaserJet 1160 series
	1f02  PhotoSmart A440 Printer series
	1f11  PSC 920
	1f12  OfficeJet Pro K5300
	1f17  color LaserJet 5550
	1f1d  un2400 Gobi Wireless Modem
	2001  Floppy
	2002  Hub
	2004  DeskJet 640c
	2005  ScanJet 3570c
	2012  OfficeJet Pro K5400
	201d  un2400 Gobi Wireless Modem (QDL mode)
	2039  Cashdrawer
	2102  PhotoSmart 7345
	2104  DeskJet 630c
	2112  OfficeJet Pro L7500
	211d  Sierra MC5725 [ev2210]
	2202  PhotoSmart 7600 series
	2205  ScanJet 3500c
	2212  OfficeJet Pro L7600
	2217  color LaserJet 9500 MFP
	2302  PhotoSmart 7600 series
	2304  DeskJet 656c
	2305  ScanJet 3970c
	2311  OfficeJet d series
	2312  OfficeJet Pro L7700
	2317  LaserJet 4350
	231d  Broadcom 2070 Bluetooth Combo
	2402  PhotoSmart 7700 series
	2404  Deskjet F2280 series
	2405  ScanJet 4070 PhotoSmart
	2417  LaserJet 4250
	241d  Gobi 2000 Wireless Modem (QDL mode)
	2424  LP1965 19" Monitor Hub
	2502  PhotoSmart 7700 series
	2504  DeskJet F4200 series
	2505  ScanJet 3770
	2512  OfficeJet Pro L7300 / Compaq LA2405 series monitor
	2514  4-port hub
	2517  LaserJet 2410
	251d  Gobi 2000 Wireless Modem
	2524  LP3065 30" Monitor Hub
	2602  PhotoSmart A520 series
	2605  ScanJet 3800c
	2611  OfficeJet 7100 series
	2617  Color LaserJet 2820 series
	2624  Pole Display (HP522 2 x 20 Line Display)
	2702  PhotoSmart A620 series
	2704  DeskJet 915
	2717  Color LaserJet 2830
	2724  Magnetic Stripe Reader IDRA-334133-HP
	2805  Scanjet G2710
	2811  PSC-2100
	2817  Color LaserJet 2840
	2902  PhotoSmart A820 series
	2911  PSC 2200
	2917  LaserJet 2420
	2a11  PSC 2150 series
	2a17  LaserJet 2430
	2a1d  Integrated Module with Bluetooth 2.1 Wireless technology
	2b11  PSC 2170 series
	2b17  LaserJet 1020
	2c12  Officejet J4680
	2c17  LaserJet 1022
	2c24  Logitech M-UAL-96 Mouse
	2d05  Scanjet 7000
	2d11  OfficeJet 6110
	2d17  Printing Support
	2e11  PSC 1000
	2e17  LaserJet 2600n
	2e24  LP2275w Monitor Hub
	2f11  PSC 1200
	2f17  EWS 2605dn
	2f24  LP2475w Monitor Hub
	3002  PhotoSmart P1000
	3004  DeskJet 980c
	3005  ScanJet 4670v
	3011  PSC 1100 series
	3017  Printing Support
	3102  PhotoSmart P1100 Printer w/ Card Reader
	3104  DeskJet 960c
	3111  OfficeJet 4100 series
	3117  EWS 2605dtn
	311d  Atheros AR9285 Malbec Bluetooth Adapter
	3202  PhotoSmart 1215
	3207  4 GB flash drive
	3211  OfficeJet 4105 series
	3217  LaserJet 3050
	3302  PhotoSmart 1218
	3304  DeskJet 990c
	3312  OfficeJet J6410
	3317  LaserJet 3052
	3402  PhotoSmart 1115
	3404  DeskJet 6122
	3417  LaserJet 3055
	3502  PhotoSmart 230
	3504  DeskJet 6127c
	3511  PSC 2300
	3517  LaserJet 3390
	3602  PhotoSmart 1315
	3611  PSC 2410 PhotoSmart
	3617  Color LaserJet 2605
	3711  PSC 2500
	3717  EWS UPD
	3724  Webcam
	3802  PhotoSmart 100
	3807  c485w Flash Drive
	3817  LaserJet P2015 series
	3902  PhotoSmart 130
	3912  Officejet Pro 8500
	3a02  PhotoSmart 7150
	3a11  OfficeJet 5500 series
	3a17  Printing Support
	3b02  PhotoSmart 7150~
	3b05  Scanjet N8460
	3b11  PSC 1300 series
	3b17  LaserJet M1005 MFP
	3c02  PhotoSmart 7350
	3c05  Scanjet Professional 1000 Mobile Scanner
	3c11  PSC 1358
	3c17  EWS UPD
	3d02  PhotoSmart 7350~
	3d11  OfficeJet 4215
	3d17  LaserJet P1005
	3e02  PhotoSmart 7550
	3e17  LaserJet P1006
	3f02  PhotoSmart 7550~
	3f11  PSC-1315/PSC-1317
	4002  PhotoSmart 635/715/720/735/935/E337 (storage)
	4004  CP1160
	4102  PhotoSmart 618
	4105  ScanJet 4370
	4111  OfficeJet 7200 series
	4117  LaserJet 1018
	4202  PhotoSmart 812
	4205  ScanJet G3010
	4211  OfficeJet 7300 series
	4217  EWS CM1015
	4302  PhotoSmart 850 (ptp)
	4305  ScanJet G3110
	4311  OfficeJet 7400 series
	4317  Color LaserJet CM1017
	4402  PhotoSmart 935 (ptp)
	4417  EWS UPD
	4502  PhotoSmart 945 (PTP mode)
	4505  ScanJet G4010
	4507  External HDD
	4511  PhotoSmart 2600
	4512  E709n [Officejet 6500 Wireless]
	4517  EWS UPD
	4605  ScanJet G4050
	4611  PhotoSmart 2700
	4717  Color LaserJet CP1215
	4811  PSC 1600
	4911  PSC 2350
	4b11  OfficeJet 6200
	4c11  PSC 1500 series
	4c17  EWS UPD
	4d11  PSC 1400
	4d17  EWS UPD
	4e11  PhotoSmart 2570 series
	4f11  OfficeJet 5600 (USBHUB)
	4f17  Color LaserJet CM1312 MFP
	5004  DeskJet 995c
	5011  PhotoSmart 3100 series
	5017  EWS UPD
	5111  PhotoSmart 3200 series
	5211  PhotoSmart 3300 series
	5307  v165w Stick
	5311  OfficeJet 6300
	5312  Officejet Pro 8500A
	5411  OfficeJet 4300
	5511  DeskJet F300 series
	5611  PhotoSmart C3180
	5617  LaserJet M1120 MFP
	5711  PhotoSmart C4100 series
	5717  LaserJet M1120n MFP
	5811  PhotoSmart C5100 series
	5817  LaserJet M1319f MFP
	5911  PhotoSmart C6180
	5912  Officejet Pro 8600
	5a11  PhotoSmart C7100 series
	5b11  OfficeJet J2100 series
	5b12  Officejet Pro 8100
	5c11  PhotoSmart C4200 Printer series
	5c12  OfficeJet 6700
	5c17  LaserJet P2055 series
	5d11  PhotoSmart C5200 series
	5e11  PhotoSmart D7400 series
	6004  DeskJet 5550
	6102  Hewlett Packard Digital Camera
	6104  DeskJet 5650c
	6117  color LaserJet 3550
	6202  PhotoSmart 215
	6204  DeskJet 5150c
	6217  Color LaserJet 4700
	6302  PhotoSmart 318/612
	6317  Color LaserJet 4730mfp
	6402  PhotoSmart 715 (ptp)
	6411  PhotoSmart C8100 series
	6417  LaserJet 5200
	6502  PhotoSmart 120 (ptp)
	6511  PhotoSmart C7200 series
	6602  PhotoSmart 320
	6611  PhotoSmart C4380 series
	6617  LaserJet 5200L
	6702  PhotoSmart 720 (ptp)
	6717  Color LaserJet 3000
	6802  PhotoSmart 620 (ptp)
	6811  PhotoSmart D5300 series
	6817  Color LaserJet 3800
	6911  PhotoSmart D7200 series
	6917  Color LaserJet 3600
	6a02  PhotoSmart 735 (ptp)
	6a11  PhotoSmart C6200 series
	6a17  LaserJet 4240
	6b02  PhotoSmart R707 (PTP mode)
	6b11  Photosmart C4500 series
	6c11  Photosmart C4480
	6c17  Color LaserJet 4610
	6f17  Color LaserJet CP6015 series
	7004  DeskJet 3320c
	7102  PhotoSmart 635 (PTP mode)
	7104  DeskJet 3420c
	7117  CM8060 Color MFP with Edgeline Technology
	7202  PhotoSmart 43x (ptp)
	7204  DeskJet 36xx
	7217  LaserJet M5035 MFP
	7302  PhotoSmart M307 (PTP mode)
	7304  DeskJet 35xx
	7311  Photosmart Premium C309
	7317  LaserJet P3005
	7404  Printing Support
	7417  LaserJet M4345 MFP
	7504  Printing Support
	7517  LaserJet M3035 MFP
	7604  DeskJet 3940
	7611  DeskJet F2492 All-in-One
	7617  LaserJet P3004
	7702  PhotoSmart R817 (PTP mode)
	7704  DeskJet D4100
	7717  CM8050 Color MFP with Edgeline Technology
	7804  DeskJet D1360
	7817  Color LaserJet CP3505
	7917  LaserJet M5025 MFP
	7a02  PhotoSmart M415 (PTP mode)
	7a04  DeskJet D2460
	7a17  LaserJet M3027 MFP
	7b02  PhotoSmart M23 (PTP mode)
	7b17  Color LaserJet CP4005
	7c17  Color LaserJet CM6040 series
	7d04  DeskJet F2100 Printer series
	7d17  Color LaserJet CM4730 MFP
	7e04  DeskJet F4100 Printer series
	8017  LaserJet P4515
	8104  Printing Support
	8117  LaserJet P4015
	811c  Ethernet HN210E
	8204  Printing Support
	8207  FHA-3510 2.4GHz Wireless Optical Mobile Mouse
	8217  LaserJet P4014
	8317  LaserJet M9050 MFP
	8404  DeskJet 6800 series
	8417  LaserJet M9040 MFP
	8504  DeskJet 6600 series
	8604  DeskJet 5440
	8607  Optical Mobile Mouse
	8704  DeskJet 5940
	8711  Deskjet 2050 J510
	8804  DeskJet 6980 series
	8904  DeskJet 6940 series
	8c07  Digital Stereo Headset
	8c11  Deskjet F4500 series
	9002  PhotoSmart M437
	9102  PhotoSmart M537
	9302  PhotoSmart R930 series
	9402  PhotoSmart R837
	9502  PhotoSmart R840 series
	9602  PhotoSmart M730 series
	9702  PhotoSmart R740 series
	9802  PhotoSmart Mz60 series
	9902  PhotoSmart M630 series
	9a02  PhotoSmart E330 series
	9b02  PhotoSmart M540 series
	9b07  Portable Drive
	9c02  PhotoSmart M440 series
	a004  DeskJet 5850c
	a011  Deskjet 3050A
	b002  PhotoSmart 7200 series
	b102  PhotoSmart 7200 series
	b107  v255w/c310w Flash Drive
	b116  Webcam
	b202  PhotoSmart 7600 series
	b302  PhotoSmart 7600 series
	b402  PhotoSmart 7700 series
	b502  PhotoSmart 7700 series
	b602  PhotoSmart 7900 series
	b702  PhotoSmart 7900 series
	b802  PhotoSmart 7400 series
	b902  PhotoSmart 7800 series
	ba02  PhotoSmart 8100 series
	bb02  PhotoSmart 8400 series
	bc02  PhotoSmart 8700 series
	bd02  PhotoSmart Pro B9100 series
	bef4  NEC Picty760
	c002  PhotoSmart 7800 series
	c102  PhotoSmart 8000 series
	c111  Deskjet 1510
	c202  PhotoSmart 8200 series
	c302  DeskJet D2300
	c402  PhotoSmart D5100 series
	c502  PhotoSmart D6100 series
	c602  PhotoSmart D7100 series
	c702  PhotoSmart D7300 series
	c802  PhotoSmart D5060 Printer
	d104  Bluetooth Dongle
	d507  39gII [NW249AA]
	efbe  NEC Picty900
	f0be  NEC Picty920
	f1be  NEC Picty800
03f1  Genoa Technology
03f2  Oak Technology, Inc.
03f3  Adaptec, Inc.
	0020  AWN-8020 WLAN [Intersil PRISM 2.5]
	0080  AVC-1100 Audio Capture
	0083  AVC-2200 Device
	0087  AVC-2210 Loader
	0088  AVC-2210 Device
	008b  AVC-2310 Loader
	008c  AVC-2310 Device
	0094  eHome Infrared Receiver
	009b  AVC-1410 GameBridge TV NTSC
	2000  USBXchange
	2001  USBXchange Adapter
	2002  USB2-Xchange
	2003  USB2-Xchange Adapter
	4000  4-port hub
	adcc  Composite Device Support
03f4  Diebold, Inc.
03f5  Siemens Electromechanical
03f8  Epson Imaging Technology Center
03f9  KeyTronic Corp.
	0100  KT-2001 Keyboard
	0101  Keyboard
	0102  Keyboard Mouse
03fb  OPTi, Inc.
03fc  Elitegroup Computer Systems
03fd  Xilinx, Inc.
	0008  Platform Cable USB II
03fe  Farallon Comunications
0400  National Semiconductor Corp.
	05dc  Rigol Technologies DS1000USB Oscilloscope
	0807  Bluetooth Dongle
	080a  Bluetooth Device
	09c4  Rigol Technologies DG1022 Arbitrary Waveform Generator
	1000  Mustek BearPaw 1200 Scanner
	1001  Mustek BearPaw 2400 Scanner
	1237  Hub
	a000  Smart Display Reference Device
	c359  Logitech Harmony
	c35b  Printing Support
	c55d  Rigol Technologies DS5000USB Oscilloscope
0401  National Registry, Inc.
0402  ALi Corp.
	5462  M5462 IDE Controller
	5602  M5602 Video Camera Controller
	5603  M5603 Video Camera Controller
	5606  M5606 Video Camera Controller [UVC]
	5621  M5621 High-Speed IDE Controller
	5623  M5623 Scanner Controller
	5627  Welland ME-740PS USB2 3.5" Power Saving Enclosure
	5632  M5632 Host-to-Host Link
	5635  M5635 Flash Card Reader
	5636  USB 2.0 Storage Device
	5637  M5637 IDE Controller
	5642  Storage Device
	5661  M5661 MP3 player
	5667  M5667 MP3 player
	9665  Gateway Webcam
0403  Future Technology Devices International, Ltd
	0000  H4SMK 7 Port Hub / Bricked Counterfeit FT232 Serial (UART) IC
	0232  Serial Converter
	1060  JTAG adapter
	1234  IronLogic RFID Adapter [Z-2 USB]
	1235  Iron Logic Z-397 RS-485/422 converter
	6001  FT232 Serial (UART) IC
	6002  Lumel PD12
	6007  Serial Converter
	6008  Serial Converter
	6009  Serial Converter
	6010  FT2232C Dual USB-UART/FIFO IC
	6011  FT4232H Quad HS USB-UART/FIFO IC
	6014  FT232H Single HS USB-UART/FIFO IC
	6015  Bridge(I2C/SPI/UART/FIFO)
	8028  Dev board JTAG (FT232H based)
	8040  4 Port Hub
	8070  7 Port Hub
	8140  Vehicle Explorer Interface
	8210  MGTimer - MGCC (Vic) Timing System
	8370  7 Port Hub
	8371  PS/2 Keyboard And Mouse
	8372  FT8U100AX Serial Port
	8a28  Rainforest Automation ZigBee Controller
	8a98  TIAO Multi-Protocol Adapter
	8b28  Alpermann+Velte TCI70
	8b29  Alpermann+Velte TC60 CLS
	8b2a  Alpermann+Velte Rubidium Q1
	8b2b  Alpermann+Velte TCD
	8b2c  Alpermann+Velte TCC70
	9090  SNAP Stick 200
	9132  LCD and Temperature Interface
	9133  CallerID
	9135  Rotary Pub alarm
	9136  Pulsecounter
	9e90  Marvell OpenRD Base/Client
	9f80  Ewert Energy Systems CANdapter
	a6d0  Texas Instruments XDS100v2 JTAG / BeagleBone A3
	a951  HCP HIT GSM/GPRS modem [Cinterion MC55i]
	a9a0  FT2232D - Dual UART/FIFO IC - FTDI
	abb8  Lego Mindstorms NXTCam
	b810  US Interface Navigator (CAT and 2nd PTT lines)
	b811  US Interface Navigator (WKEY and FSK lines)
	b812  US Interface Navigator (RS232 and CONFIG lines)
	b9b0  Fujitsu SK-16FX-100PMC V1.1
	baf8  Amontec JTAGkey
	bcd8  Stellaris Development Board
	bcd9  Stellaris Evaluation Board
	bcda  Stellaris ICDI Board
	bdc8  Egnite GmbH - JTAG/RS-232 adapter
	bfd8  OpenDCC
	bfd9  OpenDCC (Sniffer)
	bfda  OpenDCC (Throttle)
	bfdb  OpenDCC (Gateway)
	bfdc  OpenDCC (GBM)
	c630  lcd2usb interface
	c631  i2c-tiny-usb interface
	c632  xu1541 c64 floppy drive interface
	c633  TinyCrypt dongle
	c634  glcd2usb interface
	c7d0  RR-CirKits LocoBuffer-USB
	c8b8  Alpermann+Velte MTD TCU
	c8b9  Alpermann+Velte MTD TCU 1HE
	c8ba  Alpermann+Velte Rubidium H1
	c8bb  Alpermann+Velte Rubidium H3
	c8bc  Alpermann+Velte Rubidium S1
	c8bd  Alpermann+Velte Rubidium T1
	c8be  Alpermann+Velte Rubidium D1
	c8bf  Alpermann+Velte TC60 RLV
	cc48  Tactrix OpenPort 1.3 Mitsubishi
	cc49  Tactrix OpenPort 1.3 Subaru
	cc4a  Tactrix OpenPort 1.3 Universal
	cff8  Amontec JTAGkey
	d010  SCS PTC-IIusb
	d011  SCS Position-Tracker/TNC
	d012  SCS DRAGON 1
	d013  SCS DRAGON 1
	d388  Xsens converter
	d389  Xsens Wireless Receiver
	d38a  Xsens serial converter
	d38b  Xsens serial converter
	d38c  Xsens Wireless Receiver
	d38d  Xsens Awinda Station
	d38e  Xsens serial converter
	d38f  Xsens serial converter
	d491  Zolix Omni 1509 monochromator
	d578  Accesio USB-COM-4SM
	d6f8  UNI Black BOX
	d738  Propox JTAGcable II
	d739  Propox ISPcable III
	d9a9  Actisense USG-1 NMEA Serial Gateway
	d9aa  Actisense NGT-1 NMEA2000 PC Interface
	d9ab  Actisense NGT-1 NMEA2000 Gateway
	daf4  Qundis Serial Infrared Head
	e0d0  Total Phase Aardvark I2C/SPI Host Adapter
	e521  EVER Sinline XL Series UPS
	e6c8  PYRAMID Computer GmbH LCD
	e700  Elster Unicom III Optical Probe
	e729  Segway Robotic Mobility Platforms 200
	e888  Expert ISDN Control USB
	e889  USB-RS232 OptoBridge
	e88a  Expert mouseCLOCK USB II
	e88b  Precision Clock MSF USB
	e88c  Expert mouseCLOCK USB II HBG
	e8d8  Aaronia AG Spectran Spectrum Analyzer
	e8dc  Aaronia AG UBBV Preamplifier
	ea90  Eclo 1-Wire Adapter
	ecd9  miControl miCan-Stick
	ed71  HAMEG HO870 Serial Port
	ed72  HAMEG HO720 Serial Port
	ed73  HAMEG HO730 Serial Port
	ed74  HAMEG HO820 Serial Port
	ef10  FT1245BL
	f070  Serial Converter 422/485 [Vardaan VEUSB422R3]
	f0c8  SPROG Decoder Programmer
	f0c9  SPROG-DCC CAN-USB
	f0e9  Tagsys L-P101
	f1a0  Asix PRESTO Programmer
	f208  Papenmeier Braille-Display
	f3c0  4N-GALAXY Serial Converter
	f608  CTI USB-485-Mini
	f60b  CTI USB-Nano-485
	f680  Suunto Sports Instrument
	f758  GW Instek GDS-8x0 Oscilloscope
	f7c0  ZeitControl Cardsystems TagTracer MIFARE
	f850  USB-UIRT (Universal Infrared Receiver+Transmitter)
	f918  Ant8 Logic Probe
	fa00  Matrix Orbital USB Serial
	fa01  Matrix Orbital MX2 or MX3
	fa02  Matrix Orbital MX4 or MX5
	fa03  Matrix Orbital VK/LK202 Family
	fa04  Matrix Orbital VK/LK204 Family
	fa20  Ross-Tech HEX-USB
	fc08  Crystalfontz CFA-632 USB LCD
	fc09  Crystalfontz CFA-634 USB LCD
	fc0b  Crystalfontz CFA-633 USB LCD
	fc0c  Crystalfontz CFA-631 USB LCD
	fc0d  Crystalfontz CFA-635 USB LCD
	fc82  SEMC DSS-20/DSS-25 SyncStation
	fd48  ShipModul MiniPlex-4xUSB NMEA Multiplexer
	fd49  ShipModul MiniPlex-4xUSB-AIS NMEA Multiplexer
	fd4b  ShipModul MiniPlex NMEA Multiplexer
	ff08  ToolHouse LoopBack Adapter
	ff18  ScienceScope Logbook ML
	ff19  Logbook Bus
	ff1a  Logbook Bus
	ff1b  Logbook Bus
	ff1c  ScienceScope Logbook LS
	ff1d  ScienceScope Logbook HS
	ff1e  Logbook Bus
	ff1f  Logbook Bus
0404  NCR Corp.
	0202  78XX Scanner
	0203  78XX Scanner - Embedded System
	0310  K590 Printer, Self-Service
	0311  7167 Printer, Receipt/Slip
	0312  7197 Printer Receipt
	0320  5932-USB Keyboard
	0321  5953-USB Dynakey
	0322  5932-USB Enhanced Keyboard
	0323  5932-USB Enhanced Keyboard, Flash-Recovery/Download
	0324  5953-USB Enhanced Dynakey
	0325  5953-USB Enhanced Dynakey Flash-Recovery/Download
	0328  K016: USB-MSR ISO 3-track MSR: POS Standard (See HID pages)
	0329  K018: USB-MSR JIS 2-Track MSR: POS Standard
	032a  K016: USB-MSR ISO 3-Track MSR: HID Keyboard Mode
	032b  K016/K018: USB-MSR Flash-Recovery/Download
0405  Synopsys, Inc.
0406  Fujitsu-ICL Computers
0407  Fujitsu Personal Systems, Inc.
0408  Quanta Computer, Inc.
	0103  FV TouchCam N1 (Audio)
	030c  HP Webcam
	03b2  HP Webcam
	1030  FV TouchCam N1 (Video)
	3000  Optical dual-touch panel
	3001  Optical Touch Screen
0409  NEC Corp.
	0011  PC98 Series Layout Keyboard Mouse
	0012  ATerm IT75DSU ISDN TA
	0014  Japanese Keyboard
	0019  109 Japanese Keyboard with Bus-Powered Hub
	001a  PC98 Series Layout Keyboard with Bus-Powered Hub
	0025  Mini Keyboard with Bus-Powered Hub
	0027  MultiSync Monitor
	002c  Clik!-USB Drive
	0034  109 Japanese Keyboard with One-touch start buttons
	003f  Wireless Keyboard with One-touch start buttons
	0040  Floppy
	004e  SuperScript 1400 Series
	004f  Wireless Keyboard with One-touch start buttons
	0050  7-port hub
	0058  HighSpeed Hub
	0059  HighSpeed Hub
	005a  HighSpeed Hub
	006a  Conceptronic USB Harddisk Box
	007d  MINICUBE2
	007e  PG-FP5 Flash Memory Programmer
	0081  SuperScript 1400 Series
	0082  SuperScript 1400 Series
	0094  Japanese Keyboard with One-touch start buttons
	0095  Japanese Keyboard
	00a9  AtermIT21L 128K Support Standard
	00aa  AtermITX72 128K Support Standard
	00ab  AtermITX62 128K Support Standard
	00ac  AtermIT42 128K Support Standard
	00ae  INSMATEV70G-MAX Standard
	00af  AtermITX70 128K Support Standard
	00b0  AtermITX80 128K Support Standard
	00b2  AtermITX80D 128K Support Standard
	00c0  Wireless Remocon
	00f7  Smart Display PK-SD10
	011d  e228 Mobile Phone
	0203  HID Audio Controls
	021d  Aterm WL54SU2 802.11g Wireless Adapter [Atheros AR5523]
	0248  Aterm PA-WL54GU
	0249  Aterm WL300NU-G
	02b4  Aterm WL300NU-AG
	02b6  Aterm WL300NU-GS 802.11n Wireless Adapter
	02bc  Computer Monitor
	0300  LifeTouch Note
	0301  LifeTouch Note (debug mode)
	55aa  Hub
	55ab  Hub [iMac/iTouch kbd]
	8010  Intellibase Hub
	8011  Intellibase Hub
	efbe  P!cty 900 [HP DJ]
	f0be  P!cty 920 [HP DJ 812c]
040a  Kodak Co.
	0001  DVC-323
	0002  DVC-325
	0100  DC-220
	0110  DC-260
	0111  DC-265
	0112  DC-290
	0120  DC-240
	0121  DC-240 (PTP firmware)
	0130  DC-280
	0131  DC-5000
	0132  DC-3400
	0140  DC-4800
	0160  DC4800
	0170  DX3900
	0200  Digital Camera
	0300  EZ-200
	0400  MC3
	0402  Digital Camera
	0403  Z7590
	0500  DX3500
	0510  DX3600
	0525  DX3215
	0530  DX3700
	0535  EasyShare CX4230 Camera
	0540  LS420
	0550  DX4900
	0555  DX4330
	0560  CX4200
	0565  CX4210
	0566  CX4300
	0567  LS753
	0568  LS443
	0569  LS663
	0570  DX6340
	0571  CX6330
	0572  DX6440
	0573  CX6230
	0574  CX6200
	0575  DX6490
	0576  DX4530
	0577  DX7630
	0578  CX7300/CX7310
	0579  CX7220
	057a  CX7330
	057b  CX7430
	057c  CX7530
	057d  DX7440
	057e  C300
	057f  DX7590
	0580  Z730
	0581  Digital Camera
	0582  Digital Camera
	0583  Digital Camera
	0584  CX6445
	0585  Digital Camera
	0586  CX7525
	0587  Digital Camera
	0588  Digital Camera
	0589  EasyShare C360
	058a  C310
	058b  Digital Camera
	058c  C330
	058d  C340
	058e  V530
	058f  V550
	0590  Digital Camera
	0591  Digital Camera
	0592  Digital Camera
	0593  Digital Camera
	0594  Digital Camera
	0595  Digital Camera
	0596  Digital Camera
	0597  Digital Camera
	0598  EASYSHARE M1033 digital camera
	0599  Digital Camera
	059a  Digital Camera
	059b  Digital Camera
	059c  Digital Camera
	059d  Digital Camera
	059e  Digital Camera
	059f  Digital Camera
	05a0  Digital Camera
	05a1  Digital Camera
	05a2  Digital Camera
	05a3  Digital Camera
	05a4  Digital Camera
	05a5  Digital Camera
	05a6  Digital Camera
	05a7  Digital Camera
	05a8  Digital Camera
	05a9  Digital Camera
	05aa  Digital Camera
	05ab  Digital Camera
	05ac  Digital Camera
	05ad  Digital Camera
	05ae  Digital Camera
	05af  Digital Camera
	05b0  Digital Camera
	05b1  Digital Camera
	05b2  Digital Camera
	05b3  EasyShare Z710 Camera
	05b4  Digital Camera
	05b5  Digital Camera
	05b6  Digital Camera
	05b7  Digital Camera
	05b8  Digital Camera
	05b9  Digital Camera
	05ba  Digital Camera
	05bb  Digital Camera
	05bc  Digital Camera
	05bd  Digital Camera
	05be  Digital Camera
	05bf  Digital Camera
	05c0  Digital Camera
	05c1  Digital Camera
	05c2  Digital Camera
	05c3  Digital Camera
	05c4  Digital Camera
	05c5  Digital Camera
	05c8  EASYSHARE Z1485 IS Digital Camera
	05d3  EasyShare M320 Camera
	05d4  EasyShare C180 Digital Camera
	1001  EasyShare SV811 Digital Picture Frame
	4000  InkJet Color Printer
	4021  Photo Printer 6800
	4022  1400 Digital Photo Printer
	402b  Photo Printer 6850
	402e  605 Photo Printer
	4034  805 Photo Printer
	404f  305 Photo Printer
	4056  ESP 7200 Series AiO
	4109  EasyShare Printer Dock Series 3
	410d  EasyShare G600 Printer Dock
	5010  Wireless Adapter
	5012  DBT-220 Bluetooth Adapter
	6001  i30
	6002  i40
	6003  i50
	6004  i60
	6005  i80
	6029  i900
	602a  i900
040b  Weltrend Semiconductor
	0a68  Func MS-3 gaming mouse [WT6573F MCU]
	6510  Weltrend Bar Code Reader
	6520  XBOX Xploder
	6533  Speed-Link Competition Pro
	6543  Manhattan Magnetic Card Strip Reader
040c  VTech Computers, Ltd
040d  VIA Technologies, Inc.
	3184  VNT VT6656 USB-802.11 Wireless LAN Adapter
	6205  USB 2.0 Card Reader
040e  MCCI
040f  Echo Speech Corp.
0411  BUFFALO INC. (formerly MelCo., Inc.)
	0001  LUA-TX Ethernet [pegasus]
	0005  LUA-TX Ethernet
	0006  WLI-USB-L11 Wireless LAN Adapter
	0009  LUA2-TX Ethernet
	000b  WLI-USB-L11G-WR Wireless LAN Adapter
	000d  WLI-USB-L11G Wireless LAN Adapter
	0012  LUA-KTX Ethernet
	0013  USB2-IDE Adapter
	0016  WLI-USB-S11 802.11b Adapter
	0018  USB2-IDE Adapter
	001c  USB-IDE Bridge: DUB-PxxG
	0027  WLI-USB-KS11G 802.11b Adapter
	002a  SMSC USB97C202 "HD-HB300V2-EU"
	003d  LUA-U2-KTX Ethernet
	0044  WLI-USB-KB11 Wireless LAN Adapter
	004b  WLI-USB-G54 802.11g Adapter [Broadcom 4320 USB]
	004d  WLI-USB-B11 Wireless LAN Adapter
	0050  WLI2-USB2-G54 Wireless LAN Adapter
	005e  WLI-U2-KG54-YB WLAN
	0065  Python2 WDM Encoder
	0066  WLI-U2-KG54 WLAN
	0067  WLI-U2-KG54-AI WLAN
	006e  LUA-U2-GT 10/100/1000 Ethernet Adapter
	0089  RUF-C/U2 Flash Drive
	008b  Nintendo Wi-Fi
	0091  WLI-U2-KAMG54 Wireless LAN Adapter
	0092  WLI-U2-KAMG54 Bootloader
	0097  WLI-U2-KG54-BB
	00a9  WLI-U2-AMG54HP Wireless LAN Adapter
	00aa  WLI-U2-AMG54HP Bootloader
	00b3  PC-OP-RS1 RemoteStation
	00bc  WLI-U2-KG125S 802.11g Adapter [Broadcom 4320 USB]
	00ca  802.11n Network Adapter
	00cb  WLI-U2-G300N 802.11n Adapter
	00d8  WLI-U2-SG54HP
	00d9  WLI-U2-G54HP
	00da  WLI-U2-KG54L 802.11bg [ZyDAS ZD1211B]
	00db  External Hard Drive HD-PF32OU2 [Buffalo Ministation]
	00e8  WLI-UC-G300N Wireless LAN Adapter [Ralink RT2870]
	0105  External Hard Drive HD-CEU2 [Drive Station]
	012c  SATA Bridge
	012e  WLI-UC-AG300N Wireless LAN Adapter
	0148  WLI-UC-G300HP Wireless LAN Adapter
	0150  WLP-UC-AG300 Wireless LAN Adapter
	0157  External Hard Drive HD-PEU2
	0158  WLI-UC-GNHP Wireless LAN Adapter
	015d  WLI-UC-GN Wireless LAN Adapter [Ralink RT3070]
	016f  WLI-UC-G301N Wireless LAN Adapter [Ralink RT3072]
	017f  Sony UWA-BR100 802.11abgn Wireless Adapter [Atheros AR7010+AR9280]
	019e  WLI-UC-GNP Wireless LAN Adapter
	01a1  MiniStation Metro
	01a2  WLI-UC-GNM Wireless LAN Adapter [Ralink RT8070]
	01dc  Ultra-Slim Portable DVD Writer (DVSM-PC58U2V)
	01de  External Hard Drive HD-PCTU3 [Buffalo MiniStation]
	01ee  WLI-UC-GNM2 Wireless LAN Adapter [Ralink RT3070]
	01f1  SATA Adapter [HD-LBU3]
	01fd  WLI-UC-G450 Wireless LAN Adapter
0412  Award Software International
0413  Leadtek Research, Inc.
	1310  WinFast TV - NTSC + FM
	1311  WinFast TV - NTSC + MTS + FM
	1312  WinFast TV - PAL BG + FM
	1313  WinFast TV - PAL BG+TXT + FM
	1314  WinFast TV Audio - PHP PAL I
	1315  WinFast TV Audio - PHP PAL I+TXT
	1316  WinFast TV Audio - PHP PAL DK
	1317  WinFast TV Audio - PHP PAL DK+TXT
	1318  WinFast TV - PAL I/DK + FM
	1319  WinFast TV - PAL N + FM
	131a  WinFast TV Audio - PHP SECAM LL
	131b  WinFast TV Audio - PHP SECAM LL+TXT
	131c  WinFast TV Audio - PHP SECAM DK
	131d  WinFast TV - SECAM DK + TXT + FM
	131e  WinFast TV - NTSC Japan + FM
	1320  WinFast TV - NTSC
	1321  WinFast TV - NTSC + MTS
	1322  WinFast TV - PAL BG
	1323  WinFast TV - PAL BG+TXT
	1324  WinFast TV Audio - PHP PAL I
	1325  WinFast TV Audio - PHP PAL I+TXT
	1326  WinFast TV Audio - PHP PAL DK
	1327  WinFast TV Audio - PHP PAL DK+TXT
	1328  WinFast TV - PAL I/DK
	1329  WinFast TV - PAL N
	132a  WinFast TV Audio - PHP SECAM LL
	132b  WinFast TV Audio - PHP SECAM LL+TXT
	132c  WinFast TV Audio - PHP SECAM DK
	132d  WinFast TV - SECAM DK + TXT
	132e  WinFast TV - NTSC Japan
	6023  EMP Audio Device
	6024  WinFast PalmTop/Novo TV Video
	6025  WinFast DTV Dongle (cold state)
	6026  WinFast DTV Dongle (warm state)
	6029  WinFast DTV Dongle Gold
	6125  WinFast DTV Dongle
	6126  WinFast DTV Dongle BDA Driver
	6a03  RTL2832 [WinFast DTV Dongle Mini]
	6f00  WinFast DTV Dongle (STK7700P based)
0414  Giga-Byte Technology Co., Ltd
0416  Winbond Electronics Corp.
	0035  W89C35 802.11bg WLAN Adapter
	0101  Hub
	0961  AVL Flash Card Reader
	3810  Smart Card Controller
	3811  Generic Controller - Single interface
	3812  Smart Card Controller_2Interface
	3813  Panel Display
	5011  Virtual Com Port
	5518  4-Port Hub
	551a  PC Sync Keypad
	551b  PC Async Keypad
	551c  Sync Tenkey
	551d  Async Tenkey
	551e  Keyboard
	551f  Keyboard w/ Sys and Media
	5521  Keyboard
	6481  16-bit Scanner
	7721  Memory Stick Reader/Writer
	7722  Memory Stick Reader/Writer
	7723  SD Card Reader
0417  Symbios Logic
0418  AST Research
0419  Samsung Info. Systems America, Inc.
	0001  IrDA Remote Controller / Creative Cordless Mouse
	0600  Desktop Wireless 6000
	3001  Xerox P1202 Laser Printer
	3003  Olivetti PG L12L
	3201  Docuprint P8ex
	3404  SCX-5x12 series
	3406  MFP 830 series
	3407  ML-912
	3601  InkJet Color Printer
	3602  InkJet Color Printer
	4602  Remote NDIS Network Device
	8001  Hub
	8002  SyncMaster HID Monitor Control
	aa03  SDAS-3 MP3 Player
041a  Phoenix Technologies, Ltd
041b  d'TV
041d  S3, Inc.
041e  Creative Technology, Ltd
	1002  Nomad II
	1003  Blaster GamePad Cobra
	1050  GamePad Cobra
	1053  Mouse Gamer HD7600L
	200c  MuVo V100
	2020  Zen X-Fi 2
	2029  ZiiO
	2801  Prodikeys PC-MIDI multifunction keyboard
	3000  SoundBlaster Extigy
	3002  SB External Composite Device
	3010  SoundBlaster MP3+
	3014  SB External Composite Device
	3015  Sound Blaster Digital Music LX
	3020  SoundBlaster Audigy 2 NX
	3030  SB External Composite Device
	3040  SoundBlaster Live! 24-bit External SB0490
	3060  Sound Blaster Audigy 2 ZS External
	3061  SoundBlaster Audigy 2 ZS Video Editor
	3090  Sound Blaster Digital Music SX
	30d3  Sound Blaster Play!
	3100  IR Receiver (SB0540)
	3121  WoW tap chat
	3220  Sound Blaster Tactic(3D) Sigma sound card
	3f00  E-Mu Xboard 25 MIDI Controller
	3f02  E-Mu 0202
	3f04  E-Mu 0404
	3f07  E-Mu Xmidi 1x1
	4003  VideoBlaster Webcam Go Plus [W9967CF]
	4004  Nomad II MG
	4005  Webcam Blaster Go ES
	4007  Go Mini
	400a  PC-Cam 300
	400b  PC-Cam 600
	400c  Webcam 5 [pwc]
	400d  Webcam PD1001
	400f  PC-CAM 550 (Composite)
	4011  Webcam PRO eX
	4012  PC-CAM350
	4013  PC-Cam 750
	4015  CardCam Value
	4016  CardCam
	4017  Webcam Mobile [PD1090]
	4018  Webcam Vista [PD1100]
	4019  Audio Device
	401a  Webcam Vista [PD1100]
	401c  Webcam NX [PD1110]
	401d  Webcam NX Ultra
	401e  Webcam NX Pro
	401f  Webcam Notebook [PD1171]
	4020  Webcam NX
	4021  Webcam NX Ultra
	4022  Webcam NX Pro
	4028  Vista Plus cam [VF0090]
	4029  Webcam Live!
	402f  DC-CAM 3000Z
	4034  Webcam Instant
	4035  Webcam Instant
	4036  Webcam Live!/Live! Pro
	4037  Webcam Live!
	4038  ORITE CCD Webcam [PC370R]
	4039  Webcam Live! Effects
	403a  Webcam NX Pro 2
	403b  Creative Webcam Vista [VF0010]
	403c  Webcam Live! Ultra
	403d  Webcam Notebook Ultra
	403e  Webcam Vista Plus
	4041  Webcam Live! Motion
	4043  Vibra Plus Webcam
	4045  Live! Cam Voice
	4049  Live! Cam Voice
	4051  Live! Cam Notebook Pro [VF0250]
	4052  Live! Cam Vista IM
	4053  Live! Cam Video IM
	4054  Live! Cam Video IM
	4055  Live! Cam Video IM Pro
	4056  Live! Cam Video IM Pro
	4057  Live! Cam Optia
	4058  Live! Cam Optia AF
	405f  WebCam Vista (VF0330)
	4061  Live! Cam Notebook Pro [VF0400]
	4063  Live! Cam Video IM Pro
	4068  Live! Cam Notebook [VF0470]
	406c  Live! Cam Sync [VF0520]
	4083  Live! Cam Socialize [VF0640]
	4087  Live! Cam Socialize HD 1080 [VF0680]
	4088  Live! Cam Chat HD [VF0700]
	4095  Live! Cam Sync HD [VF0770]
	4100  Nomad Jukebox 2
	4101  Nomad Jukebox 3
	4102  NOMAD MuVo^2
	4106  Nomad MuVo
	4107  NOMAD MuVo
	4108  Nomad Jukebox Zen
	4109  Nomad Jukebox Zen NX
	410b  Nomad Jukebox Zen USB 2.0
	410c  Nomad MuVo NX
	410f  NOMAD MuVo^2 (Flash)
	4110  Nomad Jukebox Zen Xtra
	4111  Dell Digital Jukebox
	4116  MuVo^2
	4117  Nomad MuVo TX
	411b  Zen Touch
	411c  Nomad MuVo USB 2.0
	411d  Zen
	411e  Zen Micro
	4120  Nomad MuVo TX FM
	4123  Zen Portable Media Center
	4124  MuVo^2 FM (uHDD)
	4126  Dell DJ (2nd gen)
	4127  Dell DJ
	4128  NOMAD Jukebox Zen Xtra (mtp)
	412b  MuVo N200 with FM radio
	412f  Dell Digital Jukebox 2.Gen
	4130  Zen Micro (mtp)
	4131  DAP-HD0014 [Zen Touch] (MTP)
	4133  Mass Storage Device
	4134  Zen Neeon
	4136  Zen Sleek
	4137  Zen Sleek (mtp)
	4139  Zen Nano Plus
	413c  Zen MicroPhoto
	4150  Zen V (MTP)
	4151  Zen Vision:M (mtp)
	4152  Zen V Plus
	4153  Zen Vision W
	4154  Zen Stone
	4155  Zen Stone plus
	4157  Zen (MTP)
	500f  Broadband Blaster 8012U-V
	5015  TECOM Bluetooth Device
	ffff  Webcam Live! Ultra
041f  LCS Telegraphics
0420  Chips and Technologies
	1307  Celly SIM Card Reader
0421  Nokia Mobile Phones
	0001  E61i (PC Suite mode)
	0018  6288 GSM Smartphone
	0019  6288 GSM Smartphone (imaging mode)
	001a  6288 GSM Smartphone (file transfer mode)
	0024  5610 XpressMusic (Storage mode)
	0025  5610 XpressMusic (PC Suite mode)
	0028  5610 XpressMusic (Imaging mode)
	002d  6120 Phone (Mass storage mode)
	002e  6120 Phone (Media-Player mode)
	002f  6120 Phone (PC-Suite mode)
	0042  E51 (PC Suite mode)
	0064  3109c GSM Phone
	006b  5310 Xpress Music (PC Suite mode)
	006c  5310 Xpress music (Storage mode)
	006d  N95 (Storage mode)
	006e  N95 (Multimedia mode)
	006f  N95 (Printing mode)
	0070  N95 (PC Suite mode)
	0096  N810 Internet Tablet
	00aa  E71 (Mass storage mode)
	00ab  E71 (PC Suite mode)
	00e4  E71 (Media transfer mode)
	0103  ADL Flashing Engine AVALON Parent
	0104  ADL Re-Flashing Engine Parent
	0105  Nokia Firmware Upgrade Mode
	0106  ROM Parent
	0154  5800 XpressMusic (PC Suite mode)
	0155  5800 XpressMusic (Multimedia mode)
	0156  5800 XpressMusic (Storage mode)
	0157  5800 XpressMusic (Imaging mode)
	0199  6700 Classic (msc)
	019a  6700 Classic (PC Suite)
	019b  6700 Classic (mtp)
	01b0  6303 classic Phone (PC Suite mode)
	01b1  6303 classic Phone (Mass storage mode)
	01b2  6303 classic Phone (Printing and media mode)
	01c7  N900 (Storage Mode)
	01c8  N900 (PC-Suite Mode)
	0228  5530 XpressMusic
	023a  6730 Classic
	026a  N97 (mass storage)
	026b  N97 (Multimedia)
	026c  N97 (PC Suite)
	026d  N97 (Pictures)
	0295  660i/6600i Slide Phone (Mass Storage)
	0297  660i/6600i Slide Phone (Still Image)
	02e1  5230 (Storage mode)
	02e2  5230 (Multimedia mode)
	02e3  5230 (PC-Suite mode)
	02e4  5230 (Imaging mode)
	0360  C1-01 Ovi Suite Mode
	0396  C7-00 (Modem mode)
	03a4  C5 (Storage mode)
	03c0  C7-00 (Mass storage mode)
	03c1  C7-00 (Media transfer mode)
	03cd  C7-00 (Nokia Suite mode)
	03d1  N950
	0400  7600 Phone Parent
	0401  6650 GSM Phone
	0402  6255 Phone Parent
	0404  5510
	0405  9500 GSM Communicator
	0407  Music Player HDR-1(tm)
	040b  N-Gage GSM Phone
	040d  6620 Phone Parent
	040e  6651 Phone Parent
	040f  6230 GSM Phone
	0410  6630 Imaging Smartphone
	0411  7610 Phone Parent
	0413  6260 Phone Parent
	0414  7370
	0415  9300 GSM Smartphone
	0416  6170 Phone Parent
	0417  7270 Phone Parent
	0418  E70 (PC Suite mode)
	0419  E60 (PC Suite mode)
	041a  9500 GSM Communicator (RNDIS)
	041b  9300 GSM Smartphone (RNDIS)
	041c  7710 Phone Parent
	041d  6670 Phone Parent
	041e  6680
	041f  6235 Phone Parent
	0421  3230 Phone Parent
	0422  6681 Phone Parent
	0423  6682 Phone Parent
	0428  6230i Modem
	0429  6230i MultiMedia Card
	0431  770 Internet Tablet
	0432  N90 Phone Parent
	0435  E70 (IP Passthrough/RNDIS mode)
	0436  E60 (IP Passthrough/RNDIS mode)
	0437  6265 Phone Parent
	043a  N70 USB Phone Parent
	043b  3155 Phone Parent
	043c  6155 Phone Parent
	043d  6270 Phone Parent
	0443  N70 Phone Parent
	0444  N91
	044c  NM850iG Phone Parent
	044d  E61 (PC Suite mode)
	044e  E61 (Data Exchange mode)
	044f  E61 (IP Passthrough/RNDIS mode)
	0453  9300 Phone Parent
	0456  6111 Phone Parent
	0457  6111 Phone (Printing mode)
	045a  6280 Phone Parent
	045d  6282 Phone Parent
	046e  6110 Navigator
	0471  6110 Navigator
	0485  MTP Device
	04b9  5300
	04bc  5200 (Nokia mode)
	04bd  5200 (Storage mode)
	04be  5200 (MTP mode)
	04c3  N800 Internet Tablet
	04ce  E90 Communicator (PC Suite mode)
	04cf  E90 Communicator (Storage mode)
	04f0  Nokia N95 (PC Suite mode)
	04f9  6300 (PC Suite mode)
	0508  E65 (PC Suite mode)
	0509  E65 (Storage mode)
	0518  N9 Phone
	054d  C2-01
	0600  Digital Pen SU-1B
	0610  CS-15 (Internet Stick 3G modem)
	0661  Lumia 620/920
	069a  130 [RM-1035] (Charging only)
	0720  X (RM-980)
	0800  Connectivity Cable DKU-5
	0801  Data Cable DKU-6
	0802  CA-42 Phone Parent
0422  ADI Systems, Inc.
0423  Computer Access Technology Corp.
	000a  NetMate Ethernet
	000c  NetMate2 Ethernet
	000d  USB Chief Analyzer
	0100  Generic Universal Protocol Analyzer
	0101  UPA USBTracer
	0200  Generic 10K Universal Protocol Analyzer
	020a  PETracer ML
	0300  Generic Universal Protocol Analyzer
	0301  2500H Tracer Trainer
	030a  PETracer x1
	1237  Andromeda Hub
0424  Standard Microsystems Corp.
	0001  Integrated Hub
	0acd  Sitecom Internal Multi Memory reader/writer MD-005
	0fdc  Floppy
	10cd  Sitecom Internal Multi Memory reader/writer MD-005
	2020  USB Hub
	20cd  Sitecom Internal Multi Memory reader/writer MD-005
	20fc  6-in-1 Card Reader
	2134  Hub
	2228  9-in-2 Card Reader
	223a  8-in-1 Card Reader
	2503  USB 2.0 Hub
	2504  USB 2.0 Hub
	2507  hub
	2512  USB 2.0 Hub
	2513  2.0 Hub
	2514  USB 2.0 Hub
	2517  Hub
	2524  USB MultiSwitch Hub
	2602  USB 2.0 Hub
	2640  USB 2.0 Hub
	2660  Hub
	4060  Ultra Fast Media Reader
	4064  Ultra Fast Media Reader
	5434  Hub
	5534  Hub
	7500  LAN7500 Ethernet 10/100/1000 Adapter
	9512  SMC9512/9514 USB Hub
	9514  SMC9514 Hub
	a700  2 Port Hub
	ec00  SMSC9512/9514 Fast Ethernet Adapter
0425  Motorola Semiconductors HK, Ltd
	0101  G-Tech Wireless Mouse & Keyboard
	f102  G-Tech U+P Wireless Mouse
0426  Integrated Device Technology, Inc.
	0426  WDM Driver
0427  Motorola Electronics Taiwan, Ltd
0428  Advanced Gravis Computer Tech, Ltd
	4001  GamePad Pro
0429  Cirrus Logic
042a  Ericsson Austrian, AG
042b  Intel Corp.
	9316  8x931Hx Customer Hub
042c  Innovative Semiconductors, Inc.
042d  Micronics
042e  Acer, Inc.
	0380  MP3 Player
042f  Molex, Inc.
0430  Sun Microsystems, Inc.
	0002  109 Keyboard
	0005  Type 6 Keyboard
	000a  109 Japanese Keyboard
	000b  109 Japanese Keyboard
	0082  109 Japanese Keyboard
	0083  109 Japanese Keyboard
	00a2  Type 7 Keyboard
	0100  3-button Mouse
	100e  24.1" LCD Monitor v4 / FID-638 Mouse
	36ba  Bus Powered Hub
	a101  remote key/mouse for P3 chip
	a102  remote key/mouse/storage for P3 chip
	a103  remote storage for P3 chip
	a4a2  Ethernet (RNDIS and CDC ethernet)
	cdab  Raritan KVM dongle
0431  Itac Systems, Inc.
	0100  Mouse-Trak 3-button Track Ball
0432  Unisys Corp.
0433  Alps Electric, Inc.
	1101  IBM Game Controller
	abab  Keyboard
0434  Samsung Info. Systems America, Inc.
0435  Hyundai Electronics America
0436  Taugagreining HF
	0005  CameraMate (DPCM_USB)
0437  Framatome Connectors USA
0438  Advanced Micro Devices, Inc.
0439  Voice Technologies Group
043d  Lexmark International, Inc.
	0001  Laser Printer
	0002  Optra E310 Printer
	0003  Laser Printer
	0004  Laser Printer
	0005  Laser Printer
	0006  Laser Printer
	0007  Laser Printer
	0008  Inkjet Color Printer
	0009  Optra S2450 Printer
	000a  Laser Printer
	000b  Inkjet Color Printer
	000c  Optra E312 Printer
	000d  Laser Printer
	000e  Laser Printer
	000f  Laser Printer
	0010  Laser Printer
	0011  Laser Printer
	0012  Inkjet Color Printer
	0013  Inkjet Color Printer
	0014  InkJet Color Printer
	0015  InkJet Color Printer
	0016  Z12 Color Jetprinter
	0017  Z32 printer
	0018  Z52 Printer
	0019  Forms Printer
	001a  Z65 Printer
	001b  InkJet Photo Printer
	001c  Kodak Personal Picture Maker 200 Printer
	001d  InkJet Color Printer
	001e  InkJet Photo Printer
	001f  Kodak Personal Picture Maker 200 Card Reader
	0020  Z51 Printer
	0021  Z33 Printer
	0022  InkJet Color Printer
	0023  Laser Printer
	0024  Laser Printer
	0025  InkJet Color Printer
	0026  InkJet Color Printer
	0027  InkJet Color Printer
	0028  InkJet Color Printer
	0029  Scan Print Copy
	002a  Scan Print Copy
	002b  Scan Print Copy
	002c  Scan Print Copy
	002d  X70/X73 Scan/Print/Copy
	002e  Scan Print Copy
	002f  Scan Print Copy
	0030  Scan Print Copy
	0031  Scan Print Copy
	0032  Scan Print Copy
	0033  Scan Print Copy
	0034  Scan Print Copy
	0035  Scan Print Copy
	0036  Scan Print Copy
	0037  Scan Print Copy
	0038  Scan Print Copy
	0039  Scan Print Copy
	003a  Scan Print Copy
	003b  Scan Print Copy
	003c  Scan Print Copy
	003d  X83 Scan/Print/Copy
	003e  Scan Print Copy
	003f  Scan Print Copy
	0040  Scan Print Copy
	0041  Scan Print Copy
	0042  Scan Print Copy
	0043  Scan Print Copy
	0044  Scan Print Copy
	0045  Scan Print Copy
	0046  Scan Print Copy
	0047  Scan Print Copy
	0048  Scan Print Copy
	0049  Scan Print Copy
	004a  Scan Print Copy
	004b  Scan Print Copy
	004c  Scan Print Copy
	004d  Laser Printer
	004e  Laser Printer
	004f  InkJet Color Printer
	0050  InkJet Color Printer
	0051  Laser Printer
	0052  Laser Printer
	0053  InkJet Color Printer
	0054  InkJet Color Printer
	0057  Z35 Printer
	0058  Laser Printer
	005a  X63
	005c  InkJet Color Printer
	0060  X74/X75 Scanner
	0061  X74 Hub
	0065  X5130
	0069  X74/X75 Printer
	006d  X125
	006e  C510
	0072  X6170 Printer
	0073  InkJet Color Printer
	0078  InkJet Color Printer
	0079  InkJet Color Printer
	007a  Generic Hub
	007b  InkJet Color Printer
	007c  X1110/X1130/X1140/X1150/X1170/X1180/X1185
	007d  Photo 3150
	008a  4200 series
	008b  InkJet Color Printer
	008c  to CF/SM/SD/MS Card Reader
	008e  InkJet Color Printer
	008f  X422
	0093  X5250
	0095  E220 Printer
	0096  2200 series
	0097  P6250
	0098  7100 series
	009e  P910 series Human Interface Device
	009f  InkJet Color Printer
	00a9  IBM Infoprint 1410 MFP
	00ab  InkJet Color Printer
	00b2  3300 series
	00b8  7300 series
	00b9  8300 series
	00ba  InkJet Color Printer
	00bb  2300 series
	00bd  Printing Support
	00be  Printing Support
	00bf  Printing Support
	00c0  6300 series
	00c1  4300 series
	00c7  Printing Support
	00c8  Printing Support
	00c9  Printing Support
	00cb  Printing Support
	00cc  E120(n)
	00d0  9300 series
	00d3  X340 Scanner
	00d4  X342n Scanner
	00d5  Printing Support
	00d6  X340 Scanner
	00e8  X642e
	00e9  2400 series
	00f6  3400 series
	00f7  InkJet Color Printer
	00ff  InkJet Color Printer
	010b  2500 series
	010d  3500-4500 series
	010f  6500 series
	0142  X3650 (Printer, Scanner, Copier)
	01fa  S310 series
	4303  Xerox WorkCentre Pro 412
043e  LG Electronics USA, Inc.
	3001  AN-WF100 802.11abgn Wireless Adapter [Broadcom BCM4323]
	42bd  Flatron 795FT Plus Monitor
	4a4d  Flatron 915FT Plus Monitor
	7001  MF-PD100 Soul Digital MP3 Player
	7013  MP3 Player
	70d7  Mouse Scanner LSM-150 [LG Smart Scan Mouse]
	70f5  External HDD
	8484  LPC-U30 Webcam II
	8585  LPC-UC35 Webcam
	8888  Electronics VCS Camera II(LPC-U20)
	9800  Remote Control Receiver_iMON
	9803  eHome Infrared Receiver
	9804  DMB Receiver Control
	9c01  LGE Sync
043f  RadiSys Corp.
0440  Eizo Nanao Corp.
0441  Winbond Systems Lab.
	1456  Hub
0442  Ericsson, Inc.
	abba  Bluetooth Device
0443  Gateway, Inc.
	000e  Multimedia Keyboard
	002e  Millennium Keyboard
0445  Lucent Technologies, Inc.
0446  NMB Technologies Corp.
	6781  Keyboard with PS/2 Mouse Port
	6782  Keyboard
0447  Momentum Microsystems
044a  Shamrock Tech. Co., Ltd
044b  WSI
044c  CCL/ITRI
044d  Siemens Nixdorf AG
044e  Alps Electric Co., Ltd
	1104  Japanese Keyboard
	2002  MD-5500 Printer
	2014  Bluetooth Device
	3001  UGTZ4 Bluetooth
	3002  Bluetooth Device
	3003  Bluetooth Device
	3004  Bluetooth Adapter
	3005  Integrated Bluetooth Device
	3006  Bluetooth Adapter
	3007  Bluetooth Controller (ALPS/UGX)
	300c  Bluetooth Controller (ALPS/UGPZ6)
	300d  Bluetooth Controller (ALPS/UGPZ6)
	3010  Bluetooth Adapter
	3017  BCM2046 Bluetooth Device
	ffff  Compaq Bluetooth Multiport Module
044f  ThrustMaster, Inc.
	0400  HOTAS Cougar
	044f  GP XID
	a003  Rage 3D Game Pad
	a01b  PK-GP301 Driving Wheel
	a0a0  Top Gun Joystick
	a0a1  Top Gun Joystick (rev2)
	a0a3  Fusion Digital GamePad
	a201  PK-GP201 PlayStick
	b108  T-Flight Hotas X Flight Stick
	b10a  T.16000M Joystick
	b203  360 Modena Pro Wheel
	b300  Firestorm Dual Power
	b303  FireStorm Dual Analog 2
	b304  Firestorm Dual Power
	b307  vibrating Upad
	b30b  Wireless VibrationPad
	b315  Firestorm Dual Analog 3
	b323  Dual Trigger 3-in-1 (PC Mode)
	b324  Dual Trigger 3-in-1 (PS3 Mode)
	b603  force feedback Wheel
	b605  force feedback Racing Wheel
	b651  Ferrari GT Rumble Force Wheel
	b653  RGT Force Feedback Clutch Racing Wheel
	b654  Ferrari GT Force Feedback Wheel
	b700  Tacticalboard
0450  DFI, Inc.
0451  Texas Instruments, Inc.
	1234  Bluetooth Device
	1428  Hub
	1446  TUSB2040/2070 Hub
	16a6  BM-USBD1 BlueRobin RF heart rate sensor receiver
	2036  TUSB2036 Hub
	2046  TUSB2046 Hub
	2077  TUSB2077 Hub
	2f90  SM-USB-DIG
	3410  TUSB3410 Microcontroller
	3f00  OMAP1610
	3f02  SMC WSKP100 Wi-Fi Phone
	5409  Frontier Labs NEX IA+ Digital Audio Player
	6000  AU5 ADSL Modem (pre-reenum)
	6001  AU5 ADSL Modem
	6060  RNDIS/BeWAN ADSL2+
	6070  RNDIS/BeWAN ADSL2+
	625f  TUSB6250 ATA Bridge
	8042  Hub
	8142  TUSB8041 4-Port Hub
	926b  TUSB9260 Boot Loader
	dbc0  Device Bay Controller
	e001  GraphLink [SilverLink]
	e003  TI-84 Plus Calculator
	e004  TI-89 Titanium Calculator
	e008  TI-84 Plus Silver Calculator
	e012  TI-Nspire Calculator
	f430  MSP-FET430UIF JTAG Tool
	f432  eZ430 Development Tool
	ffff  Bluetooth Device
0452  Mitsubishi Electronics America, Inc.
	0021  HID Monitor Controls
	0050  Diamond Pro 900u CRT Monitor
	0051  Integrated Hub
	0100  Control Panel for Leica TCS SP5
0453  CMD Technology
	6781  NMB Keyboard
	6783  Chicony Composite Keyboard
0454  Vobis Microcomputer AG
0455  Telematics International, Inc.
0456  Analog Devices, Inc.
	f000  FT2232 JTAG ICE [gnICE]
	f001  FT2232H Hi-Speed JTAG ICE [gnICE+]
0457  Silicon Integrated Systems Corp.
	0150  Super Talent 1GB Flash Drive
	0151  Super Flash 1GB / GXT  64MB Flash Drive
	0162  SiS162 usb Wireless LAN Adapter
	0163  SiS163U 802.11 Wireless LAN Adapter
	0817  SiS-184-ASUS-4352.17 touch panel
	5401  Wireless Adapter RO80211GS-USB
0458  KYE Systems Corp. (Mouse Systems)
	0001  Mouse
	0002  Genius NetMouse Pro
	0003  Genius NetScroll+
	0006  Easy Mouse+
	000b  NetMouse Wheel(P+U)
	000c  TACOMA Fingerprint V1.06.01
	000e  Genius NetScroll Optical
	0013  TACOMA Fingerprint Mouse V1.06.01
	001a  Genius WebScroll+
	002e  NetScroll + Traveler / NetScroll 110
	0036  Pocket Mouse LE
	0039  NetScroll+ Superior
	003a  NetScroll+ Mini Traveler / Genius NetScroll 120
	004c  Slimstar Pro Keyboard
	0056  Ergo 300 Mouse
	0057  Enhanced Gaming Device
	0059  Enhanced Laser Device
	005a  Enhanced Device
	005b  Enhanced Device
	005c  Enhanced Laser Gaming Device
	005d  Enhanced Device
	0061  Bluetooth Dongle
	0066  Genius Traveler 1000 Wireless Mouse
	0072  Navigator 335
	0083  Bluetooth Dongle
	0087  Ergo 525V Laser Mouse
	0089  Genius Traveler 350
	00ca  Pen Mouse
	0100  EasyPen Tablet
	0101  CueCat
	011b  NetScroll T220
	1001  Joystick
	1002  Game Pad
	1003  Genius VideoCam
	1004  Flight2000 F-23 Joystick
	100a  Aashima Technology Trust Sight Fighter Vibration Feedback Joystick
	2001  ColorPage-Vivid Pro Scanner
	2004  ColorPage-HR6 V1 Scanner
	2005  ColorPage-HR6/Vivid3
	2007  ColorPage-HR6 V2 Scanner
	2008  ColorPage-HR6 V2 Scanner
	2009  ColorPage-HR6A Scanner
	2011  ColorPage-Vivid3x Scanner
	2012  Plustek Scanner
	2013  ColorPage-HR7 Scanner
	2014  ColorPage-Vivid4
	2015  ColorPage-HR7LE Scanner
	2016  ColorPage-HR6X Scanner
	2017  ColorPage-Vivid3xe
	2018  ColorPage-HR7X
	2019  ColorPage-HR6X Slim
	201a  ColorPage-Vivid4xe
	201b  ColorPage-Vivid4x
	201c  ColorPage-HR8
	201d  ColorPage-Vivid 1200 X
	201e  ColorPage-Slim 1200
	201f  ColorPage-Vivid 1200 XE
	2020  ColorPage-Slim 1200 USB2
	2021  ColorPage-SF600
	3017  SPEED WHEEL 3 Vibration
	3018  Wireless 2.4Ghz Game Pad
	3019  10-Button USB Joystick with Vibration
	301a  MaxFire G-12U Vibration
	301d  Genius MaxFire MiniPad
	400f  Genius TVGo DVB-T02Q MCE
	4012  TVGo DVB-T03 [AF9015]
	5003  G-pen 560 Tablet
	5004  G-pen Tablet
	505e  Genius iSlim 330
	6001  GF3000F Ethernet Adapter
	7004  VideoCAM Express V2
	7006  Dsc 1.3 Smart Camera Device
	7007  VideoCAM Web
	7009  G-Shot G312 Still Camera Device
	700c  VideoCAM Web V3
	700d  G-Shot G511 Composite Device
	700f  VideoCAM Web
	7012  WebCAM USB2.0
	7014  VideoCAM Live V3
	701c  G-Shot G512 Still Camera
	7020  Sim 321C
	7025  Eye 311Q Camera
	7029  Genius Look 320s (SN9C201 + HV7131R)
	702f  Genius Slim 322
	7035  i-Look 325T Camera
	7045  Genius Look 1320 V2
	704c  Genius i-Look 1321
	704d  Slim 1322AF
	7055  Slim 2020AF camera
	705a  Asus USB2.0 Webcam
	705c  Genius iSlim 1300AF
	7061  Genius iLook 1321 V2
	7066  Acer Crystal Eye Webcam
	7067  Genius iSlim 1300AF V2
	7068  Genius eFace 1325R
	706d  Genius iSlim 2000AF V2
	7076  Genius FaceCam 312
	7079  FaceCam 2025R
	707f  TVGo DVB-T03 [RTL2832]
	7088  WideCam 1050
	7089  Genius FaceCam 320
	708c  Genius WideCam F100
0459  Adobe Systems, Inc.
045a  SONICblue, Inc.
	07da  Supra Express 56K modem
	0b4a  SupraMax 2890 56K Modem [Lucent Atlas]
	0b68  SupraMax 56K Modem
	5001  Rio 600 MP3 Player
	5002  Rio 800 MP3 Player
	5003  Nike Psa/Play MP3 Player
	5005  Rio S10 MP3 Player
	5006  Rio S50 MP3 Player
	5007  Rio S35 MP3 Player
	5008  Rio 900 MP3 Player
	5009  Rio S30 MP3 Player
	500d  Fuse MP3 Player
	500e  Chiba MP3 Player
	500f  Cali MP3 Player
	5010  Rio S11 MP3 Player
	501c  Virgin MPF-1000
	501d  Rio Fuse
	501e  Rio Chiba
	501f  Rio Cali
	503f  Cali256 MP3 Player
	5202  Rio Riot MP3 Player
	5210  Rio Karma Music Player
	5220  Rio Nitrus MP3 Player
	5221  Rio Eigen
045b  Hitachi, Ltd
	0053  RX610 RX-Stick
045d  Nortel Networks, Ltd
045e  Microsoft Corp.
	0007  SideWinder Game Pad
	0008  SideWinder Precision Pro
	0009  IntelliMouse
	000b  Natural Keyboard Elite
	000e  SideWinder Freestyle Pro
	0014  Digital Sound System 80
	001a  SideWinder Precision Racing Wheel
	001b  SideWinder Force Feedback 2 Joystick
	001c  Internet Keyboard Pro
	001d  Natural Keyboard Pro
	001e  IntelliMouse Explorer
	0023  Trackball Optical
	0024  Trackball Explorer
	0025  IntelliEye Mouse
	0026  SideWinder GamePad Pro
	0027  SideWinder PnP GamePad
	0028  SideWinder Dual Strike
	0029  IntelliMouse Optical
	002b  Internet Keyboard Pro
	002d  Internet Keyboard
	002f  Integrated Hub
	0033  Sidewinder Strategic Commander
	0034  SideWinder Force Feedback Wheel
	0038  SideWinder Precision 2
	0039  IntelliMouse Optical
	003b  SideWinder Game Voice
	003c  SideWinder Joystick
	0040  Wheel Mouse Optical
	0047  IntelliMouse Explorer 3.0
	0048  Office Keyboard 1.0A
	0053  Optical Mouse
	0059  Wireless IntelliMouse Explorer
	005c  Office Keyboard (106/109)
	005f  Wireless MultiMedia Keyboard
	0061  Wireless MultiMedia Keyboard (106/109)
	0063  Wireless Natural MultiMedia Keyboard
	0065  Wireless Natural MultiMedia Keyboard (106/109)
	006a  Wireless Optical Mouse (IntelliPoint)
	006d  eHome Remote Control Keyboard keys
	006e  MN-510 802.11b Wireless Adapter [Intersil ISL3873B]
	006f  Smart Display Reference Device
	0070  Wireless MultiMedia Keyboard
	0071  Wireless MultiMedia Keyboard (106/109)
	0072  Wireless Natural MultiMedia Keyboard
	0073  Wireless Natural MultiMedia Keyboard (106/109)
	0079  IXI Ogo CT-17 handheld device
	007a  10/100 USB NIC
	007d  Notebook Optical Mouse
	007e  Wireless Transceiver for Bluetooth
	0080  Digital Media Pro Keyboard
	0083  Basic Optical Mouse
	0084  Basic Optical Mouse
	008a  Wireless Keyboard and Mouse
	008b  Dual Receiver Wireless Mouse (IntelliPoint)
	008c  Wireless Intellimouse Explorer 2.0
	0095  IntelliMouse Explorer 4.0 (IntelliPoint)
	009c  Wireless Transceiver for Bluetooth 2.0
	009d  Wireless Optical Desktop 3.0
	00a0  eHome Infrared Receiver
	00a4  Compact Optical Mouse, model 1016
	00b0  Digital Media Pro Keyboard
	00b4  Digital Media Keyboard 1.0A
	00b9  Wireless Optical Mouse 3.0
	00bb  Fingerprint Reader
	00bc  Fingerprint Reader
	00bd  Fingerprint Reader
	00c2  MN-710 802.11g Wireless Adapter [Intersil ISL3886]
	00c9  MTP Device
	00ca  Fingerprint Reader
	00cb  Basic Optical Mouse v2.0
	00ce  Generic PPC Flash device
	00d1  Optical Mouse with Tilt Wheel
	00da  eHome Infrared Receiver
	00db  Natural Ergonomic Keyboard 4000 V1.0
	00dd  Comfort Curve Keyboard 2000 V1.0
	00e1  Wireless Laser Mouse 6000 Reciever
	00f4  LifeCam VX-6000 (SN9C20x + OV9650)
	00f5  LifeCam VX-3000
	00f6  Comfort Optical Mouse 1000
	00f7  LifeCam VX-1000
	00f8  LifeCam NX-6000
	00f9  Wireless Desktop Receiver 3.1
	0202  Xbox Controller
	0280  XBox Device
	0283  Xbox Communicator
	0284  Xbox DVD Playback Kit
	0285  Xbox Controller S
	0288  Xbox Controller S Hub
	0289  Xbox Controller S
	028b  Xbox360 DVD Emulator
	028d  Xbox360 Memory Unit 64MB
	028e  Xbox360 Controller
	028f  Xbox360 Wireless Controller
	0290  Xbox360 Performance Pipe (PIX)
	0291  Xbox 360 Wireless Receiver for Windows
	0292  Xbox360 Wireless Networking Adapter
	029c  Xbox360 HD-DVD Drive
	029d  Xbox360 HD-DVD Drive
	029e  Xbox360 HD-DVD Memory Unit
	02a0  Xbox360 Big Button IR
	02a1  Xbox 360 Wireless Receiver for Windows
	02a8  Xbox360 Wireless N Networking Adapter [Atheros AR7010+AR9280]
	02ad  Xbox NUI Audio
	02ae  Xbox NUI Camera
	02b0  Xbox NUI Motor
	02b6  Xbox 360 / Bluetooth Wireless Headset
	02be  Kinect for Windows NUI Audio
	02bf  Kinect for Windows NUI Camera
	02c2  Kinect for Windows NUI Motor
	02d1  XBOX One Controller for Windows
	02d5  Xbox One Digital TV Tuner
	0400  Windows Powered Pocket PC 2002
	0401  Windows Powered Pocket PC 2002
	0402  Windows Powered Pocket PC 2002
	0403  Windows Powered Pocket PC 2002
	0404  Windows Powered Pocket PC 2002
	0405  Windows Powered Pocket PC 2002
	0406  Windows Powered Pocket PC 2002
	0407  Windows Powered Pocket PC 2002
	0408  Windows Powered Pocket PC 2002
	0409  Windows Powered Pocket PC 2002
	040a  Windows Powered Pocket PC 2002
	040b  Windows Powered Pocket PC 2002
	040c  Windows Powered Pocket PC 2002
	040d  Windows Powered Pocket PC 2002
	040e  Windows Powered Pocket PC 2002
	040f  Windows Powered Pocket PC 2002
	0410  Windows Powered Pocket PC 2002
	0411  Windows Powered Pocket PC 2002
	0412  Windows Powered Pocket PC 2002
	0413  Windows Powered Pocket PC 2002
	0414  Windows Powered Pocket PC 2002
	0415  Windows Powered Pocket PC 2002
	0416  Windows Powered Pocket PC 2002
	0417  Windows Powered Pocket PC 2002
	0432  Windows Powered Pocket PC 2003
	0433  Windows Powered Pocket PC 2003
	0434  Windows Powered Pocket PC 2003
	0435  Windows Powered Pocket PC 2003
	0436  Windows Powered Pocket PC 2003
	0437  Windows Powered Pocket PC 2003
	0438  Windows Powered Pocket PC 2003
	0439  Windows Powered Pocket PC 2003
	043a  Windows Powered Pocket PC 2003
	043b  Windows Powered Pocket PC 2003
	043c  Windows Powered Pocket PC 2003
	043d  Becker Traffic Assist Highspeed 7934
	043e  Windows Powered Pocket PC 2003
	043f  Windows Powered Pocket PC 2003
	0440  Windows Powered Pocket PC 2003
	0441  Windows Powered Pocket PC 2003
	0442  Windows Powered Pocket PC 2003
	0443  Windows Powered Pocket PC 2003
	0444  Windows Powered Pocket PC 2003
	0445  Windows Powered Pocket PC 2003
	0446  Windows Powered Pocket PC 2003
	0447  Windows Powered Pocket PC 2003
	0448  Windows Powered Pocket PC 2003
	0449  Windows Powered Pocket PC 2003
	044a  Windows Powered Pocket PC 2003
	044b  Windows Powered Pocket PC 2003
	044c  Windows Powered Pocket PC 2003
	044d  Windows Powered Pocket PC 2003
	044e  Windows Powered Pocket PC 2003
	044f  Windows Powered Pocket PC 2003
	0450  Windows Powered Pocket PC 2003
	0451  Windows Powered Pocket PC 2003
	0452  Windows Powered Pocket PC 2003
	0453  Windows Powered Pocket PC 2003
	0454  Windows Powered Pocket PC 2003
	0455  Windows Powered Pocket PC 2003
	0456  Windows Powered Pocket PC 2003
	0457  Windows Powered Pocket PC 2003
	0458  Windows Powered Pocket PC 2003
	0459  Windows Powered Pocket PC 2003
	045a  Windows Powered Pocket PC 2003
	045b  Windows Powered Pocket PC 2003
	045c  Windows Powered Pocket PC 2003
	045d  Windows Powered Pocket PC 2003
	045e  Windows Powered Pocket PC 2003
	045f  Windows Powered Pocket PC 2003
	0460  Windows Powered Pocket PC 2003
	0461  Windows Powered Pocket PC 2003
	0462  Windows Powered Pocket PC 2003
	0463  Windows Powered Pocket PC 2003
	0464  Windows Powered Pocket PC 2003
	0465  Windows Powered Pocket PC 2003
	0466  Windows Powered Pocket PC 2003
	0467  Windows Powered Pocket PC 2003
	0468  Windows Powered Pocket PC 2003
	0469  Windows Powered Pocket PC 2003
	046a  Windows Powered Pocket PC 2003
	046b  Windows Powered Pocket PC 2003
	046c  Windows Powered Pocket PC 2003
	046d  Windows Powered Pocket PC 2003
	046e  Windows Powered Pocket PC 2003
	046f  Windows Powered Pocket PC 2003
	0470  Windows Powered Pocket PC 2003
	0471  Windows Powered Pocket PC 2003
	0472  Windows Powered Pocket PC 2003
	0473  Windows Powered Pocket PC 2003
	0474  Windows Powered Pocket PC 2003
	0475  Windows Powered Pocket PC 2003
	0476  Windows Powered Pocket PC 2003
	0477  Windows Powered Pocket PC 2003
	0478  Windows Powered Pocket PC 2003
	0479  Windows Powered Pocket PC 2003
	047a  Windows Powered Pocket PC 2003
	047b  Windows Powered Pocket PC 2003
	04c8  Windows Powered Smartphone 2002
	04c9  Windows Powered Smartphone 2002
	04ca  Windows Powered Smartphone 2002
	04cb  Windows Powered Smartphone 2002
	04cc  Windows Powered Smartphone 2002
	04cd  Windows Powered Smartphone 2002
	04ce  Windows Powered Smartphone 2002
	04d7  Windows Powered Smartphone 2003
	04d8  Windows Powered Smartphone 2003
	04d9  Windows Powered Smartphone 2003
	04da  Windows Powered Smartphone 2003
	04db  Windows Powered Smartphone 2003
	04dc  Windows Powered Smartphone 2003
	04dd  Windows Powered Smartphone 2003
	04de  Windows Powered Smartphone 2003
	04df  Windows Powered Smartphone 2003
	04e0  Windows Powered Smartphone 2003
	04e1  Windows Powered Smartphone 2003
	04e2  Windows Powered Smartphone 2003
	04e3  Windows Powered Smartphone 2003
	04e4  Windows Powered Smartphone 2003
	04e5  Windows Powered Smartphone 2003
	04e6  Windows Powered Smartphone 2003
	04e7  Windows Powered Smartphone 2003
	04e8  Windows Powered Smartphone 2003
	04e9  Windows Powered Smartphone 2003
	04ea  Windows Powered Smartphone 2003
	04ec  Windows Phone (Zune)
	063e  Zune HD Media Player
	0640  KIN Phone
	0641  KIN Phone
	0642  KIN Phone
	0707  Wireless Laser Mouse 8000
	0708  Transceiver v 3.0 for Bluetooth
	070a  Charon Bluetooth Dongle (DFU)
	070f  LifeChat LX-3000 Headset
	0710  Zune Media Player
	0713  Wireless Presenter Mouse 8000
	0719  Xbox 360 Wireless Adapter
	071f  Mouse/Keyboard 2.4GHz Transceiver V2.0
	0721  LifeCam NX-3000 (UVC-compliant)
	0723  LifeCam VX-7000 (UVC-compliant)
	0724  SideWinder Mouse
	0728  LifeCam VX-5000
	0730  Digital Media Keyboard 3000
	0734  Wireless Optical Desktop 700
	0736  Sidewinder X5 Mouse
	0737  Compact Optical Mouse 500
	0745  Nano Transceiver v1.0 for Bluetooth
	0750  Wired Keyboard 600
	0752  Wired Keyboard 400
	075d  LifeCam Cinema
	0766  LifeCam VX-800
	0768  Sidewinder X4
	076c  Comfort Mouse 4500
	076d  LifeCam HD-5000
	0772  LifeCam Studio
	0779  LifeCam HD-3000
	0780  Comfort Curve Keyboard 3000
	0797  Optical Mouse 200
	07a5  Wireless Receiver 1461C
	07f8  Wired Keyboard 600 (model 1576)
	07fd  Nano Transceiver 1.1
	930a  ISOUSB.SYS Intel 82930 Isochronous IO Test Board
	ffca  Catalina
	fff8  Keyboard
	ffff  Windows CE Mass Storage
0460  Ace Cad Enterprise Co., Ltd
	0004  Tablet (5x3.75)
	0006  LCD Tablet (12x9)
	0008  Tablet (3x2.25)
0461  Primax Electronics, Ltd
	0010  HP PR1101U / Primax PMX-KPR1101U Keyboard
	0300  G2-300 Scanner
	0301  G2E-300 Scanner
	0302  G2-300 #2 Scanner
	0303  G2E-300 #2 Scanner
	0340  Colorado 9600 Scanner
	0341  Colorado 600u Scanner
	0345  Visioneer 6200 Scanner
	0346  Memorex Maxx 6136u Scanner
	0347  Primascan Colorado 2600u/Visioneer 4400 Scanner
	0360  Colorado 19200 Scanner
	0361  Colorado 1200u Scanner
	0363  VistaScan Astra 3600(ENG)
	0364  LG Electronics Scanworks 600U Scanner
	0365  VistaScan Astra 3600(ENG)
	0366  6400
	0367  VistaScan Astra 3600(ENG)
	0371  Visioneer Onetouch 8920 Scanner
	0374  UMAX Astra 2500
	0375  VistaScan Astra 3600(ENG)
	0377  Medion MD 5345 Scanner
	0378  VistaScan Astra 3600(ENG)
	037b  Medion MD 6190 Scanner
	037c  VistaScan Astra 3600(ENG)
	0380  G2-600 Scanner
	0381  ReadyScan 636i Scanner
	0382  G2-600 #2 Scanner
	0383  G2E-600 Scanner
	038a  UMAX Astra 3000/3600
	038b  Xerox 2400 Onetouch
	038c  UMAX Astra 4100
	0392  Medion/Lifetec/Tevion/Cytron MD 6190
	03a8  9420M
	0813  IBM UltraPort Camera
	0815  Micro Innovations IC200 Webcam
	0819  Fujifilm IX-30 Camera [webcam mode]
	081a  Fujifilm IX-30 Camera [storage mode]
	081c  Elitegroup ECS-C11 Camera
	081d  Elitegroup ECS-C11 Storage
	0a00  Micro Innovations Web Cam 320
	4d01  Comfort Keyboard
	4d02  Mouse-in-a-Box
	4d03  Kensington Mouse-in-a-box
	4d04  Mouse
	4d06  Balless Mouse (HID)
	4d0f  HP Optical Mouse
	4d15  Dell Optical Mouse
	4d17  Optical Mouse
	4d20  HP Optical Mouse
	4d2a  PoPo Elixir Mouse (HID)
	4d2b  Wireless Laser Mini Mouse (HID)
	4d2c  PoPo Mini Pointer Mouse (HID)
	4d2e  Optical Mobile Mouse (HID)
	4d51  0Y357C PMX-MMOCZUL (B) [Dell Laser Mouse]
	4d62  HP Laser Mobile Mini Mouse
	4d75  Rocketfish RF-FLBTAD Bluetooth Adapter
	4d81  Dell N889 Optical Mouse
	4de7  webcam
0463  MGE UPS Systems
	0001  UPS
	ffff  UPS
0464  AMP/Tycoelectronics Corp.
0467  AT&T Paradyne
0468  Wieson Technologies Co., Ltd
046a  Cherry GmbH
	0001  Keyboard
	0003  My3000 Hub
	0004  CyBoard Keyboard
	0005  XX33 SmartCard Reader Keyboard
	0008  Wireless Keyboard and Mouse
	0010  SmartBoard XX44
	0011  G83 (RS 6000) Keyboard
	0021  CyMotion Expert Combo
	0023  CyMotion Master Linux Keyboard G230
	0027  CyMotion Master Solar Keyboard
	002a  Wireless Mouse & Keyboard
	002d  SmartTerminal XX44
	003e  SmartTerminal ST-2xxx
	0041  G86 6240 Keyboard
	0080  eHealth Terminal ST 1503
	0081  eHealth Keyboard G87 1504
	0106  R-300 Wireless Mouse Receiver
	010d  MX-Board 3.0 Keyboard
046b  American Megatrends, Inc.
	0001  Keyboard
	0101  PS/2 Keyboard, Mouse & Joystick Ports
	0301  USB 1.0 Hub
	0500  Serial & Parallel Ports
	ff10  Virtual Keyboard and Mouse
046c  Toshiba Corp., Digital Media Equipment
046d  Logitech, Inc.
	0082  Acer Aspire 5672 Webcam
	0200  WingMan Extreme Joystick
	0203  M2452 Keyboard
	0301  M4848 Mouse
	0401  HP PageScan
	0402  NEC PageScan
	040f  Logitech/Storm PageScan
	0430  Mic (Cordless)
	0801  QuickCam Home
	0802  Webcam C200
	0804  Webcam C250
	0805  Webcam C300
	0807  Webcam B500
	0808  Webcam C600
	0809  Webcam Pro 9000
	080a  Portable Webcam C905
	080f  Webcam C120
	0810  QuickCam Pro
	0819  Webcam C210
	081b  Webcam C310
	081d  HD Webcam C510
	0820  QuickCam VC
	0821  HD Webcam C910
	0825  Webcam C270
	0826  HD Webcam C525
	0828  HD Webcam B990
	082b  Webcam C170
	082d  HD Pro Webcam C920
	0830  QuickClip
	0836  B525 HD Webcam
	0837  BCC950 ConferenceCam
	0840  QuickCam Express
	0843  Webcam C930e
	0850  QuickCam Web
	0870  QuickCam Express
	0890  QuickCam Traveler
	0892  OrbiCam
	0894  CrystalCam
	0895  QuickCam for Dell Notebooks
	0896  OrbiCam
	0897  QuickCam for Dell Notebooks
	0899  QuickCam for Dell Notebooks
	089d  QuickCam E2500 series
	08a0  QuickCam IM
	08a1  QuickCam IM with sound
	08a2  Labtec Webcam Pro
	08a3  QuickCam QuickCam Chat
	08a6  QuickCam IM
	08a7  QuickCam Image
	08a9  Notebook Deluxe
	08aa  Labtec Notebooks
	08ac  QuickCam Cool
	08ad  QuickCam Communicate STX
	08ae  QuickCam for Notebooks
	08af  QuickCam Easy/Cool
	08b0  QuickCam 3000 Pro [pwc]
	08b1  QuickCam Notebook Pro
	08b2  QuickCam Pro 4000
	08b3  QuickCam Zoom
	08b4  QuickCam Zoom
	08b5  QuickCam Sphere
	08b9  QuickCam IM
	08bd  Microphone (Pro 4000)
	08c0  QuickCam Pro 3000
	08c1  QuickCam Fusion
	08c2  QuickCam PTZ
	08c3  Camera (Notebooks Pro)
	08c5  QuickCam Pro 5000
	08c6  QuickCam for DELL Notebooks
	08c7  QuickCam OEM Cisco VT Camera II
	08c9  QuickCam Ultra Vision
	08ca  Mic (Fusion)
	08cb  Mic (Notebooks Pro)
	08cc  Mic (PTZ)
	08ce  QuickCam Pro 5000
	08cf  QuickCam UpdateMe
	08d0  QuickCam Express
	08d7  QuickCam Communicate STX
	08d8  QuickCam for Notebook Deluxe
	08d9  QuickCam IM/Connect
	08da  QuickCam Messanger
	08dd  QuickCam for Notebooks
	08e0  QuickCam Express
	08e1  Labtec Webcam
	08f0  QuickCam Messenger
	08f1  QuickCam Express
	08f2  Microphone (Messenger)
	08f3  QuickCam Express
	08f4  Labtec Webcam
	08f5  QuickCam Messenger Communicate
	08f6  QuickCam Messenger Plus
	0900  ClickSmart 310
	0901  ClickSmart 510
	0903  ClickSmart 820
	0905  ClickSmart 820
	0910  QuickCam Cordless
	0920  QuickCam Express
	0921  Labtec Webcam
	0922  QuickCam Live
	0928  QuickCam Express
	0929  Labtec Webcam Pro
	092a  QuickCam for Notebooks
	092b  Labtec Webcam Plus
	092c  QuickCam Chat
	092d  QuickCam Express / Go
	092e  QuickCam Chat
	092f  QuickCam Express Plus
	0950  Pocket Camera
	0960  ClickSmart 420
	0970  Pocket750
	0990  QuickCam Pro 9000
	0991  QuickCam Pro for Notebooks
	0992  QuickCam Communicate Deluxe
	0994  QuickCam Orbit/Sphere AF
	09a1  QuickCam Communicate MP/S5500
	09a2  QuickCam Communicate Deluxe/S7500
	09a4  QuickCam E 3500
	09a5  Quickcam 3000 For Business
	09a6  QuickCam Vision Pro
	09b0  Acer OrbiCam
	09b2  Fujitsu Webcam
	09c0  QuickCam for Dell Notebooks Mic
	09c1  QuickCam Deluxe for Notebooks
	0a01  USB Headset
	0a02  Premium Stereo USB Headset 350
	0a03  Logitech USB Microphone
	0a04  V20 portable speakers (USB powered)
	0a07  Z-10 Speakers
	0a0b  ClearChat Pro USB
	0a0c  Clear Chat Comfort USB Headset
	0a13  Z-5 Speakers
	0a17  G330 Headset
	0a1f  G930
	0a29  H600 [Wireless Headset]
	0a37  USB Headset H540
	0a38  Headset H340
	0a44  Headset H390
	0a4d  G430 Surround Sound Gaming Headset
	0b02  C-UV35 [Bluetooth Mini-Receiver] (HID proxy mode)
	8801  Video Camera
	b305  BT Mini-Receiver
	bfe4  Premium Optical Wheel Mouse
	c000  N43 [Pilot Mouse]
	c001  N48/M-BB48/M-UK96A [FirstMouse Plus]
	c002  M-BA47 [MouseMan Plus]
	c003  MouseMan
	c004  WingMan Gaming Mouse
	c005  WingMan Gaming Wheel Mouse
	c00b  MouseMan Wheel
	c00c  Optical Wheel Mouse
	c00d  MouseMan Wheel+
	c00e  M-BJ58/M-BJ69 Optical Wheel Mouse
	c00f  MouseMan Traveler/Mobile
	c011  Optical MouseMan
	c012  Mouseman Dual Optical
	c014  Corded Workstation Mouse
	c015  Corded Workstation Mouse
	c016  Optical Wheel Mouse
	c018  Optical Wheel Mouse
	c019  Optical Tilt Wheel Mouse
	c01a  M-BQ85 Optical Wheel Mouse
	c01b  MX310 Optical Mouse
	c01c  Optical Mouse
	c01d  MX510 Optical Mouse
	c01e  MX518 Optical Mouse
	c024  MX300 Optical Mouse
	c025  MX500 Optical Mouse
	c030  iFeel Mouse
	c031  iFeel Mouse+
	c032  MouseMan iFeel
	c033  iFeel MouseMan+
	c034  MouseMan Optical
	c035  Mouse
	c036  Mouse
	c037  Mouse
	c038  Mouse
	c03d  M-BT96a Pilot Optical Mouse
	c03e  Premium Optical Wheel Mouse (M-BT58)
	c03f  M-BT85 [UltraX Optical Mouse]
	c040  Corded Tilt-Wheel Mouse
	c041  G5 Laser Mouse
	c042  G3 Laser Mouse
	c043  MX320/MX400 Laser Mouse
	c044  LX3 Optical Mouse
	c045  Optical Mouse
	c046  RX1000 Laser Mouse
	c047  Laser Mouse M-UAL120
	c048  G9 Laser Mouse
	c049  G5 Laser Mouse
	c050  RX 250 Optical Mouse
	c051  G3 (MX518) Optical Mouse
	c053  Laser Mouse
	c054  Bluetooth mini-receiver
	c058  M115 Mouse
	c05a  M90/M100 Optical Mouse
	c05b  M-U0004 810-001317 [B110 Optical USB Mouse]
	c05d  Optical Mouse
	c05f  M115 Optical Mouse
	c061  RX1500 Laser Mouse
	c062  M-UAS144 [LS1 Laser Mouse]
	c063  DELL Laser Mouse
	c064  M110 corded optical mouse (M-B0001)
	c066  G9x Laser Mouse
	c068  G500 Laser Mouse
	c069  M-U0007 [Corded Mouse M500]
	c06a  USB Optical Mouse
	c06b  G700 Wireless Gaming Mouse
	c06c  Optical Mouse
	c077  M105 Optical Mouse
	c07c  M-R0017 [G700s Rechargeable Gaming Mouse]
	c101  UltraX Media Remote
	c110  Harmony 785/885 Remote
	c111  Harmony 525 Remote
	c112  Harmony 890 Remote
	c11f  Harmony 900/1100 Remote
	c121  Harmony One Remote
	c122  Harmony 650/700 Remote
	c124  Harmony 300/350 Remote
	c125  Harmony 200 Remote
	c126  Harmony Link
	c129  Harmony Hub
	c12b  Harmony Touch/Ultimate Remote
	c201  WingMan Extreme Joystick with Throttle
	c202  WingMan Formula
	c207  WingMan Extreme Digital 3D
	c208  WingMan Gamepad Extreme
	c209  WingMan Gamepad
	c20a  WingMan RumblePad
	c20b  WingMan Action Pad
	c20c  WingMan Precision
	c20d  WingMan Attack 2
	c20e  WingMan Formula GP
	c211  iTouch Cordless Reciever
	c212  WingMan Extreme Digital 3D
	c213  J-UH16 (Freedom 2.4 Cordless Joystick)
	c214  ATK3 (Attack III Joystick)
	c215  Extreme 3D Pro
	c216  Dual Action Gamepad
	c218  Logitech RumblePad 2 USB
	c219  Cordless RumblePad 2
	c21a  Precision Gamepad
	c21c  G13 Advanced Gameboard
	c21d  F310 Gamepad [XInput Mode]
	c21e  F510 Gamepad [XInput Mode]
	c21f  F710 Wireless Gamepad [XInput Mode]
	c221  G11/G15 Keyboard / Keyboard
	c222  G15 Keyboard / LCD
	c223  G11/G15 Keyboard / USB Hub
	c225  G11/G15 Keyboard / G keys
	c226  G15 Refresh Keyboard
	c227  G15 Refresh Keyboard
	c228  G19 Gaming Keyboard
	c229  G19 Gaming Keyboard Macro Interface
	c22a  Gaming Keyboard G110
	c22b  Gaming Keyboard G110 G-keys
	c22d  G510 Gaming Keyboard
	c22e  G510 Gaming Keyboard onboard audio
	c245  G400 Optical Mouse
	c246  Gaming Mouse G300
	c248  G105 Gaming Keyboard
	c24a  G600 Gaming Mouse
	c24c  G400s Optical Mouse
	c24d  G710 Gaming Keyboard
	c24e  G500s Laser Gaming Mouse
	c281  WingMan Force
	c283  WingMan Force 3D
	c285  WingMan Strike Force 3D
	c286  Force 3D Pro
	c287  Flight System G940
	c291  WingMan Formula Force
	c293  WingMan Formula Force GP
	c294  Driving Force
	c295  Momo Force Steering Wheel
	c298  Driving Force Pro
	c299  G25 Racing Wheel
	c29b  G27 Racing Wheel
	c29c  Speed Force Wireless Wheel for Wii
	c2a0  Wingman Force Feedback Mouse
	c2a1  WingMan Force Feedback Mouse
	c301  iTouch Keyboard
	c302  iTouch Pro Keyboard
	c303  iTouch Keyboard
	c305  Internet Keyboard
	c307  Internet Keyboard
	c308  Internet Navigator Keyboard
	c309  Y-BF37 [Internet Navigator Keyboard]
	c30a  iTouch Composite
	c30b  NetPlay Keyboard
	c30c  Internet Keys (X)
	c30d  Internet Keys
	c30e  UltraX Keyboard (Y-BL49)
	c30f  Logicool HID-Compliant Keyboard (106 key)
	c311  Y-UF49 [Internet Pro Keyboard]
	c312  DeLuxe 250 Keyboard
	c313  Internet 350 Keyboard
	c315  Classic Keyboard 200
	c316  HID-Compliant Keyboard
	c317  Wave Corded Keyboard
	c318  Illuminated Keyboard
	c31a  Comfort Wave 450
	c31b  Compact Keyboard K300
	c31c  Keyboard K120
	c31d  Media Keyboard K200
	c401  TrackMan Marble Wheel
	c402  Marble Mouse (2-button)
	c403  Turbo TrackMan Marble FX
	c404  TrackMan Wheel
	c408  Marble Mouse (4-button)
	c501  Cordless Mouse Receiver
	c502  Cordless Mouse & iTouch Keys
	c503  Cordless Mouse+Keyboard Receiver
	c504  Cordless Mouse+Keyboard Receiver
	c505  Cordless Mouse+Keyboard Receiver
	c506  MX700 Cordless Mouse Receiver
	c508  Cordless Trackball
	c509  Cordless Keyboard & Mouse
	c50a  Cordless Mouse
	c50b  Cordless Desktop Optical
	c50c  Cordless Desktop S510
	c50d  Cordless Mouse
	c50e  Cordless Mouse Receiver
	c510  Cordless Mouse
	c512  LX-700 Cordless Desktop Receiver
	c513  MX3000 Cordless Desktop Receiver
	c514  Cordless Mouse
	c515  Cordless 2.4 GHz Presenter Presentation remote control
	c517  LX710 Cordless Desktop Laser
	c518  MX610 Laser Cordless Mouse
	c51a  MX Revolution/G7 Cordless Mouse
	c51b  V220 Cordless Optical Mouse for Notebooks
	c521  Cordless Mouse Receiver
	c525  MX Revolution Cordless Mouse
	c526  Nano Receiver
	c529  Logitech Keyboard + Mice
	c52b  Unifying Receiver
	c52d  R700 Remote Presenter receiver
	c52e  MK260 Wireless Combo Receiver
	c52f  Unifying Receiver
	c531  C-U0007 [Unifying Receiver]
	c532  Unifying Receiver
	c534  Unifying Receiver
	c603  3Dconnexion Spacemouse Plus XT
	c605  3Dconnexion CADman
	c606  3Dconnexion Spacemouse Classic
	c621  3Dconnexion Spaceball 5000
	c623  3Dconnexion Space Traveller 3D Mouse
	c625  3Dconnexion Space Pilot 3D Mouse
	c626  3Dconnexion Space Navigator 3D Mouse
	c627  3Dconnexion Space Explorer 3D Mouse
	c628  3Dconnexion Space Navigator for Notebooks
	c629  3Dconnexion SpacePilot Pro 3D Mouse
	c62b  3Dconnexion Space Mouse Pro
	c640  NuLOOQ navigator
	c702  Cordless Presenter
	c703  Elite Keyboard Y-RP20 + Mouse MX900 (Bluetooth)
	c704  diNovo Wireless Desktop
	c705  MX900 Bluetooth Wireless Hub (C-UJ16A)
	c707  Bluetooth wireless hub
	c708  Bluetooth wireless hub
	c709  BT Mini-Receiver (HCI mode)
	c70a  MX5000 Cordless Desktop
	c70b  BT Mini-Receiver (HID proxy mode)
	c70c  BT Mini-Receiver (HID proxy mode)
	c70d  Bluetooth wireless hub
	c70e  MX1000 Bluetooth Laser Mouse
	c70f  Bluetooth wireless hub
	c712  Bluetooth wireless hub
	c714  diNovo Edge Keyboard
	c715  Bluetooth wireless hub
	c71a  Bluetooth wireless hub
	c71d  Bluetooth wireless hub
	c71f  diNovo Mini Wireless Keyboard
	c720  Bluetooth wireless hub
	ca03  MOMO Racing
	ca04  Formula Vibration Feedback Wheel
	cab1  Cordless Keyboard for Wii HID Receiver
	d001  QuickCam Pro
046e  Behavior Tech. Computer Corp.
	0100  Keyboard
	3001  Mass Storage Device
	3002  Mass Storage Device
	3003  Mass Storage Device
	3005  Mass Storage Device
	3008  Mass Storage Device
	5250  KeyMaestro Multimedia Keyboard
	5273  KeyMaestro Multimedia Keyboard
	52e6  Cordless Mouse
	5308  KeyMaestro Keyboard
	5408  KeyMaestro Multimedia Keyboard/Hub
	5500  Portable Keyboard 86+9 keys (Model 6100C US)
	5550  5 button optical mouse model M873U
	5720  Smart Card Reader
	6782  BTC 7932 mouse+keyboard
046f  Crystal Semiconductor
0471  Philips (or NXP)
	0101  DSS350 Digital Speaker System
	0104  DSS330 Digital Speaker System [uda1321]
	0105  UDA1321
	014f  GoGear SA9200
	0160  MP3 Player
	0161  MP3 Player
	0163  GoGear SA1100
	0164  GoGear SA1110/02
	0165  GoGear SA1330
	0201  Hub
	0222  Creative Nomad Jukebox
	0302  PCA645VC Webcam [pwc]
	0303  PCA646VC Webcam [pwc]
	0304  Askey VC010 Webcam [pwc]
	0307  PCVC675K Webcam [pwc]
	0308  PCVC680K Webcam [pwc]
	030b  PC VGA Camera (Vesta Fun)
	030c  PCVC690K Webcam [pwc]
	0310  PCVC730K Webcam [pwc]
	0311  PCVC740K ToUcam Pro [pwc]
	0312  PCVC750K Webcam [pwc]
	0314  DMVC 1000K
	0316  DMVC 2000K Video Capture
	0321  FunCam
	0322  DMVC1300K PC Camera
	0325  SPC 200NC PC Camera
	0326  SPC 300NC PC Camera
	0327  Webcam SPC 6000 NC (Webcam w/ mic)
	0328  SPC 700NC PC Camera
	0329  SPC 900NC PC Camera / ORITE CCD Webcam(PC370R)
	032d  SPC 210NC PC Camera
	032e  SPC 315NC PC Camera
	0330  SPC 710NC PC Camera
	0331  SPC 1300NC PC Camera
	0332  SPC 1000NC PC Camera
	0333  SPC 620NC PC Camera
	0334  SPC 520/525NC PC Camera
	0401  Semiconductors CICT Keyboard
	0402  PS/2 Mouse on Semiconductors CICT Keyboard
	0406  15 inch Detachable Monitor
	0407  10 inch Mobile Monitor
	0408  SG3WA1/74 802.11b WLAN Adapter [Atmel AT76C503A]
	0471  Digital Speaker System
	0601  OVU1020 IR Dongle (Kbd+Mouse)
	0602  ATI Remote Wonder II Input Device
	0603  ATI Remote Wonder II Controller
	0608  eHome Infrared Receiver
	060a  TSU9600 Remote Control
	060c  Consumer Infrared Transceiver (HP)
	060d  Consumer Infrared Transceiver (SRM5100)
	060e  RF Dongle
	060f  Consumer Infrared Transceiver
	0613  Infrared Transceiver
	0617  IEEE802.15.4 RF Dongle
	0619  TSU9400 Remote Control
	0666  Hantek DDS-3005 Arbitrary Waveform Generator
	0700  Semiconductors CICT Hub
	0701  150P1 TFT Display
	0809  AVNET Bluetooth Device
	0811  JR24 CDRW
	0814  DCCX38/P data cable
	0815  eHome Infrared Receiver
	0844  SA2111/02 1GB Flash Audio Player
	084a  GoGear SA3125
	084e  GoGear SA60xx (mtp)
	0888  Hantek DDS-3005 Arbitrary Waveform Generator
	1103  Digital Speaker System
	1120  Creative Rhomba MP3 player
	1125  Nike psa[128max Player
	1137  HDD065 MP3 player
	1201  Arima Bluetooth Device
	1230  Wireless Adapter 11g
	1232  SNU6500 Wireless Adapter
	1233  Wireless Adapter Bootloader Download
	1236  SNU5600 802.11bg
	1237  TalkTalk SNU5630NS/05 802.11bg
	1552  ISP 1581 Hi-Speed USB MPEG2 Encoder Reference Kit
	1801  Diva MP3 player
	200a  Wireless Network Adapter
	200f  802.11n Wireless Adapter
	2021  SDE3273FC/97 2.5" SATA HDD Enclosure [INIC-1608L]
	2022  GoGear SA52XX
	2034  Webcam SPC530NC
	2036  Webcam SPC1030NC
	203f  TSU9200 Remote Control
	2046  TSU9800 Remote Control
	204e  GoGear RaGa (SA1942/02)
	205e  TSU9300 Remote Control
	206c  MCE IR Receiver - Spinel plusf0r ASUS
	2070  GoGear Mix
	2076  GoGear Aria
	2079  GoGear Opus
	2088  MCE IR Receiver with ALS- Spinel plus for ASUS
	209e  PTA01 Wireless Adapter
	20b6  GoGear Vibe
	20d0  SPZ2000 Webcam [PixArt PAC7332]
	20e3  GoGear Raga
	20e4  GoGear ViBE 8GB
	2160  Mio LINK Heart Rate Monitor
	262c  SPC230NC Webcam
	485d  Senselock SenseIV v2.x
	df55  LPCXpresso LPC-Link
0472  Chicony Electronics Co., Ltd
	0065  PFU-65 Keyboard [Chicony]
	b086  Asus USB2.0 Webcam
	b091  Webcam
0473  Sanyo Information Business Co., Ltd
0474  Sanyo Electric Co., Ltd
	0110  Digital Voice Recorder R200
	0217  Xacti J2
	022f  C5 Digital Media Camera (mass storage mode)
	0230  C5 Digital Media Camera (PictBridge mode)
	0231  C5 Digital Media Camera (PC control mode)
	0401  Optical Drive
	0701  SCP-4900 Cellphone
	071f  Usb Com Port Enumerator
	0722  W33SA Camera
0475  Relisys/Teco Information System
	0100  NEC Petiscan
	0103  Eclipse 1200U/Episode
	0210  Scorpio Ultra 3
0476  AESP
0477  Seagate Technology, Inc.
0478  Connectix Corp.
	0001  QuickCam
	0002  QuickClip
	0003  QuickCam Pro
0479  Advanced Peripheral Laboratories
047a  Semtech Corp.
	0004  ScreenCoder UR7HCTS2-USB
047b  Silitek Corp.
	0001  Keyboard
	0002  Keyboard and Mouse
	0011  SK-1688U Keyboard
	00f9  SK-1789u Keyboard
	0101  BlueTooth Keyboard and Mouse
	020b  SK-3105 SmartCard Reader
	050e  Internet Compact Keyboard
	1000  Trust Office Scan USB 19200
	1002  HP ScanJet 4300c Parallel Port
047c  Dell Computer Corp.
	ffff  UPS Tower 500W LV
047d  Kensington
	1001  Mouse*in*a*Box
	1002  Expert Mouse Pro
	1003  Orbit TrackBall
	1004  MouseWorks
	1005  TurboBall
	1006  TurboRing
	1009  Orbit TrackBall for Mac
	1012  PocketMouse
	1013  Mouse*in*a*Box Optical Pro
	1014  Expert Mouse Pro Wireless
	1015  Expert Mouse
	1016  ADB/USB Orbit
	1018  Studio Mouse
	101d  Mouse*in*a*Box Optical Pro
	101e  Studio Mouse Wireless
	101f  PocketMouse Pro
	1020  Expert Mouse Trackball
	1021  Expert Mouse Wireless
	1022  Orbit Optical
	1023  Pocket Mouse Pro Wireless
	1024  PocketMouse
	1025  Mouse*in*a*Box Optical Elite Wireless
	1026  Pocket Mouse Pro
	1027  StudioMouse
	1028  StudioMouse Wireless
	1029  Mouse*in*a*Box Optical Elite
	102a  Mouse*in*a*Box Optical
	102b  PocketMouse
	102c  Iridio
	102d  Pilot Optical
	102e  Pilot Optical Pro
	102f  Pilot Optical Pro Wireless
	1042  Ci25m Notebook Optical Mouse [Diamond Eye Precision]
	1043  Ci65m Wireless Notebook Optical Mouse
	104a  PilotMouse Mini Retractable
	105d  PocketMouse Bluetooth
	105e  Bluetooth EDR Dongle
	1061  PocketMouse Grip
	1062  PocketMouse Max
	1063  PocketMouse Max Wireless
	1064  PocketMouse 2.0 Wireless
	1065  PocketMouse 2.0
	1066  PocketMouse Max Glow
	1067  ValueMouse
	1068  ValueOpt White
	1069  ValueOpt Black
	106a  PilotMouse Laser Wireless Mini
	106b  PilotMouse Laser - 3 Button
	106c  PilotMouse Laser - Gaming
	106d  PilotMouse Laser - Wired
	106e  PilotMouse Micro Laser
	1070  ValueOpt Travel
	1071  ValueOpt RF TX
	1072  PocketMouse Colour
	1073  PilotMouse Laser - 6 Button
	1074  PilotMouse Laser Wireless Mini
	1075  SlimBlade Presenter Media Mouse
	1076  SlimBlade Media Mouse
	1077  SlimBlade Presenter Mouse
	1152  Bluetooth EDR Dongle
	2002  Optical Elite Wireless
	2010  Wireless Presentation Remote
	2012  Wireless Presenter with Laser Pointer
	2021  PilotBoard Wireless
	2030  PilotBoard Wireless
	2034  SlimBlade Media Notebook Set
	2041  SlimBlade Trackball
	2048  Orbit Trackball with Scroll Ring
	4003  Gravis Xterminator Digital Gamepad
	4005  Gravis Eliminator GamePad Pro
	4006  Gravis Eliminator AfterShock
	4007  Gravis Xterminator Force
	4008  Gravis Destroyer TiltPad
	5001  Cabo I Camera
	5002  VideoCam CABO II
	5003  VideoCam
047e  Agere Systems, Inc. (Lucent)
	0300  ORiNOCO Card
	1001  USS720 Parallel Port
	2892  Systems Soft Modem
	bad1  Lucent 56k Modem
	f101  Atlas Modem
047f  Plantronics, Inc.
	0101  Bulk Driver
	0301  Bulk Driver
	0411  Savi Office Base Station
	0ca1  USB DSP v4 Audio Interface
	4254  BUA-100 Bluetooth Adapter
	ac01  Savi 7xx
	ad01  GameCom 777 5.1 Headset
	c008  Audio 655 DSP
	c00e  Blackwire C310 headset
0480  Toshiba America Inc
	0001  InTouch Module
	0004  InTouch Module
	0011  InTouch Module
	0014  InTouch Module
	0100  Stor.E Slim USB 3.0
	0200  External Disk
	a006  External Disk 1.5TB
	a007  External Disk USB 3.0
	a009  Stor.E Basics
	a00d  STOR.E BASICS 500GB
	a202  Canvio Basics HDD
	b001  Stor.E Partner
	d000  External Disk 2TB Model DT01ABA200
	d010  External Disk 3TB
	d011  Canvio Desk
0481  Zenith Data Systems
0482  Kyocera Corp.
	000e  FS-1020D Printer
	000f  FS-1920 Mono Printer
	0100  Finecam S3x
	0101  Finecam S4
	0103  Finecam S5
	0105  Finecam L3
	0106  Finecam
	0107  Digital Camera Device
	0108  Digital Camera Device
	0203  AH-K3001V
	0204  iBurst Terminal
0483  STMicroelectronics
	0137  BeWAN ADSL USB ST (blue or green)
	0138  Unicorn II (ST70138B + MTC-20174TQ chipset)
	1307  Cytronix 6in1 Card Reader
	163d  Cool Icam Digi-MP3
	2015  TouchChip Fingerprint Reader
	2016  Fingerprint Reader
	2017  Biometric Smart Card Reader
	2018  BioSimKey
	2302  Portable Flash Device (PFD)
	347b  ST-LINK/V2-1
	3744  STLINK Pseudo disk
	3747  ST Micro Connect Lite
	3748  ST-LINK/V2
	374b  ST-LINK/V2.1 (Nucleo-F103RB)
	4810  ISDN adapter
	481d  BT Digital Access adapter
	5000  ST Micro/Ergenic ERG BT-002 Bluetooth Adapter
	5001  ST Micro Bluetooth Device
	5710  Joystick in FS Mode
	5720  STM microSD Flash Device
	5721  Hantek DDS-3X25 Arbitrary Waveform Generator
	5730  STM32 Audio Streaming
	5740  STM32F407
	7270  ST Micro Serial Bridge
	7554  56k SoftModem
	91d1  Sensor Hub
	df11  STM Device in DFU Mode
	ff10  Swann ST56 Modem
0484  Specialix
0485  Nokia Monitors
0486  ASUS Computers, Inc.
	0185  EeePC T91MT HID Touch Panel
0487  Stewart Connector
0488  Cirque Corp.
0489  Foxconn / Hon Hai
	0502  SmartMedia Card Reader Firmware Loader
	0503  SmartMedia Card Reader
	d00c  Rollei Compactline (Storage Mode)
	d00e  Rollei Compactline (Video Mode)
	e000  T-Com TC 300
	e003  Pirelli DP-L10
	e00d  Broadcom Bluetooth 2.1 Device
	e00f  Foxconn T77H114 BCM2070 [Single-Chip Bluetooth 2.1 + EDR Adapter]
	e011  Acer Bluetooth module
	e016  Ubee PXU1900 WiMAX Adapter [Beceem BCSM250]
	e02c  Atheros AR5BBU12 Bluetooth Device
	e032  Broadcom BCM20702 Bluetooth
	e042  Broadcom BCM20702 Bluetooth
	e04d  Atheros AR3012 Bluetooth
048a  S-MOS Systems, Inc.
048c  Alps Electric Ireland, Ltd
048d  Integrated Technology Express, Inc.
	1165  IT1165 Flash Controller
	1336  SD/MMC Cardreader
	1345  Multi Cardreader
	9006  IT9135 BDA Afatech DVB-T HDTV Dongle
	9009  Zolid HD DVD Maker
	9135  Zolid Mini DVB-T Stick
	9306  IT930x DVB stick
	9503  ITE it9503 feature-limited DVB-T transmission chip [ccHDtv]
	9507  ITE it9507 full featured DVB-T transmission chip [ccHDtv]
048f  Eicon Tech.
0490  United Microelectronics Corp.
0491  Capetronic
	0003  Taxan Monitor Control
0492  Samsung SemiConductor, Inc.
	0140  MP3 player
	0141  MP3 Player
0493  MAG Technology Co., Ltd
0495  ESS Technology, Inc.
0496  Micron Electronics
0497  Smile International
	c001  Camera Device
0498  Capetronic (Kaohsiung) Corp.
0499  Yamaha Corp.
	1000  UX256 MIDI I/F
	1001  MU1000
	1002  MU2000
	1003  MU500
	1004  UW500
	1005  MOTIF6
	1006  MOTIF7
	1007  MOTIF8
	1008  UX96 MIDI I/F
	1009  UX16 MIDI I/F
	100a  EOS BX
	100c  UC-MX
	100d  UC-KX
	100e  S08
	100f  CLP-150
	1010  CLP-170
	1011  P-250
	1012  TYROS
	1013  PF-500
	1014  S90
	1015  MOTIF-R
	1016  MDP-5
	1017  CVP-204
	1018  CVP-206
	1019  CVP-208
	101a  CVP-210
	101b  PSR-1100
	101c  PSR-2100
	101d  CLP-175
	101e  PSR-K1
	101f  EZ-J24
	1020  EZ-250i
	1021  MOTIF ES 6
	1022  MOTIF ES 7
	1023  MOTIF ES 8
	1024  CVP-301
	1025  CVP-303
	1026  CVP-305
	1027  CVP-307
	1028  CVP-309
	1029  CVP-309GP
	102a  PSR-1500
	102b  PSR-3000
	102e  ELS-01/01C
	1030  PSR-295/293
	1031  DGX-205/203
	1032  DGX-305
	1033  DGX-505
	1037  PSR-E403
	103c  MOTIF-RACK ES
	1054  S90XS Keyboard/Music Synthesizer
	160f  P-105
	2000  DGP-7
	2001  DGP-5
	3001  YST-MS55D USB Speaker
	3003  YST-M45D USB Speaker
	4000  NetVolante RTA54i Broadband&ISDN Router
	4001  NetVolante RTW65b Broadband Wireless Router
	4002  NetVolante RTW65i Broadband&ISDN Wireless Router
	4004  NetVolante RTA55i Broadband VoIP Router
	5000  CS1D
	5001  DSP1D
	5002  DME32
	5003  DM2000
	5004  02R96
	5005  ACU16-C
	5006  NHB32-C
	5007  DM1000
	5008  01V96
	5009  SPX2000
	500a  PM5D
	500b  DME64N
	500c  DME24N
	6001  CRW2200UX Lightspeed 2 External CD-RW Drive
	7000  DTX
	7010  UB99
049a  Gandalf Technologies, Ltd
049b  Curtis Computer Products
049c  Acer Advanced Labs, Inc.
	0002  Keyboard (???)
049d  VLSI Technology
049f  Compaq Computer Corp.
	0002  InkJet Color Printer
	0003  iPAQ PocketPC
	000e  Internet Keyboard
	0012  InkJet Color Printer
	0018  PA-1/PA-2 MP3 Player
	0019  InkJet Color Printer
	001a  S4 100 Scanner
	001e  IJ650 Inkjet Printer
	001f  WL215 Adapter
	0021  S200 Scanner
	0027  Bluetooth Multiport Module by Compaq
	002a  1400P Inkjet Printer
	002b  A3000
	002c  Lexmark X125
	0032  802.11b Adapter [ipaq h5400]
	0033  Wireless LAN MultiPort W100 [Intersil PRISM 2.5]
	0036  Bluetooth Multiport Module
	0051  KU-0133 Easy Access Interner Keyboard
	0076  Wireless LAN MultiPort W200
	0080  GPRS Multiport
	0086  Bluetooth Device
	504a  Personal Jukebox PJB100
	505a  Linux-USB "CDC Subset" Device, or Itsy (experimental)
	8511  iPAQ Networking 10/100 Ethernet [pegasus2]
04a0  Digital Equipment Corp.
04a1  SystemSoft Corp.
	fff0  Telex Composite Device
04a2  FirePower Systems
04a3  Trident Microsystems, Inc.
04a4  Hitachi, Ltd
	0004  DVD-CAM DZ-MV100A Camcorder
	001e  DVDCAM USB HS Interface
04a5  Acer Peripherals Inc. (now BenQ Corp.)
	0001  Keyboard
	0002  API Ergo K/B
	0003  API Generic K/B Mouse
	12a6  AcerScan C310U
	1a20  Prisa 310U
	1a2a  Prisa 620U
	2022  Prisa 320U/340U
	2040  Prisa 620UT
	205e  ScanPrisa 640BU
	2060  Prisa 620U+/640U
	207e  Prisa 640BU
	209e  ScanPrisa 640BT
	20ae  S2W 3000U
	20b0  S2W 3300U/4300U
	20be  Prisa 640BT
	20c0  Prisa 1240UT
	20de  S2W 4300U+
	20f8  Benq 5000
	20fc  Benq 5000
	20fe  SW2 5300U
	2137  Benq 5150/5250
	2202  Benq 7400UT
	2311  Benq 5560
	3003  Benq Webcam
	3008  Benq 1500
	300a  Benq 3410
	300c  Benq 1016
	3019  Benq DC C40
	4000  P30 Composite Device
	4013  BenQ-Siemens EF82/SL91
	4044  BenQ-Siemens SF71
	4045  BenQ-Siemens E81
	4048  BenQ M7
	6001  Mass Storage Device
	6002  Mass Storage Device
	6003  ATA/ATAPI Adapter
	6004  Mass Storage Device
	6005  Mass Storage Device
	6006  Mass Storage Device
	6007  Mass Storage Device
	6008  Mass Storage Device
	6009  Mass Storage Device
	600a  Mass Storage Device
	600b  Mass Storage Device
	600c  Mass Storage Device
	600d  Mass Storage Device
	600e  Mass Storage Device
	600f  Mass Storage Device
	6010  Mass Storage Device
	6011  Mass Storage Device
	6012  Mass Storage Device
	6013  Mass Storage Device
	6014  Mass Storage Device
	6015  Mass Storage Device
	6125  MP3 Player
	6180  MP3 Player
	6200  MP3 Player
	7500  Hi-Speed Mass Storage Device
	9000  AWL300 Wireless Adapter
	9001  AWL400 Wireless Adapter
	9213  Kbd Hub
04a6  Nokia Display Products
	00b9  Audio
	0180  Hub Type P
	0181  HID Monitor Controls
04a7  Visioneer
	0100  StrobePro
	0101  Strobe Pro Scanner (1.01)
	0102  StrobePro Scanner
	0211  OneTouch 7600 Scanner
	0221  OneTouch 5300 Scanner
	0223  OneTouch 8200
	0224  OneTouch 4800 USB/Microtek Scanport 3000
	0225  VistaScan Astra 3600(ENG)
	0226  OneTouch 5300 USB
	0229  OneTouch 7100
	022a  OneTouch 6600
	022c  OneTouch 9000/9020
	0231  6100 Scanner
	0311  6200 EPP/USB Scanner
	0321  OneTouch 8100 EPP/USB Scanner
	0331  OneTouch 8600 EPP/USB Scanner
	0341  6400
	0361  VistaScan Astra 3600(ENG)
	0362  OneTouch 9320
	0371  OneTouch 8700/8920
	0380  OneTouch 7700
	0382  Photo Port 7700
	0390  9650
	03a0  Xerox 4800 One Touch
	0410  OneTouch Pro 8800/8820
	0421  9450 USB
	0423  9750 Scanner
	0424  Strobe XP 450
	0425  Strobe XP 100
	0426  Strobe XP 200
	0427  Strobe XP 100
	0444  OneTouch 7300
	0445  CardReader 100
	0446  Xerox DocuMate 510
	0447  XEROX DocuMate 520
	0448  XEROX DocuMate 250
	0449  Xerox DocuMate 252
	044a  Xerox 6400
	044c  Xerox DocuMate 262
	0474  Strobe XP 300
	0475  Xerox DocuMate 272
	0478  Strobe XP 220
	0479  Strobe XP 470
	047a  9450
	047b  9650
	047d  9420
	0480  9520
	048f  Strobe XP 470
	0491  Strobe XP 450
	0493  9750
	0494  Strobe XP 120
	0497  Patriot 430
	0498  Patriot 680
	0499  Patriot 780
	049b  Strobe XP 100
	04a0  7400
	04ac  Xerox Travel Scanner 100
04a8  Multivideo Labs, Inc.
	0101  Hub
	0303  Peripheral Switch
	0404  Peripheral Switch
04a9  Canon, Inc.
	1005  BJ Printer Hub
	1035  PD Printer Storage
	1050  BJC-8200
	1051  BJC-3000 Color Printer
	1052  BJC-6100
	1053  BJC-6200
	1054  BJC-6500
	1055  BJC-85
	1056  BJC-2110 Color Printer
	1057  LR1
	105a  BJC-55
	105b  S600 Printer
	105c  S400
	105d  S450 Printer
	105e  S800
	1062  S500 Printer
	1063  S4500
	1064  S300 Printer
	1065  S100
	1066  S630
	1067  S900
	1068  S9000
	1069  S820
	106a  S200 Printer
	106b  S520 Printer
	106d  S750 Printer
	106e  S820D
	1070  S530D
	1072  I850 Printer
	1073  I550 Printer
	1074  S330 Printer
	1076  i70
	1077  i950
	107a  S830D
	107b  i320
	107c  i470D
	107d  i9100
	107e  i450
	107f  i860
	1082  i350
	1084  i250
	1085  i255
	1086  i560
	1088  i965
	108a  i455
	108b  i900D
	108c  i475D
	108d  PIXMA iP2000
	108f  i80
	1090  i9900 Photo Printer
	1091  PIXMA iP1500
	1093  PIXMA iP4000
	1094  PIXMA iP3000x Printer
	1095  PIXMA iP6000D
	1097  PIXMA iP5000
	1098  PIXMA iP1000
	1099  PIXMA iP8500
	109c  PIXMA iP4000R
	109d  iP90
	10a0  PIXMA iP1600 Printer
	10a2  iP4200
	10a4  iP5200R
	10a5  iP5200
	10a7  iP6210D
	10a8  iP6220D
	10a9  iP6600D
	10b6  PIXMA iP4300 Printer
	10c2  PIXMA iP1800 Printer
	10c4  Pixma iP4500 Printer
	1404  W6400PG
	1405  W8400PG
	150f  BIJ2350 PCL
	1510  BIJ1350 PCL
	1512  BIJ1350D PCL
	1601  DR-2080C Scanner
	1607  DR-6080 Scanner
	1700  PIXMA MP110 Scanner
	1701  PIXMA MP130 Scanner
	1702  MP410 Composite
	1703  MP430 Composite
	1704  MP330 Composite
	1706  PIXMA MP750 Scanner
	1707  PIXMA MP780 Scanner
	1708  PIXMA MP760 Scanner
	1709  PIXMA MP150 Scanner
	170a  PIXMA MP170 Scanner
	170b  PIXMA MP450 Scanner
	170c  PIXMA MP500 Scanner
	170d  PIXMA MP800 Scanner
	170e  MP800R
	1710  MP950
	1712  MP530
	1713  PIXMA MP830 Scanner
	1714  MP160
	1715  MP180 Storage
	1716  MP460 Composite
	1717  MP510
	1718  MP600 Storage
	171a  MP810 Storage
	171b  MP960
	1721  MP210 ser
	1723  MP470 ser
	1724  PIXMA MP520 series
	1725  MP610 ser
	1726  MP970 ser
	1727  MX300 ser
	1728  MX310 ser
	1729  MX700 ser
	172b  MP140 ser
	1736  PIXMA MX320 series
	173a  MP250 series printer
	173b  PIXMA MP270 All-In-One Printer
	173e  MP560
	173f  Pixma MP640 Multifunction device
	1748  Pixma MG5150
	174d  MX360 ser
	176d  PIXMA MG2550
	1900  CanoScan LiDE 90
	1901  CanoScan 8800F
	1904  CanoScan LiDE 100
	1905  CanoScan LiDE 200
	1906  CanoScan 5600F
	1907  CanoScan LiDE 700F
	1909  CanoScan LiDE 110
	190a  CanoScan LiDE 210
	190d  CanoScan 9000F Mark II
	190e  CanoScan LiDE 120
	2200  CanoScan LiDE 25
	2201  CanoScan FB320U
	2202  CanoScan FB620U
	2204  CanoScan FB630U
	2205  CanoScan FB1210U
	2206  CanoScan N650U/N656U
	2207  CanoScan 1220U
	2208  CanoScan D660U
	220a  CanoScan D2400UF
	220b  CanoScan D646U
	220c  CanoScan D1250U2
	220d  CanoScan N670U/N676U/LiDE 20
	220e  CanoScan N1240U/LiDE 30
	220f  CanoScan 8000F
	2210  CanoScan 9900F
	2212  CanoScan 5000F
	2213  CanoScan LiDE 50/LiDE 35/LiDE 40
	2214  CanoScan LiDE 80
	2215  CanoScan 3000/3000F/3000ex
	2216  CanoScan 3200F
	2217  CanoScan 5200F
	2219  CanoScan 9950F
	221b  CanoScan 4200F
	221c  CanoScan LiDE 60
	221e  CanoScan 8400F
	221f  CanoScan LiDE 500F
	2220  CanoScan LIDE 25
	2224  CanoScan LiDE 600F
	2225  CanoScan LiDE 70
	2228  CanoScan 4400F
	2229  CanoScan 8600F
	2602  MultiPASS C555
	2603  MultiPASS C755
	260a  CAPT Printer
	260e  LBP-2000
	2610  MPC600F
	2611  SmartBase MPC400
	2612  MultiPASS C855
	2617  CAPT Printer
	261a  iR1600
	261b  iR1610
	261c  iC2300
	261f  MPC200 Printer
	2621  iR2000
	2622  iR2010
	2623  FAX-B180C
	2629  FAXPHONE L75
	262b  LaserShot LBP-1120 Printer
	262d  iR C3200
	262f  MultiPASS MP730
	2630  MultiPASS MP700
	2631  LASER CLASS 700
	2632  FAX-L2000
	2635  MPC190
	2637  iR C6800
	2638  iR C3100
	263c  Smartbase MP360
	263d  MP370
	263e  MP390 FAX
	263f  MP375
	2646  MF5530 Scanner Device V1.9.1
	2647  MF5550 Composite
	264d  PIXMA MP710
	264e  MF5630
	264f  MF5650 (FAX)
	2650  iR 6800C EUR
	2651  iR 3100C EUR
	2655  FP-L170/MF350/L380/L398
	2659  MF8100
	265b  CAPT Printer
	265c  iR C3220
	265d  MF5730
	265e  MF5750
	265f  MF5770
	2660  MF3110
	2663  iR3570/iR4570
	2664  iR2270/iR2870
	2665  iR C2620
	2666  iR C5800
	2667  iR85PLUS
	2669  iR105PLUS
	266a  CAPT Device
	266b  iR8070
	266c  iR9070
	266d  iR 5800C EUR
	266e  CAPT Device
	266f  iR2230
	2670  iR3530
	2671  iR5570/iR6570
	2672  iR C3170
	2673  iR 3170C EUR
	2674  L120
	2675  iR2830
	2676  CAPT Device
	2677  iR C2570
	2678  iR 2570C EUR
	2679  CAPT Device
	267a  iR2016
	267b  iR2020
	267d  MF7100 series
	2684  MF3200 series
	2686  MF6500 series
	2687  iR4530
	2688  LBP3460
	268c  iR C6870
	268d  iR 6870C EUR
	268e  iR C5870
	268f  iR 5870C EUR
	2691  iR7105
	26a3  MF4100 series
	26b0  MF4600 series
	26b4  MF4010 series
	26b5  MF4200 series
	26da  LBP3010B printer
	26e6  iR1024
	2736  I-SENSYS MF4550d
	2737  MF4410
	3041  PowerShot S10
	3042  CanoScan FS4000US Film Scanner
	3043  PowerShot S20
	3044  EOS D30
	3045  PowerShot S100
	3046  IXY Digital
	3047  Digital IXUS
	3048  PowerShot G1
	3049  PowerShot Pro90 IS
	304a  CP-10
	304b  IXY Digital 300
	304c  PowerShot S300
	304d  Digital IXUS 300
	304e  PowerShot A20
	304f  PowerShot A10
	3050  PowerShot unknown 1
	3051  PowerShot S110
	3052  Digital IXUS V
	3055  PowerShot G2
	3056  PowerShot S40
	3057  PowerShot S30
	3058  PowerShot A40
	3059  PowerShot A30
	305b  ZR45MC Digital Camcorder
	305c  PowerShot unknown 2
	3060  EOS D60
	3061  PowerShot A100
	3062  PowerShot A200
	3063  CP-100
	3065  PowerShot S200
	3066  Digital IXUS 330
	3067  MV550i Digital Video Camera
	3069  PowerShot G3
	306a  Digital unknown 3
	306b  MVX2i Digital Video Camera
	306c  PowerShot S45
	306d  PowerShot S45 PtP Mode
	306e  PowerShot G3 (normal mode)
	306f  PowerShot G3 (ptp)
	3070  PowerShot S230
	3071  PowerShot S230 (ptp)
	3072  PowerShot SD100 / Digital IXUS II (ptp)
	3073  PowerShot A70 (ptp)
	3074  PowerShot A60 (ptp)
	3075  IXUS 400 Camera
	3076  PowerShot A300
	3077  PowerShot S50
	3078  ZR70MC Digital Camcorder
	307a  MV650i (normal mode)
	307b  MV630i Digital Video Camera
	307c  CP-200
	307d  CP-300
	307f  Optura 20
	3080  MVX150i (normal mode) / Optura 20 (normal mode)
	3081  Optura 10
	3082  MVX100i / Optura 10
	3083  EOS 10D
	3084  EOS 300D / EOS Digital Rebel
	3085  PowerShot G5
	3087  Elura 50 (PTP mode)
	3088  Elura 50 (normal mode)
	308d  MVX3i
	308e  FV M1 (normal mode) / MVX 3i (normal mode) / Optura Xi (normal mode)
	3093  Optura 300
	3096  IXY DV M2 (normal mode) / MVX 10i (normal mode)
	3099  EOS 300D (ptp)
	309a  PowerShot A80
	309b  Digital IXUS (ptp)
	309c  PowerShot S1 IS
	309d  Powershot Pro 1
	309f  Camera
	30a0  Camera
	30a1  Camera
	30a2  Camera
	30a8  Elura 60E/Optura 40 (ptp)
	30a9  MVX25i (normal mode) / Optura 40 (normal mode)
	30b1  PowerShot S70 (normal mode) / PowerShot S70 (PTP mode)
	30b2  PowerShot S60 (normal mode) / PowerShot S60 (PTP mode)
	30b3  PowerShot G6 (normal mode) / PowerShot G6 (PTP mode)
	30b4  PowerShot S500
	30b5  PowerShot A75
	30b6  Digital IXUS II2  / Digital IXUS II2 (PTP mode) / PowerShot SD110 (PTP mode) / PowerShot SD110 Digital ELPH
	30b7  PowerShot A400 / PowerShot A400 (PTP mode)
	30b8  PowerShot A310 / PowerShot A310 (PTP mode)
	30b9  Powershot A85
	30ba  PowerShot S410 Digital Elph
	30bb  PowerShot A95
	30bd  CP-220
	30be  CP-330
	30bf  Digital IXUS 40
	30c0  Digital IXUS 30 (PTP mode) / PowerShot SD200 (PTP mode)
	30c1  Digital IXUS 50 (normal mode) / IXY Digital 55 (normal mode) / PowerShot A520 (PTP mode) / PowerShot SD400 (normal mode)
	30c2  PowerShot A510 (normal mode) / PowerShot A510 (PTP mode)
	30c4  Digital IXUS i5 (normal mode) / IXY Digital L2 (normal mode) / PowerShot SD20 (normal mode)
	30ea  EOS 1D Mark II (PTP mode)
	30eb  EOS 20D
	30ec  EOS 20D (ptp)
	30ee  EOS 350D
	30ef  EOS 350D (ptp)
	30f0  PowerShot S2 IS (PTP mode)
	30f2  Digital IXUS 700 (normal mode) / Digital IXUS 700 (PTP mode) / IXY Digital 600 (normal mode) / PowerShot SD500 (normal mode) / PowerShot SD500 (PTP mode)
	30f4  PowerShot SD30 / Ixus iZoom / IXY DIGITAL L3
	30f5  SELPHY CP500
	30f6  SELPHY CP400
	30f8  Powershot A430
	30f9  PowerShot A410 (PTP mode)
	30fa  PowerShot S80
	30fc  PowerShot A620 (PTP mode)
	30fd  PowerShot A610 (normal mode)/PowerShot A610 (PTP mode)
	30fe  Digital IXUS 65 (PTP mode)/PowerShot SD630 (PTP mode)
	30ff  Digital IXUS 55 (PTP mode)/PowerShot SD450 (PTP mode)
	3100  PowerShot TX1
	310b  SELPHY CP600
	310e  Digital IXUS 50 (PTP mode)
	310f  PowerShot A420
	3110  EOS Digital Rebel XTi
	3115  PowerShot SD900 / Digital IXUS 900 Ti / IXY DIGITAL 1000
	3116  Digital IXUS 750 / PowerShot SD550 (PTP mode)
	3117  PowerShot A700
	3119  PowerShot SD700 IS / Digital IXUS 800 IS / IXY Digital 800 IS
	311a  PowerShot S3 IS
	311b  PowerShot A540
	311c  PowerShot SD600 DIGITAL ELPH / DIGITAL IXUS 60 / IXY DIGITAL 70
	3125  PowerShot G7
	3126  PowerShot A530
	3127  SELPHY CP710
	3128  SELPHY CP510
	312d  Elura 100
	3136  PowerShot SD800 IS / Digital IXUS 850 IS / IXY DIGITAL 900 IS
	3137  PowerShot SD40 / Digital IXUS i7 IXY / DIGITAL L4
	3138  PowerShot A710 IS
	3139  PowerShot A640
	313a  PowerShot A630
	3141  SELPHY ES1
	3142  SELPHY CP730
	3143  SELPHY CP720
	3145  EOS 450D
	3146  EOS 40D
	3147  EOS 1Ds Mark III
	3148  PowerShot S5 IS
	3149  PowerShot A460
	314b  PowerShot SD850 IS DIGITAL ELPH / Digital IXUS 950 IS / IXY DIGITAL 810 IS
	314c  PowerShot A570 IS
	314d  PowerShot A560
	314e  PowerShot SD750 DIGITAL ELPH / DIGITAL IXUS 75 / IXY DIGITAL 90
	314f  PowerShot SD1000 DIGITAL ELPH / DIGITAL IXUS 70 / IXY DIGITAL 10
	3150  PowerShot A550
	3155  PowerShot A450
	315a  PowerShot G9
	315b  PowerShot A650 IS
	315d  PowerShot A720
	315e  PowerShot SX100 IS
	315f  PowerShot SD950 IS DIGITAL ELPH / DIGITAL IXUS 960 IS / IXY DIGITAL 2000 IS
	3160  Digital IXUS 860 IS
	3170  SELPHY CP750
	3171  SELPHY CP740
	3172  SELPHY CP520
	3173  PowerShot SD890 IS DIGITAL ELPH / Digital IXUS 970 IS / IXY DIGITAL 820 IS
	3174  PowerShot SD790 IS DIGITAL ELPH / Digital IXUS 90 IS / IXY DIGITAL 95 IS
	3175  IXY Digital 25 IS
	3176  PowerShot A590
	3177  PowerShot A580
	317a  PC1267 [Powershot A470]
	3184  Digital IXUS 80 IS (PTP mode)
	3185  SELPHY ES2
	3186  SELPHY ES20
	318d  PowerShot SX100 IS
	318e  PowerShot A1000 IS
	318f  PowerShot G10
	3191  PowerShot A2000 IS
	3192  PowerShot SX110 IS
	3193  PowerShot SD990 IS DIGITAL ELPH / Digital IXUS 980 IS / IXY DIGITAL 3000 IS
	3195  PowerShot SX1 IS
	3196  PowerShot SD880 IS DIGITAL ELPH / Digital IXUS 870 IS / IXY DIGITAL 920 IS
	319a  EOS 7D
	319b  EOS 50D
	31aa  SELPHY CP770
	31ab  SELPHY CP760
	31ad  PowerShot E1
	31af  SELPHY ES3
	31b0  SELPHY ES30
	31b1  SELPHY CP530
	31bc  PowerShot D10
	31bd  PowerShot SD960 IS DIGITAL ELPH / Digital IXUS 110 IS / IXY DIGITAL 510 IS
	31be  PowerShot A2100 IS
	31bf  PowerShot A480
	31c0  PowerShot SX200 IS
	31c1  PowerShot SD970 IS DIGITAL ELPH / Digital IXUS 990 IS / IXY DIGITAL 830 IS
	31c2  PowerShot SD780 IS DIGITAL ELPH / Digital IXUS 100 IS / IXY DIGITAL 210 IS
	31c3  PowerShot A1100 IS
	31c4  PowerShot SD1200 IS DIGITAL ELPH / Digital IXUS 95 IS / IXY DIGITAL 110 IS
	31cf  EOS Rebel T1i / EOS 500D / EOS Kiss X3
	31dd  SELPHY CP780
	31df  PowerShot G11
	31e0  PowerShot SX120 IS
	31e1  PowerShot S90
	31e4  PowerShot SX20 IS
	31e5  Digital IXUS 200 IS
	31e6  PowerShot SD940 IS DIGITAL ELPH / Digital IXUS 120 IS / IXY DIGITAL 220 IS
	31e7  SELPHY CP790
	31ea  EOS Rebel T2i / EOS 550D / EOS Kiss X4
	31ee  SELPHY ES40
	31ef  PowerShot A495
	31f0  PowerShot A490
	31f1  PowerShot A3100 IS / PowerShot A3150 IS
	31f2  PowerShot A3000 IS
	31f3  PowerShot Digital ELPH SD1400 IS
	31f4  PowerShot SD1300 IS / IXUS 105
	31f5  Powershot SD3500 IS / IXUS 210 IS
	31f6  PowerShot SX210 IS
	31f7  Powershot SD4000 IS / IXUS 300 HS / IXY 30S
	31f8  Powershot SD4500 IS / IXUS 1000 HS / IXY 50S
	31ff  Digital IXUS 55
	3209  Vixia HF S21 A
	320f  PowerShot G12
	3210  Powershot SX30 IS
	3211  PowerShot SX130 IS
	3212  Powershot S95
	3214  SELPHY CP800
	3218  EOS 600D / Rebel T3i (ptp)
	3223  PowerShot A3300 IS
	3224  PowerShot A3200 IS
	3225  PowerShot ELPH 500 HS / IXUS 310 HS
	3226  PowerShow A800
	3227  PowerShot ELPH 100 HS / IXUS 115 HS
	3228  PowerShot SX230 HS
	3229  PowerShot ELPH 300 HS / IXUS 220 HS
	322a  PowerShot A2200
	322b  Powershot A1200
	322c  PowerShot SX220 HS
	3233  PowerShot G1 X
	3234  PowerShot SX150 IS
	3235  PowerShot ELPH 510 HS / IXUS 1100 HS
	3236  PowerShot S100
	3237  PowerShot ELPH 310 HS / IXUS 230 HS
	3238  PowerShot SX40 HS
	323b  EOS Rebel T4i
	323e  PowerShot A1300
	323f  PowerShot A810
	3240  PowerShot ELPH 320 HS / IXUS 240 HS
	3241  PowerShot ELPH 110 HS / IXUS 125 HS
	3242  PowerShot D20
	3243  PowerShot A4000 IS
	3244  PowerShot SX260 HS
	3245  PowerShot SX240 HS
	3247  PowerShot ELPH 520 HS / IXUS 500 HS
	3248  PowerShot A3400 IS
	3249  PowerShot A2400 IS
	324a  PowerShot A2300
	3255  SELPHY CP900
	3256  SELPHY CP810
	3258  PowerShot G15
	3259  PowerShot SX50 HS
	325a  PowerShot SX160 IS
	325b  PowerShot S110
	325c  PowerShot SX500 IS
	325e  PowerShot N
	325f  PowerShot SX280 HS
	3260  PowerShot SX270 HS
	3261  PowerShot A3500 IS
	3262  PowerShot A2600
	3263  PowerShot SX275 HS
	3264  PowerShot A1400
	3265  Powershot ELPH 130 IS / IXUS 140
	3266  Powershot ELPH 120 IS / IXUS 135
	3268  PowerShot ELPH 330 HS / IXUS 255 HS
	3271  PowerShot A2500
	3276  PowerShot SX170 IS
	3277  PowerShot SX510 HS
	3278  PowerShot S200
	327d  Powershot ELPH 115 IS / IXUS 132
	327f  EOS Rebel T5 / EOS 1200D / EOS Kiss X70
	3284  PowerShot D30
	3285  PowerShot SX700 HS
	3286  PowerShot SX600 HS
	3287  PowerShot ELPH 140 IS / IXUS 150
	3288  Powershot ELPH 135 / IXUS 145
	3289  PowerShot ELPH 340 HS / IXUS 265 HS
	328a  PowerShot ELPH 150 IS / IXUS 155
	328b  PowerShot N Facebook(R) Ready
	3299  EOS M3
	329a  PowerShot SX60 HS
	329b  PowerShot SX520 HS
	329c  PowerShot SX400 IS
	329f  PowerShot SX530 HS
	32a6  PowerShot SX710 HS
	32aa  Powershot ELPH 160 / IXUS 160
	32ac  PowerShot ELPH 170 IS / IXUS 170
04aa  DaeWoo Telecom, Ltd
04ab  Chromatic Research
04ac  Micro Audiometrics Corp.
04ad  Dooin Electronics
	2501  Bluetooth Device
04af  Winnov L.P.
04b0  Nikon Corp.
	0102  Coolpix 990
	0103  Coolpix 880
	0104  Coolpix 995
	0106  Coolpix 775
	0107  Coolpix 5000
	0108  Coolpix 2500
	0109  Coolpix 2500 (ptp)
	010a  Coolpix 4500
	010b  Coolpix 4500 (ptp)
	010d  Coolpix 5700 (ptp)
	010e  Coolpix 4300 (storage)
	010f  Coolpix 4300 (ptp)
	0110  Coolpix 3500 (Sierra Mode)
	0111  Coolpix 3500 (ptp)
	0112  Coolpix 885 (ptp)
	0113  Coolpix 5000 (ptp)
	0114  Coolpix 3100 (storage)
	0115  Coolpix 3100 (ptp)
	0117  Coolpix 2100 (ptp)
	0119  Coolpix 5400 (ptp)
	011d  Coolpix 3700 (ptp)
	0121  Coolpix 3200 (ptp)
	0122  Coolpix 2200 (ptp)
	0124  Coolpix 8400 (mass storage mode)
	0125  Coolpix 8400 (ptp)
	0126  Coolpix 8800
	0129  Coolpix 4800 (ptp)
	012c  Coolpix 4100 (storage)
	012d  Coolpix 4100 (ptp)
	012e  Coolpix 5600 (ptp)
	0130  Coolpix 4600 (ptp)
	0135  Coolpix 5900 (ptp)
	0136  Coolpix 7900 (storage)
	0137  Coolpix 7900 (ptp)
	013a  Coolpix 100 (storage)
	013b  Coolpix 100 (ptp)
	0141  Coolpix P2 (storage)
	0142  Coolpix P2 (ptp)
	0163  Coolpix P5100 (ptp)
	0169  Coolpix P50 (ptp)
	0202  Coolpix SQ (ptp)
	0203  Coolpix 4200 (mass storage mode)
	0204  Coolpix 4200 (ptp)
	0205  Coolpix 5200 (storage)
	0206  Coolpix 5200 (ptp)
	0301  Coolpix 2000 (storage)
	0302  Coolpix 2000 (ptp)
	0317  Coolpix L20 (ptp)
	0402  DSC D100 (ptp)
	0403  D2H (mass storage mode)
	0404  D2H SLR (ptp)
	0405  D70 (mass storage mode)
	0406  DSC D70 (ptp)
	0408  D2X SLR (ptp)
	0409  D50 digital camera
	040a  D50 (ptp)
	040c  D2Hs
	040e  DSC D70s (ptp)
	040f  D200 (mass storage mode)
	0410  D200 (ptp)
	0413  D40 (mass storage mode)
	041e  D60 digital camera (mass storage mode)
	0422  D700 (ptp)
	0423  D5000
	0424  D3000
	0425  D300S
	0428  D7000
	0429  D5100
	042a  D800 (ptp)
	0f03  PD-10 Wireless Printer Adapter
	4000  Coolscan LS 40 ED
	4001  LS 50 ED/Coolscan V ED
	4002  Super Coolscan LS-5000 ED
04b1  Pan International
04b3  IBM Corp.
	3003  Rapid Access III Keyboard
	3004  Media Access Pro Keyboard
	300a  Rapid Access IIIe Keyboard
	3016  UltraNav Keyboard Hub
	3018  UltraNav Keyboard
	301a  2-port low-power hub
	301b  SK-8815 Keyboard
	301c  Enhanced Performance Keyboard
	3020  Enhanced Performance Keyboard
	3025  NetVista Full Width Keyboard
	3100  NetVista Mouse
	3103  ScrollPoint Pro Mouse
	3104  ScrollPoint Wireless Mouse
	3105  ScrollPoint Optical (HID)
	3107  ThinkPad 800dpi Optical Travel Mouse
	3108  800dpi Optical Mouse w/ Scroll Point
	3109  Optical ScrollPoint Pro Mouse
	310b  Red Wheel Mouse
	310c  Wheel Mouse
	4427  Portable CD ROM
	4482  Serial Converter
	4484  SMSC USB20H04 3-Port Hub [ThinkPad X4 UltraBase, Wistron S Note-3 Media Slice]
	4485  ThinkPad Dock Hub
	4524  40 Character Vacuum Fluorescent Display
	4525  Double sided CRT
	4535  4610 Suremark Printer
	4550  NVRAM (128 KB)
	4554  Cash Drawer
	4580  Hub w/ NVRAM
	4581  4800-2xx Hub w/ Cash Drawer
	4604  Keyboard w/ Card Reader
	4671  4820 LCD w/ MSR/KB
04b4  Cypress Semiconductor Corp.
	0001  Mouse
	0002  CY7C63x0x Thermometer
	0033  Mouse
	0060  Wireless optical mouse
	0100  Cino FuzzyScan F760-B
	0101  Keyboard/Hub
	0102  Keyboard with APM
	0130  MyIRC Remote Receiver
	0306  Telephone Receiver
	0407  Optical Skype Mouse
	0bad  MetaGeek Wi-Spy
	1002  CY7C63001 R100 FM Radio
	1006  Human Interface Device
	2050  hub
	2830  Opera1 DVB-S (cold state)
	3813  NANO BIOS Programmer
	4235  Monitor 02 Driver
	4381  SCAPS USC-1 Scanner Controller
	4611  Storage Adapter FX2 (CY)
	4616  Flash Disk (TPP)
	5201  Combi Keyboard-Hub (Hub)
	5202  Combi Keyboard-Hub (Keyboard)
	5500  HID->COM RS232 Adapter
	5a9b  Dacal CD/DVD Library D-101/DC-300/DC-016RW
	6370  ViewMate Desktop Mouse CC2201
	6560  CY7C65640 USB-2.0 "TetraHub"
	6830  CY7C68300A EZ-USB AT2 USB 2.0 to ATA/ATAPI
	6831  Storage Adapter ISD-300LP (CY)
	7417  Wireless PC Lock/Ultra Mouse
	8329  USB To keyboard/Mouse Converter
	8613  CY7C68013 EZ-USB FX2 USB 2.0 Development Kit
	8614  DTV-DVB UDST7020BDA DVB-S Box(DVBS for MCE2005)
	861f  Anysee E30 USB 2.0 DVB-T Receiver
	bca1  Barcode Reader
	cc04  Centor USB RACIA-ALVAR USB PORT
	cc06  Centor-P RACIA-ALVAR USB PORT
	d5d5  CY7C63x0x Zoltrix Z-Boxer GamePad
	de61  Barcode Reader
	de64  Barcode Reader
	f000  CY30700 Licorice evaluation board
	f111  CY8CKIT-002 PSoC MiniProg3 Rev A Program and debug kit
	f115  PSoC FirstTouch Programmer
	fd13  Programmable power socket
04b5  ROHM LSI Systems USA, LLC
	3064  Hantek DSO-3064
04b6  Hint Corp.
04b7  Compal Electronics, Inc.
04b8  Seiko Epson Corp.
	0001  Stylus Color 740 / Photo 750
	0002  ISD Smart Cable for Mac
	0003  ISD Smart Cable
	0004  Printer
	0005  Printer
	0006  Printer
	0007  Printer
	0015  Stylus Photo R3000
	0101  GT-7000U [Perfection 636]
	0102  GT-2200
	0103  GT-6600U [Perfection 610]
	0104  GT-7600UF [Perfection 1200U/1200U Photo]
	0105  Stylus Scan 2000
	0106  Stylus Scan 2500
	0107  ES-2000 [Expression 1600U]
	0108  CC-700
	0109  ES-8500 [Expression 1640 XL]
	010a  GT-8700/GT-8700F [Perfection 1640SU/1640SU PHOTO]
	010b  GT-7700U [Perfection 1240U]
	010c  GT-6700U [Perfection 640]
	010d  CC-500L
	010e  ES-2200 [Perfection 1680]
	010f  GT-7200U [Perfection 1250/1250 PHOTO]
	0110  GT-8200U/GT-8200UF [Perfection 1650/1650 PHOTO]
	0112  GT-9700F [Perfection 2450 PHOTO]
	0114  Perfection 660
	0116  GT-9400UF [Perfection 3170]
	0118  GT-F600 [Perfection 4180]
	0119  GT-X750 [Perfection 4490 Photo]
	011a  CC-550L [1000 ICS]
	011b  GT-9300UF [Perfection 2400 PHOTO]
	011c  GT-9800F [Perfection 3200]
	011d  GT-7300U [Perfection 1260/1260 PHOTO]
	011e  GT-8300UF [Perfection 1660 PHOTO]
	011f  GT-8400UF [Perfection 1670/1670 PHOTO]
	0120  GT-7400U [Perfection 1270]
	0121  GT-F500/GT-F550 [Perfection 2480/2580 PHOTO]
	0122  GT-F520/GT-F570 [Perfection 3590 PHOTO]
	0126  ES-7000H [GT-15000]
	0128  GT-X700 [Perfection 4870]
	0129  ES-10000G [Expression 10000XL]
	012a  GT-X800 [Perfection 4990 PHOTO]
	012b  ES-H300 [GT-2500]
	012c  GT-X900 [Perfection V700/V750 Photo]
	012d  GT-F650 [GT-S600/Perfection V10/V100]
	012e  GT-F670 [Perfection V200 Photo]
	012f  GT-F700 [Perfection V350]
	0130  GT-X770 [Perfection V500]
	0131  GT-F720 [GT-S620/Perfection V30/V300 Photo]
	0133  GT-1500 [GT-D1000]
	0135  GT-X970
	0136  ES-D400 [GT-S80]
	0137  ES-D200 [GT-S50]
	0138  ES-H7200 [GT-20000]
	013a  GT-X820 [Perfection V600 Photo]
	0142  GT-F730 [GT-S630/Perfection V33/V330 Photo]
	0143  GT-S55
	0144  GT-S85
	0202  Receipt Printer M129C/TM-T70
	0401  CP 800 Digital Camera
	0402  PhotoPC 850z
	0403  PhotoPC 3000z
	0509  JVC PIX-MC10
	0601  Stylus Photo 875DC Card Reader
	0602  Stylus Photo 895 Card Reader
	0801  CC-600PX [Stylus CX5200/CX5400/CX6600]
	0802  CC-570L [Stylus CX3100/CX3200]
	0803  Printer (Composite Device)
	0804  Storage Device
	0805  Stylus CX6300/CX6400
	0806  PM-A850 [Stylus Photo RX600/610]
	0807  Stylus Photo RX500/510
	0808  Stylus CX5200/CX5300/CX5400
	0809  Storage Device
	080a  F-3200
	080c  ME100 [Stylus CX1500]
	080d  Stylus CX4500/4600
	080e  PX-A550 [CX-3500/3600/3650 MFP]
	080f  Stylus Photo RX420/RX425/RX430
	0810  PM-A900 [Stylus Photo RX700]
	0811  PM-A870 [Stylus Photo RX620/RX630]
	0812  MFP Composite Device
	0813  Stylus CX6500/6600
	0814  PM-A700
	0815  LP-A500 [AcuLaser CX1]
	0816  Printer (Composite Device)
	0817  LP-M5500/LP-M5500F
	0818  Stylus CX3700/CX3800/DX3800
	0819  PX-A650 [Stylus CX4700/CX4800/DX4800/DX4850]
	081a  PM-A750 [Stylus Photo RX520/RX530]
	081b  MFP Composite Device
	081c  PM-A890 [Stylus Photo RX640/RX650]
	081d  PM-A950
	081e  MFP Composite Device
	081f  Stylus CX7700/7800
	0820  Stylus CX4100/CX4200/DX4200
	0821  Stylus CX5700F/CX5800F
	0822  Storage Device
	0823  MFP Composite Device
	0824  Storage Device
	0825  MFP Composite Device
	0826  Storage Device
	0827  PM-A820 [Stylus Photo RX560/RX580/RX585/RX590]
	0828  PM-A970
	0829  PM-T990
	082a  PM-A920
	082b  Stylus CX5900/CX5000/DX5000/DX5050
	082c  Storage Device
	082d  Storage Device
	082e  PX-A720 [Stylus CX5900/CX6000/DX6000]
	082f  PX-A620 [Stylus CX3900/DX4000/DX4050]
	0830  ME 200 [Stylus CX2800/CX2900]
	0831  Stylus CX6900F/CX7000F/DX7000F
	0832  MFP Composite Device
	0833  LP-M5600
	0834  LP-M6000
	0835  AcuLaser CX21
	0836  PM-T960
	0837  PM-A940 [Stylus Photo RX680/RX685/RX690]
	0838  PX-A640 [CX7300/CX7400/DX7400]
	0839  PX-A740 [CX8300/CX8400/DX8400]
	083a  PX-FA700 [CX9300F/CX9400Fax/DX9400F]
	083b  MFP Composite Device
	083c  PM-A840S [Stylus Photo RX595/RX610]
	083d  MFP Composite Device
	083e  MFP Composite Device
	083f  Stylus CX4300/CX4400/CX5500/CX5600/DX4400/DX4450
	0841  PX-401A [ME 300/Stylus NX100]
	0843  LP-M5000
	0844  EP-901A/EP-901F [Artisan 800/Stylus Photo PX800FW]
	0846  EP-801A [Artisan 700/Stylus Photo PX700W/TX700W]
	0847  PX-601F [ME Office 700FW/Stylus Office BX600FW/TX600FW]
	0848  ME Office 600F/Stylus Office BX300F/TX300F
	0849  Stylus SX205
	084a  PX-501A [Stylus NX400]
	084d  PX-402A [Stylus SX115/Stylus NX110 Series]
	084f  ME OFFICE 510
	0850  EP-702A [Stylus Photo PX650/TX650 Series]
	0851  Stylus SX410
	0852  EP-802A [Artisan 710 Series/Stylus Photo PX710W/TX720W Series]
	0853  EP-902A [Artisan 810 Series/Stylus Photo PX810FW Series]
	0854  ME OFFICE 650FN Series/Stylus Office BX310FN/TX520FN Series
	0855  PX-602F [Stylus Office BX610FW/TX620FW Series]
	0856  PX-502A [Stylus SX515W]
	085c  ME 320/330 Series [Stylus SX125]
	085d  PX-603F [ME OFFICE 960FWD Series/Stylus Office BX625FWD/TX620FWD Series]
	085e  PX-503A [ME OFFICE 900WD Series/Stylus Office BX525WD]
	085f  Stylus Office BX320FW/TX525FW Series
	0860  EP-903A/EP-903F [Artisan 835/Stylus Photo PX820FWD Series]
	0861  EP-803A/EP-803AW [Artisan 725/Stylus Photo PX720WD/TX720WD Series]
	0862  EP-703A [Stylus Photo PX660 Series]
	0863  ME OFFICE 620F Series/Stylus Office BX305F/BX305FW/TX320F
	0864  ME OFFICE 560W Series
	0865  ME OFFICE 520 Series
	0866  AcuLaser MX20DN/MX20DNF/MX21DNF
	0869  PX-1600F
	086a  PX-673F [Stylus Office BX925FWD]
	0870  Stylus Office BX305FW Plus
	0871  K200 Series
	0872  K300 Series
	0873  L200 Series
	0878  EP-704A
	0879  EP-904A/EP-904F [Artisan 837/Stylus Photo PX830FWD Series]
	087b  EP-804A/EP-804AR/EP-804AW [Stylus Photo PX730WD/Artisan 730 Series]
	087c  PX-1700F
	087d  PX-B750F/WP-4525 Series
	087f  PX-403A
	0880  PX-434A [Stylus NX330 Series]
	0881  PX-404A [ME OFFICE 535]
	0883  ME 340 Series/Stylus NX130 Series
	0884  Stylus NX430W Series
	0885  Stylus NX230/SX235W Series
	088f  Stylus Office BX635FWD
	0890  ME OFFICE 940FW Series/Stylus Office BX630FW Series
	0891  Stylus Office BX535WD
	0892  Stylus Office BX935FWD
	0893  EP-774A
04b9  Rainbow Technologies, Inc.
	0300  SafeNet USB SuperPro/UltraPro
	1000  iKey 1000 Token
	1001  iKey 1200 Token
	1002  iKey Token
	1003  iKey Token
	1004  iKey Token
	1005  iKey Token
	1006  iKey Token
	1200  iKey 2000 Token
	1201  iKey Token
	1202  iKey 2032 Token
	1203  iKey Token
	1204  iKey Token
	1205  iKey Token
	1206  iKey 4000 Token
	1300  iKey 3000 Token
	1301  iKey 3000
	1302  iKey Token
	1303  iKey Token
	1304  iKey Token
	1305  iKey Token
	1306  iKey Token
04ba  Toucan Systems, Ltd
04bb  I-O Data Device, Inc.
	0101  USB2-IDE/ATAPI Bridge Adapter
	0201  USB2-IDE/ATAPI Bridge Adapter
	0204  DVD Multi-plus unit iU-CD2
	0206  DVD Multi-plus unit DVR-UEH8
	0301  Storage Device
	0314  USB-SSMRW SD-card
	0319  USB2-IDE/ATAPI Bridge Adapter
	031a  USB2-IDE/ATAPI Bridge Adapter
	031b  USB2-IDE/ATAPI Bridge Adapter
	031e  USB-SDRW SD-card
	0502  Nogatech Live! (BT)
	0528  GV-USB Video Capture
	0901  USB ETT
	0904  ET/TX Ethernet [pegasus]
	0913  ET/TX-S Ethernet [pegasus2]
	0919  USB WN-B11
	0922  IOData AirPort WN-B11/USBS 802.11b
	0930  ETG-US2
	0937  WN-WAG/USL Wireless LAN Adapter
	0938  WN-G54/USL Wireless LAN Adapter
	093b  WN-GDN/USB
	093f  WNGDNUS2 802.11n
	0944  WHG-AGDN/US Wireless LAN Adapter
	0945  WN-GDN/US3 Wireless LAN Adapter
	0947  WN-G150U Wireless LAN Adapter
	0948  WN-G300U Wireless LAN Adapter
	0a03  Serial USB-RSAQ1
	0a07  USB2-iCN Adapter
	0a08  USB2-iCN Adapter
	0c01  FM-10 Pro Disk
04bd  Toshiba Electronics Taiwan Corp.
04be  Telia Research AB
04bf  TDK Corp.
	0100  MediaReader CF
	0115  USB-PDC Adapter UPA9664
	0116  USB-cdmaOne Adapter UCA1464
	0117  USB-PHS Adapter UHA6400
	0118  USB-PHS Adapter UPA6400
	0135  MediaReader Dual
	0202  73S1121F Smart Card Reader-
	0309  Bluetooth USB dongle
	030a  IBM Bluetooth Ultraport Module
	030b  Bluetooth Device
	030c  Ultraport Bluetooth Device
	0310  Integrated Bluetooth
	0311  Integrated Bluetooth Device
	0317  Bluetooth UltraPort Module from IBM
	0318  IBM Integrated Bluetooth
	0319  Bluetooth Adapter
	0320  Bluetooth Adapter
	0321  Bluetooth Device
	0a28  INDI AV-IN Device
04c1  U.S. Robotics (3Com)
	0020  56K Voice Pro
	0022  56K Voice Pro
	007e  ISDN TA
	0082  OfficeConnect Analog Modem
	008f  Pro ISDN TA
	0097  OfficeConnect Analog
	009d  HomeConnect Webcam [vicam]
	00a9  ISDN Pro TA-U
	00b9  HomeConnect IDSL Modem
	3021  56k Voice FaxModem Pro
04c2  Methode Electronics Far East PTE, Ltd
04c3  Maxi Switch, Inc.
	1102  Mouse
	2102  Mouse
04c4  Lockheed Martin Energy Research
04c5  Fujitsu, Ltd
	1029  fi-4010c Scanner
	1033  fi-4110CU
	1041  fi-4120c Scanner
	1042  fi-4220c Scanner
	105b  AH-F401U Air H device
	1084  PalmSecure Sensor V2
	1096  fi-5110EOX
	1097  fi-5110C
	10ae  fi-4120C2
	10af  fi-4220C2
	10c7  fi-60f scanner
	10e0  fi-5120c Scanner
	10e1  fi-5220C
	10e7  fi-5900C
	10fe  S500
	1150  fi-6230
	201d  SATA 3.0 6Gbit/s Adaptor [GROOVY]
04c6  Toshiba America Electronic Components
04c7  Micro Macro Technologies
04c8  Konica Corp.
	0720  Digital Color Camera
	0721  e-miniD Camera
	0722  e-mini
	0723  KD-200Z Camera
	0726  KD-310Z Camera
	0728  Revio C2 Mass Storage Device
	0729  Revio C2 Digital Camera
	072c  Revio KD20M
	072d  Revio KD410Z
04ca  Lite-On Technology Corp.
	004f  SK-9020 keyboard
	1766  HID Monitor Controls
	2004  Bluetooth 4.0 [Broadcom BCM20702A0]
	2006  Broadcom BCM43142A0 Bluetooth Device
	300b  Atheros AR3012 Bluetooth
	300d  Atheros AR3012 Bluetooth
	300f  Atheros AR3012 Bluetooth
	3014  Qualcoom Atheros Bluetooth
	7025  HP HD Webcam
	7046  TOSHIBA Web Camera - HD
	9304  Hub
	f01c  TT1280DA DVB-T TV Tuner
04cb  Fuji Photo Film Co., Ltd
	0100  FinePix 30i/40i/50i, A101/201, 1300/2200, 1400/2400/2600/2800/4500/4700/4800/4900/6800/6900 Zoom
	0103  FinePix NX-500/NX-700 printer
	0104  FinePix A101, 2600/2800/4800/6800 Zoom (PC CAM)
	0108  FinePix F601 Zoom (DSC)
	0109  FinePix F601 Zoom (PC CAM)
	010a  FinePix S602 (Pro) Zoom (DSC)
	010b  FinePix S602 (Pro) Zoom (PC CAM)
	010d  FinePix Digital Camera 020531
	010e  FinePix F402 Zoom (DSC)
	010f  FinePix F402 Zoom (PC CAM)
	0110  FinePix M603 Zoom (DSC)
	0111  FinePix M603 Zoom (PC CAM)
	0112  FinePix A202, A200 Zoom (DSC)
	0113  FinePix A202, A200 Zoom (PC CAM)
	0114  FinePix F401 Zoom (DSC)
	0115  FinePix F401 Zoom (PC CAM)
	0116  FinePix A203 Zoom (DSC)
	0117  FinePix A203 Zoom (PC CAM)
	0118  FinePix A303 Zoom (DSC)
	0119  FinePix A303 Zoom (PC CAM)
	011a  FinePix S304/3800 Zoom (DSC)
	011b  FinePix S304/3800 Zoom (PC CAM)
	011c  FinePix A204/2650 Zoom (DSC)
	011d  FinePix A204/2650 Zoom (PC CAM)
	0120  FinePix F700 Zoom (DSC)
	0121  FinePix F700 Zoom (PC CAM)
	0122  FinePix F410 Zoom (DSC)
	0123  FinePix F410 Zoom (PC CAM)
	0124  FinePix A310 Zoom (DSC)
	0125  FinePix A310 Zoom (PC CAM)
	0126  FinePix A210 Zoom (DSC)
	0127  FinePix A210 Zoom (PC CAM)
	0128  FinePix A205(S) Zoom (DSC)
	0129  FinePix A205(S) Zoom (PC CAM)
	012a  FinePix F610 Zoom (DSC)
	012b  FinePix Digital Camera 030513
	012c  FinePix S7000 Zoom (DSC)
	012d  FinePix S7000 Zoom (PC CAM)
	012f  FinePix Digital Camera 030731
	0130  FinePix S5000 Zoom (DSC)
	0131  FinePix S5000 Zoom (PC CAM)
	013b  FinePix Digital Camera 030722
	013c  FinePix S3000 Zoom (DSC)
	013d  FinePix S3000 Zoom (PC CAM)
	013e  FinePix F420 Zoom (DSC)
	013f  FinePix F420 Zoom (PC CAM)
	0142  FinePix S7000 Zoom (PTP)
	0148  FinePix A330 Zoom (DSC)
	0149  FinePix A330 Zoom (UVC)
	014a  FinePix A330 Zoom (PTP)
	014b  FinePix A340 Zoom (DSC)
	014c  FinePix A340 Zoom (UVC)
	0159  FinePix F710 Zoom (DSC)
	0165  FinePix S3500 Zoom (DSC)
	0168  FinePix E500 Zoom (DSC)
	0169  FinePix E500 Zoom (UVC)
	016b  FinePix E510 Zoom (DSC)
	016c  FinePix E510 Zoom (PC CAM)
	016e  FinePix S5500 Zoom (DSC)
	016f  FinePix S5500 Zoom (UVC)
	0171  FinePix E550 Zoom (DSC)
	0172  FinePix E550 Zoom (UVC)
	0177  FinePix F10 (DSC)
	0179  Finepix F10 (PTP)
	0186  FinePix S5200/S5600 Zoom (DSC)
	0188  FinePix S5200/S5600 Zoom (PTP)
	018e  FinePix S9500 Zoom (DSC)
	018f  FinePix S9500 Zoom (PTP)
	0192  FinePix E900 Zoom (DSC)
	0193  FinePix E900 Zoom (PTP)
	019b  FinePix F30 (PTP)
	01af  FinePix A700 (PTP)
	01bf  FinePix F6000fd/S6500fd Zoom (PTP)
	01c0  FinePix F20 (PTP)
	01c1  FinePix F31fd (PTP)
	01c4  FinePix S5700 Zoom (PTP)
	01c5  FinePix F40fd (PTP)
	01c6  FinePix A820 Zoom (PTP)
	01d2  FinePix A800 Zoom (PTP)
	01d3  FinePix A920 (PTP)
	01d4  FinePix F50fd (PTP)
	01d5  FinePix F47 (PTP)
	01f7  FinePix J250 (PTP)
	01fd  A160
	023e  FinePix AX300
	0240  FinePix S2950 Digital Camera
	0241  FinePix S3200 Digital Camera
	0278  FinePix JV300
04cc  ST-Ericsson
	1122  Hub
	1520  USB 2.0 Hub (Avocent KVM)
	1521  USB 2.0 Hub
	1a62  GW Instek GSP-830 Spectrum Analyzer (HID)
	2323  Ux500 serial debug port
	2533  NFC device (PN533)
	8116  Camera
04cd  Tatung Co. Of America
04ce  ScanLogic Corp.
	0002  SL11R-IDE IDE Bridge
	0100  USB2PRN Printer Class
	0300  Phantom 336CX - C3 scanner
	04ce  SL11DEMO, VID: 0x4ce, PID: 0x4ce
	07d1  SL11R, VID: 0x4ce, PID: 0x07D1
04cf  Myson Century, Inc.
	0022  OCZ Alchemy Series Elixir II Keyboard
	0800  MTP800 Mass Storage Device
	8810  CS8810 Mass Storage Device
	8811  CS8811 Mass Storage Device
	8813  CS8813 Mass Storage Device
	8818  USB2.0 to ATAPI Bridge Controller
	8819  USB 2.0 SD/MMC Reader
	9920  CS8819A2-114 Mass Storage Device
04d0  Digi International
04d1  ITT Canon
04d2  Altec Lansing Technologies
	0070  ADA70 Speakers
	0305  Non-Compliant Audio Device
	0311  ADA-310 Speakers
	2060  Claritel-i750 - vp
	ff05  ADA-305 Speakers
	ff47  Lansing HID Audio Controls
	ff49  Lansing HID Audio Controls
04d3  VidUS, Inc.
04d4  LSI Logic, Inc.
04d5  Forte Technologies, Inc.
04d6  Mentor Graphics
04d7  Oki Semiconductor
	1be4  Bluetooth Device
04d8  Microchip Technology, Inc.
	0002  PicoLCD 20x2
	0003  PICkit 2 Microcontroller Programmer
	000a  CDC RS-232 Emulation Demo
	000b  PIC18F2550 (32K Flashable 10 Channel, 10 Bit A/D USB Microcontroller)
	0032  PICkit1
	0033  PICkit2
	0036  PICkit Serial Analyzer
	00e0  PIC32 Starter Board
	04cd  28Cxxx EEPROM Programmer
	0a04  AGP LIN Serial Analyzer
	8000  In-Circuit Debugger
	8001  ICD2 in-circuit debugger
	8101  PIC24F Starter Kit
	8107  Microstick II
	8108  ChipKit Pro MX7 (PIC32MX)
	9004  Microchip REAL ICE
	900a  PICkit3
	c001  PicoLCD 20x4
	e11c  TL866CS EEPROM Programmer [MiniPRO]
	f2c4  Macareux-labs Hygrometry Temperature Sensor
	f3aa  Macareux-labs Usbce Bootloader mode
	f437  SBE Tech Ultrasonic Anemometer
	f4b5  SmartScope
	f8da  Hughski Ltd. ColorHug
	f8e8  Harmony 300/350 Remote
	f91c  SPROG IIv3
	faff  Dangerous Prototypes BusPirate v4 Bootloader mode
	fb00  Dangerous Prototypes BusPirate v4
	fbb2  GCUSB-nStep stepper motor controller
	fbba  DiscFerret Magnetic Disc Analyser (bootloader mode)
	fbbb  DiscFerret Magnetic Disc Analyser (active mode)
	fc1e  Bachrus Speedometer Interface
	fc92  Open Bench Logic Sniffer
	ffee  Devantech USB-ISS
	ffef  PICoPLC [APStech]
04d9  Holtek Semiconductor, Inc.
	0022  Portable Keyboard
	048e  Optical Mouse
	0499  Optical Mouse
	1203  Keyboard
	1400  PS/2 keyboard + mouse controller
	1503  Keyboard
	1603  Keyboard
	1702  Keyboard LKS02
	2011  Keyboard [Diatec Filco Majestouch 1]
	2013  Keyboard [Das Keyboard]
	2221  Keyboard
	2323  Keyboard
	2519  Shenzhen LogoTech 2.4GHz receiver
	2832  HT82A832R Audio MCU
	2834  HT82A834R Audio MCU
	a055  Keyboard
04da  Panasonic (Matsushita)
	0901  LS-120 Camera
	0912  SDR-S10
	0b01  CD-R/RW Drive
	0b03  SuperDisk 240MB
	0d01  CD-R Drive KXL-840AN
	0d09  CD-R Drive KXL-RW32AN
	0d0a  CD-R Drive KXL-CB20AN
	0d0d  CDRCB03
	0d0e  DVD-ROM & CD-R/RW
	0f40  Printer
	104d  Elite Panaboard UB-T880 (HID)
	104e  Elite Panaboard Pen Adaptor (HID)
	1500  MFSUSB Driver
	1800  DY-WL10 802.11abgn Adapter [Broadcom BCM4323]
	1b00  MultiMediaCard
	2121  EB-VS6
	2316  DVC Mass Storage Device
	2317  DVC USB-SERIAL Driver for WinXP
	2318  NV-GS11/230/250 (webcam mode)
	2319  NV-GS15 (webcam mode)
	231a  NV-GS11/230/250 (DV mode)
	231d  DVC Web Camera Device
	231e  DVC DV Stream Device
	2372  Lumix Camera (Storage mode)
	2374  Lumix Camera (PTP mode)
	2451  HDC-SD9
	245b  HC-X920K (3MOS Full HD video camcorder)
	2497  HDC-TM700
	250c  Gobi Wireless Modem (QDL mode)
	250d  Gobi Wireless Modem
	3904  N5HBZ0000055 802.11abgn Wireless Adapter [Atheros AR7010+AR9280]
	3c04  JT-P100MR-20 [ePassport Reader]
04db  Hypertec Pty, Ltd
04dc  Huan Hsin Holdings, Ltd
04dd  Sharp Corp.
	13a6  MFC2000
	6006  AL-1216
	6007  AL-1045
	6008  AL-1255
	6009  AL-1530CS
	600a  AL-1540CS
	600b  AL-1456
	600c  AL-1555
	600d  AL-1225
	600e  AL-1551CS
	600f  AR-122E
	6010  AR-152E
	6011  AR-157E
	6012  SN-1045
	6013  SN-1255
	6014  SN-1456
	6015  SN-1555
	6016  AR-153E
	6017  AR-122E N
	6018  AR-153E N
	6019  AR-152E N
	601a  AR-157E N
	601b  AL-1217
	601c  AL-1226
	601d  AR-123E
	6021  IS01
	7002  DVC Ver.1.0
	7004  VE-CG40U Digital Still Camera
	7005  VE-CG30 Digital Still Camera
	7007  VL-Z7S Digital Camcorder
	8004  Zaurus SL-5000D/SL-5500 PDA
	8005  Zaurus A-300
	8006  Zaurus SL-B500/SL-5600 PDA
	8007  Zaurus C-700 PDA
	9009  AR-M160
	9014  IM-DR80 Portable NetMD Player
	9031  Zaurus C-750/C-760/C-860/SL-C3000 PDA
	9032  Zaurus SL-6000
	903a  GSM GPRS
	9050  Zaurus C-860 PDA
	9056  Viewcam Z
	9073  AM-900
	9074  GSM GPRS
	90a9  Sharp Composite
	90d0  USB-to-Serial Comm. Port
	90f2  Sharp 3G GSM USB Control
	9120  WS004SH
	9122  WS007SH
	9123  W-ZERO3 ES Smartphone
	91a3  922SH Internet Machine
	939a  IS03
04de  MindShare, Inc.
04df  Interlink Electronics
04e1  Iiyama North America, Inc.
	0201  Monitor Hub
04e2  Exar Corp.
	1410  XR21V1410 USB-UART IC
04e3  Zilog, Inc.
04e4  ACC Microelectronics
04e5  Promise Technology
04e6  SCM Microsystems, Inc.
	0001  E-USB ATA Bridge
	0002  eUSCSI SCSI Bridge
	0003  eUSB SmartMedia Card Reader
	0005  eUSB SmartMedia/CompactFlash Card Reader
	0006  eUSB SmartMedia Card Reader
	0007  Hifd
	0009  eUSB ATA/ATAPI Adapter
	000a  eUSB CompactFlash Adapter
	000b  eUSCSI Bridge
	000c  eUSCSI Bridge
	000d  Dazzle MS
	0012  Dazzle SD/MMC
	0101  eUSB ATA Bridge (Sony Spressa USB CDRW)
	0311  Dazzle DM-CF
	0312  Dazzle DM-SD/MMC
	0313  Dazzle SM
	0314  Dazzle MS
	0322  e-Film Reader-5
	0325  eUSB ORCA Quad Reader
	0327  Digital Media Reader
	03fe  DMHS2 DFU Adapter
	0406  eUSB SmartDM Reader
	04e6  eUSB DFU Adapter
	04e7  STCII DFU Adapter
	04e8  eUSBDM DFU Adapter
	04e9  DM-E DFU Adapter
	0500  Veridicom 5thSense Fingerprint Sensor and eUSB SmartCard
	0701  DCS200 Loader Device
	0702  DVD Creation Station 200
	0703  DVC100 Loader Device
	0704  Digital Video Creator 100
	1001  SCR300 Smart Card Reader
	1010  USBAT-2 CompactFlash Card Reader
	1014  e-Film Reader-3
	1020  USBAT ATA/ATAPI Adapter
	2007  RSA SecurID ComboReader
	2009  Citibank Smart Card Reader
	200a  Reflex v.2 Smart Card Reader
	200d  STR391 Reader
	5111  SCR331-DI SmartCard Reader
	5113  SCR333 SmartCard Reader
	5114  SCR331-DI SmartCard Reader
	5115  SCR335 SmartCard Reader
	5116  SCR331-LC1 / SCR3310 SmartCard Reader
	5117  SCR3320 - Smart Card Reader
	5118  Expresscard SIM Card Reader
	5119  SCR3340 - ExpressCard54 Smart Card Reader
	511b  SmartCard Reader
	511d  SCR3311 Smart Card Reader
	5120  SCR331-DI SmartCard Reader
	5121  SDI010 Smart Card Reader
	5151  SCR338 Keyboard Smart Card Reader
	5292  SCL011 RFID reader
	5410  SCR35xx Smart Card Reader
	5591  SCL3711-NFC&RW
	e000  SCRx31 Reader
	e001  SCR331 SmartCard Reader
	e003  SPR532 PinPad SmartCard Reader
04e7  Elo TouchSystems
	0001  TouchScreen
	0002  Touchmonitor Interface 2600 Rev 2
	0004  4000U CarrollTouch Touchmonitor Interface
	0007  2500U IntelliTouch Touchmonitor Interface
	0008  3000U AccuTouch Touchmonitor Interface
	0009  4000U CarrollTouch Touchmonitor Interface
	0020  Touchscreen Interface (2700)
	0021  Touchmonitor Interface
	0030  4500U CarrollTouch Touchmonitor Interface
	0032  Touchmonitor Interface
	0033  Touchmonitor Interface
	0041  5010 Surface Capacitive Touchmonitor Interface
	0042  Touchmonitor Interface
	0050  2216 AccuTouch Touchmonitor Interface
	0071  Touchmonitor Interface
	0072  Touchmonitor Interface
	0081  Touchmonitor Interface
	0082  Touchmonitor Interface
	00ff  Touchmonitor Interface
04e8  Samsung Electronics Co., Ltd
	0100  Kingston Flash Drive (128MB)
	0110  Connect3D Flash Drive
	0111  Connect3D Flash Drive
	0300  E2530 / GT-C3350 Phones (Mass storage mode)
	1003  MP3 Player and Recorder
	1006  SDC-200Z
	130c  NX100
	1f05  S2 Portable [JMicron] (500GB)
	1f06  HX-MU064DA portable harddisk
	2018  WIS09ABGN LinkStick Wireless LAN Adapter
	2035  Digital Photo Frame Mass Storage
	2036  Digital Photo Frame Mini Monitor
	3004  ML-4600
	3005  Docuprint P1210
	3008  ML-6060 laser printer
	300c  ML-1210 Printer
	300e  Laser Printer
	3104  ML-3550N
	3210  ML-5200A Laser Printer
	3226  Laser Printer
	3228  Laser Printer
	322a  Laser Printer
	322c  Laser Printer
	3230  ML-1440
	3232  Laser Printer
	3236  ML-1450
	3238  ML-1430
	323a  ML-1710 Printer
	323b  Phaser 3130
	323c  Laser Printer
	323d  Phaser 3120
	323e  Laser Printer
	3240  Laser Printer
	3242  ML-1510 Laser Printer
	3248  Color Laser Printer
	324a  Laser Printer
	324c  ML-1740 Printer
	324d  Phaser 3121
	3256  ML-1520 Laser Printer
	325b  Xerox Phaser 3117 Laser Printer
	325f  Phaser 3425 Laser Printer
	3260  CLP-510 Color Laser Printer
	3268  ML-1610 Mono Laser Printer
	326c  ML-2010P Mono Laser Printer
	3276  ML-3050/ML-3051 Laser Printer
	328e  CLP-310 Color Laser Printer
	3292  ML-1640 Series Laser Printer
	3296  ML-2580N Mono Laser Printer
	3297  ML-191x/ML-252x Laser Printer
	329f  CLP-325 Color Laser Printer
	3301  ML-1660 Series
	330c  ML-1865
	3310  ML-331x Series Laser Printer
	3315  ML-2540 Series Laser Printer
	331e  M262x/M282x Xpress Series Laser Printer
	3409  SCX-4216F Scanner
	340c  SCX-5x15 series
	340d  SCX-6x20 series
	340e  MFP 560 series
	340f  Printing Support
	3412  SCX-4x20 series
	3413  SCX-4100 Scanner
	3415  Composite Device
	3419  Composite Device
	341a  Printing Support
	341b  SCX-4200 series
	341c  Composite Device
	341d  Composite Device
	341f  Composite Device
	3420  Composite Device
	3426  SCX-4500 Laser Printer
	342d  SCX-4x28 Series
	344f  SCX-3400 Series
	3605  InkJet Color Printer
	3606  InkJet Color Printer
	3609  InkJet Color Printer
	3902  InkJet Color Printer
	3903  Xerox WorkCentre XK50cx
	390f  InkJet Color Printer
	3911  SCX-1020 series
	4005  GT-S8000 Jet (msc)
	4f1f  GT-S8000 Jet (mtp)
	5000  YP-MF series
	5001  YP-100
	5002  YP-30
	5003  YP-700
	5004  YP-30
	5005  YP-300
	5006  YP-750
	500d  MP3 Player
	5010  Yepp YP-35
	5011  YP-780
	5013  YP-60
	5015  yepp upgrade
	501b  MP3 Player
	5021  Yepp YP-ST5
	5026  YP-MT6V
	5027  YP-T7
	502b  YP-F1
	5032  YP-J70
	503b  YP-U1 MP3 Player
	503d  YP-T7F
	5041  YP-Z5
	5050  YP-U2 MP3 Player
	5051  YP-F2R
	5055  YP-T9
	507d  YP-U3 (mtp)
	507f  YP-T9J
	5080  Yepp YP-K3 (msc)
	5081  Yepp YP-K3 (mtp)
	5082  YP-P2 (msc)
	5083  YP-P2 (mtp)
	508a  YP-T10
	508b  YP-S5 MP3 Player
	508c  YP-S5
	5090  YP-S3 (msc)
	5091  YP-S3 (mtp)
	5092  YP-U4 (msc)
	5093  YP-U4 (mtp)
	5095  YP-S2
	510f  YP-R1
	5119  Yepp YP-P3
	511c  YP-Q2
	5121  YP-U5
	5123  Yepp YP-M1
	5a00  YP-NEU
	5a01  YP-NDU
	5a03  Yepp MP3 Player
	5a04  YP-800
	5a08  YP-90
	5a0f  Meizu M6 MiniPlayer
	5b01  Memory Stick Reader/Writer
	5b02  Memory Stick Reader/Writer
	5b03  Memory Stick Reader/Writer
	5b04  Memory Stick Reader/Writer
	5b05  Memory Stick Reader/Writer
	5b11  SEW-2001u Card
	5f00  NEXiO Sync
	5f01  NEXiO Sync
	5f02  NEXiO Sync
	5f03  NEXiO Sync
	5f04  NEXiO Sync
	5f05  STORY Station 1TB
	6032  G2 Portable hard drive
	6033  G2 Portable device
	6034  G2 Portable hard drive
	60b3  M2 Portable Hard Drive
	60c4  M2 Portable Hard Drive USB 3.0
	6124  D3 Station External Hard Drive
	6125  D3 Station External Hard Drive
	61b5  M3 Portable Hard Drive 2TB
	61b6  M3 Portable Hard Drive 1TB
	6601  Mobile Phone
	6602  Galaxy
	6603  Galaxy
	6611  MITs Sync
	6613  MITs Sync
	6615  MITs Sync
	6617  MITs Sync
	6619  MITs Sync
	661b  MITs Sync
	661e  Handheld
	6620  Handheld
	6622  Handheld
	6624  Handheld
	662e  MITs Sync
	6630  MITs Sync
	6632  MITs Sync
	663e  D900e/B2100 Phone
	663f  SGH-E720/SGH-E840
	6640  Usb Modem Enumerator
	6651  i8510 Innov8
	6702  X830
	6708  U600 Phone
	6709  U600
	6734  Juke
	6759  D900e/B2100 Media Player
	675a  D900e/B2100 Mass Storage
	675b  D900e Camera
	6772  Standalone LTE device (Trial)
	6795  S5230
	6802  Standalone HSPA device
	6806  Composite LTE device (Trial)
	6807  Composite HSPA device
	681c  Galaxy Portal/Spica/S
	681d  Galaxy Portal/Spica Android Phone
	6843  E2530 Phone (Samsung Kies mode)
	684e  Wave (GT-S8500)
	685b  GT-I9100 Phone [Galaxy S II] (mass storage mode)
	685c  GT-I9250 Phone [Galaxy Nexus] (Mass storage mode)
	685d  GT-I9100 Phone [Galaxy S II] (Download mode)
	685e  GT-I9100 / GT-C3350 Phones (USB Debugging mode)
	6860  Galaxy (MTP)
	6863  GT-I9500 [Galaxy S4] / GT-I9250 [Galaxy Nexus] (network tethering)
	6864  GT-I9070 (network tethering, USB debugging enabled)
	6865  GT-I9300 Phone [Galaxy S III] (PTP mode)
	6866  GT-I9300 Phone [Galaxy S III] (debugging mode)
	6868  Escape Composite driver for Android Phones: Modem+Diagnostic+ADB
	6875  GT-B3710 Standalone LTE device (Commercial)
	6876  GT-B3710 LTE Modem
	6877  Galaxy S
	687a  GT-E2370 mobile phone
	6888  GT-B3730 Composite LTE device (Commercial)
	6889  GT-B3730 Composite LTE device (Commercial)
	689a  LTE Storage Driver [CMC2xx]
	689e  GT-S5670 [Galaxy Fit]
	68aa  Reality
	7011  SEW-2003U Card
	7021  Bluetooth Device
	7061  eHome Infrared Receiver
	7080  Anycall SCH-W580
	7081  Human Interface Device
	8001  Handheld
	e020  SERI E02 SCOM 6200 UMTS Phone
	e021  SERI E02 SCOM 6200 Virtual UARTs
	e022  SERI E02 SCOM 6200 Flash Load Disk
	f000  Intensity 3 (Mass Storage Mode)
	ff30  SG_iMON
04e9  PC-Tel, Inc.
04ea  Brooktree Corp.
04eb  Northstar Systems, Inc.
	e004  eHome Infrared Transceiver
04ec  Tokyo Electron Device, Ltd
04ed  Annabooks
04ef  Pacific Electronic International, Inc.
04f0  Daewoo Electronics Co., Ltd
04f1  Victor Company of Japan, Ltd
	0001  GC-QX3 Digital Still Camera
	0004  GR-DVL815U Digital Video Camera
	0006  DV Camera Storage
	0008  GZ-MG30AA/MC500E Digital Video Camera
	0009  GR-DX25EK Digital Video Camera
	000a  GR-D72 Digital Video Camera
	1001  GC-A50 Camera Device
	3008  MP-PRX1 Ethernet
	3009  MP-XP7250 WLAN Adapter
04f2  Chicony Electronics Co., Ltd
	0001  KU-8933 Keyboard
	0002  NT68P81 Keyboard
	0110  KU-2971 Keyboard
	0111  KU-9908 Keyboard
	0112  KU-8933 Keyboard with PS/2 Mouse port
	0116  KU-2971/KU-0325 Keyboard
	0200  KBR-0108
	0201  Gaming Keyboard KPD0250
	0220  Wireless HID Receiver
	0402  Genius LuxeMate i200 Keyboard
	0403  KU-0420 keyboard
	0418  KU-0418 Tactical Pad
	0618  RG-0618U Wireless HID Receiver & KG-0609 Wireless Keyboard with Touchpad
	0760  Acer KU-0760 Keyboard
	0841  HP Multimedia Keyboard
	0860  2.4G Multimedia Wireless Kit
	1061  HP KG-1061 Wireless Keyboard+Mouse
	1121  Periboard 717 Mini Wireless Keyboard
	a001  E-Video DC-100 Camera
	a120  ORITE CCD Webcam(PC370R)
	a121  ORITE CCD Webcam(PC370R)
	a122  ORITE CCD Webcam(PC370R)
	a123  ORITE CCD Webcam(PC370R)
	a124  ORITE CCD Webcam(PC370R)
	a128  PC Camera (SN9C202 + OV7663 + EEPROM)
	a133  Gateway Webcam
	a136  LabTec Webcam 5500
	a147  Medion Webcam
	a204  DSC WIA Device (1300)
	a208  DSC WIA Device (2320)
	a209  Labtec DC-2320
	a20a  DSC WIA Device (3310)
	a20c  DSC WIA Device (3320)
	a210  Audio Device
	b008  USB 2.0 Camera
	b009  Integrated Camera
	b010  Integrated Camera
	b012  1.3 MPixel UVC Webcam
	b013  USB 2.0 Camera
	b015  VGA 24fps UVC Webcam
	b016  VGA 30fps UVC Webcam
	b018  2M UVC Webcam
	b021  ViewSonic 1.3M, USB2.0 Webcam
	b022  Gateway USB 2.0 Webcam
	b023  Gateway USB 2.0 Webcam
	b024  USB 2.0 Webcam
	b025  Camera
	b027  Gateway USB 2.0 Webcam
	b028  VGA UVC Webcam
	b029  1.3M UVC Webcam
	b036  Asus Integrated 0.3M UVC Webcam
	b044  Acer CrystalEye Webcam
	b057  integrated USB webcam
	b059  CKF7037 HP webcam
	b064  CNA7137 Integrated Webcam
	b070  Camera
	b071  2.0M UVC Webcam / CNF7129
	b083  CKF7063 Webcam (HP)
	b091  Webcam
	b104  CNF7069 Webcam
	b107  CNF7070 Webcam
	b14c  CNF8050 Webcam
	b15c  Sony Vaio Integrated Camera
	b175  4-Port Hub
	b1aa  Webcam-101
	b1b4  Lenovo Integrated Camera
	b1b9  Asus Integrated Webcam
	b1cf  Lenovo Integrated Camera
	b1d6  CNF9055 Toshiba Webcam
	b1e4  Toshiba Integrated Webcam
	b213  Fujitsu Integrated Camera
	b217  Lenovo Integrated Camera (0.3MP)
	b221  integrated camera
	b230  Integrated HP HD Webcam
	b257  Lenovo Integrated Camera
	b26b  Sony Visual Communication Camera
	b272  Lenovo EasyCamera
	b2b0  Camera
	b2b9  Lenovo Integrated Camera UVC
	b2da  thinkpad t430s camera
	b2ea  Integrated Camera [ThinkPad]
	b330  Asus 720p CMOS webcam
	b354  UVC 1.00 device HD UVC WebCam
	b394  Integrated Camera
	b3f6  HD WebCam (Acer)
	b40e  HP Truevision HD camera
	b444  Lenovo Integrated Webcam
04f3  Elan Microelectronics Corp.
	000a  Touchscreen
	0103  ActiveJet K-2024 Multimedia Keyboard
	01a4  Wireless Keyboard
	0201  Touchscreen
	0210  Optical Mouse
	0212  Laser Mouse
	0214  Lynx M9 Optical Mouse
	0230  3D Optical Mouse
	0232  Mouse
	02f4  2.4G Cordless Mouse
	0381  Touchscreen
	04a0  Dream Cheeky Stress/Panic Button
04f4  Harting Elektronik, Inc.
04f5  Fujitsu-ICL Systems, Inc.
04f6  Norand Corp.
04f7  Newnex Technology Corp.
04f8  FuturePlus Systems
04f9  Brother Industries, Ltd
	0002  HL-1050 Laser Printer
	0005  Printer
	0006  HL-1240 Laser Printer
	0007  HL-1250 Laser Printer
	0008  HL-1270 Laser Printer
	0009  Printer
	000a  P2500 series
	000b  Printer
	000c  Printer
	000d  HL-1440 Laser Printer
	000e  HL-1450 series
	000f  HL-1470N series
	0010  Printer
	0011  Printer
	0012  Printer
	0013  Printer
	0014  Printer
	0015  Printer
	0016  Printer
	0017  Printer
	0018  Printer
	001a  HL-1430 Laser Printer
	001c  Printer
	001e  Printer
	0020  HL-5130 series
	0021  HL-5140 series
	0022  HL-5150D series
	0023  HL-5170DN series
	0024  Printer
	0025  Printer
	0027  HL-2030 Laser Printer
	0028  Printer
	0029  Printer
	002a  HL-52x0 series
	002b  HL-5250DN Printer
	002c  Printer
	002d  Printer
	0039  HL-5340 series
	0042  HL-2270DW Laser Printer
	0100  MFC8600/9650 series
	0101  MFC9600/9870 series
	0102  MFC9750/1200 series
	0104  MFC-8300J
	0105  MFC-9600J
	0106  MFC-7300C
	0107  MFC-7400C
	0108  MFC-9200C
	0109  MFC-830
	010a  MFC-840
	010b  MFC-860
	010c  MFC-7400J
	010d  MFC-9200J
	010e  MFC-3100C Scanner
	010f  MFC-5100C
	0110  MFC-4800 Scanner
	0111  MFC-6800
	0112  DCP1000 Port(FaxModem)
	0113  MFC-8500
	0114  MFC9700 Port(FaxModem)
	0115  MFC-9800 Scanner
	0116  DCP1400 Scanner
	0119  MFC-9660
	011a  MFC-9860
	011b  MFC-9880
	011c  MFC-9760
	011d  MFC-9070
	011e  MFC-9180
	011f  MFC-9160
	0120  MFC580 Port(FaxModem)
	0121  MFC-590
	0122  MFC-5100J
	0124  MFC-4800J
	0125  MFC-6800J
	0127  MFC-9800J
	0128  MFC-8500J
	0129  Imagistics 2500 (MFC-8640D clone)
	012b  MFC-9030
	012e  FAX4100e IntelliFax 4100e
	012f  FAX-4750e
	0130  FAX-5750e
	0132  MFC-5200C RemovableDisk
	0135  MFC-100 Scanner
	0136  MFC-150CL Scanner
	013c  MFC-890 Port
	013d  MFC-5200J
	013e  MFC-4420C RemovableDisk
	013f  MFC-4820C RemovableDisk
	0140  DCP-8020
	0141  DCP-8025D
	0142  MFC-8420
	0143  MFC-8820D
	0144  DCP-4020C RemovableDisk
	0146  MFC-3220C
	0147  FAX-1820C Printer
	0148  MFC-3320CN
	0149  FAX-1920CN Printer
	014a  MFC-3420C
	014b  MFC-3820CN
	014c  DCP-3020C
	014d  FAX-1815C Printer
	014e  MFC-8820J
	014f  DCP-8025J
	0150  MFC-8220 Port(FaxModem)
	0151  MFC-8210J
	0153  DCP-1000J
	0157  MFC-3420J Printer
	0158  MFC-3820JN Port(FaxModem)
	015d  MFC Composite Device
	015e  DCP-8045D
	015f  MFC-8440
	0160  MFC-8840D
	0161  MFC-210C
	0162  MFC-420CN Remote Setup Port
	0163  MFC-410CN RemovableDisk
	0165  MFC-620CN
	0166  MFC-610CLN RemovableDisk
	0168  MFC-620CLN
	0169  DCP-110C RemovableDisk
	016b  DCP-310CN RemovableDisk
	016c  FAX-2440C Printer
	016d  MFC-5440CN
	016e  MFC-5840CN Remote Setup Port
	0170  FAX-1840C Printer
	0171  FAX-1835C Printer
	0172  FAX-1940CN Printer
	0173  MFC-3240C Remote Setup Port
	0174  MFC-3340CN RemovableDisk
	017b  Imagistics sx2100
	0180  MFC-7420
	0181  MFC-7820N Port(FaxModem)
	0182  DCP-7010
	0183  DCP-7020
	0184  DCP-7025 Printer
	0185  MFC-7220 Printer
	0186  Composite Device
	0187  FAX-2820 Printer
	0188  FAX-2920 Printer
	018a  MFC-9420CN
	018c  DCP-115C
	018d  DCP-116C
	018e  DCP-117C
	018f  DCP-118C
	0190  DCP-120C
	0191  DCP-315CN
	0192  DCP-340CW
	0193  MFC-215C
	0194  MFC-425CN
	0195  MFC-820CW Remote Setup Port
	0196  MFC-820CN Remote Setup Port
	0197  MFC-640CW
	019a  MFC-840CLN Remote Setup Port
	01a2  MFC-8640D
	01a3  Composite Device
	01a4  DCP-8065DN Printer
	01a5  MFC-8460N Port(FaxModem)
	01a6  MFC-8860DN Port(FaxModem)
	01a7  MFC-8870DW Printer
	01a8  DCP-130C
	01a9  DCP-330C
	01aa  DCP-540CN
	01ab  MFC-240C
	01ae  DCP-750CW RemovableDisk
	01af  MFC-440CN
	01b0  MFC-660CN
	01b1  MFC-665CW
	01b2  MFC-845CW
	01b4  MFC-460CN
	01b5  MFC-630CD
	01b6  MFC-850CDN
	01b7  MFC-5460CN
	01b8  MFC-5860CN
	01ba  MFC-3360C
	01bd  MFC-8660DN
	01be  DCP-750CN RemovableDisk
	01bf  MFC-860CDN
	01c0  DCP-128C
	01c1  DCP-129C
	01c2  DCP-131C
	01c3  DCP-329C
	01c4  DCP-331C
	01c5  MFC-239C
	01c9  DCP-9040CN
	01ca  MFC-9440CN
	01cb  DCP-9045CDN
	01cc  MFC-9840CDW
	01ce  DCP-135C
	01cf  DCP-150C
	01d0  DCP-350C
	01d1  DCP-560CN
	01d2  DCP-770CW
	01d3  DCP-770CN
	01d4  MFC-230C
	01d5  MFC-235C
	01d6  MFC-260C
	01d7  MFC-465CN
	01d8  MFC-680CN
	01d9  MFC-685CW
	01da  MFC-885CW
	01db  MFC-480CN
	01dc  MFC-650CD
	01dd  MFC-870CDN
	01de  MFC-880CDN
	01df  DCP-155C
	01e0  MFC-265C
	01e1  DCP-153C
	01e2  DCP-157C
	01e3  DCP-353C
	01e4  DCP-357C
	01e7  MFC-7340
	01e9  DCP-7040
	01ea  DCP-7030
	01eb  MFC-7320
	01ec  MFC-9640CW
	01f4  MFC-5890CN
	020a  MFC-8670DN
	020c  DCP-9042CDN
	020d  MFC-9450CDN
	0216  MFC-8880DN
	0217  MFC-8480DN
	0219  MFC-8380DN
	021a  MFC-8370DN
	021b  DCP-8070D
	021c  MFC-9320CW
	021d  MFC-9120CN
	021e  DCP-9010CN
	0220  MFC-9010CN
	0222  DCP-195C
	0223  DCP-365CN
	0224  DCP-375CW
	0225  DCP-395CN
	0227  DCP-595CN
	0228  MFC-255CW
	0229  MFC-295CN
	022a  MFC-495CW
	022b  MFC-495CN
	022c  MFC-795CW
	022d  MFC-675CD
	022e  MFC-695CDN
	022f  MFC-735CD
	0230  MFC-935CDN
	0234  DCP-373CW
	0235  DCP-377CW
	0236  DCP-390CN
	0239  MFC-253CW
	023a  MFC-257CW
	023e  DCP-197C
	023f  MFC-8680DN
	0240  MFC-J950DN
	0248  DCP-7055 scanner/printer
	0253  DCP-J125
	0254  DCP-J315W
	0255  DCP-J515W
	0256  DCP-J515N
	0257  DCP-J715W
	0258  DCP-J715N
	0259  MFC-J220
	025a  MFC-J410
	025b  MFC-J265W
	025c  MFC-J415W
	025d  MFC-J615W
	025e  MFC-J615N
	025f  MFC-J700D
	0260  MFC-J800D
	0261  MFC-J850DN
	026b  MFC-J630W
	026d  MFC-J805D
	026e  MFC-J855DN
	026f  MFC-J270W
	0273  DCP-7057 scanner/printer
	0276  MFC-5895CW
	0278  MFC-J410W
	0279  DCP-J525W
	027a  DCP-J525N
	027b  DCP-J725DW
	027c  DCP-J725N
	027d  DCP-J925DW
	027e  MFC-J955DN
	027f  MFC-J280W
	0280  MFC-J435W
	0281  MFC-J430W
	0282  MFC-J625DW
	0283  MFC-J825DW
	0284  MFC-J825N
	0285  MFC-J705D
	0287  MFC-J860DN
	0288  MFC-J5910DW
	0289  MFC-J5910CDW
	028a  DCP-J925N
	028d  MFC-J835DW
	028f  MFC-J425W
	0290  MFC-J432W
	0291  DCP-8110DN
	0292  DCP-8150DN
	0293  DCP-8155DN
	0294  DCP-8250DN
	0295  MFC-8510DN
	0296  MFC-8520DN
	0298  MFC-8910DW
	0299  MFC-8950DW
	029a  MFC-8690DW
	029c  MFC-8515DN
	029e  MFC-9125CN
	029f  MFC-9325CW
	02a0  DCP-J140W
	02a5  MFC-7240
	02a6  FAX-2940
	02a7  FAX-2950
	02a8  MFC-7290
	02ab  FAX-2990
	02ac  DCP-8110D
	02ad  MFC-9130CW
	02ae  MFC-9140CDN
	02af  MFC-9330CDW
	02b0  MFC-9340CDW
	02b1  DCP-9020CDN
	02b2  MFC-J810DN
	02b3  MFC-J4510DW
	02b4  MFC-J4710DW
	02b5  DCP-8112DN
	02b6  DCP-8152DN
	02b7  DCP-8157DN
	02b8  MFC-8512DN
	02ba  MFC-8912DW
	02bb  MFC-8952DW
	02bc  DCP-J540N
	02bd  DCP-J740N
	02be  MFC-J710D
	02bf  MFC-J840N
	02c0  DCP-J940N
	02c1  MFC-J960DN
	02c2  DCP-J4110DW
	02c3  MFC-J4310DW
	02c4  MFC-J4410DW
	02c5  MFC-J4610DW
	02c6  DCP-J4210N
	02c7  MFC-J4510N
	02c8  MFC-J4910CDW
	02c9  MFC-J4810DN
	02ca  MFC-8712DW
	02cb  MFC-8710DW
	02cc  MFC-J2310
	02cd  MFC-J2510
	02ce  DCP-7055W
	02cf  DCP-7057W
	02d0  DCP-1510
	02d1  MFC-1810
	02d3  DCP-9020CDW
	02d4  MFC-8810DW
	02dd  DCP-J4215N
	02de  DCP-J132W
	02df  DCP-J152W
	02e0  DCP-J152N
	02e1  DCP-J172W
	02e2  DCP-J552DW
	02e3  DCP-J552N
	02e4  DCP-J752DW
	02e5  DCP-J752N
	02e6  DCP-J952N
	02e7  MFC-J245
	02e8  MFC-J470DW
	02e9  MFC-J475DW
	02ea  MFC-J285DW
	02eb  MFC-J650DW
	02ec  MFC-J870DW
	02ed  MFC-J870N
	02ee  MFC-J720D
	02ef  MFC-J820DN
	02f0  MFC-J980DN
	02f1  MFC-J890DN
	02f2  MFC-J6520DW
	02f3  MFC-J6570CDW
	02f4  MFC-J6720DW
	02f5  MFC-J6920DW
	02f6  MFC-J6970CDW
	02f7  MFC-J6975CDW
	02f8  MFC-J6770CDW
	02f9  DCP-J132N
	02fa  MFC-J450DW
	02fb  MFC-J875DW
	02fc  DCP-J100
	02fd  DCP-J105
	02fe  MFC-J200
	02ff  MFC-J3520
	0300  MFC-J3720
	030f  DCP-L8400CDN
	0310  DCP-L8450CDW
	0311  MFC-L8600CDW
	0312  MFC-L8650CDW
	0313  MFC-L8850CDW
	0314  MFC-L9550CDW
	0318  MFC-7365DN
	0320  MFC-L2740DW
	0321  DCP-L2500D
	0322  DCP-L2520DW
	0324  DCP-L2520D
	0326  DCP-L2540DN
	0328  DCP-L2540DW
	0329  DCP-L2560DW
	0330  HL-L2380DW
	0331  MFC-L2700DW
	0335  FAX-L2700DN
	0337  MFC-L2720DW
	0338  MFC-L2720DN
	0339  DCP-J4120DW
	033a  MFC-J4320DW
	033c  MFC-J2320
	033d  MFC-J4420DW
	0340  MFC-J4620DW
	0341  MFC-J2720
	0342  MFC-J4625DW
	0343  MFC-J5320DW
	0346  MFC-J5620DW
	0347  MFC-J5720DW
	0349  DCP-J4220N
	034b  MFC-J4720N
	034e  MFC-J5720CDW
	034f  MFC-J5820DN
	0350  MFC-J5620CDW
	0351  DCP-J137N
	0353  DCP-J557N
	0354  DCP-J757N
	0355  DCP-J957N
	0356  MFC-J877N
	0357  MFC-J727D
	0358  MFC-J987DN
	0359  MFC-J827DN
	035a  MFC-J897DN
	035b  DCP-1610W
	035c  DCP-1610NW
	035d  MFC-1910W
	035e  MFC-1910NW
	0360  DCP-1618W
	0361  MFC-1919NW
	0364  MFC-J5625DW
	0365  MFC-J4520DW
	0366  MFC-J5520DW
	0367  DCP-7080D
	0368  DCP-7080
	0369  DCP-7180DN
	036a  DCP-7189DW
	036b  MFC-7380
	036c  MFC-7480D
	036d  MFC-7880DN
	036e  MFC-7889DW
	036f  DCP-9022CDW
	0370  MFC-9142CDN
	0371  MFC-9332CDW
	0372  MFC-9342CDW
	0373  MFC-L2700D
	0376  DCP-1600
	0377  MFC-1900
	0378  DCP-1608
	0379  DCP-1619
	037a  MFC-1906
	037b  MFC-1908
	037c  ADS-2000e
	037d  ADS-2100e
	037e  ADS-2500We
	037f  ADS-2600We
	0380  DCP-J562DW
	0381  DCP-J562N
	0383  DCP-J962N
	0384  MFC-J480DW
	0385  MFC-J485DW
	0386  MFC-J460DW
	0388  MFC-J680DW
	0389  MFC-J880DW
	038a  MFC-J885DW
	038b  MFC-J880N
	038c  MFC-J730DN
	038d  MFC-J990DN
	038e  MFC-J830DN
	038f  MFC-J900DN
	0390  MFC-J5920DW
	0392  MFC-L2705DW
	0393  DCP-T300
	0394  DCP-T500W
	0395  DCP-T700W
	0396  MFC-T800W
	0397  DCP-J963N
	03b3  MFC-J6925DW
	03b4  MFC-J6573CDW
	03b5  MFC-J6973CDW
	03b6  MFC-J6990CDW
	03bb  MFC-L2680W
	03bc  MFC-L2700DN
	03bd  DCP-J762N
	1000  Printer
	1002  Printer
	2002  PTUSB Printing
	2004  PT-2300/2310 p-Touch Laber Printer
	2015  QL-500 P-touch label printer
	2016  QL-550 P-touch label printer
	201a  PT-18R P-touch label printer
	201b  QL-650TD P-Touch Label Printer
	2027  QL-560 P-Touch Label Printer
	202b  PT-7600 P-Touch Label Printer
	2100  Card Reader Writer
	60a0  ADS-2000
	60a1  ADS-2100
	60a4  ADS-2500W
	60a5  ADS-2600W
	60a6  ADS-1000W
	60a7  ADS-1100W
	60a8  ADS-1500W
	60a9  ADS-1600W
04fa  Dallas Semiconductor
	2490  DS1490F 2-in-1 Fob, 1-Wire adapter
	4201  DS4201 Audio DAC
04fb  Biostar Microtech International Corp.
04fc  Sunplus Technology Co., Ltd
	0003  CM1092 / Wintech CM-5098 Optical Mouse
	0005  USB OpticalWheel Mouse
	0013  ViewMate Desktop Mouse CC2201
	0015  ViewMate Desktop Mouse CC2201
	00d3  00052486 / Laser Mouse M1052 [hama]
	0171  SPCA1527A/SPCA1528 SD card camera (Mass Storage mode)
	0201  SPCP825 RS232C Adapter
	0232  Fingerprint
	0538  Wireless Optical Mouse 2.4G [Bright]
	0561  Flexcam 100
	05d8  Wireless keyboard/mouse
	0c15  SPIF215A SATA bridge
	0c25  SATALink SPIF225A
	1528  SPCA1527A/SPCA1528 SD card camera (webcam mode)
	1533  Mass Storage
	2080  ASUS Webcam
	500c  CA500C Digital Camera
	504a  Aiptek Mini PenCam 1.3
	504b  Aiptek Mega PockerCam 1.3/Maxell MaxPocket LE 1.3
	5330  Digitrex 2110
	5331  Vivitar Vivicam 10
	5360  Sunplus Generic Digital Camera
	5563  Digital Media Player MP3/WMA [The Sharper Image]
	5720  Card Reader Driver
	6333  Siri A9 UVC chipset
	7333  Finet Technology Palmpix DC-85
	757a  Aiptek, MP315 MP3 Player
	ffff  PureDigital Ritz Disposable
04fd  Soliton Systems, K.K.
	0003  Smart Card Reader II
04fe  PFU, Ltd
04ff  E-CMOS Corp.
0500  Siam United Hi-Tech
	0001  DART Keyboard Mouse
	0002  DART-2 Keyboard
0501  Fujikura DDK, Ltd
0502  Acer, Inc.
	0001  Handheld
	0736  Handheld
	15b1  PDA n311
	1631  c10 Series
	1632  c20 Series
	16e1  n10 Handheld Sync
	16e2  n20 Pocket PC Sync
	16e3  n30 Handheld Sync
	2008  Liquid Gallant Duo E350 (preloader)
	3202  Liquid
	3203  Liquid (Debug mode)
	3230  BeTouch E120
	3317  Liquid
	3325  Iconia tablet A500
	3341  Iconia tablet A500
	33c3  Liquid Gallant Duo E350
	33c4  Liquid Gallant Duo E350 (debug mode)
	33c7  Liquid Gallant Duo E350 (USB tethering)
	33c8  Liquid Gallant Duo E350 (debug mode, USB tethering)
	d001  Divio NW801/DVC-V6+ Digital Camera
0503  Hitachi America, Ltd
0504  Hayes Microcomputer Products
0506  3Com Corp.
	009d  HomeConnect Camera
	00a0  3CREB96 Bluetooth Adapter
	00a1  Bluetooth Device
	00a2  Bluetooth Device
	00df  3Com Home Connect lite
	0100  HomeConnect ADSL Modem Driver
	03e8  3C19250 Ethernet [klsi]
	0a01  3CRSHEW696 Wireless Adapter
	0a11  3CRWE254G72 802.11g Adapter
	11f8  HomeConnect 3C460
	2922  HomeConnect Cable Modem External with
	3021  U.S.Robotics 56000 Voice FaxModem Pro
	4601  3C460B 10/100 Ethernet Adapter
	f002  3CP4218 ADSL Modem (pre-init)
	f003  3CP4218 ADSL Modem
	f100  3CP4218 ADSL Modem (pre-init)
0507  Hosiden Corp.
	0011  Konami ParaParaParadise Controller
0508  Clarion Co., Ltd
0509  Aztech Systems, Ltd
	0801  ADSL Modem
	0802  ADSL Modem (RFC1483)
	0806  DSL Modem
	080f  Binatone ADSL500 Modem Network Interface
	0812  Pirelli ADSL Modem Network Interface
050a  Cinch Connectors
050b  Cable System International
050c  InnoMedia, Inc.
050d  Belkin Components
	0004  Direct Connect
	0012  F8T012 Bluetooth Adapter
	0013  F8T013 Bluetooth Adapter
	0017  B8T017 Bluetooth+EDR 2.1 / F4U017 USB 2.0 7-port Hub
	003a  Universal Media Reader
	0050  F5D6050 802.11b Wireless Adapter v2000 [Atmel at76c503a]
	0081  F8T001v2 Bluetooth
	0083  Bluetooth Device
	0084  F8T003v2 Bluetooth
	0102  Flip KVM
	0103  F5U103 Serial Adapter [etek]
	0106  VideoBus II Adapter, Video
	0108  F1DE108B KVM
	0109  F5U109/F5U409 PDA Adapter
	0115  SCSI Adapter
	0119  F5U120-PC Dual PS/2 Ports / F5U118-UNV ADB Adapter
	0121  F5D5050 100Mbps Ethernet
	0122  Ethernet Adapter
	0131  Bluetooth Device with trace filter
	016a  Bluetooth Mini Dongle
	0200  Nostromo SpeedPad n52te Gaming Keyboard
	0201  Peripheral Switch
	0208  USBView II Video Adapter [nt1004]
	0210  F5U228 Hi-Speed USB 2.0 DVD Creator
	0211  F5U211 USB 2.0 15-in-1 Media Reader & Writer
	0224  F5U224 USB 2.0 4-Port Hub
	0234  F5U234 USB 2.0 4-Port Hub
	0237  F5U237 USB 2.0 7-Port Hub
	0240  F5U240 USB 2.0 CF Card Reader
	0249  USB 2 Flash Media Device
	0257  F5U257 Serial
	0304  FSU304 USB 2.0 - 4 Ports Hub
	0307  USB 2.0 - 7 ports Hub [FSU307]
	0409  F5U409 Serial
	0416  Staples 12416 7 port desktop hub
	0551  F6C550-AVR UPS
	065a  F8T065BF Mini Bluetooth 4.0 Adapter
	0706  2-N-1 7-Port Hub (Lower half)
	0802  Nostromo n40 Gamepad
	0803  Nostromo 1745 GamePad
	0805  Nostromo N50 GamePad
	0815  Nostromo n52 HID SpeedPad Mouse Wheel
	0826  ErgoFit Wireless Optical Mouse (HID)
	0980  HID UPS Battery
	1004  F9L1004 802.11n Surf N300 XR Wireless Adapter [Realtek RTL8192CU]
	1102  F7D1102 N150/Surf Micro Wireless Adapter v1000 [Realtek RTL8188CUS]
	1103  F9L1103 N750 DB 802.11abgn 2x3:3 [Ralink RT3573]
	1106  F9L1106v1 802.11a/b/g/n/ac Wireless Adapter [Broadcom BCM43526]
	1109  F9L1109v1 802.11a/b/g/n/ac Wireless Adapter [Realtek RTL8812AU]
	110a  F9L1101v2 802.11abgn Wireless Adapter [Realtek RTL8192DU]
	11f2  ISY Wireless Micro Adapter IWL 2000 [RTL8188CUS]
	1202  F5U120-PC Parallel Printer Port
	1203  F5U120-PC Serial Port
	2103  F7D2102 802.11n N300 Micro Wireless Adapter v3000 [Realtek RTL8192CU]
	21f1  N300 WLAN N Adapter [ISY]
	21f2  RTL8192CU 802.11n WLAN Adapter [ISY IWL 4000]
	258a  F5U258 Host to Host cable
	3101  F1DF102U/F1DG102U Flip Hub
	3201  F1DF102U/F1DG102U Flip KVM
	4050  ZD1211B
	5055  F5D5055 Gigabit Network Adapter [AX88xxx]
	6050  F6D6050 802.11abgn Wireless Adapter [Broadcom BCM4323]
	6051  F5D6051 802.11b Wireless Network Adapter [ZyDAS ZD1201]
	615a  F7D4101 / F9L1101v1 802.11abgn Wireless Adapter [Broadcom BCM4323]
	7050  F5D7050 Wireless G Adapter v1000/v2000 [Intersil ISL3887]
	7051  F5D7051 802.11g Adapter v1000 [Broadcom 4320 USB]
	705a  F5D7050 Wireless G Adapter v3000 [Ralink RT2571W]
	705b  Wireless G Adapter
	705c  F5D7050 Wireless G Adapter v4000 [Zydas ZD1211B]
	705e  F5D7050 Wireless G Adapter v5000 [Realtek RTL8187B]
	706a  2-N-1 7-Port Hub (Upper half)
	8053  F5D8053 N Wireless USB Adapter v1000/v4000 [Ralink RT2870]
	805c  F5D8053 N Wireless Adapter v3000 [Ralink RT2870]
	805e  F5D8053 N Wireless USB Adapter v5000 [Realtek RTL8192U]
	815c  F5D8053 N Wireless USB Adapter v3000 [Ralink RT2870]
	815f  F5D8053 N Wireless USB Adapter v6000 [Realtek RTL8192SU]
	825a  F5D8055 N+ Wireless Adapter v1000 [Ralink RT2870]
	825b  F5D8055 N+ Wireless Adapter v2000 [Ralink RT3072]
	845a  F7D2101 802.11n Surf & Share Wireless Adapter v1000 [Realtek RTL8192SU]
	905b  F5D9050 Wireless G+ MIMO Network Adapter v3000 [Ralink RT2573]
	905c  F5D9050 Wireless G+ MIMO Network Adapter v4000 [Ralink RT2573]
	935a  F6D4050 N150 Enhanced Wireless Network Adapter v1000 [Ralink RT3070]
	935b  F6D4050 N150 Enhanced Wireless Network Adapter v2000 [Ralink RT3070]
	945a  F7D1101 v1 Basic Wireless Adapter [Realtek RTL8188SU]
	945b  F7D1101 v2 Basic Wireless Adapter [Ralink RT3370]
	d321  Dynex DX-NUSB 802.11bgn Wireless Adapter [Broadcom BCM43231]
050e  Neon Technology, Inc.
050f  KC Technology, Inc.
	0001  Hub
	0003  KC82C160S Hub
	0180  KC-180 IrDA Dongle
	0190  KC2190 USB Host-to-Host cable
0510  Sejin Electron, Inc.
	0001  Keyboard
	1000  Keyboard with PS/2 Mouse Port
	e001  Mouse
0511  N'Able (DataBook) Technologies, Inc.
	002b  AOC DVB
0512  Hualon Microelectronics Corp.
0513  digital-X, Inc.
0514  FCI Electronics
0515  ACTC
0516  Longwell Electronics
0517  Butterfly Communications
0518  EzKEY Corp.
	0001  USB to PS2 Adaptor v1.09
	0002  EZ-9900C Keyboard
0519  Star Micronics Co., Ltd
	0003  TSP100ECO/TSP100II
	c002  Xlive Bluetooth XBM-100S MP3 Player
051a  WYSE Technology
	a005  Smart Display Version 9973
051b  Silicon Graphics
051c  Shuttle, Inc.
	0005  VFD Module
	c001  eHome Infrared Receiver
	c002  eHome Infrared Receiver
051d  American Power Conversion
	0001  UPS
	0002  Uninterruptible Power Supply
	0003  UPS
051e  Scientific Atlanta, Inc.
051f  IO Systems (Elite Electronics), Inc.
0520  Taiwan Semiconductor Manufacturing Co.
0521  Airborn Connectors
0522  Advanced Connectek, Inc.
0523  ATEN GmbH
0524  Sola Electronics
0525  Netchip Technology, Inc.
	100d  RFMD Bluetooth Device
	1080  NET1080 USB-USB Bridge
	1200  SSDC Adapter II
	1265  File-backed Storage Gadget
	3424  Lumidigm Venus fingerprint sensor
	a0f0  Cambridge Electronic Devices Power1401 mk 2
	a140  USB Clik! 40
	a141  (OME) PocketZip 40 MP3 Player Driver
	a220  GVC Bluetooth Wireless Adapter
	a4a0  Linux-USB "Gadget Zero"
	a4a1  Linux-USB Ethernet Gadget
	a4a2  Linux-USB Ethernet/RNDIS Gadget
	a4a3  Linux-USB user-mode isochronous source/sink
	a4a4  Linux-USB user-mode bulk source/sink
	a4a5  Pocketbook Pro 903
	a4a6  Linux-USB Serial Gadget
	a4a7  Linux-USB Serial Gadget (CDC ACM mode)
	a4a8  Linux-USB Printer Gadget
	a4a9  Linux-USB OBEX Gadget
	a4aa  Linux-USB CDC Composite Gadge (Ethernet and ACM)
0526  Temic MHS S.A.
0527  ALTRA
0528  ATI Technologies, Inc.
	7561  TV Wonder
	7562  TV Wonder, Edition (FN5)
	7563  TV Wonder, Edition (FI)
	7564  TV Wonder, Edition (FQ)
	7565  TV Wonder, Edition (NTSC+)
	7566  TV Wonder, Edition (FN5)
	7567  TV Wonder, Edition (FI)
	7568  TV Wonder, Edition (FQ)
	7569  Live! Pro (A)
	756a  Live! Pro Audio (O)
0529  Aladdin Knowledge Systems
	0001  HASP copy protection dongle
	030b  eToken R1 v3.1.3.x
	0313  eToken R1 v3.2.3.x
	031b  eToken R1 v3.3.3.x
	0323  eToken R1 v3.4.3.x
	0412  eToken R2 v2.2.4.x
	041a  eToken R2 v2.2.4.x
	0422  eToken R2 v2.4.4.x
	042a  eToken R2 v2.5.4.x
	050c  eToken Pro v4.1.5.x
	0514  eToken Pro v4.2.5.4
	0600  eToken Pro 64k (4.2)
	0620  Token JC
052a  Crescent Heart Software
052b  Tekom Technologies, Inc.
	0102  Ca508A HP1020 Camera v.1.3.1.6
	0801  Yakumo MegaImage 37
	1512  Yakumo MegaImage IV
	1513  Aosta CX100 Webcam
	1514  Aosta CX100 Webcam Storage
	1905  Yakumo MegaImage 47
	1911  Yakumo MegaImage 47 SL
	2202  WDM Still Image Capture
	2203  Sound Vision Stream Driver
	3a06  DigiLife DDV-5120A
	d001  P35U Camera Capture
052c  Canon Information Systems, Inc.
052d  Avid Electronics Corp.
052e  Standard Microsystems Corp.
052f  Unicore Software, Inc.
0530  American Microsystems, Inc.
0531  Wacom Technology Corp.
0532  Systech Corp.
0533  Alcatel Mobile Phones
0534  Motorola, Inc.
0535  LIH TZU Electric Co., Ltd
0536  Hand Held Products (Welch Allyn, Inc.)
	01a0  PDT
0537  Inventec Corp.
0538  Caldera International, Inc. (SCO)
0539  Shyh Shiun Terminals Co., Ltd
053a  PrehKeyTec GmbH
	0b00  Hub
	0b01  Preh MCI 3100
053b  Global Village Communication
053c  Institut of Microelectronic & Mechatronic Systems
053d  Silicon Architect
053e  Mobility Electronics
053f  Synopsys, Inc.
0540  UniAccess AB
	0101  Panache Surf ISDN TA
0541  Sirf Technology, Inc.
0543  ViewSonic Corp.
	00fe  G773 Monitor Hub
	00ff  P815 Monitor Hub
	0bf2  airpanel V150 Wireless Smart Display
	0bf3  airpanel V110 Wireless Smart Display
	0ed9  Color Pocket PC V35
	0f01  airsync Wi-Fi Wireless Adapter
	1527  Color Pocket PC V36
	1529  Color Pocket PC V37
	152b  Color Pocket PC V38
	152e  Pocket PC
	1921  Communicator Pocket PC
	1922  Smartphone
	1923  Pocket PC V30
	1a11  Wireless 802.11g Adapter
	1e60  TA310 - ATSC/NTSC/PAL Driver(PCM4)
	4153  ViewSonic G773 Control (?)
0544  Cristie Electronics, Ltd
0545  Xirlink, Inc.
	7333  Trution Web Camera
	8002  IBM NetCamera
	8009  Veo PC Camera
	800c  Veo Stingray
	800d  Veo PC Camera
	8080  IBM C-It Webcam
	808a  Veo PC Camera
	808b  Veo Stingray
	808d  Veo PC Camera
	810a  Veo Advanced Connect Webcam
	810b  Veo PC Camera
	810c  Veo PC Camera
	8135  Veo Mobile/Advanced Web Camera
	813a  Veo PC Camera
	813b  Veo PC Camera
	813c  Veo Mobile/Advanced Web Camera
	8333  Veo Stingray/Connect Web Camera
	888c  eVision 123 digital camera
	888d  eVision 123 digital camera
0546  Polaroid Corp.
	0daf  PDC 2300Z
	1bed  PDC 1320 Camera
	3097  PDC 310
	3155  PDC 3070 Camera
	3187  Digital Camera
	3191  Ion 80 Camera
	3273  PDC 2030 Camera
	3304  a500 Digital Camera
	dccf  Sound Vision Stream Driver
0547  Anchor Chips, Inc.
	0001  ICSI Bluetooth Device
	1002  Python2 WDM Encoder
	1006  Hantek DSO-2100 UF
	2131  AN2131 EZUSB Microcontroller
	2235  AN2235 EZUSB-FX Microcontroller
	2710  EZ-Link Loader (EZLNKLDR.SYS)
	2720  AN2720 USB-USB Bridge
	2727  Xircom PGUNET USB-USB Bridge
	2750  EZ-Link (EZLNKUSB.SYS)
	2810  Cypress ATAPI Bridge
	4d90  AmScope MD1900 camera
	7000  PowerSpec MCE460 Front Panel LED Display
	7777  Bluetooth Device
	9999  AN2131 uninitialized (?)
0548  Tyan Computer Corp.
	1005  EZ Cart II GameBoy Flash Programmer
0549  Pixera Corp.
054a  Fujitsu Microelectronics, Inc.
054b  New Media Corp.
054c  Sony Corp.
	0001  HUB
	0002  Standard HUB
	0010  DSC-S30/S70/S75/F505V/F505/FD92/W1 Cybershot/Mavica Digital Camera
	0014  Nogatech USBVision (SY)
	0022  Storage Adapter V2 (TPP)
	0023  CD Writer
	0024  Mavica CD-1000 Camera
	0025  NW-MS7 Walkman MemoryStick Reader
	002b  Portable USB Harddrive V2
	002c  USB Floppy Disk Drive
	002d  MSAC-US1 MemoryStick Reader
	002e  HandyCam MemoryStick Reader
	0030  Storage Adapter V2 (TPP)
	0032  MemoryStick MSC-U01 Reader
	0035  Network Walkman (E)
	0036  Net MD
	0037  MG Memory Stick Reader/Writer
	0038  Clie PEG-S300/D PalmOS PDA
	0039  Network Walkman (MS)
	003c  VAIO-MX LCD Control
	0045  Digital Imaging Video
	0046  Network Walkman
	004a  Memory Stick Hi-Fi System
	004b  Memory Stick Reader/Writer
	004e  DSC-xxx (ptp)
	0056  MG Memory Stick Reader/Writer
	0058  Clie PEG-N7x0C PalmOS PDA Mass Storage
	0066  Clie PEG-N7x0C/PEG-T425 PalmOS PDA Serial
	0067  CMR-PC3 Webcam
	0069  Memorystick MSC-U03 Reader
	006c  FeliCa S310 [PaSoRi]
	006d  Clie PEG-T425 PDA Mass Storage
	006f  Network Walkman (EV)
	0073  Storage CRX1750U
	0075  Net MD
	0076  Storage Adapter ACR-U20
	007c  Net MD
	007f  IC Recorder (MS)
	0080  Net MD
	0081  Net MD
	0084  Net MD
	0085  Net MD
	0086  Net MD
	008b  Micro Vault 64M Mass Storage
	0095  Clie s360
	0099  Clie NR70 PDA Mass Storage
	009a  Clie NR70 PDA Serial
	00ab  Visual Communication Camera (PCGA-UVC10)
	00af  DPP-EX Series Digital Photo Printer
	00bf  IC Recorder (S)
	00c0  Handycam DCR-30
	00c6  Net MD
	00c7  Net MD
	00c8  MZ-N710 Minidisc Walkman
	00c9  Net MD
	00ca  MZ-DN430 Minidisc Walkman
	00cb  MSAC-US20 Memory Stick Reader
	00da  Clie nx60
	00e8  Network Walkman (MS)
	00e9  Handheld
	00eb  Net MD
	0101  Net MD
	0103  IC Recorder (ST)
	0105  Micro Vault Hub
	0107  VCC-U01 Visual Communication Camera
	0110  Digital Imaging Video
	0113  Net MD
	0116  IC Recorder (P)
	0144  Clie PEG-TH55 PDA
	0147  Visual Communication Camera (PCGA-UVC11)
	014c  Aiwa AM-NX9 Net MD Music Recorder MDLP
	014d  Memory Stick Reader/Writer
	0154  Eyetoy Audio Device
	015f  IC Recorder (BM)
	0169  Clie PEG-TJ35 PDA Serial
	016a  Clie PEG-TJ35 PDA Mass Storage
	016b  Mobile HDD
	016d  IC Recorder (SX)
	016e  DPP-EX50 Digital Photo Printer
	0171  Fingerprint Sensor 3500
	017e  Net MD
	017f  Hi-MD WALKMAN
	0180  Net MD
	0181  Hi-MD WALKMAN
	0182  Net MD
	0183  Hi-MD WALKMAN
	0184  Net MD
	0185  Hi-MD WALKMAN
	0186  Net MD
	0187  Hi-MD MZ-NH600 WALKMAN
	0188  Net MD
	018a  Net MD
	018b  Hi-MD SOUND GATE
	019e  Micro Vault 1.0G Mass Storage
	01ad  ATRAC HDD PA
	01bb  FeliCa S320 [PaSoRi]
	01bd  MRW62E Multi-Card Reader/Writer
	01c3  NW-E55 Network Walkman
	01c6  MEMORY P-AUDIO
	01c7  Printing Support
	01c8  PSP Type A
	01c9  PSP Type B
	01d0  DVD+RW External Drive DRU-700A
	01d5  IC RECORDER
	01de  VRD-VC10 [Video Capture]
	01e8  UP-DR150 Photo Printer
	01e9  Net MD
	01ea  Hi-MD WALKMAN
	01ee  IC RECORDER
	01fa  IC Recorder (P)
	01fb  NW-E405 Network Walkman
	020f  Device
	0210  ATRAC HDD PA
	0219  Net MD
	021a  Hi-MD WALKMAN
	021b  Net MD
	021c  Hi-MD WALKMAN
	021d  Net MD
	0227  Printing Support
	022c  Net MD
	022d  Hi-MD AUDIO
	0233  ATRAC HDD PA
	0236  Mobile HDD
	023b  DVD+RW External Drive DRU-800UL
	023c  Net MD
	023d  Hi-MD WALKMAN
	0243  MicroVault Flash Drive
	024b  Vaio VGX Mouse
	0257  IFU-WLM2 USB Wireless LAN Module (Wireless Mode)
	0258  IFU-WLM2 USB Wireless LAN Module (Memory Mode)
	0259  IC RECORDER
	0267  Tachikoma Device
	0268  Batoh Device / PlayStation 3 Controller
	0269  HDD WALKMAN
	026a  HDD WALKMAN
	0271  IC Recorder (P)
	027c  NETWORK WALKMAN
	027e  SONY Communicator
	027f  IC RECORDER
	0286  Net MD
	0287  Hi-MD WALKMAN
	0290  VGP-UVC100 Visual Communication Camera
	029b  PRS-500 eBook reader
	02a5  MicroVault Flash Drive
	02af  Handycam DCR-DVD306E
	02c4  Device
	02d1  DVD RW
	02d2  PSP Slim
	02d8  SBAC-US10 SxS PRO memory card reader/writer
	02e1  FeliCa S330 [PaSoRi]
	02ea  PlayStation 3 Memory Card Adaptor
	02f9  DSC-H9
	0317  WALKMAN
	031a  Walkman NWD-B103F
	031e  PRS-300/PRS-505 eBook reader
	0325  NWZ-A818
	033e  DSC-W120/W290
	0346  Handycam DCR-SR55E
	0348  HandyCam HDR-TG3E
	035b  Walkman NWZ-A828
	035c  NWZ-A726/A728/A729
	035f  UP-DR200 Photo Printer
	0382  Memory Stick PRO-HG Duo Adaptor (MSAC-UAH1)
	0385  Walkman NWZ-E436F
	0387  IC Recorder (P)
	03bc  Webbie HD - MHS-CM1
	03d1  DPF-X95
	03d3  DR-BT100CX
	03d5  PlayStation Move motion controller
	03fc  WALKMAN [NWZ-E345]
	03fd  Walkman NWZ-E443
	042f  PlayStation Move navigation controller
	0440  DSC-H55
	0485  MHS-PM5 HD camcorder
	04cb  WALKMAN NWZ-E354
	0541  DSC-HX100V [Cybershot Digital Still Camera]
	05c4  DualShock 4
	0689  Walkman NWZ-B173F
	06bb  WALKMAN NWZ-F805
	088c  Portable Headphone Amplifier
	1000  Wireless Buzz! Receiver
054d  Try Corp.
054e  Proside Corp.
054f  WYSE Technology Taiwan
0550  Fuji Xerox Co., Ltd
	0002  InkJet Color Printer
	0004  InkJet Color Printer
	0005  InkJet Color Printer
	000b  Workcentre 24
	014e  CM215b Printer
	0165  DocuPrint M215b
0551  CompuTrend Systems, Inc.
0552  Philips Monitors
0553  STMicroelectronics Imaging Division (VLSI Vision)
	0001  TerraCAM
	0002  CPiA Webcam
	0100  STV0672 Camera
	0140  Video Camera
	0150  CDE CAM 100
	0151  Digital Blue QX5 Microscope
	0200  Dual-mode Camera0
	0201  Dual-mode Camera1
	0202  STV0680 Camera
	0674  Multi-mode Camera
	0679  NMS Video Camera (Webcam)
	1002  Che-ez! Splash
0554  Dictaphone Corp.
0555  ANAM S&T Co., Ltd
0556  Asahi Kasei Microsystems Co., Ltd
	0001  AK5370 I/F A/D Converter
0557  ATEN International Co., Ltd
	2001  UC-1284 Printer Port
	2002  10Mbps Ethernet [klsi]
	2004  UC-100KM PS/2 Mouse and Keyboard adapter
	2006  UC-1284B Printer Port
	2007  UC-110T 100Mbps Ethernet [pegasus]
	2008  UC-232A Serial Port [pl2303]
	2009  UC-210T Ethernet
	2011  UC-2324 4xSerial Ports [mos7840]
	2202  CS124U Miniview II KVM Switch
	2213  CS682 2-Port USB 2.0 DVI KVM Switch
	2221  Winbond Hermon
	2404  4-port switch
	2600  IDE Bridge
	2701  CE700A KVM Extender
	4000  DSB-650 10Mbps Ethernet [klsi]
	7000  Hub
	7820  UC-2322 2xSerial Ports [mos7820]
	8021  CS1764A [CubiQ DVI KVMP Switch]
0558  Truevision, Inc.
	1009  GW Instek GDS-1000 Oscilloscope
	100a  GW Instek GDS-1000A Oscilloscope
	2009  GW Instek GDS-2000 Oscilloscope
0559  Cadence Design Systems, Inc.
055a  Kenwood USA
055b  KnowledgeTek, Inc.
055c  Proton Electronic Ind.
055d  Samsung Electro-Mechanics Co.
	0001  Keyboard
	0bb1  Bluetooth Device
	1030  Optical Wheel Mouse (OMS3CB/OMGB30)
	1031  Optical Wheel Mouse (OMA3CB/OMGI30)
	1040  Mouse HID Device
	1050  E-Mail Optical Wheel Mouse (OMS3CE)
	1080  Optical Wheel Mouse (OMS3CH)
	2020  Floppy Disk Drive
	6780  Keyboard V1
	6781  Keyboard Mouse
	8001  E.M. Hub
	9000  AnyCam [pwc]
	9001  MPC-C30 AnyCam Premium for Notebooks [pwc]
	a000  SWL-2100U
	a010  WLAN Adapter(SWL-2300)
	a011  Boot Device
	a012  WLAN Adapter(SWL-2300)
	a013  WLAN Adapter(SWL-2350)
	a230  Boot Device
	b000  11Mbps WLAN Mini Adapter
	b230  Netopia 802.11b WLAN Adapter
	b231  LG Wireless LAN 11b Adapter
055e  CTX Opto-Electronics Corp.
055f  Mustek Systems, Inc.
	0001  ScanExpress 1200 CU
	0002  ScanExpress 600 CU
	0003  ScanExpress 1200 USB
	0006  ScanExpress 1200 UB
	0007  ScanExpress 1200 USB Plus
	0008  ScanExpress 1200 CU Plus
	0010  BearPaw 1200F
	0210  ScanExpress A3 USB
	0218  BearPaw 2400 TA
	0219  BearPaw 2400 TA Plus
	021a  BearPaw 2448 TA Plus
	021b  BearPaw 1200 CU Plus
	021c  BearPaw 1200 CU Plus
	021d  BearPaw 2400 CU Plus
	021e  BearPaw 1200 TA/CS
	021f  SNAPSCAN e22
	0400  BearPaw 2400 TA Pro
	0401  P 3600 A3 Pro
	0408  BearPaw 2448 CU Pro
	0409  BearPaw 2448 TA Pro
	040b  ScanExpress A3 USB 1200 PRO
	0873  ScanExpress 600 USB
	1000  BearPaw 4800 TA Pro
	a350  gSmart 350 Camera
	a800  MDC 800 Camera
	b500  MDC 3000 Camera
	c005  PC CAM 300A
	c200  gSmart 300
	c211  Kowa Bs888e Microcamera
	c220  gSmart mini
	c230  Digicam 330K
	c232  MDC3500 Camera
	c360  DV 4000 Camera
	c420  gSmart mini 2 Camera
	c430  gSmart LCD 2 Camera
	c440  DV 3000 Camera
	c520  gSmart mini 3 Camera
	c530  gSmart LCD 2 Camera
	c540  gSmart D30 Camera
	c630  MDC 4000 Camera
	c631  MDC 4000 Camera
	c650  MDC 5500Z Camera
	d001  WCam 300
	d003  WCam 300A
	d004  WCam 300AN
0560  Interface Corp.
0561  Oasis Design, Inc.
0562  Telex Communications, Inc.
	0001  Enhanced Microphone
	0002  Telex Microphone
0563  Immersion Corp.
0564  Kodak Digital Product Center, Japan Ltd. (formerly Chinon Industries Inc.)
0565  Peracom Networks, Inc.
	0001  Serial Port [etek]
	0002  Enet Ethernet [klsi]
	0003  @Home Networks Ethernet [klsi]
	0005  Enet2 Ethernet [klsi]
	0041  Peracom Remote NDIS Ethernet Adapter
0566  Monterey International Corp.
	0110  ViewMate Desktop Mouse CC2201
	1001  ViewMate Desktop Mouse CC2201
	1002  ViewMate Desktop Mouse CC2201
	1003  ViewMate Desktop Mouse CC2201
	1004  ViewMate Desktop Mouse CC2201
	1005  ViewMate Desktop Mouse CC2201
	1006  ViewMate Desktop Mouse CC2201
	1007  ViewMate Desktop Mouse CC2201
	2800  MIC K/B
	2801  MIC K/B Mouse
	2802  Kbd Hub
	3002  Keyboard
	3004  Genius KB-29E
	3107  Keyboard
0567  Xyratex International, Ltd
0568  Quartz Ingenierie
0569  SegaSoft
056a  Wacom Co., Ltd
	0000  PenPartner
	0001  PenPartner 4x5
	0002  PenPartner 6x8
	0003  PTU-600 [Cintiq Partner]
	0010  ET-0405 [Graphire]
	0011  ET-0405A [Graphire2 (4x5)]
	0012  ET-0507A [Graphire2 (5x7)]
	0013  CTE-430 [Graphire3 (4x5)]
	0014  CTE-630 [Graphire3 (6x8)]
	0015  CTE-440 [Graphire4 (4x5)]
	0016  CTE-640 [Graphire4 (6x8)]
	0017  CTE-450 [Bamboo Fun (small)]
	0018  CTE-650 [Bamboo Fun (medium)]
	0019  CTE-631 [Bamboo One]
	0020  GD-0405 [Intuos (4x5)]
	0021  GD-0608 [Intuos (6x8)]
	0022  GD-0912 [Intuos (9x12)]
	0023  GD-1212 [Intuos (12x12)]
	0024  GD-1218 [Intuos (12x18)]
	0026  PTH-450 [Intuos5 touch (S)]
	0027  PTH-650 [Intuos5 touch (M)]
	0028  PTH-850 [Intuos5 touch (L)]
	0029  PTK-450 [Intuos5 (S)]
	002a  PTK-650 [Intuos5 (M)]
	0030  PL400
	0031  PL500
	0032  PL600
	0033  PL600SX
	0034  PL550
	0035  PL800
	0037  PL700
	0038  PL510
	0039  DTU-710
	003f  DTZ-2100 [Cintiq 21UX]
	0041  XD-0405-U [Intuos2 (4x5)]
	0042  XD-0608-U [Intuos2 (6x8)]
	0043  XD-0912-U [Intuos2 (9x12)]
	0044  XD-1212-U [Intuos2 (12x12)]
	0045  XD-1218-U [Intuos2 (12x18)]
	0047  Intuos2 6x8
	0057  DTK-2241
	0059  DTH-2242 tablet
	005b  DTH-2200 [Cintiq 22HD Touch] tablet
	005d  DTH-2242 touchscreen
	005e  DTH-2200 [Cintiq 22HD Touch] touchscreen
	0060  FT-0405 [Volito, PenPartner, PenStation (4x5)]
	0061  FT-0203 [Volito, PenPartner, PenStation (2x3)]
	0062  CTF-420 [Volito2]
	0063  CTF-220 [BizTablet]
	0064  CTF-221 [PenPartner2]
	0065  MTE-450 [Bamboo]
	0069  CTF-430 [Bamboo One]
	006a  CTE-460 [Bamboo One Pen (S)]
	006b  CTE-660 [Bamboo One Pen (M)]
	0081  CTE-630BT [Graphire Wireless (6x8)]
	0084  Wireless adapter for Bamboo tablets
	0090  TPC90
	0093  TPC93
	0097  TPC97
	009a  TPC9A
	00b0  PTZ-430 [Intuos3 (4x5)]
	00b1  PTZ-630 [Intuos3 (6x8)]
	00b2  PTZ-930 [Intuos3 (9x12)]
	00b3  PTZ-1230 [Intuos3 (12x12)]
	00b4  PTZ-1231W [Intuos3 (12x19)]
	00b5  PTZ-631W [Intuos3 (6x11)]
	00b7  PTZ-431W [Intuos3 (4x6)]
	00b8  PTK-440 [Intuos4 (4x6)]
	00b9  PTK-640 [Intuos4 (6x9)]
	00ba  PTK-840 [Intuos4 (8x13)]
	00bb  PTK-1240 [Intuos4 (12x19)]
	00c0  DTF-521
	00c4  DTF-720
	00c5  DTZ-20WSX [Cintiq 20WSX]
	00c6  DTZ-12WX [Cintiq 12WX]
	00c7  DTU-1931
	00cc  DTK-2100 [Cintiq 21UX]
	00ce  DTU-2231
	00d0  CTT-460 [Bamboo Touch]
	00d1  CTH-460 [Bamboo Pen & Touch]
	00d2  CTH-461 [Bamboo Fun/Craft/Comic Pen & Touch (S)]
	00d3  CTH-661 [Bamboo Fun/Comic Pen & Touch (M)]
	00d4  CTL-460 [Bamboo Pen (S)]
	00d5  CTL-660 [Bamboo Pen (M)]
	00d6  CTH-460 [Bamboo Pen & Touch]
	00d7  CTH-461 [Bamboo Fun/Craft/Comic Pen & Touch (S)]
	00d8  CTH-661 [Bamboo Fun/Comic Pen & Touch (M)]
	00d9  CTT-460 [Bamboo Touch]
	00da  CTH-461SE [Bamboo Pen & Touch Special Edition (S)]
	00db  CTH-661SE [Bamboo Pen & Touch Special Edition (M)]
	00dc  CTT-470 [Bamboo Touch]
	00dd  CTL-470 [Bamboo Connect]
	00de  CTH-470 [Bamboo Fun Pen & Touch]
	00df  CTH-670 [Bamboo Create/Fun]
	00e2  TPCE2
	00e3  TPCE3
	00e5  TPCE5
	00e6  TPCE6
	00ec  TPCEC
	00ed  TPCED
	00ef  TPCEF
	00f4  DTK-2400 [Cintiq 24HD] tablet
	00f6  DTH-2400 [Cintiq 24HD touch] touchscreen
	00f8  DTH-2400 [Cintiq 24HD touch] tablet
	00fa  DTK-2200 [Cintiq 22HD] tablet
	00fb  DTU-1031
	0100  TPC100
	0101  TPC101
	010d  TPC10D
	010e  TPC10E
	010f  TPC10F
	0116  TPC116
	012c  TPC12C
	0300  CTL-471 [Bamboo Splash, One by Wacom (S)]
	0301  CTL-671 [One by Wacom (M)]
	0302  CTH-480 [Intuos Pen & Touch (S)]
	0303  CTH-680 [Intuos Pen & Touch (M)]
	0304  DTK-1300 [Cintiq 13HD]
	0307  DTH-A1300 [Cintiq Companion Hybrid] tablet
	0309  DTH-A1300 [Cintiq Companion Hybrid] touchscreen
	030e  CTL-480 [Intuos Pen (S)]
	0314  PTH-451 [Intuos pro (S)]
	0315  PTH-651 [Intuos pro (M)]
	0317  PTH-851 [Intuos pro (L)]
	032f  DTU-1031X
	0400  PenPartner 4x5
	4001  TPC4001
	4004  TPC4004
	4850  PenPartner 6x8
	5000  TPC5000
	5002  TPC5002
	5010  TPC5010
056b  Decicon, Inc.
056c  eTEK Labs
	0006  KwikLink Host-Host Connector
	8007  Kwik232 Serial Port
	8100  KwikLink Host-Host Connector
	8101  KwikLink USB-USB Bridge
056d  EIZO Corp.
	0000  Hub
	0001  Monitor
	0002  HID Monitor Controls
	0003  Device Bay Controller
056e  Elecom Co., Ltd
	0002  29UO Mouse
	0072  Mouse
	200c  LD-USB/TX
	4002  Laneed 100Mbps Ethernet LD-USB/TX [pegasus]
	4005  LD-USBL/TX
	400b  LD-USB/TX
	4010  LD-USB20
	5003  UC-SGT
	5004  UC-SGT
	6008  Flash Disk
	abc1  LD-USB/TX
056f  Korea Data Systems Co., Ltd
	cd00  CDM-751 CD organizer
0570  Epson America
0571  Interex, Inc.
	0002  echoFX InterView Lite
0572  Conexant Systems (Rockwell), Inc.
	0001  Ezcam II Webcam
	0002  Ezcam II Webcam
	0040  Wondereye CP-115 Webcam
	0041  Webcam Notebook
	0042  Webcam Notebook
	0320  DVBSky T330 DVB-T2/C tuner
	1232  V.90 modem
	1234  Typhoon Redfun Modem V90 56k
	1252  HCF V90 Data Fax Voice Modem
	1253  Zoom V.92 Faxmodem
	1300  SoftK56 Data Fax Voice CARP
	1301  Modem Enumerator
	1328  TrendNet TFM-561 modem
	2000  SoftGate 802.11 Adapter
	2002  SoftGate 802.11 Adapter
	262a  tm5600 Video & Audio Grabber Capture
	680c  DVBSky T680C DVB-T2/C tuner
	6831  DVBSky S960 DVB-S2 tuner
	8390  WinFast PalmTop/Novo TV Video
	8392  WinFast PalmTop/Novo TV Video
	960c  DVBSky S960C DVB-S2 tuner
	c686  Geniatech T220A DVB-T2 TV Stick
	c688  Geniatech T230 DVB-T2 TV Stick
	cafc  CX861xx ROM Boot Loader
	cafd  CX82310 ROM Boot Loader
	cafe  AccessRunner ADSL Modem
	cb00  ADSL Modem
	cb01  ADSL Modem
	cb06  StarModem Network Interface
0573  Zoran Co. Personal Media Division (Nogatech)
	0003  USBGear USBG-V1
	0400  D-Link V100
	0600  Dazzle USBVision (1006)
	1300  leadtek USBVision (1006)
	2000  X10 va10a Wireless Camera
	2001  Dazzle EmMe (2001)
	2101  Zoran Co. PMD (Nogatech) AV-grabber Manhattan
	2d00  Osprey 50
	2d01  Hauppauge USB-Live Model 600
	3000  Dazzle MicroCam (NTSC)
	3001  Dazzle MicroCam (PAL)
	4000  Nogatech TV! (NTSC)
	4001  Nogatech TV! (PAL)
	4002  Nogatech TV! (PAL-I-)
	4003  Nogatech TV! (MF-)
	4008  Nogatech TV! (NTSC) (T)
	4009  Nogatech TV! (PAL) (T)
	4010  Nogatech TV! (NTSC) (A)
	4100  USB-TV FM (NTSC)
	4110  PNY USB-TV (NTSC) FM
	4400  Nogatech TV! Pro (NTSC)
	4401  Nogatech TV! Pro (PAL)
	4450  PixelView PlayTv-USB PRO (PAL) FM
	4451  Nogatech TV! Pro (PAL+)
	4452  Nogatech TV! Pro (PAL-I+)
	4500  Nogatech TV! Pro (NTSC)
	4501  Nogatech TV! Pro (PAL)
	4550  ZTV ZT-721 2.4GHz A/V Receiver
	4551  Dazzle TV! Pro Audio (P+)
	4d00  Hauppauge WinTV-USB USA
	4d01  Hauppauge WinTV-USB
	4d02  Hauppauge WinTV-USB UK
	4d03  Hauppauge WinTV-USB France
	4d04  Hauppauge WinTV (PAL D/K)
	4d10  Hauppauge WinTV-USB with FM USA radio
	4d11  Hauppauge WinTV-USB (PAL) with FM radio
	4d12  Hauppauge WinTV-USB UK with FM Radio
	4d14  Hauppauge WinTV (PAL D/K FM)
	4d20  Hauppauge WinTV-USB II (PAL) with FM radio
	4d21  Hauppauge WinTV-USB II (PAL)
	4d22  Hauppauge WinTV-USB II (PAL) Model 566
	4d23  Hauppauge WinTV-USB France 4D23
	4d24  Hauppauge WinTV Pro (PAL D/K)
	4d25  Hauppauge WinTV-USB Model 40209 rev B234
	4d26  Hauppauge WinTV-USB Model 40209 rev B243
	4d27  Hauppauge WinTV-USB Model 40204 Rev B281
	4d28  Hauppauge WinTV-USB Model 40204 rev B283
	4d29  Hauppauge WinTV-USB Model 40205 rev B298
	4d2a  Hauppague WinTV-USB Model 602 Rev B285
	4d2b  Hauppague WinTV-USB Model 602 Rev B282
	4d2c  Hauppauge WinTV Pro (PAL/SECAM)
	4d30  Hauppauge WinTV-USB FM Model 40211 Rev B123
	4d31  Hauppauge WinTV-USB III (PAL) with FM radio Model 568
	4d32  Hauppauge WinTV-USB III (PAL) FM Model 573
	4d34  Hauppauge WinTV Pro (PAL D/K FM)
	4d35  Hauppauge WinTV-USB III (PAL) FM Model 597
	4d36  Hauppauge WinTV Pro (PAL B/G FM)
	4d37  Hauppauge WinTV-USB Model 40219 rev E189
	4d38  Hauppauge WinTV Pro (NTSC FM)
0574  City University of Hong Kong
0575  Philips Creative Display Solutions
0576  BAFO/Quality Computer Accessories
0577  ELSA
0578  Intrinsix Corp.
0579  GVC Corp.
057a  Samsung Electronics America
057b  Y-E Data, Inc.
	0000  FlashBuster-U Floppy
	0001  Tri-Media Reader Floppy
	0006  Tri-Media Reader Card Reader
	0010  Memory Stick Reader Writer
	0020  HEXA Media Drive 6-in-1 Card Reader Writer
	0030  Memory Card Viewer (TV)
057c  AVM GmbH
	0b00  ISDN-Controller B1 Family
	0c00  ISDN-Controller FRITZ!Card
	1000  ISDN-Controller FRITZ!Card v2.0
	1900  ISDN-Controller FRITZ!Card v2.1
	2000  ISDN-Connector FRITZ!X
	2200  BlueFRITZ!
	2300  Teledat X130 DSL
	2800  ISDN-Connector TA
	3200  Teledat X130 DSL
	3500  FRITZ!Card DSL SL
	3701  FRITZ!Box SL
	3702  FRITZ!Box
	3800  BlueFRITZ! Bluetooth Stick
	3a00  FRITZ!Box Fon
	3c00  FRITZ!Box WLAN
	3d00  Fritz!Box
	3e01  FRITZ!Box (Annex A)
	4001  FRITZ!Box Fon (Annex A)
	4101  FRITZ!Box WLAN (Annex A)
	4201  FRITZ!Box Fon WLAN (Annex A)
	4601  Eumex 5520PC (WinXP/2000)
	4602  Eumex 400 (WinXP/2000)
	4701  AVM FRITZ!Box Fon ata
	5401  Eumex 300 IP
	5601  AVM Fritz!WLAN [Texas Instruments TNETW1450]
	6201  AVM Fritz!WLAN v1.1 [Texas Instruments TNETW1450]
	62ff  AVM Fritz!WLAN USB (in CD-ROM-mode)
	8401  Fritz!WLAN N [Atheros AR9001U]
	8402  Fritz!WLAN N 2.4 [Atheros AR9001U]
	8403  Fritz!WLAN N v2 [Atheros AR9271]
	84ff  AVM Fritz!WLAN USB N (in CD-ROM-mode)
	8501  FRITZ WLAN N v2 [RT5572/rt2870.bin]
057d  Shark Multimedia, Inc.
057e  Nintendo Co., Ltd
	0305  Broadcom BCM2045A Bluetooth Radio [Nintendo Wii]
	0306  Wii Remote Controller RVL-003
057f  QuickShot, Ltd
	6238  USB StrikePad
0580  Denron, Inc.
0581  Racal Data Group
0582  Roland Corp.
	0000  UA-100(G)
	0002  UM-4/MPU-64 MIDI Interface
	0003  SoundCanvas SC-8850
	0004  U-8
	0005  UM-2(C/EX)
	0007  SoundCanvas SC-8820
	0008  PC-300
	0009  UM-1(E/S/X)
	000b  SK-500
	000c  SC-D70
	0010  EDIROL UA-5
	0011  Edirol UA-5 Sound Capture
	0012  XV-5050
	0013  XV-5050
	0014  EDIROL UM-880 MIDI I/F (native)
	0015  EDIROL UM-880 MIDI I/F (generic)
	0016  EDIROL SD-90
	0017  EDIROL SD-90
	0018  UA-1A
	001b  MMP-2
	001c  MMP-2
	001d  V-SYNTH
	001e  V-SYNTH
	0023  EDIROL UM-550
	0024  EDIROL UM-550
	0025  EDIROL UA-20
	0026  EDIROL UA-20
	0027  EDIROL SD-20
	0028  EDIROL SD-20
	0029  EDIROL SD-80
	002a  EDIROL SD-80
	002b  EDIROL UA-700
	002c  EDIROL UA-700
	002d  XV-2020 Synthesizer
	002e  XV-2020 Synthesizer
	002f  VariOS
	0030  VariOS
	0033  EDIROL PCR
	0034  EDIROL PCR
	0035  M-1000
	0037  Digital Piano
	0038  Digital Piano
	003b  BOSS GS-10
	003c  BOSS GS-10
	0040  GI-20
	0041  GI-20
	0042  RS-70
	0043  RS-70
	0044  EDIROL UA-1000
	0047  EDIROL UR-80 WAVE
	0048  EDIROL UR-80 MIDI
	0049  EDIROL UR-80 WAVE
	004a  EDIROL UR-80 MIDI
	004b  EDIROL M-100FX
	004c  EDIROL PCR-A WAVE
	004d  EDIROL PCR-A MIDI
	004e  EDIROL PCR-A WAVE
	004f  EDIROL PCR-A MIDI
	0050  EDIROL UA-3FX
	0052  EDIROL UM-1SX
	0054  Digital Piano
	0060  EXR Series
	0064  EDIROL PCR-1 WAVE
	0065  EDIROL PCR-1 MIDI
	0066  EDIROL PCR-1 WAVE
	0067  EDIROL PCR-1 MIDI
	006a  SP-606
	006b  SP-606
	006d  FANTOM-X
	006e  FANTOM-X
	0073  EDIROL UA-25
	0074  EDIROL UA-25
	0075  BOSS DR-880
	0076  BOSS DR-880
	007a  RD
	007b  RD
	007d  EDIROL UA-101
	0080  G-70
	0081  G-70
	0084  V-SYNTH XT
	0089  BOSS GT-PRO
	008b  EDIROL PC-50
	008c  EDIROL PC-50
	008d  EDIROL UA-101 USB1
	0092  EDIROL PC-80 WAVE
	0093  EDIROL PC-80 MIDI
	0096  EDIROL UA-1EX
	009a  EDIROL UM-3EX
	009d  EDIROL UM-1
	00a0  MD-P1
	00a2  Digital Piano
	00a3  EDIROL UA-4FX
	00a6  Juno-G
	00a9  MC-808
	00ad  SH-201
	00b2  VG-99
	00b3  VG-99
	00b7  BK-7m/VIMA JM-5/8
	00c2  SonicCell
	00c4  EDIROL M-16DX
	00c5  SP-555
	00c7  V-Synth GT
	00d1  Music Atelier
	00d3  M-380/400
	00da  BOSS GT-10
	00db  BOSS GT-10 Guitar Effects Processor
	00dc  BOSS GT-10B
	00de  Fantom G
	00e6  EDIROL UA-25EX (Advanced mode)
	00e7  EDIROL UA-25EX
	00e9  UA-1G
	00eb  VS-100
	00f6  GW-8/AX-Synth
	00f8  JUNO Series
	00fc  VS-700C
	00fd  VS-700
	00fe  VS-700 M1
	00ff  VS-700 M2
	0100  VS-700
	0101  VS-700 M2
	0102  VB-99
	0104  UM-1G
	0106  UM-2G
	0108  UM-3G
	0109  eBand JS-8
	010d  A-500S
	010f  A-PRO
	0110  A-PRO
	0111  GAIA SH-01
	0113  ME-25
	0114  SD-50
	0116  WAVE/MP3 RECORDER R-05
	0117  VS-20
	0119  OCTAPAD SPD-30
	011c  Lucina AX-09
	011e  BR-800
	0120  OCTA-CAPTURE
	0121  OCTA-CAPTURE
	0123  JUNO-Gi
	0124  M-300
	0127  GR-55
	012a  UM-ONE
	012b  DUO-CAPTURE
	012f  QUAD-CAPTURE
	0130  MICRO BR BR-80
	0132  TRI-CAPTURE
	0134  V-Mixer
	0138  Boss RC-300 (Audio mode)
	0139  Boss RC-300 (Storage mode)
	013a  JUPITER-80
	013e  R-26
	0145  SPD-SX
	014b  eBand JS-10
	014d  GT-100
	0150  TD-15
	0151  TD-11
	0154  JUPITER-50
	0156  A-Series
	0158  TD-30
	0159  DUO-CAPTURE EX
	015b  INTEGRA-7
	015d  R-88
	0505  EDIROL UA-101
0583  Padix Co., Ltd (Rockfire)
	0001  4 Axis 12 button +POV
	0002  4 Axis 12 button +POV
	2030  RM-203 USB Nest [mode 1]
	2031  RM-203 USB Nest [mode 2]
	2032  RM-203 USB Nest [mode 3]
	2033  RM-203 USB Nest [mode 4]
	2050  PX-205 PSX Bridge
	205f  PSX/USB converter
	206f  USB, 2-axis 8-button gamepad
	3050  QF-305u Gamepad
	3379  Rockfire X-Force
	337f  Rockfire USB RacingStar Vibra
	509f  USB,4-Axis,12-Button with POV
	5259  Rockfire USB SkyShuttle Vibra
	525f  USB Vibration Pad
	5308  USB Wireless VibrationPad
	5359  Rockfire USB SkyShuttle Pro
	535f  USB,real VibrationPad
	5659  Rockfire USB SkyShuttle Vibra
	565f  USB VibrationPad
	6009  Revenger
	600f  USB,GameBoard II
	6258  USB, 4-axis, 6-button joystick w/view finder
	6889  Windstorm Pro
	688f  QF-688uv Windstorm Pro Joystick
	7070  QF-707u Bazooka Joystick
	a000  MaxFire G-08XU Gamepad
	a015  4-Axis,16-Button with POV
	a019  USB, Vibration ,4-axis, 8-button joystick w/view finder
	a020  USB,4-Axis,10-Button with POV
	a021  USB,4-Axis,12-Button with POV
	a022  USB,4-Axis,14-Button with POV
	a023  USB,4-Axis,16-Button with POV
	a024  4axis,12button vibrition audio gamepad
	a025  4axis,12button vibrition audio gamepad
	a130  USB Wireless 2.4GHz Gamepad
	a131  USB Wireless 2.4GHz Joystick
	a132  USB Wireless 2.4GHz Wheelpad
	a133  USB Wireless 2.4GHz Wheel&Gamepad
	a202  ForceFeedbackWheel
	a209  MetalStrike FF
	b000  USB,4-Axis,12-Button with POV
	b001  USB,4-Axis,12-Button with POV
	b002  Vibration,12-Button USB Wheel
	b005  USB,12-Button Wheel
	b008  USB Wireless 2.4GHz Wheel
	b009  USB,12-Button  Wheel
	b00a  PSX/USB converter
	b00b  PSX/USB converter
	b00c  PSX/USB converter
	b00d  PSX/USB converter
	b00e  4-Axis,12-Button with POV
	b00f  USB,5-Axis,10-Button with POV
	b010  MetalStrike Pro
	b012  Wireless MetalStrike
	b013  USB,Wiress  2.4GHZ Joystick
	b016  USB,5-Axis,10-Button with POV
	b018  TW6 Wheel
	ff60  USB Wireless VibrationPad
0584  RATOC System, Inc.
	0008  Fujifilm MemoryCard ReaderWriter
	0220  U2SCX SCSI Converter
	0304  U2SCX-LVD (SCSI Converter)
	b000  REX-USB60
	b020  REX-USB60F
0585  FlashPoint Technology, Inc.
	0001  Digital Camera
	0002  Digital Camera
	0003  Digital Camera
	0004  Digital Camera
	0005  Digital Camera
	0006  Digital Camera
	0007  Digital Camera
	0008  Digital Camera
	0009  Digital Camera
	000a  Digital Camera
	000b  Digital Camera
	000c  Digital Camera
	000d  Digital Camera
	000e  Digital Camera
	000f  Digital Camera
0586  ZyXEL Communications Corp.
	0025  802.11b/g/n USB Wireless Network Adapter
	0100  omni.net
	0102  omni.net II ISDN TA [HFC-S]
	0110  omni.net Plus
	1000  omni.net LCD Plus - ISDN TA
	1500  Omni 56K Plus
	2011  Scorpion-980N keyboard
	3304  LAN Modem
	3309  ADSL Modem Prestige 600 series
	330a  ADSL Modem Interface
	330e  USB Broadband ADSL Modem Rev 1.10
	3400  ZyAIR B-220 IEEE 802.11b Adapter
	3401  ZyAIR G-220 802.11bg
	3402  ZyAIR G-220F 802.11bg
	3403  AG-200 802.11abg Wireless Adapter [Atheros AR5523]
	3407  G-200 v2 802.11bg
	3408  G-260 802.11bg
	3409  AG-225H 802.11bg
	340a  M-202 802.11bg
	340c  G-270S 802.11bg Wireless Adapter [Atheros AR5523]
	340f  G-220 v2 802.11bg
	3410  ZyAIR G-202 802.11bg
	3412  802.11bg
	3413  ZyAIR AG-225H v2 802.11bg
	3415  G-210H 802.11g Wireless Adapter
	3416  NWD-210N 802.11b/g/n-draft wireless adapter
	3417  NWD271N 802.11n Wireless Adapter [Atheros AR9001U-(2)NG]
	3418  NWD211AN 802.11abgn Wireless Adapter [Ralink RT2870]
	3419  G-220 v3 802.11bg Wireless Adapter [ZyDAS ZD1211B]
	341a  NWD-270N Wireless N-lite USB Adapter
	341e  NWD2105 802.11bgn Wireless Adapter [Ralink RT3070]
	341f  NWD2205 802.11n Wireless N Adapter [Realtek RTL8192CU]
	3425  NWD6505 802.11a/b/g/n/ac Wireless Adapter [MediaTek MT7610U]
	343e  N220 802.11bgn Wireless Adapter
0587  America Kotobuki Electronics Industries, Inc.
0588  Sapien Design
0589  Victron
058a  Nohau Corp.
058b  Infineon Technologies
	0015  Flash Loader utility
	001c  Flash Drive
	0041  Flash Loader utility
058c  In Focus Systems
	0007  Flash
	0008  LP130
	000a  LP530
	0010  Projector
	0011  Projector
	0012  Projector
	0013  Projector
	0014  Projector
	0015  Projector
	0016  Projector
	0017  Projector
	0018  Projector
	0019  Projector
	001a  Projector
	001b  Projector
	001c  Projector
	001d  Projector
	001e  Projector
	001f  Projector
	ffe5  IN34 Projector
058d  Micrel Semiconductor
058e  Tripath Technology, Inc.
058f  Alcor Micro Corp.
	1234  Flash Drive
	2412  SCard R/W CSR-145
	2802  Monterey Keyboard
	5492  Hub
	6232  Hi-Speed 16-in-1 Flash Card Reader/Writer
	6254  USB Hub
	6331  SD/MMC/MS Card Reader
	6332  Multi-Function Card Reader
	6335  SD/MMC Card Reader
	6360  Multimedia Card Reader
	6361  Multimedia Card Reader
	6362  Flash Card Reader/Writer
	6364  AU6477 Card Reader Controller
	6366  Multi Flash Reader
	6377  AU6375 4-LUN card reader
	6386  Memory Card
	6387  Flash Drive
	6390  USB 2.0-IDE bridge
	6391  IDE Bridge
	9213  MacAlly Kbd Hub
	9215  AU9814 Hub
	9254  Hub
	9310  Mass Storage (UID4/5A & UID7A)
	9320  Micro Storage Driver for Win98
	9321  Micro Storage Driver for Win98
	9330  SD Reader
	9331  Micro Storage Driver for Win98
	9340  Delkin eFilm Reader-32
	9350  Delkin eFilm Reader-32
	9360  8-in-1 Media Card Reader
	9361  Multimedia Card Reader
	9368  Multimedia Card Reader
	9380  Flash Drive
	9381  Flash Drive
	9382  Acer/Sweex Flash drive
	9384  qdi U2Disk T209M
	9410  Keyboard
	9472  Keyboard Hub
	9510  ChunghwaTL USB02 Smartcard Reader
	9520  EMV Certified Smart Card Reader
	9540  AU9540 Smartcard Reader
	9720  USB-Serial Adapter
	a014  Asus Integrated Webcam
	b002  Acer Integrated Webcam
0590  Omron Corp.
	0004  Cable Modem
	000b  MR56SVS
	0028  HJ-720IT / HEM-7080IT-E / HEM-790IT
0591  Questra Consulting
0592  Powerware Corp.
	0002  UPS (X-Slot)
0593  Incite
0594  Princeton Graphic Systems
0595  Zoran Microelectronics, Ltd
	1001  Digitrex DSC-1300/DSC-2100 (mass storage mode)
	2002  DIGITAL STILL CAMERA 6M 4X
	4343  Digital Camera EX-20 DSC
0596  MicroTouch Systems, Inc.
	0001  Touchscreen
	0002  Touch Screen Controller
	0500  PCT Multitouch HID Controller
	0543  DELL XPS touchscreen
0597  Trisignal Communications
0598  Niigata Canotec Co., Inc.
0599  Brilliance Semiconductor, Inc.
059a  Spectrum Signal Processing, Inc.
059b  Iomega Corp.
	0001  Zip 100 (Type 1)
	000b  Zip 100 (Type 2)
	0021  Win98 Disk Controller
	0030  Zip 250 (Ver 1)
	0031  Zip 100 (Type 3)
	0032  Zip 250 (Ver 2)
	0034  Zip 100 Driver
	0037  Zip 750 MB
	0040  SCSI Bridge
	0042  Rev 70 GB
	0050  Zip CD 650 Writer
	0053  CDRW55292EXT CD-RW External Drive
	0056  External CD-RW Drive Enclosure
	0057  Mass Storage Device
	005d  Mass Storage Device
	005f  CDRW64892EXT3-C CD-RW 52x24x52x External Drive
	0060  PCMCIA PocketZip Dock
	0061  Varo PocketZip 40 MP3 Player
	006d  HipZip MP3 Player
	0070  eGo Portable Hard Drive
	007c  Ultra Max USB/1394
	007d  HTC42606 0G9AT00 [Iomega HDD]
	007e  Mini 256MB/512MB Flash Drive [IOM2D5]
	00db  FotoShow Zip 250 Driver
	0150  Mass Storage Device
	015d  Super DVD Writer
	0173  Hi-Speed USB-to-IDE Bridge Controller
	0174  Hi-Speed USB-to-IDE Bridge Controller
	0176  Hi-Speed USB-to-IDE Bridge Controller
	0177  Hi-Speed USB-to-IDE Bridge Controller
	0178  Hi-Speed USB-to-IDE Bridge Controller
	0179  Hi-Speed USB-to-IDE Bridge Controller
	017a  HDD
	017b  HDD/1394A
	017c  HDD/1394B
	0251  Optical
	0252  Optical
	0278  LDHD-UPS [Professional Desktop Hard Drive eSATA / USB2.0]
	027a  LPHD250-U [Portable Hard Drive Silver Series 250 Go]
	0470  Prestige Portable Hard Drive
	047a  Select Portable Hard Drive
	0571  Prestige Portable Hard Drive
	0579  eGo Portable Hard Drive
	1052  DVD+RW External Drive
059c  A-Trend Technology Co., Ltd
059d  Advanced Input Devices
059e  Intelligent Instrumentation
059f  LaCie, Ltd
	0201  StudioDrive USB2
	0202  StudioDrive USB2
	0203  StudioDrive USB2
	0211  PocketDrive
	0212  PocketDrive
	0213  PocketDrive USB2
	0323  LaCie d2 Drive USB2
	0421  Big Disk G465
	0525  BigDisk Extreme 500
	0641  Mobile Hard Drive
	0829  BigDisk Extreme+
	100c  Rugged Triple Interface Mobile Hard Drive
	1010  Desktop Hard Drive
	1018  Desktop Hard Drive
	1019  Desktop Hard Drive
	1021  Little Disk
	1027  iamaKey V2
	102a  Rikiki Hard Drive
	1049  rikiki Harddrive
	1052  P'9220 Mobile Drive
	1064  Rugged 16 and 32 GB
	106e  Porsche Design Desktop Drive
	a601  HardDrive
	a602  CD R/W
05a0  Vetronix Corp.
05a1  USC Corp.
05a2  Fuji Film Microdevices Co., Ltd
05a3  ARC International
	8388  Marvell 88W8388 802.11a/b/g WLAN
05a4  Ortek Technology, Inc.
	1000  WKB-1000S Wireless Ergo Keyboard with Touchpad
	2000  WKB-2000 Wireless Keyboard with Touchpad
	9720  Keyboard Mouse
	9722  Keyboard
	9731  MCK-600W/MCK-800USB Keyboard
	9783  Wireless Keypad
	9837  Targus Number Keypad
	9862  Targus Number Keypad (Composite Device)
	9881  IR receiver [VRC-1100 Vista MCE Remote Control]
05a5  Sampo Technology Corp.
05a6  Cisco Systems, Inc.
	0001  CVA124 Cable Voice Adapter (WDM)
	0002  CVA122 Cable Voice Adapter (WDM)
	0003  CVA124E Cable Voice Adapter (WDM)
	0004  CVA122E Cable Voice Adapter (WDM)
05a7  Bose Corp.
	4000  Bluetooth Headset
	4001  Bluetooth Headset in DFU mode
	4002  Bluetooth Headset Series 2
	4003  Bluetooth Headset Series 2 in DFU mode
	bc50  SoundLink Wireless Mobile speaker
	bc51  SoundLink Wireless Mobile speaker in DFU mode
05a8  Spacetec IMC Corp.
05a9  OmniVision Technologies, Inc.
	0511  OV511 Webcam
	0518  OV518 Webcam
	0519  OV519 Microphone
	1550  VEHO Filmscanner
	2640  OV2640 Webcam
	2643  Monitor Webcam
	264b  Monitor Webcam
	2800  SuperCAM
	4519  Webcam Classic
	7670  OV7670 Webcam
	8065  GAIA Sensor FPGA Demo Board
	8519  OV519 Webcam
	a511  OV511+ Webcam
	a518  D-Link DSB-C310 Webcam
05aa  Utilux South China, Ltd
05ab  In-System Design
	0002  Parallel Port
	0030  Storage Adapter V2 (TPP)
	0031  ATA Bridge
	0060  USB 2.0 ATA Bridge
	0061  Storage Adapter V3 (TPP-I)
	0101  Storage Adapter (TPP)
	0130  Compact Flash and Microdrive Reader (TPP)
	0200  USS725 ATA Bridge
	0201  Storage Adapter (TPP)
	0202  ATA Bridge
	0300  Portable Hard Drive (TPP)
	0301  Portable Hard Drive V2
	0350  Portable Hard Drive (TPP)
	0351  Portable Hard Drive V2
	081a  ATA Bridge
	0cda  ATA Bridge for CD-R/RW
	1001  BAYI Printer Class Support
	5700  Storage Adapter V2 (TPP)
	5701  USB Storage Adapter V2
	5901  Smart Board (TPP)
	5a01  ATI Storage Adapter (TPP)
	5d01  DataBook Adapter (TPP)
05ac  Apple, Inc.
	0201  USB Keyboard [Alps or Logitech, M2452]
	0202  Keyboard [ALPS]
	0205  Extended Keyboard [Mitsumi]
	0206  Extended Keyboard [Mitsumi]
	020b  Pro Keyboard [Mitsumi, A1048/US layout]
	020c  Extended Keyboard [Mitsumi]
	020d  Pro Keyboard [Mitsumi, A1048/JIS layout]
	020e  Internal Keyboard/Trackpad (ANSI)
	020f  Internal Keyboard/Trackpad (ISO)
	0214  Internal Keyboard/Trackpad (ANSI)
	0215  Internal Keyboard/Trackpad (ISO)
	0216  Internal Keyboard/Trackpad (JIS)
	0217  Internal Keyboard/Trackpad (ANSI)
	0218  Internal Keyboard/Trackpad (ISO)
	0219  Internal Keyboard/Trackpad (JIS)
	021a  Internal Keyboard/Trackpad (ANSI)
	021b  Internal Keyboard/Trackpad (ISO)
	021c  Internal Keyboard/Trackpad (JIS)
	021d  Aluminum Mini Keyboard (ANSI)
	021e  Aluminum Mini Keyboard (ISO)
	021f  Aluminum Mini Keyboard (JIS)
	0220  Aluminum Keyboard (ANSI)
	0221  Aluminum Keyboard (ISO)
	0222  Aluminum Keyboard (JIS)
	0223  Internal Keyboard/Trackpad (ANSI)
	0224  Internal Keyboard/Trackpad (ISO)
	0225  Internal Keyboard/Trackpad (JIS)
	0229  Internal Keyboard/Trackpad (ANSI)
	022a  Internal Keyboard/Trackpad (MacBook Pro) (ISO)
	022b  Internal Keyboard/Trackpad (MacBook Pro) (JIS)
	0230  Internal Keyboard/Trackpad (MacBook Pro 4,1) (ANSI)
	0231  Internal Keyboard/Trackpad (MacBook Pro 4,1) (ISO)
	0232  Internal Keyboard/Trackpad (MacBook Pro 4,1) (JIS)
	0236  Internal Keyboard/Trackpad (ANSI)
	0237  Internal Keyboard/Trackpad (ISO)
	0238  Internal Keyboard/Trackpad (JIS)
	023f  Internal Keyboard/Trackpad (ANSI)
	0240  Internal Keyboard/Trackpad (ISO)
	0241  Internal Keyboard/Trackpad (JIS)
	0242  Internal Keyboard/Trackpad (ANSI)
	0243  Internal Keyboard/Trackpad (ISO)
	0244  Internal Keyboard/Trackpad (JIS)
	0245  Internal Keyboard/Trackpad (ANSI)
	0246  Internal Keyboard/Trackpad (ISO)
	0247  Internal Keyboard/Trackpad (JIS)
	024a  Internal Keyboard/Trackpad (MacBook Air) (ISO)
	024d  Internal Keyboard/Trackpad (MacBook Air) (ISO)
	0250  Aluminium Keyboard (ISO)
	0252  Internal Keyboard/Trackpad (ANSI)
	0253  Internal Keyboard/Trackpad (ISO)
	0254  Internal Keyboard/Trackpad (JIS)
	0263  Apple Internal Keyboard / Trackpad (MacBook Retina)
	0301  USB Mouse [Mitsumi, M4848]
	0302  Optical Mouse [Fujitsu]
	0304  Mighty Mouse [Mitsumi, M1152]
	0306  Optical USB Mouse [Fujitsu]
	030a  Internal Trackpad
	030b  Internal Trackpad
	030d  Magic Mouse
	030e  MC380Z/A [Magic Trackpad]
	1000  Bluetooth HCI MacBookPro (HID mode)
	1001  Keyboard Hub [ALPS]
	1002  Extended Keyboard Hub [Mitsumi]
	1003  Hub in Pro Keyboard [Mitsumi, A1048]
	1006  Hub in Aluminum Keyboard
	1008  Mini DisplayPort to Dual-Link DVI Adapter
	1101  Speakers
	1105  Audio in LED Cinema Display
	1107  Thunderbolt Display Audio
	1112  FaceTime HD Camera (Display)
	1201  3G iPod
	1202  iPod 2G
	1203  iPod 4.Gen Grayscale 40G
	1204  iPod [Photo]
	1205  iPod Mini 1.Gen/2.Gen
	1206  iPod '06'
	1207  iPod '07'
	1208  iPod '08'
	1209  iPod Video
	120a  iPod Nano
	1223  iPod Classic/Nano 3.Gen (DFU mode)
	1224  iPod Nano 3.Gen (DFU mode)
	1225  iPod Nano 4.Gen (DFU mode)
	1227  Mobile Device (DFU Mode)
	1231  iPod Nano 5.Gen (DFU mode)
	1240  iPod Nano 2.Gen (DFU mode)
	1242  iPod Nano 3.Gen (WTF mode)
	1243  iPod Nano 4.Gen (WTF mode)
	1245  iPod Classic 3.Gen (WTF mode)
	1246  iPod Nano 5.Gen (WTF mode)
	1255  iPod Nano 4.Gen (DFU mode)
	1260  iPod Nano 2.Gen
	1261  iPod Classic
	1262  iPod Nano 3.Gen
	1263  iPod Nano 4.Gen
	1265  iPod Nano 5.Gen
	1266  iPod Nano 6.Gen
	1267  iPod Nano 7.Gen
	1281  Apple Mobile Device [Recovery Mode]
	1290  iPhone
	1291  iPod Touch 1.Gen
	1292  iPhone 3G
	1293  iPod Touch 2.Gen
	1294  iPhone 3GS
	1296  iPod Touch 3.Gen (8GB)
	1297  iPhone 4
	1299  iPod Touch 3.Gen
	129a  iPad
	129c  iPhone 4(CDMA)
	129e  iPod Touch 4.Gen
	129f  iPad 2
	12a0  iPhone 4S
	12a2  iPad 2 (3G; 64GB)
	12a3  iPad 2 (CDMA)
	12a4  iPad 3 (wifi)
	12a5  iPad 3 (CDMA)
	12a6  iPad 3 (3G, 16 GB)
	12a8  iPhone5/5C/5S/6
	12a9  iPad 2
	12aa  iPod Touch 5.Gen [A1421]
	12ab  iPad 4/Mini1
	1300  iPod Shuffle
	1301  iPod Shuffle 2.Gen
	1302  iPod Shuffle 3.Gen
	1303  iPod Shuffle 4.Gen
	1401  Modem
	1402  Ethernet Adapter [A1277]
	1500  SuperDrive [A1379]
	8005  OHCI Root Hub Simulation
	8006  EHCI Root Hub Simulation
	8007  XHCI Root Hub USB 2.0 Simulation
	8202  HCF V.90 Data/Fax Modem
	8203  Bluetooth HCI
	8204  Built-in Bluetooth 2.0+EDR HCI
	8205  Bluetooth HCI
	8206  Bluetooth HCI
	820a  Bluetooth HID Keyboard
	820b  Bluetooth HID Mouse
	820f  Bluetooth HCI
	8213  Bluetooth Host Controller
	8215  Built-in Bluetooth 2.0+EDR HCI
	8216  Bluetooth USB Host Controller
	8217  Bluetooth USB Host Controller
	8218  Bluetooth Host Controller
	821a  Bluetooth Host Controller
	821f  Built-in Bluetooth 2.0+EDR HCI
	8240  Built-in IR Receiver
	8241  Built-in IR Receiver
	8242  Built-in IR Receiver
	8281  Bluetooth Host Controller
	8286  Bluetooth Host Controller
	828c  Bluetooth Host Controller
	8300  Built-in iSight (no firmware loaded)
	8403  Internal Memory Card Reader
	8404  Internal Memory Card Reader
	8501  Built-in iSight [Micron]
	8502  Built-in iSight
	8505  Built-in iSight
	8507  Built-in iSight
	8508  iSight in LED Cinema Display
	8509  FaceTime HD Camera
	850a  FaceTime Camera
	8510  FaceTime HD Camera (Built-in)
	911c  Hub in A1082 [Cinema HD Display 23"]
	9127  Hub in Thunderbolt Display
	912f  Hub in 30" Cinema Display
	9215  Studio Display 15"
	9217  Studio Display 17"
	9218  Cinema Display 23"
	9219  Cinema Display 20"
	921c  A1082 [Cinema HD Display 23"]
	921e  Cinema Display 24"
	9221  30" Cinema Display
	9226  LED Cinema Display
	9227  Thunderbolt Display
	9232  Cinema HD Display 30"
	ffff  Bluetooth in DFU mode - Driver
05ad  Y.C. Cable U.S.A., Inc.
05ae  Synopsys, Inc.
05af  Jing-Mold Enterprise Co., Ltd
	0806  HP SK806A Keyboard
	0809  Wireless Keyboard and Mouse
	0821  IDE to
	3062  Cordless Keyboard
	9167  KB 9151B - 678
	9267  KB 9251B - 678 Mouse
05b0  Fountain Technologies, Inc.
05b1  First International Computer, Inc.
	1389  Bluetooth Wireless Adapter
05b4  LG Semicon Co., Ltd
	4857  M-Any DAH-210
	6001  HYUNDAI GDS30C6001 SSFDC / MMC I/F Controller
05b5  Dialogic Corp.
05b6  Proxima Corp.
05b7  Medianix Semiconductor, Inc.
05b8  Agiler, Inc.
	3002  Scroll Mouse
05b9  Philips Research Laboratories
05ba  DigitalPersona, Inc.
	0007  Fingerprint Reader
	0008  Fingerprint Reader
	000a  Fingerprint Reader
05bb  Grey Cell Systems
05bc  3G Green Green Globe Co., Ltd
	0004  Trackball
05bd  RAFI GmbH & Co. KG
05be  Tyco Electronics (Raychem)
05bf  S & S Research
05c0  Keil Software
05c1  Kawasaki Microelectronics, Inc.
05c2  Media Phonics (Suisse) S.A.
05c5  Digi International, Inc.
	0002  AccelePort USB 2
	0004  AccelePort USB 4
	0008  AccelePort USB 8
05c6  Qualcomm, Inc.
	0114  Select RW-200 CDMA Wireless Modem
	1000  Mass Storage Device
	3100  CDMA Wireless Modem/Phone
	3196  CDMA Wireless Modem
	3197  CDMA Wireless Modem/Phone
	6000  Siemens SG75
	6503  AnyData APE-540H
	6613  Onda H600/N501HS ZTE MF330
	6764  A0001 Phone [OnePlus One]
	9000  SIMCom SIM5218 modem
	9001  Gobi Wireless Modem
	9002  Gobi Wireless Modem
	9003  Quectel UC20
	9008  Gobi Wireless Modem (QDL mode)
	9018  Qualcomm HSUSB Device
	9025  Qualcomm HSUSB Device
	9201  Gobi Wireless Modem (QDL mode)
	9202  Gobi Wireless Modem
	9203  Gobi Wireless Modem
	9205  Gobi 2000
	9211  Acer Gobi Wireless Modem (QDL mode)
	9212  Acer Gobi Wireless Modem
	9214  Acer Gobi 2000 Wireless Modem (QDL mode)
	9215  Acer Gobi 2000 Wireless Modem
	9221  Gobi Wireless Modem (QDL mode)
	9222  Gobi Wireless Modem
	9224  Sony Gobi 2000 Wireless Modem (QDL mode)
	9225  Sony Gobi 2000 Wireless Modem
	9231  Gobi Wireless Modem (QDL mode)
	9234  Top Global Gobi 2000 Wireless Modem (QDL mode)
	9235  Top Global Gobi 2000 Wireless Modem
	9244  Samsung Gobi 2000 Wireless Modem (QDL mode)
	9245  Samsung Gobi 2000 Wireless Modem
	9264  Asus Gobi 2000 Wireless Modem (QDL mode)
	9265  Asus Gobi 2000 Wireless Modem
	9274  iRex Technologies Gobi 2000 Wireless Modem (QDL mode)
	9275  iRex Technologies Gobi 2000 Wireless Modem
05c7  Qtronix Corp.
	0113  PC Line Mouse
	1001  Lynx Mouse
	2001  Keyboard
	2011  SCorpius Keyboard
	6001  Ten-Keypad
05c8  Cheng Uei Precision Industry Co., Ltd (Foxlink)
	0103  FO13FF-65 PC-CAM
	021a  HP Webcam
	0318  Webcam
	0361  SunplusIT INC. HP Truevision HD Webcam
	036e  Webcam
	0403  Webcam
	041b  HP 2.0MP High Definition Webcam
05c9  Semtech Corp.
05ca  Ricoh Co., Ltd
	0101  RDC-5300 Camera
	0325  Caplio GX (ptp)
	032d  Caplio GX 8 (ptp)
	032f  Caplio R3 (ptp)
	03a1  IS200e
	0403  Printing Support
	0405  Type 101
	0406  Type 102
	1803  V5 camera [R5U870]
	1810  Pavilion Webcam [R5U870]
	1812  Pavilion Webcam
	1814  HD Webcam
	1815  Dell Laptop Integrated Webcam
	1820  Integrated Webcam
	1830  Visual Communication Camera VGP-VCC2 [R5U870]
	1832  Visual Communication Camera VGP-VCC3 [R5U870]
	1833  Visual Communication Camera VGP-VCC2 [R5U870]
	1834  Visual Communication Camera VGP-VCC2 [R5U870]
	1835  Visual Communication Camera VGP-VCC5 [R5U870]
	1836  Visual Communication Camera VGP-VCC4 [R5U870]
	1837  Visual Communication Camera VGP-VCC4 [R5U870]
	1839  Visual Communication Camera VGP-VCC6 [R5U870]
	183a  Visual Communication Camera VGP-VCC7 [R5U870]
	183b  Visual Communication Camera VGP-VCC8 [R5U870]
	183d  Sony Vaio Integrated Webcam
	183e  Visual Communication Camera VGP-VCC9 [R5U870]
	1841  Fujitsu F01/ Lifebook U810 [R5U870]
	1870  Webcam 1000
	18b0  Sony Vaio Integrated Webcam
	18b1  Sony Vaio Integrated Webcam
	18b3  Sony Vaio Integrated Webcam
	18b5  Sony Vaio Integrated Webcam
	2201  RDC-7 Camera
	2202  Caplio RR30
	2203  Caplio 300G
	2204  Caplio G3
	2205  Caplio RR30 / Medion MD 6126 Camera
	2206  Konica DG-3Z
	2207  Caplio Pro G3
	2208  Caplio G4
	2209  Caplio 400G wide
	220a  KONICA MINOLTA DG-4Wide
	220b  Caplio RX
	220c  Caplio GX
	220d  Caplio R1/RZ1
	220e  Sea & Sea 5000G
	220f  Rollei dr5 / Rollei dr5 (PTP mode)
	2211  Caplio R1S
	2212  Caplio R1v Camera
	2213  Caplio R2
	2214  Caplio GX 8
	2215  DSC 725
	2216  Caplio R3
	2222  RDC-i500
05cb  PowerVision Technologies, Inc.
	1483  PV8630 interface (scanners, webcams)
05cc  ELSA AG
	2100  MicroLink ISDN Office
	2219  MicroLink ISDN
	2265  MicroLink 56k
	2267  MicroLink 56k (V.250)
	2280  MicroLink 56k Fun
	3000  Micolink USB2Ethernet [pegasus]
	3100  AirLancer USB-11
	3363  MicroLink ADSL Fun
05cd  Silicom, Ltd
05ce  sci-worx GmbH
05cf  Sung Forn Co., Ltd
05d0  GE Medical Systems Lunar
05d1  Brainboxes, Ltd
	0003  Bluetooth Adapter BL-554
05d2  Wave Systems Corp.
05d3  Tohoku Ricoh Co., Ltd
05d5  Super Gate Technology Co., Ltd
05d6  Philips Semiconductors, CICT
05d7  Thomas & Betts Corp.
	0099  10Mbps Ethernet [klsi]
05d8  Ultima Electronics Corp.
	4001  Artec Ultima 2000
	4002  Artec Ultima 2000 (GT6801 based)/Lifetec LT9385/ScanMagic 1200 UB Plus Scanner
	4003  Artec E+ 48U
	4004  Artec E+ Pro
	4005  MEM48U
	4006  TRUST EASY WEBSCAN 19200
	4007  TRUST 240H EASY WEBSCAN GOLD
	4008  Trust Easy Webscan 19200
	4009  Umax Astraslim
	4013  IT Scan 1200
	8105  Artec T1 USB TVBOX (cold)
	8106  Artec T1 USB TVBOX (warm)
	8107  Artec T1 USB TVBOX with AN2235 (cold)
	8108  Artec T1 USB TVBOX with AN2235 (warm)
	8109  Artec T1 USB2.0 TVBOX (cold
05d9  Axiohm Transaction Solutions
	a225  A225 Printer
	a758  A758 Printer
	a794  A794 Printer
05da  Microtek International, Inc.
	0091  ScanMaker X6u
	0093  ScanMaker V6USL
	0094  Phantom 336CX/C3
	0099  ScanMaker X6/X6U
	009a  Phantom C6
	00a0  Phantom 336CX/C3 (#2)
	00a3  ScanMaker V6USL
	00ac  ScanMaker V6UL
	00b6  ScanMaker V6UPL
	00ef  ScanMaker V6UPL
	1006  Jenoptik JD350 entrance
	1011  NHJ Che-ez! Kiss Digital Camera
	1018  Digital Dream Enigma 1.3
	1020  Digital Dream l'espion xtra
	1025  Take-it Still Camera Device
	1026  Take-it
	1043  Take-It 1300 DSC Bulk Driver
	1045  Take-it D1
	1047  Take-it Camera Composite Device
	1048  Take-it Q3
	1049  3M Still Camera Device
	1051  Camcorder Series
	1052  Mass Storage Device
	1053  Take-it DV Composite Device
	1054  Mass Storage Device
	1055  Digital Camera Series(536)
	1056  Mass Storage Device
	1057  Take-it DSC Camera Device(536)
	1058  Mass Storage Device
	1059  Camcorder DSC Series
	1060  Microtek Take-it MV500
	2007  ArtixScan DI 1210
	200c  1394_USB2 Scanner
	200e  ArtixScan DI 810
	2017  UF ICE Scanner
	201c  4800 Scanner
	201d  ArtixScan DI 1610
	201f  4800 Scanner-ICE
	202e  ArtixScan DI 2020
	208b  ScanMaker 6800
	208f  ArtixScan DI 2010
	209e  ScanMaker 4700LP
	20a7  ScanMaker 5600
	20b0  ScanMaker X12USL
	20b1  ScanMaker 8700
	20b4  ScanMaker 4700
	20bd  ScanMaker 5700
	20c9  ScanMaker 6700
	20d2  Microtek ArtixScan 1800f
	20d6  PS4000
	20de  ScanMaker 9800XL
	20e0  ScanMaker 9700XL
	20ed  ScanMaker 4700
	20ee  Micortek ScanMaker X12USL
	2838  RT2832U
	3008  Scanner
	300a  4800 ICE Scanner
	300b  4800 Scanner
	300f  MiniScan C5
	3020  4800dpi Scanner
	3021  1200dpi Scanner
	3022  Scanner 4800dpi
	3023  USB1200II Scanner
	30c1  USB600 Scanner
	30ce  ScanMaker 3800
	30cf  ScanMaker 4800
	30d4  USB1200 Scanner
	30d8  Scanner
	30d9  USB2400 Scanner
	30e4  ScanMaker 4100
	30e5  USB3200 Scanner
	30e6  ScanMaker i320
	40b3  ScanMaker 3600
	40b8  ScanMaker 3700
	40c7  ScanMaker 4600
	40ca  ScanMaker 3600
	40cb  ScanMaker 3700
	40dd  ScanMaker 3750i
	40ff  ScanMaker 3600
	5003  Goya
	5013  3200 Scanner
	6072  XT-3500 A4 HD Scanner
	80a3  ScanMaker V6USL (#2)
	80ac  ScanMaker V6UL/SpicyU
05db  Sun Corp. (Suntac?)
	0003  SUNTAC U-Cable type D2
	0005  SUNTAC U-Cable type P1
	0009  SUNTAC Slipper U
	000a  SUNTAC Ir-Trinity
	000b  SUNTAC U-Cable type A3
	0011  SUNTAC U-Cable type A4
05dc  Lexar Media, Inc.
	0001  jumpSHOT CompactFlash Reader
	0002  JumpShot
	0003  JumpShot
	0080  Jumpdrive Secure 64MB
	0081  RBC Compact Flash Drive
	00a7  JumpDrive Impact
	0100  JumpDrive PRO
	0200  JumpDrive 2.0 Pro
	0300  Jumpdrive Geysr
	0301  JumpDrive Classic
	0302  JD Micro
	0303  JD Micro Pro
	0304  JD Secure II
	0310  JumpDrive
	0311  JumpDrive Classic
	0312  JD Micro
	0313  JD Micro Pro
	0320  JumpDrive
	0321  JD Micro
	0322  JD Micro Pro
	0323  UFC
	0330  JumpDrive Expression
	0340  JumpDrive TAD
	0350  Express Card
	0400  UFDC
	0401  UFDC
	0403  Locked B Device
	0405  Locked C Device
	0407  Locked D Device
	0409  Locked E Device
	040b  Locked F Device
	040d  Locked G Device
	040f  Locked H Device
	0410  JumpDrive
	0411  JumpDrive
	0413  Locked J Device
	0415  Locked K Device
	0417  Locked L Device
	0419  Locked M Device
	041b  Locked N Device
	041d  Locked O Device
	041f  Locked P Device
	0420  JumpDrive
	0421  JumpDrive
	0423  Locked R Device
	0425  Locked S Device
	0427  Locked T Device
	0429  Locked U Device
	042b  Locked V Device
	042d  Locked W Device
	042f  Locked X Device
	0431  Locked Y Device
	0433  Locked Z Device
	4d02  MP3 Player
	4d12  MP3 Player
	4d30  MP3 Player
	a209  JumpDrive S70
	a300  JumpDrive2
	a400  JumpDrive trade; Pro 40-501
	a410  JumpDrive 128MB/256MB
	a411  JumpDrive Traveler
	a420  JumpDrive Pro
	a421  JumpDrive Pro II
	a422  JumpDrive Micro Pro
	a430  JumpDrive Secure
	a431  JumpDrive Secure II
	a432  JumpDrive Classic
	a440  JumpDrive Lightning
	a450  JumpDrive TouchGuard
	a460  JD Mercury
	a501  JumpDrive Classic
	a510  JumpDrive Sport
	a530  JumpDrive Expression
	a531  JumpDrive Secure II
	a560  JumpDrive FireFly
	a701  JumpDrive FireFly
	a731  JumpDrive FireFly
	a768  JumpDrive Retrax
	a790  JumpDrive 2GB
	a811  16GB Gizmo!
	a813  16gB flash thumb drive
	a815  JumpDrive V10
	a833  JumpDrive S23 64GB
	b002  USB CF Reader
	b018  Multi-Card Reader
	b047  SDHC Reader [RW047-7000]
	ba02  Workflow CFR1
	c753  JumpDrive TwistTurn
05dd  Delta Electronics, Inc.
	ff31  AWU-120
	ff32  FriendlyNET AeroLAN AL2011
	ff35  PCW 100 - Wireless 802.11b Adapter
	ff91  2Wire PC Port Phoneline 10Mbps Adapter
05df  Silicon Vision, Inc.
05e0  Symbol Technologies
	0700  Bar Code Scanner (CS1504)
	0800  Spectrum24 Wireless LAN Adapter
	1200  Bar Code Scanner
	1701  Bar Code Scanner (CDC)
	1900  SNAPI Imaging Device
	2000  MC3090 Rugged Mobile Computer
	200d  MC70 Rugged Mobile Computer
05e1  Syntek Semiconductor Co., Ltd
	0100  802.11g + Bluetooth Wireless Adapter
	0408  STK1160 Video Capture Device
	0500  DC-112X Webcam
	0501  DC-1125 Webcam
	0890  STK011 Camera
	0892  STK013 Camera
	0895  STK016 Camera
	0896  STK017 Camera
	2010  ARCTIC Sound P261 Headphones
05e2  ElecVision, Inc.
05e3  Genesys Logic, Inc.
	000a  Keyboard with PS/2 Port
	000b  Mouse
	0100  Nintendo Game Boy Advance SP
	0120  Pacific Image Electronics PrimeFilm 1800u slide/negative scanner
	0131  CF/SM Reader/Writer
	0142  Multiple Slides Scanner-3600
	0143  Multiple Frames Film Scanner-36series
	0180  Plustek Scanner
	0182  Wize Media 1000
	0189  ScanJet 4600 series
	018a  Xerox 6400
	0300  GLUSB98PT Parallel Port
	0301  USB2LPT Cable Release2
	0406  Hub
	0501  GL620USB Host-Host interface
	0502  GL620USB-A GeneLink USB-USB Bridge
	0503  Webcam
	0504  HID Keyboard Filter
	0604  USB 1.1 Hub
	0605  USB 2.0 Hub
	0606  USB 2.0 Hub / D-Link DUB-H4 USB 2.0 Hub
	0607  Logitech G110 Hub
	0608  Hub
	0610  4-port hub
	0616  hub
	0660  USB 2.0 Hub
	0700  SIIG US2256 CompactFlash Card Reader
	0701  USB 2.0 IDE Adapter
	0702  USB 2.0 IDE Adapter [GL811E]
	0703  Card Reader
	0704  Card Reader
	0705  Card Reader
	0706  Card Reader
	0707  Card Reader
	0708  Card Reader
	0709  Card Reader
	070a  Pen Flash
	070b  DMHS1B Rev 3 DFU Adapter
	070e  USB 2.0 Card Reader
	070f  Pen Flash
	0710  USB 2.0 33-in-1 Card Reader
	0711  Card Reader
	0712  Delkin Mass Storage Device
	0715  USB 2.0 microSD Reader
	0716  USB 2.0 Multislot Card Reader/Writer
	0717  All-in-1 Card Reader
	0718  IDE/SATA Adapter
	0719  SATA adapter
	0722  SD/MMC card reader
	0723  GL827L SD/MMC/MS Flash Card Reader
	0726  SD Card Reader
	0727  microSD Reader/Writer
	0731  GL3310 SATA 3Gb/s Bridge Controller
	0732  All-in-One Cardreader
	0736  microSD Reader/Writer
	0741  microSD Card Reader
	0743  SDXC and microSDXC CardReader
	0745  Logilink CR0012
	0760  USB 2.0 Card Reader/Writer
	0761  Genesys Mass Storage Device
	0780  USBFS DFU Adapter
	07a0  Pen Flash
	0880  Wasp (SL-6612)
	0927  Card Reader
	1205  Afilias Optical Mouse H3003 / Trust Optical USB MultiColour Mouse MI-2330
	a700  Pen Flash
	f102  VX7012 TV Box
	f103  VX7012 TV Box
	f104  VX7012 TV Box
	fd21  3M TL20 Temperature Logger
	fe00  Razer Mouse
05e4  Red Wing Corp.
05e5  Fuji Electric Co., Ltd
05e6  Keithley Instruments
05e8  ICC, Inc.
05e9  Kawasaki LSI
	0008  KL5KUSB101B Ethernet [klsi]
	0009  Sony 10Mbps Ethernet [pegasus]
	000c  USB-to-RS-232
	000d  USB-to-RS-232
	0014  RS-232 J104
	0040  Ethernet Adapter
	2008  Ethernet Adapter
05eb  FFC, Ltd
05ec  COM21, Inc.
05ee  Cytechinfo Inc.
05ef  AVB, Inc. [anko?]
	020a  Top Shot Pegasus Joystick
	8884  Mag Turbo Force Wheel
	8888  Top Shot Force Feedback Racing Wheel
05f0  Canopus Co., Ltd
	0101  DA-Port DAC
05f1  Compass Communications
05f2  Dexin Corp., Ltd
	0010  AQ Mouse
05f3  PI Engineering, Inc.
	0007  Kinesis Advantage PRO MPC/USB Keyboard
	0081  Kinesis Integrated Hub
	00ff  VEC Footpedal
	0203  Y-mouse Keyboard & Mouse Adapter
	020b  PS2 Adapter
	0232  X-Keys Switch Interface, Programming Mode
	0261  X-Keys Switch Interface, SPLAT Mode
	0264  X-Keys Switch Interface, Composite Mode
05f5  Unixtar Technology, Inc.
05f6  AOC International
05f7  RFC Distribution(s) PTE, Ltd
05f9  PSC Scanning, Inc.
	1104  Magellan 2200VS
	1206  Gryphon series (OEM mode)
	2202  Point of Sale Handheld Scanner
	2206  Gryphon series (keyboard emulation mode)
	220c  Datalogic Gryphon GD4430
	2601  Datalogic Magellan 1000i Barcode Scanner
	2602  Datalogic Magellan 1100i Barcode Scanner
	4204  Gryphon series (RS-232 emulation mode)
	5204  Datalogic Gryphon GFS4170 (config mode)
05fa  Siemens Telecommunications Systems, Ltd
	3301  Keyboard with PS/2 Mouse Port
	3302  Keyboard
	3303  Keyboard with PS/2 Mouse Port
05fc  Harman
	0001  Soundcraft Si Multi Digital Card
	7849  Harman/Kardon SoundSticks
05fd  InterAct, Inc.
	0239  SV-239 HammerHead Digital
	0251  Raider Pro
	0253  ProPad 8 Digital
	0286  SV-286 Cyclone Digital
	107a  PowerPad Pro X-Box pad
	262a  3dfx HammerHead FX
	262f  HammerHead Fx
	daae  Game Shark
05fe  Chic Technology Corp.
	0001  Mouse
	0003  Cypress USB Mouse
	0005  Viewmaster 4D Browser Mouse
	0007  Twinhead Mouse
	0009  Inland Pro 4500/5000 Mouse
	0011  Browser Mouse
	0014  Gamepad
	1010  Optical Wireless
	2001  Microsoft Wireless Receiver 700
05ff  LeCroy Corp.
0600  Barco Display Systems
0601  Jazz Hipster Corp.
	0003  Internet Security Co., Ltd. SecureKey
0602  Vista Imaging, Inc.
	1001  ViCam Webcam
0603  Novatek Microelectronics Corp.
	00f1  Keyboard (Labtec Ultra Flat Keyboard)
	00f2  Keyboard (Labtec Ultra Flat Keyboard)
	6871  Mouse
0604  Jean Co., Ltd
0605  Anchor C&C Co., Ltd
0606  Royal Information Electronics Co., Ltd
0607  Bridge Information Co., Ltd
0608  Genrad Ads
0609  SMK Manufacturing, Inc.
	031d  eHome Infrared Receiver
	0322  eHome Infrared Receiver
	0334  eHome Infrared Receiver
	ff12  SMK Bluetooth Device
060a  Worthington Data Solutions, Inc.
060b  Solid Year
	0001  MacAlly Keyboard
	0230  KSK-8003 UX Keyboard
	1006  Japanese Keyboard - 260U
	2101  Keyboard
	2231  KSK-6001 UELX Keyboard
	2270  Gigabyte K8100 Aivia Gaming Keyboard
	5253  Thermaltake MEKA G-Unit Gaming Keyboard
	5811  ACK-571U Wireless Keyboard
	5903  Japanese Keyboard - 595U
	6001  SolidTek USB 2p HUB
	6002  SolidTek USB Keyboard
	6003  Japanese Keyboard - 600HM
	6231  Thermaltake eSPORTS Meka Keyboard
	8007  P-W1G1F12 VER:1 [Macally MegaCam]
	a001  Maxwell Compact Pc PM3
060c  EEH Datalink GmbH
060d  Auctor Corp.
060e  Transmonde Technologies, Inc.
060f  Joinsoon Electronics Mfg. Co., Ltd
0610  Costar Electronics, Inc.
0611  Totoku Electric Co., Ltd
0613  TransAct Technologies, Inc.
0614  Bio-Rad Laboratories
0615  Quabbin Wire & Cable Co., Inc.
0616  Future Techno Designs PVT, Ltd
0617  Swiss Federal Insitute of Technology
0618  MacAlly
	0101  Mouse
0619  Seiko Instruments, Inc.
	0101  SLP-100 Driver
	0102  SLP-200 Driver
	0103  SLP-100N Driver
	0104  SLP-200N Driver
	0105  SLP-240 Driver
	0501  SLP-440 Driver
	0502  SLP-450 Driver
061a  Veridicom International, Inc.
	0110  5thSense Fingerprint Sensor
	0200  FPS200 Fingerprint Sensor
	8200  VKI-A Fingerprint Sensor/Flash Storage (dumb)
	9200  VKI-B Fingerprint Sensor/Flash Storage (smart)
061b  Promptus Communications, Inc.
061c  Act Labs, Ltd
061d  Quatech, Inc.
	c020  SSU-100
061e  Nissei Electric Co.
	0001  nissei 128DE-USB -
	0010  nissei 128DE-PNA -
0620  Alaris, Inc.
	0004  QuickVideo weeCam
	0007  QuickVideo weeCam
	000a  QuickVideo weeCam
	000b  QuickVideo weeCam
0621  ODU-Steckverbindungssysteme GmbH & Co. KG
0622  Iotech, Inc.
0623  Littelfuse, Inc.
0624  Avocent Corp.
	0248  Virtual Hub
	0249  Virtual Keyboard/Mouse
	0251  Virtual Mass Storage
	0294  Dell 03R874 KVM dongle
	0402  Cisco Virtual Keyboard and Mouse
	0403  Cisco Virtual Mass Storage
0625  TiMedia Technology Co., Ltd
0626  Nippon Systems Development Co., Ltd
0627  Adomax Technology Co., Ltd
0628  Tasking Software, Inc.
0629  Zida Technologies, Ltd
062a  Creative Labs
	0000  Optical mouse
	0001  Notebook Optical Mouse
	0102  Wireless Keyboard/Mouse Combo [MK1152WC]
	0201  Defender Office Keyboard (K7310) S Zodiak KM-9010
	0252  Emerge Uni-retractable Laser Mouse
	3286  Nano Receiver [Sandstrom Laser Mouse SMWLL11]
	4101  Wireless Keyboard/Mouse
	6301  Trust Wireless Optical Mouse MI-4150K
	9003  VoIP Conference Hub (A16GH)
	9004  USR9602 USB Internet Mini Phone
062b  Greatlink Electronics Taiwan, Ltd
062c  Institute for Information Industry
062d  Taiwan Tai-Hao Enterprises Co., Ltd
062e  Mainsuper Enterprises Co., Ltd
062f  Sin Sheng Terminal & Machine, Inc.
0631  JUJO Electronics Corp.
0633  Cyrix Corp.
0634  Micron Technology, Inc.
	0655  Embedded Mass Storage Drive [RealSSD]
0635  Methode Electronics, Inc.
0636  Sierra Imaging, Inc.
	0003  Vivicam 35Xx
0638  Avision, Inc.
	0268  iVina 1200U Scanner
	026a  Minolta Dimage Scan Dual II AF-2820U (2886)
	0a10  iVina FB1600/UMAX Astra 4500
	0a13  AV600U
	0a15  Konica Minolta SC-110
	0a16  Konica Minolta SC-215
	0a30  UMAX Astra 6700 Scanner
	0a41  Avision AM3000/MF3000 Series
	0f01  fi-4010CU
# typo?
	4004  Minolta Dimage Scan Elite II AF-2920 (2888)
0639  Chrontel, Inc.
063a  Techwin Corp.
063b  Taugagreining HF
063c  Yamaichi Electronics Co., Ltd (Sakura)
063d  Fong Kai Industrial Co., Ltd
063e  RealMedia Technology, Inc.
063f  New Technology Cable, Ltd
0640  Hitex Development Tools
	0026  LPC-Stick
0641  Woods Industries, Inc.
0642  VIA Medical Corp.
0644  TEAC Corp.
	0000  Floppy
	0200  All-In-One Multi-Card Reader CA200/B/S
	1000  CD-ROM Drive
	800d  TASCAM Portastudio DP-01FX
	800e  TASCAM US-122L
	801d  TASCAM DR-100
	8021  TASCAM US-122mkII
	d001  CD-R/RW Unit
	d002  CD-R/RW Unit
	d010  CD-RW/DVD Unit
0645  Who? Vision Systems, Inc.
0646  UMAX
0647  Acton Research Corp.
	0100  ARC SpectraPro UV/VIS/IR Monochromator/Spectrograph
	0101  ARC AM-VM Mono Airpath/Vacuum Monochromator/Spectrograph
	0102  ARC Inspectrum Mono
	0103  ARC Filterwheel
	03e9  Inspectrum 128x1024 F VIS Spectrograph
	03ea  Inspectrum 256x1024 F VIS Spectrograph
	03eb  Inspectrum 128x1024 B VIS Spectrograph
	03ec  Inspectrum 256x1024 B VIS Spectrograph
0648  Inside Out Networks
0649  Weli Science Co., Ltd
064b  Analog Devices, Inc. (White Mountain DSP)
	0165  Blackfin 535 [ADZS HPUSB ICE]
064c  Ji-Haw Industrial Co., Ltd
064d  TriTech Microelectronics, Ltd
064e  Suyin Corp.
	2100  Sony Visual Communication Camera
	9700  Asus Integrated Webcam
	a100  Acer OrbiCam
	a101  Acer CrystalEye Webcam
	a102  Acer/Lenovo Webcam [CN0316]
	a103  Acer/HP Integrated Webcam [CN0314]
	a110  HP Webcam
	a114  Lemote Webcam
	a116  UVC 1.3MPixel WebCam
	a136  Asus Integrated Webcam [CN031B]
	a219  1.3M WebCam (notebook emachines E730, Acer sub-brand)
	c107  HP webcam [dv6-1190en]
	c335  HP TrueVision HD
	d101  Acer CrystalEye Webcam
	d217  HP TrueVision HD
	e201  Lenovo Integrated Webcam
	e203  Lenovo Integrated Webcam
	e258  HP TrueVision HD Integrated Webcam
	f102  Lenovo Integrated Webcam [R5U877]
	f103  Lenovo Integrated Webcam [R5U877]
	f300  UVC 0.3M Webcam
064f  WIBU-Systems AG
	03e9  CmStick (article no. 1001)
	03f2  CmStick/M (article no. 1010)
	03f3  CmStick/M (article no. 1011)
	0bd7  BOX/U
	0bd8  BOX/RU
0650  Dynapro Systems
0651  Likom Technology Sdn. Bhd.
0652  Stargate Solutions, Inc.
0653  CNF, Inc.
0654  Granite Microsystems, Inc.
	0005  Device Bay Controller
	0006  Hub
	0007  Device Bay Controller
	0016  Hub
0655  Space Shuttle Hi-Tech Co., Ltd
0656  Glory Mark Electronic, Ltd
0657  Tekcon Electronics Corp.
0658  Sigma Designs, Inc.
0659  Aethra
065a  Optoelectronics Co., Ltd
	0001  Opticon OPR-2001 / NLV-1001 (keyboard mode)
	0009  NLV-1001 (serial mode) / OPN-2001 [Opticon]
065b  Tracewell Systems
065e  Silicon Graphics
065f  Good Way Technology Co., Ltd & GWC technology Inc.
0660  TSAY-E (BVI) International, Inc.
0661  Hamamatsu Photonics K.K.
0662  Kansai Electric Co., Ltd
0663  Topmax Electronic Co., Ltd
	0103  CobraPad
0664  ET&T Technology Co., Ltd.
	0301  Groovy Technology Corp. GTouch Touch Screen
	0302  Groovy Technology Corp. GTouch Touch Screen
	0303  Groovy Technology Corp. GTouch Touch Screen
	0304  Groovy Technology Corp. GTouch Touch Screen
	0305  Groovy Technology Corp. GTouch Touch Screen
	0306  Groovy Technology Corp. GTouch Touch Screen
	0307  Groovy Technology Corp. GTouch Touch Screen
	0309  Groovy Technology Corp. GTouch Touch Screen
0665  Cypress Semiconductor
	5161  USB to Serial
0667  Aiwa Co., Ltd
	0fa1  TD-U8000 Tape Drive
0668  WordWand
0669  Oce' Printing Systems GmbH
066a  Total Technologies, Ltd
066b  Linksys, Inc.
	0105  SCM eUSB SmartMedia Card Reader
	010a  Melco MCR-U2 SmartMedia / CompactFlash Reader
	200c  USB10TX
	2202  USB10TX Ethernet [pegasus]
	2203  USB100TX Ethernet [pegasus]
	2204  USB100TX HomePNA Ethernet [pegasus]
	2206  USB Ethernet [pegasus]
	2207  HomeLink Phoneline 10M Network Adapter
	2211  WUSB11 802.11b Adapter
	2212  WUSB11v2.5 802.11b Adapter
	2213  WUSB12v1.1 802.11b Adapter
	2219  Instant Wireless Network Adapter
	400b  USB10TX
066d  Entrega, Inc.
066e  Acer Semiconductor America, Inc.
066f  SigmaTel, Inc.
	003b  MP3 Player
	003e  MP3 Player
	003f  MP3 Player
	0040  MP3 Player
	0041  MP3 Player
	0042  MP3 Player
	0043  MP3 Player
	004b  A-Max PA11 MP3 Player
	3400  STMP3400 D-Major MP3 Player
	3410  STMP3410 D-Major MP3 Player
	3500  Player Recovery Device
	3780  STMP3780/i.MX23 SystemOnChip in RecoveryMode
	4200  STIr4200 IrDA Bridge
	4210  STIr4210 IrDA Bridge
	8000  MSCN MP3 Player
	8001  SigmaTel MSCN Audio Player
	8004  MSCNMMC MP3 Player
	8008  i-Bead 100 MP3 Player
	8020  MP3 Player
	8034  MP3 Player
	8036  MP3 Player
	8038  MP3 Player
	8056  MP3 Player
	8060  MP3 Player
	8066  MP3 Player
	807e  MP3 Player
	8092  MP3 Player
	8096  MP3 Player
	809a  MP3 Player
	80aa  MP3 Player
	80ac  MP3 Player
	80b8  MP3 Player
	80ba  MP3 Player
	80bc  MP3 Player
	80bf  MP3 Player
	80c5  MP3 Player
	80c8  MP3 Player
	80ca  MP3 Player
	80cc  MP3 Player
	8104  MP3 Player
	8106  MP3 Player
	8108  MP3 Player
	810a  MP3 Player
	810c  MP3 Player
	8122  MP3 Player
	8124  MP3 Player
	8126  MP3 Player
	8128  MP3 Player
	8134  MP3 Player
	8136  MP3 Player
	8138  MP3 Player
	813a  MP3 Player
	813e  MP3 Player
	8140  MP3 Player
	8142  MP3 Player
	8144  MP3 Player
	8146  MP3 Player
	8148  MP3 Player
	814c  MP3 Player
	8201  MP3 Player
	8202  Jens of Sweden / I-BEAD 150M/150H MP3 player
	8203  MP3 Player
	8204  MP3 Player
	8205  MP3 Player
	8206  Digital MP3 Music Player
	8207  MP3 Player
	8208  MP3 Player
	8209  MP3 Player
	820a  MP3 Player
	820b  MP3 Player
	820c  MP3 Player
	820d  MP3 Player
	820e  MP3 Player
	820f  MP3 Player
	8210  MP3 Player
	8211  MP3 Player
	8212  MP3 Player
	8213  MP3 Player
	8214  MP3 Player
	8215  MP3 Player
	8216  MP3 Player
	8217  MP3 Player
	8218  MP3 Player
	8219  MP3 Player
	821a  MP3 Player
	821b  MP3 Player
	821c  MP3 Player
	821d  MP3 Player
	821e  MP3 Player
	821f  MP3 Player
	8220  MP3 Player
	8221  MP3 Player
	8222  MP3 Player
	8223  MP3 Player
	8224  MP3 Player
	8225  MP3 Player
	8226  MP3 Player
	8227  MP3 Player
	8228  MP3 Player
	8229  MP3 Player
	8230  MP3 Player
	829c  MP3 Player
	82e0  MP3 Player
	8320  TrekStor i.Beat fun
	835d  MP3 Player
	9000  MP3 Player
	9001  MP3 Player
	9002  MP3 Player
0670  Sequel Imaging
	0001  Calibrator
	0005  Enable Cable
0672  Labtec, Inc.
	1041  LCS1040 Speaker System
	5000  SpaceBall 4000 FLX
0673  HCL
	5000  Keyboard
0674  Key Mouse Electronic Enterprise Co., Ltd
0675  DrayTek Corp.
	0110  Vigor 128 ISDN TA
	0530  Vigor530 IEEE 802.11G Adapter (ISL3880+NET2280)
	0550  Vigor550
	1688  miniVigor 128 ISDN TA [HFC-S]
	6694  miniVigor 128 ISDN TA
0676  Teles AG
0677  Aiwa Co., Ltd
	07d5  TM-ED1285(USB)
	0fa1  TD-U8000 Tape Drive
0678  ACard Technology Corp.
067b  Prolific Technology, Inc.
	0000  PL2301 USB-USB Bridge
	0001  PL2302 USB-USB Bridge
	0307  Motorola Serial Adapter
	04bb  PL2303 Serial (IODATA USB-RSAQ2)
	0600  IDE Bridge
	0610  Onext EG210U MODEM
	0611  AlDiga AL-11U Quad-band GSM/GPRS/EDGE modem
	2303  PL2303 Serial Port
	2305  PL2305 Parallel Port
	2306  Raylink Bridge Controller
	2307  PL2307 USB-ATAPI4 Bridge
	2313  FITEL PHS U Cable Adaptor
	2315  Flash Disk Embedded Hub
	2316  Flash Disk Security Device
	2317  Mass Storage Device
	2501  PL2501 USB-USB Bridge (USB 2.0)
	2506  Kaser 8gB micro hard drive
	2507  PL2507 Hi-speed USB to IDE bridge controller
	2515  Flash Disk Embedded Hub
	2517  Flash Disk Mass Storage Device
	2528  Storage device (8gB thumb drive)
	25a1  PL25A1 Host-Host Bridge
	2773  PL2773 SATAII bridge controller
	3400  Hi-Speed Flash Disk with TruePrint AES3400
	3500  Hi-Speed Flash Disk with TruePrint AES3500
	3507  PL3507 ATAPI6 Bridge
	aaa0  Prolific Pharos
	aaa2  PL2303 Serial Adapter (IODATA USB-RSAQ3)
067c  Efficient Networks, Inc.
	1001  Siemens SpeedStream 100MBps Ethernet
	1022  Siemens SpeedStream 1022 802.11b Adapter
	1023  SpeedStream Wireless
	4020  SpeedStream 4020 ATM/ADSL Installer
	4031  Efficient ADSL Modem
	4032  SpeedStream 4031 ATM/ADSL Installer
	4033  SpeedStream 4031 ATM/ADSL Installer
	4060  Alcatel Speedstream 4060 ADSL Modem
	4062  Efficient Networks 4060 Loader
	5667  Efficient Networks Virtual Bus for ADSL Modem
	c031  SpeedStream 4031 ATM/ADSL Installer
	c032  SpeedStream 4031 ATM/ADSL Installer
	c033  SpeedStream 4031 ATM/ADSL Installer
	c060  SpeedStream 4060 Miniport ATM/ADSL Adapter
	d667  Efficient Networks Virtual Bus for ADSL Modem
	e240  Speedstream Ethernet Adapter E240
	e540  Speedstream Ethernet Adapter E240
067d  Hohner Corp.
067e  Intermec Technologies Corp.
	0801  HID Keyboard, Barcode scanner
	0803  VCP, Barcode scanner
	0805  VCP + UVC, Barcode scanner
	1001  Mobile Computer
067f  Virata, Ltd
	4552  DSL-200 ADSL Modem
	6542  DSL Modem
	6549  DSL Modem
	7541  DSL Modem
0680  Realtek Semiconductor Corp., CPP Div. (Avance Logic)
	0002  Arowana Optical Wheel Mouse MSOP-01
0681  Siemens Information and Communication Products
	0001  Dect Base
	0002  Gigaset 3075 Passive ISDN
	0005  ID-Mouse with Fingerprint Reader
	0012  I-Gate 802.11b Adapter
	001b  WLL013
	001d  Hipath 1000
	0022  Gigaset SX353 ISDN
	0026  DECT Data - Gigaset M34
	002b  A-100-I ADSL Modem
	002e  ADSL Router_S-141
	0034  GSM module MC35/ES75 USB Modem
	3c06  54g USB Network Adapter
0682  Victor Company of Japan, Ltd
0684  Actiontec Electronics, Inc.
0685  ZD Incorporated
	7000  HSDPA Modem
0686  Minolta Co., Ltd
	2001  PagePro 4110W
	2004  PagePro 1200W
	2005  Magicolor 2300 DL
	3001  PagePro 4100
	3005  PagePro 1250E
	3006  PagePro 1250W
	3009  Magicolor 2300W
	300b  PagePro 1350W
	300c  PagePro 1300W
	302e  Develop D 1650iD PCL
	3034  Develop D 2050iD PCL
	4001  Dimage 2300
	4003  Dimage 2330 Zoom Camera
	4004  Dimage Scan Elite II AF-2920 (2888)
	4005  Minolta DiMAGE E201 Mass Storage Device
	4006  Dimage 7 Camera
	4007  Dimage S304 Camera
	4008  Dimage 5 Camera
	4009  Dimage X Camera
	400a  Dimage S404 Camera
	400b  Dimage 7i Camera
	400c  Dimage F100 Camera
	400d  Dimage Scan Dual III AF-2840 (2889)
	400e  Dimage Scan Elite 5400 (2890)
	400f  Dimage 7Hi Camera
	4010  Dimage Xi Camera
	4011  Dimage F300 Camera
	4012  Dimage F200 Camera
	4014  Dimage S414 Camera
	4015  Dimage XT Camera [storage]
	4016  Dimage XT Camera [remote mode]
	4017  Dimage E223
	4018  Dimage Z1  Camera
	4019  Dimage A1 Camera [remote mode]
	401a  Dimage A1 Camera [storage]
	401c  Dimage X20 Camera
	401e  Dimage E323 Camera
068a  Pertech, Inc.
068b  Potrans International, Inc.
068e  CH Products, Inc.
	00d3  OEM 3 axis 5 button joystick
	00e2  HFX OEM Joystick
	00f0  Multi-Function Panel
	00f1  Pro Throttle
	00f2  Flight Sim Pedals
	00f3  Fighterstick
	00f4  Combatstick
	00fa  Ch Throttle Quadrant
	00ff  Flight Sim Yoke
	0500  GameStick 3D
	0501  CH Pro Pedals
	0504  F-16 Combat Stick
0690  Golden Bridge Electech, Inc.
0693  Hagiwara Sys-Com Co., Ltd
	0002  FlashGate SmartMedia Card Reader
	0003  FlashGate CompactFlash Card Reader
	0005  FlashGate
	0006  SM PCCard R/W and SPD
	0007  FlashGate ME (Authenticated)
	000a  SDCard/MMC Reader/Writer
0694  Lego Group
	0001  Mindstorms Tower
	0002  Mindstorms NXT
0698  Chuntex (CTX)
	1786  1300ex Monitor
	2003  CTX M730V built in Camera
	9999  VLxxxx Monitor+Hub
0699  Tektronix, Inc.
	0347  AFG 3022B
069a  Askey Computer Corp.
	0001  VC010 Webcam [pwc]
	0303  Cable Modem
	0311  ADSL Router Remote NDIS Device
	0318  Remote NDIS Device
	0319  220V Remote NDIS Device
	0320  IEEE 802.11b Wireless LAN Card
	0321  Dynalink WLL013 / Compex WLU11A 802.11b Adapter
	0402  Scientific Atlanta WebSTAR 100 & 200 series Cable Modem
	0811  BT Virtual Bus for Helium
	0821  BT Voyager 1010 802.11b Adapter
	4402  Scientific Atlanta WebSTAR 2000 series Cable Modem
	4403  Scientific Atlanta WebSTAR 300 series Cable Modem
	4501  Scientific-Atlanta WebSTAR 2000 series Cable Modem
069b  Thomson, Inc.
	0704  DCM245 Cable Modem
	0705  THG540K Cable Modem
	0709  Lyra PDP2424
	070c  MP3 Player
	070d  MP3 Player
	070e  MP3 Player
	070f  RCA Lyra RD1071 MP3 Player
	0731  Lyra M200E256
	0761  RCA H100A
	0778  PEARL USB Device
	2220  RCA Kazoo RD1000 MP3 Player
	300a  RCA Lyra MP3 Player
	3012  MP3 Player
	3013  MP3 Player
	5557  RCA CDS6300
069d  Hughes Network Systems (HNS)
	0001  Satellite Receiver Device
	0002  Satellite Device
069e  Welcat Inc.
	0005  Marx CryptoBox v1.2
069f  Allied Data Technologies BV
	0010  Tornado Speakerphone FaxModem 56.0
	0011  Tornado Speakerphone FaxModem 56.0
	1000  ADT VvBus for CopperJet
	1004  CopperJet 821 RouterPlus
06a2  Topro Technology, Inc.
	0033  USB Mouse
06a3  Saitek PLC
	0006  Cyborg Gold Joystick
	0109  P880 Pad
	0160  ST290 Pro
	0200  Xbox Adrenalin Hub
	0241  Xbox Adrenalin Gamepad
	0255  X52 Flight Controller
	040b  P990 Dual Analog Pad
	040c  P2900 Wireless Pad
	0422  ST90 Joystick
	0460  ST290 Pro Flight Stick
	0463  ST290
	0464  Cyborg Evo
	0471  Cyborg Graphite Stick
	0501  R100 Sports Wheel
	0502  ST200 Stick
	0506  R220 Digital Wheel
	051e  Cyborg Digital II Stick
	052d  P750 Gamepad
	053c  X45 Flight Controller
	053f  X36F Flightstick
	056c  P2000 Tilt Pad
	056f  P2000 Tilt Pad
	05d2  PC Dash 2
	075c  X52 Flight Controller
	0762  Saitek X52 Pro Flight Control System
	0763  Pro Flight Rudder Pedals
	0764  Flight Pro Combat Rudder
	0805  R440 Force Wheel
	0b4e  Pro Flight Backlit Information Panel
	0bac  Pro Flight Yoke
	0c2d  Pro Flight Quadrant
	0d05  Pro Flight Radio Panel
	0d06  Flight Pro Multi Panel
	0d67  Pro Flight Switch Panel
	1003  GM2 Action Pad
	1009  Action Pad
	100a  SP550 Pad and Joystick Combo
	100b  SP550 Pad
	1509  P3000 Wireless Pad
	1589  P3000 Wireless Pad
	2541  X45 Flight Controller
	3509  P3000 RF GamePad
	353e  Cyborg Evo Wireless
	3589  P3000 Wireless Pad
	35be  Cyborg Evo
	5509  P3000 Wireless Pad
	712c  Pro Flight Yoke integrated hub
	8000  Gamers' Keyboard
	801e  Cyborg 3D Digital Stick II
	8020  Eclipse Keyboard
	8021  Eclipse II Keyboard
	802d  P750 Pad
	803f  X36 Flight Controller
	806f  P2000 Tilt Pad
	80c0  Pro Gamer Command Unit
	80c1  Cyborg Command Pad Unit
	a2ae  Pro Flight Instrument Panel
	a502  Gaming Mouse
	f518  P3200 Rumble Force Game Pad
	ff04  R440 Force Wheel
	ff0c  Cyborg Force Rumble Pad
	ff0d  P2600 Rumble Force Pad
	ff12  Cyborg 3D Force Stick
	ff17  ST 330 Rumble Force Stick
	ff52  Cyborg 3D Rumble Force Joystick
	ffb5  Cyborg Evo Force Joystick
06a4  Xiamen Doowell Electron Co., Ltd
06a5  Divio
	0000  Typhoon Webcam 100k [nw8000]
	d001  ProLink DS3303u Webcam
	d800  Chicony TwinkleCam
	d820  Wize Media 1000
06a7  MicroStore, Inc.
06a8  Topaz Systems, Inc.
	0042  SignatureGem 1X5 Pad
	0043  SignatureGem 1X5-HID Pad
06a9  Westell
	0005  WireSpeed Dual Connect Modem
	0006  WireSpeed Dual Connect Modem
	000a  WireSpeed Dual Connect Modem
	000b  WireSpeed Dual Connect Modem
	000e  A90-211WG-01 802.11g Adapter [Intersil ISL3887]
06aa  Sysgration, Ltd
06ac  Fujitsu Laboratories of America, Inc.
06ad  Greatland Electronics Taiwan, Ltd
06ae  Professional Multimedia Testing Centre
06af  Harting, Inc. of North America
06b8  Pixela Corp.
06b9  Alcatel Telecom
	0120  SpeedTouch 120g 802.11g Wireless Adapter [Intersil ISL3886]
	0121  SpeedTouch 121g Wireless Dongle
	2001  SPEED TOUCH Card
	4061  SpeedTouch ISDN or ADSL Modem
	4062  SpeedTouch ISDN or ADSL router
	a5a5  DynaMiTe Modem
06ba  Smooth Cord & Connector Co., Ltd
06bb  EDA, Inc.
06bc  Oki Data Corp.
	000b  Okipage 14ex Printer
	0027  Okipage 14e
	00f7  OKI B4600 Mono Printer
	015e  OKIPOS 411/412 POS Printer
	01c9  OKI B430 Mono Printer
	020b  OKI ES4140 Mono Printer
	02bb  OKI PT390 POS Printer
	0a91  B2500MFP (printer+scanner)
	3801  B6100 Laser Printer
06bd  AGFA-Gevaert NV
	0001  SnapScan 1212U
	0002  SnapScan 1236U
	0100  SnapScan Touch
	0101  SNAPSCAN ELITE
	0200  ScanMaker 8700
	02bf  DUOSCAN f40
	0400  CL30
	0401  Mass Storage
	0403  ePhoto CL18 Camera
	0404  ePhoto CL20 Camera
	2061  SnapScan 1212U (?)
	208d  Snapscan e40
	208f  SnapScan e50
	2091  SnapScan e20
	2093  SnapScan e10
	2095  SnapScan e25
	2097  SnapScan e26
	20fd  SnapScan e52
	20ff  SnapScan e42
06be  AME Optimedia Technology Co., Ltd
	0800  Optimedia Camera
	1005  Dazzle DPVM! (1005)
	d001  P35U Camera Capture
06bf  Leoco Corp.
06c2  Phidgets Inc. (formerly GLAB)
	0030  PhidgetRFID
	0031  RFID reader
	0038  4-Motor PhidgetServo v3.0
	0039  1-Motor PhidgetServo v3.0
	003a  8-Motor PhidgetAvancedServo
	0040  PhidgetInterface Kit 0-0-4
	0044  PhidgetInterface Kit 0-16-16
	0045  PhidgetInterface Kit 8-8-8
	0048  PhidgetStepper (Under Development)
	0049  PhidgetTextLED Ver 1.0
	004a  PhidgetLED Ver 1.0
	004b  PhidgetEncoder Ver 1.0
	0051  PhidgetInterface Kit 0-5-7 (Custom)
	0052  PhidgetTextLCD
	0053  PhidgetInterfaceKit 0-8-8
	0058  PhidgetMotorControl Ver 1.0
	0070  PhidgetTemperatureSensor Ver 1.0
	0071  PhidgetAccelerometer Ver 1.0
	0072  PhidgetWeightSensor Ver 1.0
	0073  PhidgetHumiditySensor
	0074  PhidgetPHSensor
	0075  PhidgetGyroscope
06c4  Bizlink International Corp.
06c5  Hagenuk, GmbH
06c6  Infowave Software, Inc.
06c8  SIIG, Inc.
06c9  Taxan (Europe), Ltd
	0005  Monitor Control
	0007  Monitor Control
	0009  Monitor Control
06ca  Newer Technology, Inc.
	2003  uSCSI
06cb  Synaptics, Inc.
	0001  TouchPad
	0002  Integrated TouchPad
	0003  cPad
	0005  Touchpad/FPS
	0006  TouchScreen
	0007  USB Styk
	0008  WheelPad
	0009  Composite TouchPad and TrackPoint
	000e  HID Device
	0010  Wireless TouchPad
	0013  DisplayPad
	2970  touchpad
06cc  Terayon Communication Systems
	0101  Cable Modem
	0102  Cable Modem
	0103  Cable Modem
	0104  Cable Modem
	0304  Cable Modem
06cd  Keyspan
	0101  USA-28 PDA [no firmware]
	0102  USA-28X PDA [no firmware]
	0103  USA-19 PDA [no firmware]
	0104  PDA [prerenum]
	0105  USA-18X PDA [no firmware]
	0106  USA-19W PDA [no firmware]
	0107  USA-19 PDA
	0108  USA-19W PDA
	0109  USA-49W serial adapter [no firmware]
	010a  USA-49W serial adapter
	010b  USA-19Qi serial adapter [no firmware]
	010c  USA-19Qi serial adapter
	010d  USA-19Q serial Adapter (no firmware)
	010e  USA-19Q serial Adapter
	010f  USA-28 PDA
	0110  USA-28Xb PDA
	0111  USA-18 serial Adapter
	0112  USA-18X PDA
	0113  USA-28Xb PDA [no firmware]
	0114  USA-28Xa PDA [no firmware]
	0115  USA-28Xa PDA
	0116  USA-18XA serial Adapter (no firmware)
	0117  USA-18XA serial Adapter
	0118  USA-19QW PDA [no firmware]
	0119  USA-19QW PDA
	011a  USA-49Wlc serial adapter [no firmware]
	011b  MPR Serial Preloader (MPRQI)
	011c  MPR Serial (MPRQI)
	011d  MPR Serial Preloader (MPRQ)
	011e  MPR Serial (MPRQ)
	0121  USA-19hs serial adapter
	012a  USA-49Wlc serial adapter
	0201  UIA-10 Digital Media Remote [Cypress AN2131SC]
	0202  UIA-11 Digital Media Remote
06ce  Contec
	8311  COM-1(USB)H
06cf  SpheronVR AG
	1010  PanoCam 10
	1012  PanoCam 12/12X
06d0  LapLink, Inc.
	0622  LapLink Gold USB-USB Bridge [net1080]
06d1  Daewoo Electronics Co., Ltd
06d3  Mitsubishi Electric Corp.
	0284  FX-USB-AW/-BD RS482 Converters
	0380  CP8000D Port
	0381  CP770D Port
	0385  CP900D Port
	0387  CP980D Port
	038b  CP3020D Port
	038c  CP900DW(ID) Port
	0393  CP9500D/DW Port
	0394  CP9000D/DW Port
	03a1  CP9550D/DW Port
	3b30  CP-D70DW / CP-D707DW
	3b31  CP-K60DW-S
06d4  Cisco Systems
06d5  Toshiba
	4000  Japanese Keyboard
06d6  Aashima Technology B.V.
	0025  Gamepad
	0026  Predator TH 400 Gamepad
	002d  Trust PowerC@m 350FT
	002e  Trust PowerC@m 350FS
	0030  Trust 710 LCD POWERC@M ZOOM - MSD
	0031  Trust 610/710 LCD POWERC@M ZOOM
	003a  Trust PowerC@m 770Z (mass storage mode)
	003b  Trust PowerC@m 770Z (webcam mode)
	003c  Trust 910z PowerC@m
	003f  Trust 735S POWERC@M ZOOM, WDM DSC Bulk Driver
	0050  Trust 738AV LCD PV Digital Camera
	0062  TRUST 782AV LCD P. V. Video Capture
	0066  TRUST Digital PCTV and Movie Editor
	0067  Trust 350FS POWERC@M FLASH
	006b  TRUST AUDIO VIDEO EDITOR
06d7  Network Computing Devices (NCD)
06d8  Technical Marketing Research, Inc.
06da  Phoenixtec Power Co., Ltd
	0002  UPS
	0003  1300VA UPS
06db  Paradyne
06dc  Foxlink Image Technology Co., Ltd
	0012  Scan 1200c Scanner
	0014  Prolink Winscan Pro 2448U
06de  Heisei Electronics Co., Ltd
06e0  Multi-Tech Systems, Inc.
	0319  MT9234ZBA-USB MultiModem ZBA
	f101  MT5634ZBA-USB MultiModemUSB (old firmware)
	f103  MT5634MU MultiMobileUSB
	f104  MT5634ZBA-USB MultiModemUSB (new firmware)
	f107  MT5634ZBA-USB-V92 MultiModemUSB
	f120  MT9234ZBA-USB-CDC-ACM-XR MultiModem ZBA CDC-ACM-XR
06e1  ADS Technologies, Inc.
	0008  UBS-10BT Ethernet [klsi]
	0009  UBS-10BT Ethernet
	0833  Mass Storage Device
	a155  FM Radio Receiver/Instant FM Music (RDX-155-EF)
	a160  Instant Video-To-Go RDX-160 (no firmware)
	a161  Instant Video-To-Go RDX-160
	a190  Instand VCD Capture
	a191  Instant VideoXpress
	a337  Mini DigitalTV
	a701  DVD Xpress
	a708  saa7114H video input card (Instant VideoMPX)
	b337  Mini DigitalTV
	b701  DVD Xpress B
06e4  Alcatel Microelectronics
06e6  Tiger Jet Network, Inc.
	0200  Internet Phone
	0201  Internet Phone
	0202  Composite Device
	0203  Internet Phone
	0210  Composite Device
	0211  Internet Phone
	0212  Internet Phone
	031c  Internet Phone
	031d  Internet Phone
	031e  Internet Phone
	3200  Composite Device
	3201  Internet Phone
	3202  Composite Device
	3203  Composite Device
	7200  Composite Device
	7210  Composite Device
	7250  Composite Device
	825c  Internet Phone
	831c  Internet Phone
	831d  Composite Device
	831e  Composite Device
	b200  Composite Device
	b201  Composite Device
	b202  Internet Phone
	b210  Internet Phone
	b211  Composite Device
	b212  Composite Device
	b250  Composite Device
	b251  Internet Phone
	b252  Internet Phone
	c200  Internet Phone
	c201  Internet Phone
	c202  Composite Device
	c203  Internet Phone
	c210  Personal PhoneGateway
	c211  Personal PhoneGateway
	c212  Personal PhoneGateway
	c213  PPG Device
	c25c  Composite Device
	c290  PPG Device
	c291  PPG Device
	c292  PPG Device
	c293  Personal PhoneGateway
	c31c  Composite Device
	c39c  Personal PhoneGateway
	c39d  PPG Device
	c39e  PPG Device
	c39f  PPG Device
	c700  Internet Phone
	c701  Internet Phone
	c702  Composite Device
	c703  Internet Phone
	c710  VoIP Combo Device
	c711  VoIP Combo
	c712  VoIP Combo Device
	c713  VoIP Combo Device
	cf00  Composite Device
	cf01  Internet Phone
	cf02  Internet Phone
	cf03  Composite Device
	d210  Personal PhoneGateway
	d211  PPG Device
	d212  PPG Device
	d213  Personal PhoneGateway
	d700  Composite Device
	d701  Composite Device
	d702  Internet Phone
	d703  Composite Device
	d710  VoIP Combo
	d711  VoIP Combo Device
	d712  VoIP Combo
	d713  VoIP Combo
	df00  Composite Device
	df01  Composite Device
	df02  Internet Phone
	df03  Internet Phone
	f200  Internet Phone
	f201  Internet Phone
	f202  Composite Device
	f203  Composite Device
	f210  Internet Phone
	f250  Composite Device
	f252  Internet Phone
	f310  Internet Phone
	f350  Composite Device
06ea  Sirius Technologies
	0001  NetCom Roadster II 56k
	0002  Roadster II 56k
06eb  PC Expert Tech. Co., Ltd
06ef  I.A.C. Geometrische Ingenieurs B.V.
06f0  T.N.C Industrial Co., Ltd
	de01  DualCam Video Camera
	de02  DualCam Still Camera
06f1  Opcode Systems, Inc.
	a011  SonicPort
	a021  SonicPort Optical
06f2  Emine Technology Co.
	0011  KVM Switch Keyboard
06f6  Wintrend Technology Co., Ltd
06f7  Wailly Technology Ltd
	0003  USB->Din 4 Adaptor
06f8  Guillemot Corp.
	3002  Hercules Blog Webcam
	3004  Hercules Classic Silver
	3005  Hercules Dualpix Exchange
	3007  Hercules Dualpix Chat and Show
	3020  Hercules Webcam EC300
	a300  Dual Analog Leader GamePad
	b000  Hercules DJ Console
	c000  Hercules Muse Pocket
	d002  Hercules DJ Console
	e000  HWGUSB2-54 WLAN
	e010  HWGUSB2-54-LB
	e020  HWGUSB2-54V2-AP
	e031  Hercules HWNUm-300 Wireless N mini [Realtek RTL8191SU]
	e032  HWGUm-54 [Hercules Wireless G Ultra Mini Key]
	e033  Hercules HWNUp-150 802.11n Wireless N Pico [Realtek RTL8188CUS]
06f9  ASYST electronic d.o.o.
06fa  HSD S.r.L
06fc  Motorola Semiconductor Products Sector
06fd  Boston Acoustics
	0101  Audio Device
	0102  Audio Device
	0201  2-piece Audio Device
06fe  Gallant Computer, Inc.
0701  Supercomal Wire & Cable SDN. BHD.
0703  Bvtech Industry, Inc.
0705  NKK Corp.
0706  Ariel Corp.
0707  Standard Microsystems Corp.
	0100  2202 Ethernet [klsi]
	0200  2202 Ethernet [pegasus]
	0201  EZ Connect USB Ethernet
	ee04  SMCWUSB32 802.11b Wireless LAN Card
	ee06  SMC2862W-G v1 EZ Connect 802.11g Adapter [Intersil ISL3886]
	ee13  SMC2862W-G v2 EZ Connect 802.11g Adapter [Intersil ISL3887]
0708  Putercom Co., Ltd
	047e  USB-1284 BRIDGE
0709  Silicon Systems, Ltd (SSL)
070a  Oki Electric Industry Co., Ltd
	4002  Bluetooth Device
	4003  Bluetooth Device
070d  Comoss Electronic Co., Ltd
070e  Excel Cell Electronic Co., Ltd
0710  Connect Tech, Inc.
	0001  WhiteHeat (fake ID)
	8001  WhiteHeat
0711  Magic Control Technology Corp.
	0100  Hub
	0180  IRXpress Infrared Device
	0181  IRXpress Infrared Device
	0200  BAY-3U1S1P Serial Port
	0210  MCT1S Serial Port
	0230  MCT-232 Serial Port
	0231  PS/2 Mouse Port
	0232  Serial On Port
	0240  PS/2 to USB Converter
	0300  BAY-3U1S1P Parallel Port
	0302  Parallel Port
	0900  SVGA Adapter
	5001  Trigger UV-002BD[Startech USBVGAE]
	5100  Magic Control Technology Corp. (USB2VGA dongle)
0713  Interval Research Corp.
0714  NewMotion, Inc.
	0003  ADB to USB convertor
0717  ZNK Corp.
0718  Imation Corp.
	0002  SuperDisk 120MB
	0003  SuperDisk 120MB (Authenticated)
	0060  Flash Drive
	0061  Flash Drive
	0062  Flash Drive
	0063  Swivel Flash Drive
	0064  Flash Drive
	0065  Flash Drive
	0066  Flash Drive
	0067  Flash Drive
	0068  Flash Drive
	0084  Flash Drive Mini
	043c  Flash drive 16GB [Nano Pro]
	0582  Revo Flash Drive
	0622  TDK Trans-It 4GB
	0624  TDK Trans-It 16GB
	1120  RDX External dock (redbud)
	d000  Disc Stakka CD/DVD Manager
0719  Tremon Enterprises Co., Ltd
071b  Domain Technologies, Inc.
	0002  DTI-56362-USB Digital Interface Unit
	0101  Audio4-USB DSP Data Acquisition Unit
	0201  Audio4-5410 DSP Data Acquisition Unit
	0301  SB-USB JTAG Emulator
	3203  Rockchip Media Player
	32bb  Music Mediatouch
071c  Xionics Document Technologies, Inc.
071d  Eicon Networks Corp.
	1000  Diva 2.01 S/T [PSB2115F]
	1003  Diva ISDN 2.0
	1005  Diva ISDN 4.0 [HFC-S]
	2000  Teledat Surf
071e  Ariston Technologies
0723  Centillium Communications Corp.
	0002  Palladia 300/400 Adsl Modem
0726  Vanguard International Semiconductor-America
0729  Amitm
	1000  USC-1000 Serial Port
072e  Sunix Co., Ltd
072f  Advanced Card Systems, Ltd
	0001  AC1030-based SmartCard Reader
	0008  ACR 80 Smart Card Reader
	0100  AET65
	0101  AET65
	0102  AET62
	0103  AET62
	0901  ACR1281U-C4 (BSI)
	1000  PLDT Drive
	1001  PLDT Drive
	2011  ACR88U
	2100  ACR128U
	2200  ACR122U
	220a  ACR1281U-C5 (BSI)
	220c  ACR1283 Bootloader
	220f  ACR1281U-C2 (qPBOC)
	2211  ACR1261 1S Dual Reader
	2214  ACR1222 1SAM PICC Reader
	2215  ACR1281 2S CL Reader
	221a  ACR1251U-A1
	221b  ACR1251U-C
	2224  ACR1281 1S Dual Reader
	222b  ACR1222U-C8
	222c  ACR1283L-D2
	222d  [OEM Reader]
	222e  ACR123U
	2242  ACR1251 1S Dual Reader
	8002  AET63 BioTRUSTKey
	8003  ACR120
	8103  ACR120
	8201  APG8201
	8900  ACR89U-A1
	8901  ACR89U-A2
	8902  ACR89U-A3
	9000  ACR38 AC1038-based Smart Card Reader
	9006  CryptoMate
	90cc  ACR38 SmartCard Reader
	90ce  [OEM Reader]
	90cf  ACR38 SAM Smart Card Reader
	90d0  PertoSmart EMV - Card Reader
	90d2  ACR83U
	90d8  ACR3801
	90db  CryptoMate64
	b000  ACR3901U
	b100  ACR39U
	b101  ACR39K
	b102  ACR39T
	b103  ACR39F
	b104  ACR39U-SAM
	b106  ACOS5T2
	b200  ACOS5T1
	b301  ACR32-A1
0731  Susteen, Inc.
	0528  SonyEricsson DCU-11 Cable
0732  Goldfull Electronics & Telecommunications Corp.
0733  ViewQuest Technologies, Inc.
	0101  Digital Video Camera
	0110  VQ110 Video Camera
	0401  CS330 Webcam
	0402  M-318B Webcam
	0430  Intel Pro Share Webcam
	0630  VQ630 Dual Mode Digital Camera(Bulk)
	0631  Hercules Dualpix
	0780  Smart Cam Deluxe(composite)
	1310  Epsilon 1.3/Jenoptik JD C1.3/UMAX AstraPix 470
	1311  Digital Dream Epsilon 1.3
	1314  Mercury 2.1MEG Deluxe Classic Cam
	2211  Jenoptik jdc 21 LCD Camera
	2221  Mercury Digital Pro 3.1p
	3261  Concord 3045 spca536a Camera
	3281  Cyberpix S550V
0734  Lasat Communications A/S
	0001  560V Modem
	0002  Lasat 560V Modem
	043a  DVS Audio
	043b  3DeMon USB Capture
0735  Asuscom Network
	2100  ISDN Adapter
	2101  ISDN Adapter
	6694  ISDNlink 128K
	c541  ISDN TA 280
0736  Lorom Industrial Co., Ltd
0738  Mad Catz, Inc.
	4507  XBox Device
	4516  XBox Device
	4520  XBox Device
	4526  XBox Device
	4536  XBox Device
	4540  XBox Device
	4556  XBox Device
	4566  XBox Device
	4576  XBox Device
	4586  XBox Device
	4588  XBox Device
	8818  Street Fighter IV Arcade FightStick (PS3)
073a  Chaplet Systems, Inc.
	2230  infrared dongle for remote
073b  Suncom Technologies
073c  Industrial Electronic Engineers, Inc.
	0305  Pole Display (PC305-3415  2 x 20 Line Display)
	0322  Pole Display (PC322-3415  2 x 20 Line Display)
	0324  Pole Display (LB324-USB   4 x 20 Line Display)
	0330  Pole Display (P330-3415   2 x 20 Line Display)
	0424  Pole Display (SP324-4415  4 x 20 Line Display)
	0450  Pole Display (L450-USB   Graphic Line Display)
	0505  Pole Display (SPC505-3415 2 x 20 Line Display)
	0522  Pole Display (SPC522-3415 2 x 20 Line Display)
	0624  Pole Display (SP324-3415  4 x 20 Line Display)
073d  Eutron S.p.a.
	0005  Crypto Token
	0007  CryptoIdentity CCID
	0025  SmartKey 3
	0c00  Pocket Reader
	0d00  StarSign Bio Token 3.0 EU
073e  NEC, Inc.
	0301  Game Pad
0742  Stollmann
	2008  ISDN TA [HFC-S]
	2009  ISDN TA [HFC-S]
	200a  ISDN TA [HFC-S]
0745  Syntech Information Co., Ltd
0746  Onkyo Corp.
	5500  SE-U55 Audio Device
0747  Labway Corp.
0748  Strong Man Enterprise Co., Ltd
0749  EVer Electronics Corp.
074a  Ming Fortune Industry Co., Ltd
074b  Polestar Tech. Corp.
074c  C-C-C Group PLC
074d  Micronas GmbH
	3553  Composite USB-Device
	3554  Composite USB-Device
	3556  Composite USB-Device
074e  Digital Stream Corp.
	0001  PS/2 Adapter
	0002  PS/2 Adapter
0755  Aureal Semiconductor
0757  Network Technologies, Inc.
075b  Sophisticated Circuits, Inc.
	0001  Kick-off! Watchdog
0763  Midiman
	0115  O2 / KeyRig 25
	0117  Trigger Finger
	0119  MidAir
	0150  M-Audio Uno
	0160  M-Audio 1x1
	0192  M-Audio Keystation 88es
	0193  ProKeys 88
	0194  ProKeys 88sx
	0195  Oxygen 8 v2
	0196  Oxygen 49
	0197  Oxygen 61
	0198  Axiom 25
	0199  Axiom 49
	019a  Axiom 61
	019b  KeyRig 49
	019c  KeyStudio
	1001  MidiSport 2x2
	1002  MidiSport 2x2
	1003  MidiSport 2x2
	1010  MidiSport 1x1
	1011  MidiSport 1x1
	1014  M-Audio Keystation Loader
	1015  M-Audio Keystation
	1020  Midisport 4x4
	1021  MidiSport 4x4
	1030  Midisport 8x8
	1031  MidiSport 8x8/s Loader
	1033  MidiSport 8x8/s
	1040  M-Audio MidiSport 2x4 Loader
	1041  M-Audio MidiSport 2x4
	1110  MidiSport 1x1
	2001  M Audio Quattro
	2002  M Audio Duo
	2003  M Audio AudioPhile
	2004  M-Audio MobilePre
	2006  M-Audio Transit
	2007  M-Audio Sonica Theater
	2008  M-Audio Ozone
	200d  M-Audio OmniStudio
	200f  M-Audio MobilePre
	2010  M-Audio Fast Track
	2012  M-Audio Fast Track Pro
	2013  M-Audio JamLab
	2015  M-Audio RunTime DFU
	2016  M-Audio RunTime DFU
	2019  M-Audio Ozone Academic
	201a  M-Audio Micro
	201b  M-Audio RunTime DFU
	201d  M-Audio Producer
	2024  M-Audio Fast Track MKII
	2080  M-Audio RunTime DFU
	2081  M-Audio RunTime DFU / Fast Track Ultra 8R
	2803  M-Audio Audiophile DFU
	2804  M-Audio MobilePre DFU
	2806  M-Audio Transit DFU
	2815  M-Audio DFU
	2816  M-Audio DFU
	281b  M-Audio DFU
	2880  M-Audio DFU
	2881  M-Audio DFU
0764  Cyber Power System, Inc.
	0005  Cyber Power UPS
	0501  CP1500 AVR UPS
	0601  PR1500LCDRT2U UPS
0765  X-Rite, Inc.
	5001  Huey PRO Colorimeter
	5020  i1 Display Pro
	6003  ColorMunki Smile
	d094  X-Rite DTP94 [Quato Silver Haze Pro]
0766  Jess-Link Products Co., Ltd
	001b  Packard Bell Go
	0204  TopSpeed Cyberlink Remote Control
0767  Tokheim Corp.
0768  Camtel Technology Corp.
	0006  Camtel Technology USB TV Genie Pro FM Model TVB330
	0023  eHome Infrared Receiver
0769  Surecom Technology Corp.
	11f2  EP-9001-g 802.11g 54M WLAN Adapter
	11f3  RT2570
	11f7  802.11g 54M WLAN Adapter
	31f3  RT2573
076a  Smart Technology Enablers, Inc.
076b  OmniKey AG
	0596  CardMan 2020
	1021  CardMan 1021
	1221  CardMan 1221
	1784  CardMan 6020
	3021  CardMan 3121
	3022  CardMan 3021
	3610  CardMan 3620
	3621  CardMan 3621
	3821  CardMan 3821
	4321  CardMan 4321
	5121  CardMan 5121
	5125  CardMan 5125
	5321  CardMan 5321
	5340  CardMan 5021 CL
	6622  CardMan 6121
	a011  CCID Smart Card Reader Keyboard
	a021  CCID Smart Card Reader
	a022  CardMan Smart@Link
	c000  CardMan 3x21 CS
	c001  CardMan 5121 CS
076c  Partner Tech
076d  Denso Corp.
076e  Kuan Tech Enterprise Co., Ltd
076f  Jhen Vei Electronic Co., Ltd
0770  Welch Allyn, Inc - Medical Division
0771  Observator Instruments BV
	4455  OMC45III
	ae0f  OMC45III
0772  Your data Our Care
0774  AmTRAN Technology Co., Ltd
0775  Longshine Electronics Corp.
0776  Inalways Corp.
0777  Comda Enterprise Corp.
0778  Volex, Inc.
0779  Fairchild Semiconductor
077a  Sankyo Seiki Mfg. Co., Ltd
077b  Linksys
	08be  BEFCMU10 v4 Cable Modem
	2219  WUSB11 V2.6 802.11b Adapter
	2226  USB200M 100baseTX Adapter
	2227  Network Everywhere NWU11B
077c  Forward Electronics Co., Ltd
	0005  NEC Keyboard
077d  Griffin Technology
	0223  IMic Audio In/Out
	0405  iMate, ADB Adapter
	0410  PowerMate
	041a  PowerWave
	04aa  SoundKnob
	07af  iMic
	1016  AirClick
	627a  Radio SHARK
077f  Well Excellent & Most Corp.
0780  Sagem Monetel GmbH
	1202  ORGA 900 Smart Card Terminal Virtual Com Port
	1302  ORGA 6000 Smart Card Terminal Virtual Com Port
	1303  ORGA 6000 Smart Card Terminal USB RNDIS
	df55  ORGA 900/6000 Smart Card Terminal DFU
0781  SanDisk Corp.
	0001  SDDR-05a ImageMate CompactFlash Reader
	0002  SDDR-31 ImageMate II CompactFlash Reader
	0005  SDDR-05b (CF II) ImageMate CompactFlash Reader
	0100  ImageMate SDDR-12
	0200  SDDR-09 (SSFDC) ImageMate SmartMedia Reader [eusb]
	0400  SecureMate SD/MMC Reader
	0621  SDDR-86 Imagemate 6-in-1 Reader
	0720  Sansa C200 series in recovery mode
	0729  Sansa E200 series in recovery mode
	0810  SDDR-75 ImageMate CF-SM Reader
	0830  ImageMate CF/MMC/SD Reader
	1234  Cruzer Mini Flash Drive
	5150  SDCZ2 Cruzer Mini Flash Drive (thin)
	5151  Cruzer Micro Flash Drive
	5153  Cruzer Flash Drive
	5204  Cruzer Crossfire
	5402  U3 Cruzer Micro
	5406  Cruzer Micro U3
	5408  Cruzer Titanium U3
	540e  Cruzer Contour Flash Drive
	5530  Cruzer
	5567  Cruzer Blade
	556c  Ultra
	556d  Memory Vault
	5571  Cruzer Fit
	5576  Cruzer Facet
	557d  Cruzer Force (64GB)
	5580  SDCZ80 Flash Drive
	5581  Ultra
	5e10  Encrypted
	6100  Ultra II SD Plus 2GB
	7100  Cruzer Mini
	7101  Pen Flash
	7102  Cruzer Mini
	7103  Cruzer Mini
	7104  Cruzer Micro Mini 256MB Flash Drive
	7105  Cruzer Mini
	7106  Cruzer Mini
	7112  Cruzer Micro 128MB Flash Drive
	7113  Cruzer Micro 256MB Flash Drive
	7114  Cruzer Mini
	7115  Cruzer Mini
	7301  Sansa e100 series (mtp)
	7302  Sansa e100 series (msc)
	7400  Sansa M200 series (mtp)
	7401  Sansa M200 series (msc)
	7420  Sansa E200 series (mtp)
	7421  Sansa E200 Series (msc)
	7422  Sansa E200 series v2 (mtp)
	7423  Sansa E200 series v2 (msc)
	7430  Sansa M200 series
	7431  Sansa M200 series V4 (msc)
	7432  Sansa Clip (mtp)
	7433  Sansa Clip (msc)
	7434  Sansa Clip V2 (mtp)
	7435  Sansa Clip V2 (msc)
	7450  Sansa C250
	7451  Sansa C240
	7460  Sansa Express
	7480  Sansa Connect
	7481  Sansa Connect (in recovery mode)
	74b0  Sansa View (msc)
	74b1  Sansa View (mtp)
	74c0  Sansa Fuze (mtp)
	74c1  Sansa Fuze (msc)
	74c2  Sansa Fuze V2 (mtp)
	74c3  Sansa Fuze V2 (msc)
	74d0  Sansa Clip+ (mtp)
	74d1  Sansa Clip+ (msc)
	74e5  Sansa Clip Zip
	8181  Pen Flash
	8183  Hi-Speed Mass Storage Device
	8185  SDCZ2 Cruzer Mini Flash Drive (older, thick)
	8888  Card Reader
	8889  SDDR-88 Imagemate 8-in-1 Reader
	8919  Card Reader
	8989  ImageMate 12-in-1 Reader
	9191  ImageMate CF
	9219  Card Reader
	9292  ImageMate CF Reader/Writer
	9393  ImageMate SD-MMC
	9595  ImageMate xD-SM
	9797  ImageMate MS-PRO
	9919  Card Reader
	9999  SDDR-99 5-in-1 Reader
	a7c1  Storage device (SD card reader)
	a7e8  SDDR-113 MicroMate SDHC Reader
	b2b3  SDDR-103 MobileMate SD+ Reader
	b4b5  SDDR-89 V4 ImageMate 12-in-1 Reader
0782  Trackerball
0783  C3PO
	0003  LTC31 SmartCard Reader
	0006  LTC31v2
	0009  KBR36
	0010  LTC32
0784  Vivitar, Inc.
	0100  Vivicam 2655
	1310  Vivicam 3305
	1688  Vivicam 3665
	1689  Gateway DC-M42/Labtec DC-505/Vivitar Vivicam 3705
	2620  AOL Photocam Plus
	2888  Polaroid DC700
	3330  Nytec ND-3200 Camera
	4300  Traveler D1
	5260  Werlisa Sport PX 100 / JVC GC-A33 Camera
	5300  Pretec dc530
0785  NTT-ME
	0001  MN128mini-V ISDN TA
	0003  MN128mini-J ISDN TA
0789  Logitec Corp.
	0026  LHD Device
	0033  DVD Multi-plus unit LDR-H443SU2
	0063  LDR Device
	0064  LDR-R Device
	00b3  DVD Multi-plus unit LDR-H443U2
	0105  LAN-TX/U1H2 10/100 Ethernet Adapter [pegasus II]
	010c  Realtek RTL8187 Wireless 802.11g 54Mbps Network Adapter
	0160  LAN-GTJ/U2A
	0162  LAN-WN22/U2 Wireless LAN Adapter
	0163  LAN-WN12/U2 Wireless LAN Adapter
	0164  LAN-W150/U2M Wireless LAN Adapter
	0166  LAN-W300N/U2 Wireless LAN Adapter
	0168  LAN-W150N/U2 Wireless LAN Adapter
	0170  LAN-W300AN/U2 Wireless LAN Adapter
078b  Happ Controls, Inc.
	0010  Driving UGCI
	0020  Flying UGCI
	0030  Fighting UGCI
078c  GTCO/CalComp
	0090  Tablet Adapter
	0100  Tablet Adapter
	0200  Tablet Adapter
	0300  Tablet Adapter
	0400  Digitizer (Whiteboard)
078e  Brincom, Inc.
0790  Pro-Image Manufacturing Co., Ltd
0791  Copartner Wire and Cable Mfg. Corp.
0792  Axis Communications AB
0793  Wha Yu Industrial Co., Ltd
0794  ABL Electronics Corp.
0795  RealChip, Inc.
0796  Certicom Corp.
0797  Grandtech Semiconductor Corp.
	6801  Flatbed Scanner
	6802  InkJet Color Printer
	8001  SmartCam
	801a  Typhoon StyloCam
	801c  Meade Binoculars/Camera
	8901  ScanHex SX-35a
	8909  ScanHex SX-35b
	8911  ScanHex SX-35c
0798  Optelec
	0001  Braille Voyager
	0640  BC640
	0680  BC680
0799  Altera
	7651  Programming Unit
079b  Sagem
	0024  MSO300/MSO301 Fingerprint Sensor
	0026  MSO350/MSO351 Fingerprint Sensor & SmartCard Reader
	0027  USB-Serial Controller
	002f  Mobile
	0030  Mobile Communication Device
	0042  Mobile
	0047  CBM/MSO1300 Fingerprint Sensor
	004a  XG-760A 802.11bg
	004b  Wi-Fi 11g adapter
	0052  MSO1350 Fingerprint Sensor & SmartCard Reader
	0056  Agfa AP1100 Photo Printer
	005d  Mobile Mass Storage
	0062  XG-76NA 802.11bg
	0078  Laser Pro Monochrome MFP
079d  Alfadata Computer Corp.
	0201  GamePort Adapter
07a1  Digicom S.p.A.
	d952  Palladio USB V.92 Modem
07a2  National Technical Systems
07a3  Onnto Corp.
07a4  Be, Inc.
07a6  ADMtek, Inc.
	07c2  AN986A Ethernet
	0986  AN986 Pegasus Ethernet
	8266  Infineon WildCard-USB Wireless LAN Adapter
	8511  ADM8511 Pegasus II Ethernet
	8513  AN8513 Ethernet
	8515  AN8515 Ethernet
07aa  Corega K.K.
	0001  Ether USB-T Ethernet [klsi]
	0004  FEther USB-TX Ethernet [pegasus]
	000c  WirelessLAN USB-11
	000d  FEther USB-TXS
	0011  Wireless LAN USB-11 mini
	0012  Stick-11 802.11b Adapter
	0017  FEther USB2-TX
	0018  Wireless LAN USB-11 mini 2
	001a  ULUSB-11 Key
	001c  CG-WLUSB2GT 802.11g Wireless Adapter [Intersil ISL3880]
	0020  CG-WLUSB2GTST 802.11g Wireless Adapter [Intersil ISL3887]
	002e  CG-WLUSB2GPX [Ralink RT2571W]
	002f  CG-WLUSB2GNL
	0031  CG-WLUSB2GS 802.11bg [Atheros AR5523]
	003c  CG-WLUSB2GNL
	003f  CG-WLUSB300AGN
	0041  CG-WLUSB300GNS
	0042  CG-WLUSB300GNM
	0043  CG-WLUSB300N rev A2 [Realtek RTL8192U]
	0047  CG-WLUSBNM
	0051  CG-WLUSB300NM
	7613  Stick-11 V2 802.11b Adapter
	9601  FEther USB-TXC
07ab  Freecom Technologies
	fc01  IDE bridge
	fc02  Cable II USB-2
	fc03  USB2-IDE IDE bridge
	fcd6  Freecom HD Classic
	fcf6  DataBar
	fcf8  Freecom Classic SL Network Drive
	fcfe  Hard Drive 80GB
07af  Microtech
	0004  SCSI-DB25 SCSI Bridge [shuttle]
	0005  SCSI-HD50 SCSI Bridge [shuttle]
	0006  CameraMate SmartMedia and CompactFlash Card Reader [eusb/shuttle]
	fc01  Freecom USB-IDE
07b0  Trust Technologies
	0001  ISDN TA
	0002  ISDN TA128 Plus
	0003  ISDN TA128 Deluxe
	0005  ISDN TA128 SE
	0006  ISDN TA 128 [HFC-S]
	0007  ISDN TA [HFC-S]
	0008  ISDN TA
07b1  IMP, Inc.
07b2  Motorola BCS, Inc.
	0100  SURFboard Voice over IP Cable Modem
	0900  SURFboard Gateway
	0950  SURFboard SBG950 Gateway
	1000  SURFboard SBG1000 Gateway
	4100  SurfBoard SB4100 Cable Modem
	4200  SurfBoard SB4200 Cable Modem
	4210  SurfBoard 4210 Cable Modem
	4220  SURFboard SB4220 Cable Modem
	4500  CG4500 Communications Gateway
	450b  CG4501 Communications Gateway
	450e  CG4500E Communications Gateway
	5100  SurfBoard SB5100 Cable Modem
	5101  SurfBoard SB5101 Cable Modem
	5120  SurfBoard SB5120 Cable Modem (RNDIS)
	5121  Surfboard 5121 Cable Modem
	7030  WU830G 802.11bg Wireless Adapter [Envara WiND512]
07b3  Plustek, Inc.
	0001  OpticPro 1212U Scanner
	0003  Scanner
	0010  OpticPro U12 Scanner
	0011  OpticPro U24 Scanner
	0013  OpticPro UT12 Scanner
	0014  Scanner
	0015  OpticPro U24 Scanner
	0017  OpticPro UT12/16/24 Scanner
	0204  Scanner
	0400  OpticPro 1248U Scanner
	0401  OpticPro 1248U Scanner #2
	0403  OpticPro U16B Scanner
	0404  Scanner
	0405  A8 Namecard-s Controller
	0406  A8 Namecard-D Controller
	0410  Scanner
	0412  Scanner
	0413  OpticSlim 1200 Scanner
	0601  OpticPro ST24 Scanner
	0800  OpticPro ST48 Scanner
	0900  OpticBook 3600 Scanner
	090c  OpticBook 3600 Plus Scanner
	0a06  TVcam VD100
	0b00  SmartPhoto F50
	0c00  OpticPro ST64 Scanner
	0c03  OpticPro ST64+ Scanner
	0c04  Optic Film 7200i scanner
	0c0c  PL806 Scanner
	0c26  OpticBook 4600 Scanner
	0c2b  Mobile Office D428 Scanner
	0e08  OpticBook A300 Scanner
	1300  OpticBook 3800 Scanner
	1301  OpticBook 4800 Scanner
07b4  Olympus Optical Co., Ltd
	0100  Camedia C-2100/C-3000 Ultra Zoom Camera
	0102  Camedia E-10/C-220/C-50 Camera
	0105  Camedia C-310Z/C-700/C-750UZ/C-755/C-765UZ/C-3040/C-4000/C-5050Z/D-560/C-3020Z Zoom Camera
	0109  C-370Z/C-500Z/D-535Z/X-450
	010a  MAUSB-10 xD and SmartMedia Card Reader
	0112  MAUSB-100 xD Card Reader
	0113  Mju 500 / Stylus Digital Camera (PTP)
	0114  C-350Z Camera
	0118  Mju Mini Digital/Mju Digital 500 Camera / Stylus 850 SW
	0125  Tough TG-1 Camera
	0184  P-S100 port
	0202  Foot Switch RS-26
	0203  Digital Voice Recorder DW-90
	0206  Digital Voice Recorder DS-330
	0207  Digital Voice Recorder & Camera W-10
	0209  Digital Voice Recorder DM-20
	020b  Digital Voice Recorder DS-4000
	020d  Digital Voice Recorder VN-240PC
	0211  Digital Voice Recorder DS-2300
	0218  Foot Switch RS-28
	0244  Digital Voice Recorder VN-8500PC
	024f  Digital Voice Recorder DS-7000
	0280  m:robe 100
07b5  Mega World International, Ltd
	0017  Joystick
	0213  Thrustmaster Firestorm Digital 3 Gamepad
	0312  Gamepad
	9902  GamePad
07b6  Marubun Corp.
07b7  TIME Interconnect, Ltd
07b8  AboCom Systems Inc
	110c  XX1
	1201  IEEE 802.11b Adapter
	200c  XX2
	2573  Wireless LAN Card
	2770  802.11n/b/g Mini Wireless LAN USB2.0 Adapter
	2870  802.11n/b/g Wireless LAN USB2.0 Adapter
	3070  802.11n/b/g Mini Wireless LAN USB2.0 Adapter
	3071  802.11n/b/g Mini Wireless LAN USB2.0 Adapter
	3072  802.11n/b/g Mini Wireless LAN USB2.0 Adapter
	4000  DU-E10 Ethernet [klsi]
	4002  DU-E100 Ethernet [pegasus]
	4003  1/10/100 Ethernet Adapter
	4004  XX4
	4007  XX5
	400b  XX6
	400c  XX7
	401a  RTL8151
	4102  USB 1.1 10/100M Fast Ethernet Adapter
	4104  XX9
	420a  UF200 Ethernet
	5301  GW-US54ZGL 802.11bg
	6001  802.11bg
	8188  AboCom Systems Inc [WN2001 Prolink Wireless-N Nano Adapter]
	a001  WUG2200 802.11g Wireless Adapter [Envara WiND512]
	abc1  DU-E10 Ethernet [pegasus]
	b000  BWU613
	b02a  AboCom Bluetooth Device
	b02b  Bluetooth dongle
	b02c  BCM92045DG-Flash with trace filter
	b02d  BCM92045DG-Flash with trace filter
	b02e  BCM92045DG-Flash with trace filter
	b030  BCM92045DG-Flash with trace filter
	b031  BCM92045DG-Flash with trace filter
	b032  BCM92045DG-Flash with trace filter
	b033  BCM92045DG-Flash with trace filter
	b21a  WUG2400 802.11g Wireless Adapter [Texas Instruments TNETW1450]
	b21b  HWU54DM
	b21c  RT2573
	b21d  RT2573
	b21e  RT2573
	b21f  WUG2700
	d011  MP3 Player
	e001  Mass Storage Device
	e002  Mass Storage Device
	e003  Mass Storage Device
	e004  Mass Storage Device
	e005  Mass Storage Device
	e006  Mass Storage Device
	e007  Mass Storage Device
	e008  Mass Storage Device
	e009  Mass Storage Device
	e00a  Mass Storage Device
	e4f0  Card Reader Driver
	f101  DSB-560 Modem [atlas]
07bc  Canon Computer Systems, Inc.
07bd  Webgear, Inc.
07be  Veridicom
07c0  Code Mercenaries Hard- und Software GmbH
	1113  JoyWarrior24F8
	1116  JoyWarrior24F14
	1121  The Claw
	1500  IO-Warrior 40
	1501  IO-Warrior 24
	1502  IO-Warrior 48
	1503  IO-Warrior 28
	1511  IO-Warrior 24 Power Vampire
	1512  IO-Warrior 24 Power Vampire
07c1  Keisokugiken
	0068  HKS-0200 USBDAQ
07c4  Datafab Systems, Inc.
	0102  USB to LS120
	0103  USB to IDE
	1234  USB to ATAPI
	a000  CompactFlash Card Reader
	a001  CompactFlash & SmartMedia Card Reader [eusb]
	a002  Disk Drive
	a003  Datafab-based Reader
	a004  USB to MMC Class Drive
	a005  CompactFlash & SmartMedia Card Reader
	a006  SmartMedia Card Reader
	a007  Memory Stick Class Drive
	a103  MDSM-B reader
	a107  USB to Memory Stick (LC1) Drive
	a109  LC1 CompactFlash & SmartMedia Card Reader
	a10b  USB to CF+MS(LC1)
	a200  DF-UT-06 Hama MMC/SD Reader
	a400  CompactFlash & Microdrive Reader
	a600  Card Reader
	a604  12-in-1 Card Reader
	ad01  Mass Storage Device
	ae01  Mass Storage Device
	af01  Mass Storage Device
	b000  USB to CF(LC1)
	b001  USB to CF+PCMCIA
	b004  MMC/SD Reader
	b006  USB to PCMCIA
	b00a  USB to CF+SD Drive(LC1)
	b00b  USB to Memory Stick(LC1)
	c010  Kingston FCR-HS2/ATA Card Reader
07c5  APG Cash Drawer
	0500  Cash Drawer
07c6  ShareWave, Inc.
	0002  Bodega Wireless Access Point
	0003  Bodega Wireless Network Adapter
07c7  Powertech Industrial Co., Ltd
07c8  B.U.G., Inc.
	0202  MN128-SOHO PAL
07c9  Allied Telesyn International
	b100  AT-USB100
07ca  AVerMedia Technologies, Inc.
	0002  AVerTV PVR USB/EZMaker Pro Device
	0026  AVerTV
	0337  A867 DVB-T dongle
	0837  H837 Hybrid ATSC/QAM
	1228  MPEG-2 Capture Device (M038)
	1830  AVerTV Volar Video Capture (H830)
	3835  AVerTV Volar Green HD (A835B)
	850a  AverTV Volar Black HD (A850)
	850b  AverTV Red HD+ (A850T)
	a309  AVerTV DVB-T (A309)
	a801  AVerTV DVB-T (A800)
	a815  AVerTV DVB-T Volar X (A815)
	a827  AVerTV Hybrid Volar HX (A827)
	a867  AVerTV DVB-T (A867)
	b300  A300 DVB-T TV receiver
	b800  MR800 FM Radio
	e880  MPEG-2 Capture Device (E880)
	e882  MPEG-2 Capture Device (E882)
07cb  Kingmax Technology, Inc.
07cc  Carry Computer Eng., Co., Ltd
	0000  CF Card Reader
	0001  Reader (UICSE)
	0002  Reader (UIS)
	0003  SM Card Reader
	0004  SM/CF/PCMCIA Card Reader
	0005  Reader (UISA2SE)
	0006  SM/CF/PCMCIA Card Reader
	0007  Reader (UISA6SE)
	000c  SM/CF Card Reader
	000d  SM/CF Card Reader
	000e  Reader (UISDA)
	000f  Reader (UICLIK)
	0010  Reader (UISMA)
	0012  Reader (UISC6SE-FLASH)
	0014  Litronic Fortezza Reader
	0030  Mass Storage (UISDMC12S)
	0040  Mass Storage (UISDMC13S)
	0100  Reader (UID)
	0101  Reader (UIM)
	0102  Reader (UISDMA)
	0103  Reader (UISDMC)
	0104  Reader (UISDM)
	0200  6-in-1 Card Reader
	0201  Mass Storage (UISDMC1S & UISDMC3S)
	0202  Mass Storage (UISDMC5S)
	0203  Mass Storage (UISMC5S)
	0204  Mass Storage (UIM4/5S & UIM7S)
	0205  Mass Storage (UIS4/5S & UIS7S)
	0206  Mass Storage (UISDMC10S & UISDMC11S)
	0207  Mass Storage (UPIDMA)
	0208  Mass Storage (UCFC II)
	0210  Mass Storage (UPIXXA)
	0213  Mass Storage (UPIDA)
	0214  Mass Storage (UPIMA)
	0215  Mass Storage (UPISA)
	0217  Mass Storage (UPISDMA)
	0223  Mass Storage (UCIDA)
	0224  Mass Storage (UCIMA)
	0225  Mass Storage (UIS7S)
	0227  Mass Storage (UCIDMA)
	0234  Mass Storage (UIM7S)
	0235  Mass Storage (UIS4S-S)
	0237  Velper (UISDMC4S)
	0300  6-in-1 Card Reader
	0301  6-in-1 Card Reader
	0303  Mass Storage (UID10W)
	0304  Mass Storage (UIM10W)
	0305  Mass Storage (UIS10W)
	0308  Mass Storage (UIC10W)
	0309  Mass Storage (UISC3W)
	0310  Mass Storage (UISDMA2W)
	0311  Mass Storage (UISDMC14W)
	0320  Mass Storage (UISDMC4W)
	0321  Mass Storage (UISDMC37W)
	0330  WINTERREADER Reader
	0350  9-in-1 Card Reader
	0500  Mass Storage
	0501  Mass Storage
07cd  Elektor
	0001  USBuart Serial Port
07cf  Casio Computer Co., Ltd
	1001  QV-8000SX/5700/3000EX Digicam; Exilim EX-M20
	1003  Exilim EX-S500
	1004  Exilim EX-Z120
	1011  USB-CASIO PC CAMERA
	1116  EXILIM EX-Z19
	1125  Exilim EX-H10 Digital Camera (mass storage mode)
	1133  Exilim EX-Z350 Digital Camera (mass storage mode)
	1225  Exilim EX-H10 Digital Camera (PictBridge mode)
	1233  Exilim EX-Z350 Digital Camera (PictBridge mode)
	2002  E-125 Cassiopeia Pocket PC
	3801  WMP-1 MP3-Watch
	4001  Label Printer KL-P1000
	4007  CW50 Device
	4104  Cw75 Device
	4107  CW-L300 Device
	4500  LV-20 Digital Camera
	6101  fx-9750gII
	6801  PL-40R
	6802  MIDI Keyboard
07d0  Dazzle
	0001  Digital Video Creator I
	0002  Global Village VideoFX Grabber
	0003  Fusion Model DVC-50 Rev 1 (NTSC)
	0004  DVC-800 (PAL) Grabber
	0005  Fusion Video and Audio Ports
	0006  DVC 150 Loader Device
	0007  DVC 150
	0327  Fusion Digital Media Reader
	1001  DM-FLEX DFU Adapter
	1002  DMHS2 DFU Adapter
	1102  CF Reader/Writer
	1103  SD Reader/Writer
	1104  SM Reader/Writer
	1105  MS Reader/Writer
	1106  xD/SM Reader/Writer
	1202  MultiSlot Reader/Writer
	2000  FX2 DFU Adapter
	2001  eUSB CompactFlash Reader
	4100  Kingsun SF-620 Infrared Adapter
	4101  Connectivity Cable (CA-42 clone)
	4959  Kingsun KS-959 Infrared Adapter
07d1  D-Link System
	13ec  VvBus for Helium 2xx
	13ed  VvBus for Helium 2xx
	13f1  DSL-302G Modem
	13f2  DSL-502G Router
	3300  DWA-130 802.11n Wireless N Adapter(rev.E) [Realtek RTL8191SU]
	3302  DWA-130 802.11n Wireless N Adapter(rev.C2) [Realtek RTL8191SU]
	3303  DWA-131 802.11n Wireless N Nano Adapter(rev.A1) [Realtek RTL8192SU]
	3304  FR-300USB 802.11bgn Wireless Adapter
	3a07  WUA-2340 RangeBooster G Adapter(rev.A) [Atheros AR5523]
	3a08  WUA-2340 RangeBooster G Adapter(rev.A) (no firmware) [Atheros AR5523]
	3a09  DWA-160 802.11abgn Xtreme N Dual Band Adapter(rev.A2) [Atheros AR9170+AR9104]
	3a0d  DWA-120 802.11g Wireless 108G Adapter [Atheros AR5523]
	3a0f  DWA-130 802.11n Wireless N Adapter(rev.D) [Atheros AR9170+AR9102]
	3a10  DWA-126 802.11n Wireless Adapter [Atheros AR9271]
	3b01  AirPlus G DWL-G122 Wireless Adapter(rev.D) [Marvell 88W8338+88W8010]
	3b10  DWA-142 RangeBooster N Adapter [Marvell 88W8362+88W8060]
	3b11  DWA-130 802.11n Wireless N Adapter(rev.A1) [Marvell 88W8362+88W8060]
	3c03  AirPlus G DWL-G122 Wireless Adapter(rev.C1) [Ralink RT2571W]
	3c04  WUA-1340
	3c05  EH103 Wireless G Adapter
	3c06  DWA-111 802.11bg Wireless Adapter [Ralink RT2571W]
	3c07  DWA-110 Wireless G Adapter(rev.A1) [Ralink RT2571W]
	3c09  DWA-140 RangeBooster N Adapter(rev.B1) [Ralink RT2870]
	3c0a  DWA-140 RangeBooster N Adapter(rev.B2) [Ralink RT3072]
	3c0b  DWA-110 Wireless G Adapter(rev.B) [Ralink RT2870]
	3c0d  DWA-125 Wireless N 150 Adapter(rev.A1) [Ralink RT3070]
	3c0e  WUA-2340 RangeBooster G Adapter(rev.B) [Ralink RT2070]
	3c0f  AirPlus G DWL-G122 Wireless Adapter(rev.E1) [Ralink RT2070]
	3c10  DWA-160 802.11abgn Xtreme N Dual Band Adapter(rev.A1) [Atheros AR9170+AR9104]
	3c11  DWA-160 Xtreme N Dual Band USB Adapter(rev.B) [Ralink RT2870]
	3c13  DWA-130 802.11n Wireless N Adapter(rev.B) [Ralink RT2870]
	3c15  DWA-140 RangeBooster N Adapter(rev.B3) [Ralink RT2870]
	3c16  DWA-125 Wireless N 150 Adapter(rev.A2) [Ralink RT3070]
	3e02  DWM-156 3.75G HSUPA Adapter
	5100  Remote NDIS Device
	a800  DWM-152 3.75G HSUPA Adapter
	f101  DBT-122 Bluetooth
	fc01  DBT-120 Bluetooth Adapter
07d2  Aptio Products, Inc.
07d3  Cyberdata Corp.
07d5  Radiant Systems
07d7  GCC Technologies, Inc.
07da  Arasan Chip Systems
07de  Diamond Multimedia
	2820  VC500 Video Capture Dongle
07df  David Electronics Co., Ltd
07e1  Ambient Technologies, Inc.
	5201  V.90 Modem
07e2  Elmeg GmbH & Co., Ltd
07e3  Planex Communications, Inc.
07e4  Movado Enterprise Co., Ltd
	0967  SCard R/W CSR-145
	0968  SCard R/W CSR-145
07e5  QPS, Inc.
	05c2  IDE-to-USB2.0 PCA
	5c01  Que! CDRW
07e6  Allied Cable Corp.
07e7  Mirvo Toys, Inc.
07e8  Labsystems
07ea  Iwatsu Electric Co., Ltd
07eb  Double-H Technology Co., Ltd
07ec  Taiyo Electric Wire & Cable Co., Ltd
07ee  Torex Retail (formerly Logware)
	0002  Cash Drawer I/F
07ef  STSN
	0001  Internet Access Device
07f2  Microcomputer Applications, Inc.
	0001  KEYLOK II
07f6  Circuit Assembly Corp.
07f7  Century Corp.
	0005  ScanLogic/Century Corporation uATA
	011e  Century USB Disk Enclosure
07f9  Dotop Technology, Inc.
07fa  DrayTek Corp.
	0778  miniVigor 128 ISDN TA
	0846  ISDN TA [HFC-S]
	0847  ISDN TA [HFC-S]
	1012  BeWAN ADSL USB ST (grey)
	1196  BWIFI-USB54AR 802.11bg
	a904  BeWAN ADSL
	a905  BeWAN ADSL ST
07fd  Mark of the Unicorn
	0000  FastLane MIDI Interface
	0001  FastLane Quad MIDI Interface
	0002  MOTU Audio for 64 bit
07ff  Unknown
	00ff  Portable Hard Drive
0801  MagTek
	0001  Mini Swipe Reader (Keyboard Emulation)
	0002  Mini Swipe Reader
	0003  Magstripe Insert Reader
0802  Mako Technologies, LLC
0803  Zoom Telephonics, Inc.
	1300  V92 Faxmodem
	3095  V.92 56K Mini External Modem Model 3095
	4310  4410a Wireless-G Adapter [Intersil ISL3887]
	4410  4410b Wireless-G Adapter [ZyDAS ZD1211B]
	5241  Cable Modem
	5551  DSL Modem
	9700  2986L FaxModem
	9800  Cable Modem
	a312  Wireless-G
0809  Genicom Technology, Inc.
080a  Evermuch Technology Co., Ltd
080b  Cross Match Technologies
	0002  Fingerprint Scanner (After ReNumeration)
	0010  300LC Series Fingerprint Scanner (Before ReNumeration)
080c  Datalogic S.p.A.
	0300  Gryphon D120 Barcode Scanner
	0400  Gryphon D120 Barcode Scanner
	0500  Gryphon D120 Barcode Scanner
	0600  Gryphon M100 Barcode Scanner
080d  Teco Image Systems Co., Ltd
	0102  Hercules Scan@home 48
	0104  3.2Slim
	0110  UMAX AstraSlim 1200 Scanner
0810  Personal Communication Systems, Inc.
	0001  Dual PSX Adaptor
	0002  Dual PCS Adaptor
	0003  PlayStation Gamepad
0813  Mattel, Inc.
	0001  Intel Play QX3 Microscope
	0002  Dual Mode Camera Plus
0819  eLicenser
	0101  License Management and Copy Protection
081a  MG Logic
	1000  Duo Pen Tablet
081b  Indigita Corp.
	0600  Storage Adapter
	0601  Storage Adapter
081c  Mipsys
081e  AlphaSmart, Inc.
	df00  Handheld
0822  Reudo Corp.
	2001  IRXpress Infrared Device
0825  GC Protronics
0826  Data Transit
0827  BroadLogic, Inc.
0828  Sato Corp.
0829  DirecTV Broadband, Inc. (Telocity)
082d  Handspring
	0100  Visor
	0200  Treo
	0300  Treo 600
	0400  Handheld
	0500  Handheld
	0600  Handheld
0830  Palm, Inc.
	0001  m500
	0002  m505
	0003  m515
	0004  Handheld
	0005  Handheld
	0006  Handheld
	0010  Handheld
	0011  Handheld
	0012  Handheld
	0013  Handheld
	0014  Handheld
	0020  i705
	0021  Handheld
	0022  Handheld
	0023  Handheld
	0024  Handheld
	0030  Handheld
	0031  Tungsten W
	0032  Handheld
	0033  Handheld
	0034  Handheld
	0040  m125
	0041  Handheld
	0042  Handheld
	0043  Handheld
	0044  Handheld
	0050  m130
	0051  Handheld
	0052  Handheld
	0053  Handheld
	0054  Handheld
	0060  Tungsten C/E/T/T2/T3 / Zire 71
	0061  Lifedrive / Treo 650/680 / Tunsten E2/T5/TX / Centro / Zire 21/31/72 / Z22
	0062  Handheld
	0063  Handheld
	0064  Handheld
	0070  Zire
	0071  Handheld
	0072  Handheld
	0080  Serial Adapter [for Palm III]
	0081  Handheld
	0082  Handheld
	00a0  Treo 800w
	0101  Pre
0832  Kouwell Electronics Corp.
	5850  Cable
0833  Sourcenext Corp.
	012e  KeikaiDenwa 8 with charger
	039f  KeikaiDenwa 8
0835  Action Star Enterprise Co., Ltd
0836  TrekStor
	2836  i.Beat mood
0839  Samsung Techwin Co., Ltd
	0005  Digimax Camera
	0008  Digimax 230 Camera
	0009  Digimax 340
	000a  Digimax 410
	000e  Digimax 360
	0010  Digimax 300
	1003  Digimax 210SE
	1005  Digimax 220
	1009  Digimax V4
	1012  6500 Document Camera
	1058  S730 Camera
	1064  Digimax D830 Camera
	1542  Digimax 50 Duo
	3000  Digimax 35 MP3
083a  Accton Technology Corp.
	1046  10/100 Ethernet [pegasus]
	1060  HomeLine Adapter
	1f4d  SMC8013WG Broadband Remote NDIS Device
	3046  10/100 Series Adapter
	3060  1/10/100 Adapter
	3501  2664W
	3502  WN3501D Wireless Adapter
	3503  T-Sinus 111 Wireless Adapter
	4501  T-Sinus 154data
	4502  Siemens S30853-S1016-R107 802.11g Wireless Adapter [Intersil ISL3886]
	4505  SMCWUSB-G 802.11bg
	4507  SMCWUSBT-G2 802.11g Wireless Adapter [Atheros AR5523]
	4521  Siemens S30863-S1016-R107-2 802.11g Wireless Adapter [Intersil ISL3887]
	4531  T-Com Sinus 154 data II [Intersil ISL3887]
	5046  SpeedStream 10/100 Ethernet [pegasus]
	5501  Wireless Adapter 11g
	6500  Cable Modem
	6618  802.11n Wireless Adapter
	7511  Arcadyan 802.11N Wireless Adapter
	7512  Arcadyan 802.11N Wireless Adapter
	7522  Arcadyan 802.11N Wireless Adapter
	8522  Arcadyan 802.11N Wireless Adapter
	8541  WN4501F 802.11g Wireless Adapter [Intersil ISL3887]
	a512  Arcadyan 802.11N Wireless Adapter
	a618  SMCWUSBS-N EZ Connect N Draft 11n Wireless Adapter [Ralink RT2870]
	a701  SMCWUSBS-N3 EZ Connect N Wireless Adapter [Ralink RT3070]
	b004  CPWUE001 USB/Ethernet Adapter
	b522  SMCWUSBS-N2 EZ Connect N Wireless Adapter [Ralink RT2870]
	bb01  BlueExpert Bluetooth Device
	c003  802.11b Wireless Adapter
	c501  Zoom 4410 Wireless-G [Intersil ISL3887]
	c561  802.11a/g Wireless Adapter
	d522  Speedport W 102 Stick IEEE 802.11n USB 2.0 Adapter
	e501  ZD1211B
	e503  Arcadyan WN4501 802.11b/g
	e506  WUS-201 802.11bg
	f501  802.11g Wireless Adapter
	f502  802.11g Wireless Adapter
	f522  Arcadyan WN7512 802.11n
083f  Global Village
	b100  TelePort V.90 Fax/Modem
0840  Argosy Research, Inc.
	0060  Storage Adapter Bridge Module
0841  Rioport.com, Inc.
	0001  Rio 500
0844  Welland Industrial Co., Ltd
0846  NetGear, Inc.
	1001  EA101 10 Mbps 10BASE-T Ethernet [Kawasaki LSI KL5KLUSB101B]
	1002  Ethernet
	1020  FA101 Fast Ethernet USB 1.1
	1040  FA120 Fast Ethernet USB 2.0 [Asix AX88172 / AX8817x]
	1100  Managed Switch M4100 series, M5300 series, M7100 series
	4110  MA111(v1) 802.11b Wireless [Intersil Prism 3.0]
	4200  WG121(v1) 54 Mbps Wireless [Intersil ISL3886]
	4210  WG121(v2) 54 Mbps Wireless [Intersil ISL3886]
	4220  WG111(v1) 54 Mbps Wireless [Intersil ISL3886]
	4230  MA111(v2) 802.11b Wireless [SIS SIS 162]
	4240  WG111(v1) rev 2 54 Mbps Wireless [Intersil ISL3887]
	4260  WG111v3 54 Mbps Wireless [realtek RTL8187B]
	4300  WG111U Double 108 Mbps Wireless [Atheros AR5004X / AR5005UX]
	4301  WG111U (no firmware) Double 108 Mbps Wireless [Atheros AR5004X / AR5005UX]
	5f00  WPN111 802.11g Wireless Adapter [Atheros AR5523]
	6a00  WG111v2 54 Mbps Wireless [RealTek RTL8187L]
	7100  WN121T RangeMax Next Wireless-N [Marvell TopDog]
	9000  WN111(v1) RangeMax Next Wireless [Marvell 88W8362+88W8060]
	9001  WN111(v2) RangeMax Next Wireless [Atheros AR9170+AR9101]
	9010  WNDA3100v1 802.11abgn [Atheros AR9170+AR9104]
	9011  WNDA3100v2 802.11abgn [Broadcom BCM4323]
	9012  WNDA4100 802.11abgn 3x3:3 [Ralink RT3573]
	9018  WNDA3200 802.11abgn Wireless Adapter [Atheros AR7010+AR9280]
	9020  WNA3100(v1) Wireless-N 300 [Broadcom BCM43231]
	9021  WNA3100M(v1) Wireless-N 300 [Realtek RTL8192CU]
	9030  WNA1100 Wireless-N 150 [Atheros AR9271]
	9040  WNA1000 Wireless-N 150 [Atheros AR9170+AR9101]
	9041  WNA1000M 802.11bgn [Realtek RTL8188CUS]
	9042  On Networks N150MA 802.11bgn [Realtek RTL8188CUS]
	9043  WNA1000Mv2 802.11bgn [Realtek RTL8188CUS?]
	9050  A6200 802.11a/b/g/n/ac Wireless Adapter [Broadcom BCM43526]
	9052  A6100 AC600 DB Wireless Adapter [Realtek RTL8811AU]
	a001  PA101 10 Mbps HPNA Home Phoneline RJ-1
	f001  On Networks N300MA 802.11bgn [Realtek RTL8192CU]
084d  Minton Optic Industry Co., Inc.
	0001  Jenoptik JD800i
	0003  S-Cam F5/D-Link DSC-350 Digital Camera
	0011  Argus DC3500 Digital Camera
	0014  Praktica DC 32
	0019  Praktica DPix3000
	0025  Praktica DC 60
	1001  ScanHex SX-35d
084e  KB Gear
	0001  JamCam Camera
	1001  Jam Studio Tablet
	1002  Pablo Tablet
084f  Empeg
	0001  Empeg-Car Mark I/II Player
0850  Fast Point Technologies, Inc.
0851  Macronix International Co., Ltd
	1542  SiPix Blink
	1543  Maxell WS30 Slim Digital Camera, or Pandigital PI8004W01 digital photo frame
	a168  MXIC
0852  CSEM
0853  Topre Corporation
	0100  HHKB Professional
0854  ActiveWire, Inc.
	0100  I/O Board
	0101  I/O Board, rev1
0856  B&B Electronics
	ac01  uLinks USOTL4 RS422/485 Adapter
0858  Hitachi Maxell, Ltd
	3102  Bluetooth Device
	ffff  Maxell module with BlueCore in DFU mode
0859  Minolta Systems Laboratory, Inc.
085a  Xircom
	0001  Portstation Dual Serial Port
	0003  Portstation Paraller Port
	0008  Ethernet
	0009  Ethernet
	000b  Portstation Dual PS/2 Port
	0021  1 port to Serial Converter
	0022  Parallel Port
	0023  2 port to Serial Converter
	0024  Parallel Port
	0026  PortGear SCSI
	0027  1 port to Serial Converter
	0028  PortGear to SCSI Converter
	0032  PortStation SCSI Module
	003c  Bluetooth Adapter
	0299  Colorvision, Inc. Monitor Spyder
	8021  1 port to Serial
	8023  2 port to Serial
	8027  PGSDB9 Serial Port
085c  ColorVision, Inc.
	0100  Spyder 1
	0200  Spyder 2
	0300  Spyder 3
	0400  Spyder 4
0862  Teletrol Systems, Inc.
0863  Filanet Corp.
0864  NetGear, Inc.
	4100  MA101 802.11b Adapter
	4102  MA101 802.11b Adapter
0867  Data Translation, Inc.
	9812  ECON Data acquisition unit
	9816  DT9816 ECON data acquisition module
	9836  DT9836 data acquisition card
086a  Emagic Soft- und Hardware GmbH
	0001  Unitor8
	0002  AMT8
	0003  MT4
086c  DeTeWe - Deutsche Telephonwerke AG & Co.
	1001  Eumex 504PC ISDN TA
	1002  Eumex 504PC (FlashLoad)
	1003  TA33 ISDN TA
	1004  TA33 (FlashLoad)
	1005  Eumex 604PC HomeNet
	1006  Eumex 604PC HomeNet (FlashLoad)
	1007  Eumex 704PC DSL
	1008  Eumex 704PC DSL (FlashLoad)
	1009  Eumex 724PC DSL
	100a  Eumex 724PC DSL (FlashLoad)
	100b  OpenCom 30
	100c  OpenCom 30 (FlashLoad)
	100d  BeeTel Home 100
	100e  BeeTel Home 100 (FlashLoad)
	1011  USB2DECT
	1012  USB2DECT (FlashLoad)
	1013  Eumex 704PC LAN
	1014  Eumex 704PC LAN (FlashLoad)
	1019  Eumex 504 SE
	101a  Eumex 504 SE (Flash-Mode)
	1021  OpenCom 40
	1022  OpenCom 40 (FlashLoad)
	1023  OpenCom 45
	1024  OpenCom 45 (FlashLoad)
	1025  Sinus 61 data
	1029  dect BOX
	102c  Eumex 604PC HomeNet [FlashLoad]
	1030  Eumex 704PC DSL [FlashLoad]
	1032  OpenCom 40 [FlashLoad]
	1033  OpenCom 30 plus
	1034  OpenCom 30 plus (FlashLoad)
	1041  Eumex 220PC
	1042  Eumex 220PC (FlashMode)
	1055  Eumex 220 Version 2 ISDN TA
	1056  Eumex 220 Version 2 ISDN TA (Flash-Mode)
	2000  OpenCom 1000
086e  System TALKS, Inc.
	1920  SGC-X2UL
086f  MEC IMEX, Inc.
0870  Metricom
	0001  Ricochet GS
0871  SanDisk, Inc.
	0001  SDDR-01 Compact Flash Reader
	0002  SDDR-31 Compact Flash Reader
	0005  SDDR-05 Compact Flash Reader
0873  Xpeed, Inc.
0874  A-Tec Subsystem, Inc.
0879  Comtrol Corp.
087c  Adesso/Kbtek America, Inc.
087d  Jaton Corp.
	5704  Ethernet
087e  Fujitsu Computer Products of America
087f  QualCore Logic Inc.
0880  APT Technologies, Inc.
0883  Recording Industry Association of America (RIAA)
0885  Boca Research, Inc.
0886  XAC Automation Corp.
	0630  Intel PC Camera CS630
0887  Hannstar Electronics Corp.
088a  TechTools
	1002  DigiView DV3100
088b  MassWorks, Inc.
	4944  MassWorks ID-75 TouchScreen
088c  Swecoin AB
	2030  Ticket Printer TTP 2030
088e  iLok
	5036  Portable secure storage for software licenses
0892  DioGraphy, Inc.
	0101  Smartdio Reader/Writer
0894  TSI Incorporated
	0010  Remote NDIS Network Device
0897  Lauterbach
	0002  Power Debug/Power Debug II
089c  United Technologies Research Cntr.
089d  Icron Technologies Corp.
089e  NST Co., Ltd
089f  Primex Aerospace Co.
08a5  e9, Inc.
08a6  Toshiba TEC
	0051  B-SV4
08a8  Andrea Electronics
08a9  CWAV Inc.
	0005  USBee ZX
	0009  USBee SX
	0012  USBee AX-Standard
	0013  USBee AX-Plus
	0014  USBee AX-Pro
	0015  USBee DX
08ac  Macraigor Systems LLC
	2024  usbWiggler
08ae  Macally (Mace Group, Inc.)
08b0  Metrohm
	0006  814 Sample Processor
	0015  857 Titrando
	001a  852 Titrando
08b4  Sorenson Vision, Inc.
08b7  NATSU
	0001  Playstation adapter
08b8  J. Gordon Electronic Design, Inc.
	01f4  USBSIMM1
08b9  RadioShack Corp. (Tandy)
08bb  Texas Instruments
	2702  Speakers
	2704  Audio Codec
	2706  PCM2706 Audio Codec
	2900  PCM2900 Audio Codec
	2901  PCM2901 Audio Codec
	2902  PCM2902 Audio Codec
	2904  PCM2904 Audio Codec
	2910  PCM2912 Audio Codec
	29b0  PCM2900B Audio CODEC
	29b2  PCM2902 Audio CODEC
	29b3  PCM2903B Audio CODEC
	29b6  PCM2906B Audio CODEC
	29c0  PCM2900C Audio CODEC
	29c2  PCM2902C Audio CODEC
	29c3  PCM2903C Audio CODEC
	29c6  PCM2906C Audio CODEC
08bd  Citizen Watch Co., Ltd
	0208  CLP-521 Label Printer
	1100  X1-USB Floppy
08c3  Precise Biometrics
	0001  100 SC
	0002  100 A
	0003  100 SC BioKeyboard
	0006  100 A BioKeyboard
	0100  100 MC ISP
	0101  100 MC FingerPrint and SmartCard Reader
	0300  100 AX
	0400  100 SC
	0401  150 MC
	0402  200 MC FingerPrint and SmartCard Reader
	0404  100 SC Upgrade
	0405  150 MC Upgrade
	0406  100 MC Upgrade
08c4  Proxim, Inc.
	0100  Skyline 802.11b Wireless Adapter
	02f2  Farallon Home Phoneline Adapter
08c7  Key Nice Enterprise Co., Ltd
08c8  2Wire, Inc.
08c9  Nippon Telegraph and Telephone Corp.
08ca  Aiptek International, Inc.
	0001  Tablet
	0010  Tablet
	0020  APT-6000U Tablet
	0021  APT-2 Tablet
	0022  Tablet
	0023  Tablet
	0024  Tablet
	0100  Pen Drive
	0102  DualCam
	0103  Pocket DV Digital Camera
	0104  Pocket DVII
	0105  Mega DV(Disk)
	0106  Pocket DV3100+
	0107  Pocket DV3100
	0109  Nisis DV4 Digital Camera
	010a  Trust 738AV LCD PV Mass Storage
	0111  PenCam VGA Plus
	2008  Mini PenCam 2
	2010  Pocket CAM 3 Mega (webcam)
	2011  Pocket CAM 3 Mega (storage)
	2016  PocketCam 2 Mega
	2018  Pencam SD 2M
	2020  Slim 3000F
	2022  Slim 3200
	2024  Pocket DV3500
	2028  Pocket Cam4M
	2040  Pocket DV4100M
	2042  Pocket DV5100M Composite Device
	2043  Pocket DV5100M (Disk)
	2060  Pocket DV5300
08cd  Jue Hsun Ind. Corp.
08ce  Long Well Electronics Corp.
08cf  Productivity Enhancement Products
08d1  smartBridges, Inc.
	0001  smartNIC Ethernet [catc]
	0003  smartNIC 2 PnP Ethernet
08d3  Virtual Ink
08d4  Fujitsu Siemens Computers
	0009  SCR SmartCard Reader
08d8  IXXAT Automation GmbH
	0002  USB-to-CAN compact
	0003  USB-to-CAN II
	0100  USB-to-CAN
08d9  Increment P Corp.
08dd  Billionton Systems, Inc.
	0112  Wireless LAN Adapter
	0113  Wireless LAN Adapter
	0986  USB-100N Ethernet [pegasus]
	0987  USBLP-100 HomePNA Ethernet [pegasus]
	0988  USBEL-100 Ethernet [pegasus]
	1986  10/100 LAN Adapter
	2103  DVB-T TV-Tuner Card-R
	8511  USBE-100 Ethernet [pegasus2]
	90ff  USB2AR Ethernet
08de  ???
	7a01  802.11b Adapter
08df  Spyrus, Inc.
	0001  Rosetta Token V1
	0002  Rosetta Token V2
	0003  Rosetta Token V3
	0a00  Lynks Interface
08e3  Olitec, Inc.
	0002  USB-RS232 Bridge
	0100  Interface ADSL
	0101  Interface ADSL
	0102  ADSL
	0301  RNIS ISDN TA [HFC-S]
08e4  Pioneer Corp.
	0184  DDJ-WeGO
	0185  DDJ-WeGO2
08e5  Litronic
08e6  Gemalto (was Gemplus)
	0001  GemPC-Touch 430
	0430  GemPC430 SmartCard Reader
	0432  GemPC432 SmartCard Reader
	0435  GemPC435 SmartCard Reader
	0437  GemPC433 SL SmartCard Reader
	1359  UA SECURE STORAGE TOKEN
	2202  Gem e-Seal Pro Token
	3437  GemPC Twin SmartCard Reader
	3438  GemPC Key SmartCard Reader
	3478  PinPad Smart Card Reader
	34ec  Compact Smart Card Reader Writer
	4433  GemPC433-Swap
	5501  GemProx-PU Contactless Smart Card Reader
	5503  Prox-DU Contactless Interface
	ace0  UA HYBRID TOKEN
08e7  Pan-International Wire & Cable
08e8  Integrated Memory Logic
08e9  Extended Systems, Inc.
	0100  XTNDAccess IrDA Dongle
08ea  Ericsson, Inc., Blue Ridge Labs
	00c9  ADSL Modem HM120dp Loader
	00ca  ADSL WAN Modem HM120dp
	00ce  HM230d Virtual Bus for Helium
	abba  USB Driver for Bluetooth Wireless Technology
	abbb  Bluetooth Device in DFU State
08ec  M-Systems Flash Disk Pioneers
	0001  TravelDrive 2C
	0002  TravelDrive 2C
	0005  TravelDrive 2C
	0008  TravelDrive 2C
	0010  DiskOnKey
	0011  DiskOnKey
	0012  TravelDrive 2C
	0014  TravelDrive 2C
	0015  Kingston DataTraveler ELITE
	0016  Kingston DataTraveler U3
	0020  TravelDrive Intuix U3 2GB
	0021  TravelDrive
	0022  TravelDrive
	0023  TravelDrive
	0024  TravelDrive
	0025  TravelDrive
	0026  TravelDrive
	0027  TravelDrive
	0028  TravelDrive
	0029  TravelDrive
	0030  TravelDrive
	0822  TravelDrive 2C
	0832  Hi-Speed Mass Storage Device
	0834  M-Disk 220
	0998  Kingston Data Traveler2.0 Disk Driver
	0999  Kingston Data Traveler2.0 Disk Driver
	1000  TravelDrive 2C
	2000  TravelDrive 2C
	2038  TravelDrive
	2039  TravelDrive
	204a  TravelDrive
	204b  TravelDrive
08ed  MediaTek Inc.
	0002  CECT M800 memory card
08ee  CCSI/Hesso
08f0  Corex Technologies
08f1  CTI Electronics Corp.
08f2  Gotop Information Inc.
	007f  Super Q2 Tablet
08f5  SysTec Co., Ltd
08f6  Logic 3 International, Ltd
08f7  Vernier
	0001  LabPro
	0002  EasyTemp/Go!Temp
	0003  Go!Link
	0004  Go!Motion
08f8  Keen Top International Enterprise Co., Ltd
08f9  Wipro Technologies
08fa  Caere
08fb  Socket Communications
08fc  Sicon Cable Technology Co., Ltd
08fd  Digianswer A/S
	0001  Bluetooth Device
08ff  AuthenTec, Inc.
	1600  AES1600
	1610  AES1600
	1660  AES1660 Fingerprint Sensor
	1680  AES1660 Fingerprint Sensor
	168f  AES1660 Fingerprint Sensor
	2500  AES2501
	2501  AES2501
	2502  AES2501
	2503  AES2501
	2504  AES2501
	2505  AES2501
	2506  AES2501
	2507  AES2501
	2508  AES2501
	2509  AES2501
	250a  AES2501
	250b  AES2501
	250c  AES2501
	250d  AES2501
	250e  AES2501
	250f  AES2501
	2510  AES2510
	2550  AES2550 Fingerprint Sensor
	2580  AES2501 Fingerprint Sensor
	2588  AES2501
	2589  AES2501
	258a  AES2501
	258b  AES2501
	258c  AES2501
	258d  AES2501
	258e  AES2501
	258f  AES2501
	2660  AES2660 Fingerprint Sensor
	2680  AES2660 Fingerprint Sensor
	268f  AES2660 Fingerprint Sensor
	2810  AES2810
	3400  AES3400 TruePrint Sensor
	3401  AES3400 Sensor
	3402  AES3400 Sensor
	3403  AES3400 Sensor
	3404  AES3400 TruePrint Sensor
	3405  AES3400 TruePrint Sensor
	3406  AES3400 TruePrint Sensor
	3407  AES3400 TruePrint Sensor
	4902  BioMV with TruePrint AES3500
	4903  BioMV with TruePrint AES3400
	5500  AES4000
	5501  AES4000 TruePrint Sensor
	5503  AES4000 TruePrint Sensor
	5505  AES4000 TruePrint Sensor
	5507  AES4000 TruePrint Sensor
	55ff  AES4000 TruePrint Sensor.
	5700  AES3500 Fingerprint Reader
	5701  AES3500 TruePrint Sensor
	5702  AES3500 TruePrint Sensor
	5703  AES3500 TruePrint Sensor
	5704  AES3500-BZ TruePrint Sensor
	5705  AES3500-BZ TruePrint Sensor
	5706  AES3500-BZ TruePrint Sensor
	5707  AES3500-BZ TruePrint Sensor
	5710  AES3500 TruePrint Sensor
	5711  AES3500 TruePrint Sensor
	5712  AES3500 TruePrint Sensor
	5713  AES3500 TruePrint Sensor
	5714  AES3500-BZ TruePrint Sensor
	5715  AES3500-BZ TruePrint Sensor
	5716  AES3500-BZ TruePrint Sensor
	5717  AES3500-BZ TruePrint Sensor
	5730  AES3500 TruePrint Sensor
	5731  AES3500 TruePrint Sensor
	5732  AES3500 TruePrint Sensor
	5733  AES3500 TruePrint Sensor
	5734  AES3500-BZ TruePrint Sensor
	5735  AES3500-BZ TruePrint Sensor
	5736  AES3500-BZ TruePrint Sensor
	5737  AES3500-BZ TruePrint Sensor
	afe3  FingerLoc Sensor Module (Anchor)
	afe4  FingerLoc Sensor Module (Anchor)
	afe5  FingerLoc Sensor Module (Anchor)
	afe6  FingerLoc Sensor Module (Anchor)
	fffd  AES2510 Sensor (USB Emulator)
	ffff  Sensor (Emulator)
0900  Pinnacle Systems, Inc.
0901  VST Technologies
	0001  Hard Drive Adapter (TPP)
	0002  SigmaDrive Adapter (TPP)
0906  Faraday Technology Corp.
0908  Siemens AG
	01f4  SIMATIC NET CP 5711
	01fe  SIMATIC NET PC Adapter A2
	04b1  MediSET
	04b2  NC interface
	2701  ShenZhen SANZHAI Technology Co.,Ltd Spy Pen VGA
0909  Audio-Technica Corp.
090a  Trumpion Microelectronics, Inc.
	1001  T33520 Flash Card Controller
	1100  Comotron C3310 MP3 player
	1200  MP3 player
	1540  Digitex Container Flash Disk
090b  Neurosmith
090c  Silicon Motion, Inc. - Taiwan (formerly Feiya Technology Corp.)
	0371  Silicon Motion SM371 Camera
	0373  Silicon Motion Camera
	037a  Silicon Motion Camera
	037b  Silicon Motion Camera
	037c  300k Pixel Camera
	1000  Flash Drive
	1132  5-in-1 Card Reader
	337b  Silicon Motion Camera
	3710  Silicon Motion Camera
	3720  Silicon Motion Camera
	37bc  HP Webcam-101 Integrated Camera
	37c0  Silicon Motion Camera
	6000  SD/SDHC Card Reader (SG365 / FlexiDrive XC+)
	6200  microSD card reader
	71b3  SM731 Camera
	837b  Silicon Motion Camera
	937b  Silicon Motion Camera
	b370  Silicon Motion SM370 Camera
	b371  Silicon Motion SM371 Camera
	f37d  Endoscope camera
090d  Multiport Computer Vertriebs GmbH
090e  Shining Technology, Inc.
090f  Fujitsu Devices, Inc.
0910  Alation Systems, Inc.
0911  Philips Speech Processing
	149a  SpeechMike II Pro Plus LFH5276
	2512  SpeechMike Pro
0912  Voquette, Inc.
0915  GlobeSpan, Inc.
	0001  DSL Modem
	0002  ADSL ATM Modem
	0005  LAN Modem
	2000  802.11 Adapter
	2002  802.11 Adapter
	8000  ADSL LAN Modem
	8005  DSL-302G Modem
	8101  ADSL WAN Modem
	8102  DSL-200 ADSL Modem
	8103  DSL-200 ADSL Modem
	8104  DSL-200 Modem
	8400  DSL Modem
	8401  DSL Modem
	8402  DSL Modem
	8500  DSL Modem
	8501  DSL Modem
0917  SmartDisk Corp.
	0001  eFilm Reader-11 SM/CF
	0002  eFilm Reader-11 SM
	0003  eFilm Reader-11 CF
	0200  FireFly
	0201  FireLite
	0202  STORAGE ADAPTER (FirePower)
	0204  FlashTrax Storage
	0205  STORAGE ADAPTER (CrossFire)
	0206  FireFly 20G HDD
	0207  FireLite
	020f  STORAGE ADAPTER (FireLite)
	da01  eFilm Reader-11 Test
	ffff  eFilm Reader-11 (Class/PDR)
0919  Tiger Electronics
	0100  Fast Flicks Digital Camera
091e  Garmin International
	0003  GPS (various models)
	0004  iQue 3600
	0200  Data Card Programmer (install)
	1200  Data Card Programmer
	21a5  etrex Cx (msc)
	2236  nuvi 360
	2271  Edge 605/705
	2295  Colorado 300
	22b6  eTrex Vista HCx (Mass Storage mode)
	231b  Oregon 400t
	2353  Nvi 205T
	2380  Oregon series
	23cc  nvi 1350
	2459  GPSmap 62/78 series
	2491  Edge 800
	2519  eTrex 30
	2535  Edge 800
	253c  GPSmap 62sc
	255b  Nuvi 2505LM
0920  Echelon Co.
	7500  Network Interface
0921  GoHubs, Inc.
	1001  GoCOM232 Serial
0922  Dymo-CoStar Corp.
	0007  LabelWriter 330
	0009  LabelWriter 310
	0019  LabelWriter 400
	001a  LabelWriter 400 Turbo
	0020  LabelWriter 450
	1001  LabelManager PnP
	8004  M25 Digital Postal Scale
0923  IC Media Corp.
	010f  SIIG MobileCam
0924  Xerox
	23dd  DocuPrint M760 (X760_USB)
	3ce8  Phaser 3428 Printer
	3d5b  Phaser 6115MFP TWAIN Scanner
	3d6d  WorkCentre 6015N/NI
	420f  WorkCentre PE220 Series
	421f  M20 Scanner
	423b  Printing Support
	4274  Xerox Phaser 3635MFPX
	ffef  WorkCenter M15
	fffb  DocuPrint M750 (X750_USB)
0925  Lakeview Research
	0005  Gamtec.,Ltd SmartJoy PLUS Adapter
	03e8  Wii Classic Controller Adapter
	3881  Saleae Logic
	8101  Phidgets, Inc., 1-Motor PhidgetServo v2.0
	8104  Phidgets, Inc., 4-Motor PhidgetServo v2.0
	8800  WiseGroup Ltd, MP-8800 Quad Joypad
	8866  WiseGroup Ltd, MP-8866 Dual Joypad
0927  Summus, Ltd
0928  PLX Technology, Inc. (formerly Oxford Semiconductor, Ltd)
	8000  Firmware uploader
	ffff  Blank Oxford Device
0929  American Biometric Co.
092a  Toshiba Information & Industrial Sys. And Services
092b  Sena Technologies, Inc.
092f  Northern Embedded Science/CAVNEX
	0004  JTAG-4
	0005  JTAG-5
0930  Toshiba Corp.
	0009  Gigabeat F/X (HDD audio player)
	000c  Gigabeat F (mtp)
	0010  Gigabeat S (mtp)
	0200  Integrated Bluetooth (Taiyo Yuden)
	021c  Atheros AR3012 Bluetooth
	0301  PCX1100U Cable Modem (WDM)
	0302  PCX2000 Cable Modem (WDM)
	0305  Cable Modem PCX3000
	0307  Cable Modem PCX2500
	0308  PCX2200 Cable Modem (WDM)
	0309  PCX5000 Cable Modem (WDM)
	030b  Cable Modem PCX2600
	0501  Bluetooth Controller
	0502  Integrated Bluetooth
	0503  Bluetooth Controller
	0505  Integrated Bluetooth
	0506  Integrated Bluetooth
	0507  Bluetooth Adapter
	0508  Integrated Bluetooth HCI
	0509  BT EDR Dongle
	0706  PocketPC e740
	0707  Pocket PC e330 Series
	0708  Pocket PC e350Series
	0709  Pocket PC e750 Series
	070a  Pocket PC e400 Series
	070b  Pocket PC e800 Series
	0a07  WLM-10U1 802.11abgn Wireless Adapter [Ralink RT3572]
	0a13  AX88179 Gigabit Ethernet [Toshiba]
	0b05  PX1220E-1G25 External hard drive
	0b09  PX1396E-3T01 External hard drive
	0b1a  STOR.E ALU 2S
	1300  Wireless Broadband (CDMA EV-DO) SM-Bus Minicard Status Port
	1301  Wireless Broadband (CDMA EV-DO) Minicard Status Port
	1302  Wireless Broadband (3G HSDPA) SM-Bus Minicard Status Port
	1303  Wireless Broadband (3G HSDPA) Minicard Status Port
	1308  Broadband (3G HSDPA) SM-Bus Minicard Diagnostics Port
	130b  F3507g Mobile Broadband Module
	130c  F3607gw Mobile Broadband Module
	1311  F3607gw v2 Mobile Broadband Module
	1400  Memory Stick 2GB
	642f  TravelDrive
	6506  TravelDrive 2C
	6507  TravelDrive 2C
	6508  TravelDrive 2C
	6509  TravelDrive 2C
	6510  TravelDrive 2C
	6517  TravelDrive 2C
	6518  TravelDrive 2C
	6519  Kingston DataTraveler 2.0 USB Stick
	651a  TravelDrive 2C
	651b  TravelDrive 2C
	651c  TravelDrive 2C
	651d  TravelDrive 2C
	651e  TravelDrive 2C
	651f  TravelDrive 2C
	6520  TravelDrive 2C
	6521  TravelDrive 2C
	6522  TravelDrive 2C
	6523  TravelDrive
	6524  TravelDrive
	6525  TravelDrive
	6526  TravelDrive
	6527  TravelDrive
	6528  TravelDrive
	6529  TravelDrive
	652a  TravelDrive
	652b  TravelDrive
	652c  TravelDrive
	652d  TravelDrive
	652f  TravelDrive
	6530  TravelDrive
	6531  TravelDrive
	6532  256M Stick
	6533  512M Stick
	6534  TravelDrive
	653c  Kingston DataTraveler 2.0 Stick (512M)
	653d  Kingston DataTraveler 2.0 Stick (1GB)
	653e  Flash Memory
	6540  TransMemory Flash Memory
	6544  TransMemory-Mini / Kingston DataTraveler 2.0 Stick (2GB)
	6545  Kingston DataTraveler 102/2.0 / HEMA Flash Drive 2 GB / PNY Attache 4GB Stick
0931  Harmonic Data Systems, Ltd
0932  Crescentec Corp.
	0300  VideoAdvantage
	0302  Syntek DC-112X
	0320  VideoAdvantage
	0482  USB2.0 TVBOX
	1100  DC-1100 Video Enhamcement Device
	1112  Veo Web Camera
	a311  Video Enhancement Device
0933  Quantum Corp.
0934  Spirent Communications
0936  NuTesla
	000c  Rhythmedics 6 BioData Integrator
	0030  Composite Device, Mass Storage Device (Flash Drive) amd HID
	003c  Rhythmedics HID Bootloader
0939  Lumberg, Inc.
	0b15  Toshiba Stor.E Alu 2
093a  Pixart Imaging, Inc.
	0007  CMOS 100K-R Rev. 1.90
	010e  Digital camera, CD302N/Elta Medi@ digi-cam/HE-501A
	010f  Argus DC-1610/DC-1620/Emprex PCD3600/Philips P44417B keychain camera/Precision Mini,Model HA513A/Vivitar Vivicam 55
	020f  Bullet Line Photo Viewer
	050f  Mars-Semi Pc-Camera
	2460  Q-TEC WEBCAM 100
	2468  SoC PC-Camera
	2470  SoC PC-Camera
	2471  SoC PC-Camera
	2500  USB Optical Mouse
	2510  Optical Mouse
	2521  Optical Mouse
	2600  Typhoon Easycam USB 330K (newer)/Typhoon Easycam USB 2.0 VGA 1.3M/Sansun SN-508
	2601  SPC 610NC Laptop Camera
	2603  PAC7312 Camera
	2608  PAC7311 Trust WB-3300p
	260e  PAC7311 Gigaware VGA PC Camera:Trust WB-3350p:SIGMA cam 2350
	260f  PAC7311 SnakeCam
	2621  PAC731x Trust Webcam
	2622  Webcam Genius
	2624  Webcam
093b  Plextor Corp.
	0010  Storage Adapter
	0011  PlexWriter 40/12/40U
	0041  PX-708A DVD RW
	0042  PX-712UF DVD RW
	a002  ConvertX M402U XLOADER
	a003  ConvertX AV100U A/V Capture Audio
	a004  ConvertX TV402U XLOADER
	a005  ConvertX TV100U A/V Capture
	a102  ConvertX M402U A/V Capture
	a104  ConvertX PX-TV402U/NA
093c  Intrepid Control Systems, Inc.
	0601  ValueCAN
	0701  NeoVI Blue vehicle bus interface
093d  InnoSync, Inc.
093e  J.S.T. Mfg. Co., Ltd
093f  Olympia Telecom Vertriebs GmbH
0940  Japan Storage Battery Co., Ltd
0941  Photobit Corp.
0942  i2Go.com, LLC
0943  HCL Technologies India Private, Ltd
0944  KORG, Inc.
	0001  PXR4 4-Track Digital Recorder
	0020  KAOSS Pad KP3 Dynamic Effect/Sampler
	0023  KAOSSILATOR PRO Dynamic Phrase Synthesizer
	010d  nanoKEY MIDI keyboard
	010e  nanoPAD pad controller
	010f  nanoKONTROL studio controller
	0117  nanoKONTROL2 MIDI Controller
	0f03  K-Series K61P MIDI studio controller
0945  Pasco Scientific
0948  Kronauer music in digital
	0301  USB Pro (24/48)
	0302  USB Pro (24/96 playback)
	0303  USB Pro (24/96 record)
	0304  USB Pro (16/48)
	1105  USB One
094b  Linkup Systems Corp.
	0001  neonode N2
094d  Cable Television Laboratories
094f  Yano
	0101  U640MO-03
	05fc  METALWEAR-HDD
0951  Kingston Technology
	0008  Ethernet
	000a  KNU101TX 100baseTX Ethernet
	1600  DataTraveler II Pen Drive
	1601  DataTraveler II+ Pen Drive
	1602  DataTraveler Mini
	1603  DataTraveler 1GB/2GB Pen Drive
	1606  Eee PC 701 SD Card Reader [ENE UB6225]
	1607  DataTraveler 100
	160d  DataTraveler Vault Privacy
	160e  DT110P/1GB Capless
	1613  DataTraveler DT101C Flash Drive
	1616  DataTraveler Locker 4GB
	161a  Dell HyperVisor internal flash drive
	1621  DataTraveler 150 (32GB)
	1624  DataTraveler G2
	1625  DataTraveler 101 II
	162a  DataTraveler 112 4GB Pen Drive
	162d  DataTraveler 102
	1630  DataTraveler 200 (32GB)
	1642  DT101 G2
	1643  DataTraveler G3
	1653  Data Traveler 100 G2 8 GiB
	1656  DataTraveler Ultimate G2
	1665  Digital DataTraveler SE9 64GB
	1666  DataTraveler G4
	1689  DataTraveler SE9
	168a  DataTraveler Micro
	168c  DT Elite 3.0
0954  RPM Systems Corp.
0955  NVidia Corp.
	7030  Tegra 3 (recovery mode)
	7100  Tegra Device
	7820  Tegra 2 AC100 developer mode
	b400  SHIELD (debug)
	b401  SHIELD
	cf05  SHIELD Tablet (debug)
	cf06  SHIELD Tablet
	cf07  SHIELD Tablet
	cf08  SHIELD Tablet
	cf09  SHIELD Tablet
0956  BSquare Corp.
0957  Agilent Technologies, Inc.
	0200  E-Video DC-350 Camera
	0202  E-Video DC-350 Camera
	0407  33220A Waveform Generator
	0518  82357B GPIB Interface
	0a07  34411A Multimeter
	1507  33210A Waveform Generator
	1745  Test and Measurement Device (IVI)
	2918  U2702A oscilloscope
	fb18  LC Device
0958  CompuLink Research, Inc.
0959  Cologne Chip AG
	2bd0  Intelligent ISDN (Ver. 3.60.04) [HFC-S]
095a  Portsmith
	3003  Express Ethernet
095b  Medialogic Corp.
095c  K-Tec Electronics
095d  Polycom, Inc.
	0001  Polycom ViaVideo
0967  Acer NeWeb Corp.
	0204  WarpLink 802.11b Adapter
0968  Catalyst Enterprises, Inc.
096e  Feitian Technologies, Inc.
	0120  Microcosm Ltd Dinkey
	0802  ePass2000 (G&D STARCOS SPK 2.4)
	0807  ePass2003
0971  Gretag-Macbeth AG
	2000  i1 Pro
	2001  i1 Monitor
	2003  Eye-One display
	2005  Huey
	2007  ColorMunki Photo
0973  Schlumberger
	0001  e-gate Smart Card
0974  Datagraphix, a business unit of Anacomp
0975  OL'E Communications, Inc.
0976  Adirondack Wire & Cable
0977  Lightsurf Technologies
0978  Beckhoff GmbH
0979  Jeilin Technology Corp., Ltd
	0222  Keychain Display
	0224  JL2005A Toy Camera
	0226  JL2005A Toy Camera
	0227  JL2005B/C/D Toy Camera
097a  Minds At Work LLC
	0001  Digital Wallet
097b  Knudsen Engineering, Ltd
097c  Marunix Co., Ltd
097d  Rosun Technologies, Inc.
097e  Biopac Systems Inc.
	0035  MP35 v1.0
097f  Barun Electronics Co., Ltd
0981  Oak Technology, Ltd
0984  Apricorn
	0040  SATA Wire (2.5")
	0200  Hard Drive Storage (TPP)
0985  cab Produkttechnik GmbH & Co KG
	0045  Mach4/200 Label Printer
	00a3  A3/200 or A3/300 Label Printer
0986  Matsushita Electric Works, Ltd.
098c  Vitana Corp.
098d  INDesign
098e  Integrated Intellectual Property, Inc.
098f  Kenwood TMI Corp.
0993  Gemstar eBook Group, Ltd
	0001  REB1100 eBook Reader
	0002  eBook
0996  Integrated Telecom Express, Inc.
099a  Zippy Technology Corp.
	0638  Sanwa Supply Inc. Small Keyboard
	610c  EL-610 Super Mini Electron luminescent Keyboard
	713a  WK-713 Multimedia Keyboard
	7160  Hyper Slim Keyboard
09a3  PairGain Technologies
09a4  Contech Research, Inc.
09a5  VCON Telecommunications
09a6  Poinchips
	8001  Mass Storage Device
09a7  Data Transmission Network Corp.
09a8  Lin Shiung Enterprise Co., Ltd
09a9  Smart Card Technologies Co., Ltd
09aa  Intersil Corp.
	1000  Prism GT 802.11b/g Adapter
	3642  Prism 2.x 802.11b Adapter
09ab  Japan Cash Machine Co., Ltd.
09ae  Tripp Lite
09b2  Franklin Electronic Publishers, Inc.
	0001  eBookman Palm Computer
09b3  Altius Solutions, Inc.
09b4  MDS Telephone Systems
09b5  Celltrix Technology Co., Ltd
09bc  Grundig
	0002  MPaxx MP150 MP3 Player
09be  MySmart.Com
	0001  MySmartPad
09bf  Auerswald GmbH & Co. KG
	00c0  COMpact 2104 ISDN PBX
	00db  COMpact 4410/2206 ISDN
	00dc  COMpact 4406 DSL (PBX)
	00dd  COMpact 2204 (PBX)
	00de  COMpact 2104 (Rev.2 PBX)
	00e0  COMmander Business (PBX)
	00e2  COMmander Basic.2 (PBX)
	00f1  COMfort 2000 (System telephone)
	00f2  COMfort 1200 (System telephone)
	00f5  COMfortel 2500 (System telephone)
	8000  COMpact 2104 DSL (DSL modem)
	8001  COMpact 4406 DSL (DSL modem)
	8002  Analog/ISDN Converter (Line converter)
	8005  WG-640 (Automatic event dialer)
09c0  Genpix Electronics, LLC
	0136  Axon CNS, MultiClamp 700B
	0202  8PSK DVB-S tuner
	0203  Skywalker-1 DVB-S tuner
	0204  Skywalker-CW3K DVB-S tuner
	0205  Skywalker-CW3K DVB-S tuner
	0206  Skywalker-2 DVB-S tuner
09c1  Arris Interactive LLC
	1337  TOUCHSTONE DEVICE
09c2  Nisca Corp.
09c3  ActivCard, Inc.
	0007  Reader V2
	0008  ZFG-9800-AC SmartCard Reader
	0014  ActivIdentity ActivKey SIM USB Token
09c4  ACTiSYS Corp.
	0011  ACT-IR2000U IrDA Dongle
09c5  Memory Corp.
09ca  BMC Messsysteme GmbH
	5544  PIO
09cc  Workbit Corp.
	0404  BAFO USB-ATA/ATAPI Bridge Controller
09cd  Psion Dacom Home Networks, Ltd
	2001  Psion WaveFinder DAB radio receiver
09ce  City Electronics, Ltd
09cf  Electronics Testing Center, Taiwan
09d1  NeoMagic, Inc.
09d2  Vreelin Engineering, Inc.
09d3  Com One
	0001  ISDN TA / Light Rider 128K
	000b  Bluetooth Adapter class 1 [BlueLight]
09d7  Novatel Wireless
	0100  NovAtel FlexPack GPS receiver
09d9  KRF Tech, Ltd
09da  A4Tech Co., Ltd.
	0006  Optical Mouse WOP-35 / Trust 450L Optical Mouse
	000a  Optical Mouse Opto 510D / OP-620D
	000e  X-F710F Optical Mouse 3xFire Gaming Mouse
	0018  Trust Human Interface Device
	001a  Wireless Mouse & RXM-15 Receiver
	002a  Wireless Optical Mouse NB-30
	022b  Wireless Mouse (Battery Free)
	024f  RF Receiver and G6-20D Wireless Optical Mouse
	0260  KV-300H Isolation Keyboard
	032b  Wireless Mouse (Battery Free)
	8090  X-718BK Oscar Optical Gaming Mouse
	9033  X-718BK Optical Mouse
	9066  F3 V-Track Gaming Mouse
	9090  XL-730K / XL-750BK / XL-755BK Mice
09db  Measurement Computing Corp.
	0075  MiniLab 1008
	0076  PMD-1024
	007a  PMD-1208LS
	0081  USB-1616FS
	0082  USB-1208FS
	0088  USB-1616FS internal hub
09dc  Aimex Corp.
09dd  Fellowes, Inc.
09df  Addonics Technologies Corp.
09e1  Intellon Corp.
	5121  MicroLink dLAN
09e5  Jo-Dan International, Inc.
09e6  Silutia, Inc.
09e7  Real 3D, Inc.
09e8  AKAI  Professional M.I. Corp.
	0062  MPD16 MIDI Pad Controller Unit
	006d  EWI electronic wind instrument
	0071  MPK25 MIDI Keyboard
	0076  LPK25 MIDI Keyboard
09e9  Chen-Source, Inc.
09eb  IM Networks, Inc.
	4331  iRhythm Tuner Remote
09ef  Xitel
	0101  MD-Port DG2 MiniDisc Interface
09f3  GoFlight, Inc.
	0018  GF-46 Multi-Mode Display Module
	0028  RP-48 Combination Pushbutton-Rotary Module
	0048  LGTII - Landing Gear and Trim Control Module
	0064  MCPPro - Airliner Mode Control Panel (Autopilot)
	0300  EFIS - Electronic Flight Information System
09f5  AresCom
	0168  Network Adapter
	0188  LAN Adapter
	0850  Adapter
09f6  RocketChips, Inc.
09f7  Edu-Science (H.K.), Ltd
09f8  SoftConnex Technologies, Inc.
09f9  Bay Associates
09fa  Mtek Vision
09fb  Altera
	6001  Blaster
09ff  Gain Technology Corp.
0a00  Liquid Audio
0a01  ViA, Inc.
0a05  Unknown Manufacturer
	0001  Hub
	7211  hub
0a07  Ontrak Control Systems Inc.
	0064  ADU100 Data Acquisition Interface
	0078  ADU120 Data Acquisition Interface
	0082  ADU130 Data Acquisition Interface
	00c8  ADU200 Relay I/O Interface
	00d0  ADU208 Relay I/O Interface
	00da  ADU218 Solid-State Relay I/O Interface
0a0b  Cybex Computer Products Co.
0a0d  Servergy, Inc
	2514  CTS-1000 Internal Hub
0a11  Xentec, Inc.
0a12  Cambridge Silicon Radio, Ltd
	0001  Bluetooth Dongle (HCI mode)
	0002  Frontline Test Equipment Bluetooth Device
	0003  Nanosira
	0004  Nanosira WHQL Reference Radio
	0005  Nanosira-Multimedia
	0006  Nanosira-Multimedia WHQL Reference Radio
	0007  Nanosira3-ROM
	0008  Nanosira3-ROM
	0009  Nanosira4-EDR WHQL Reference Radio
	000a  Nanosira4-EDR-ROM
	000b  Nanosira5-ROM
	0042  SPI Converter
	0043  Bluetooth Device
	0100  Casira with BlueCore2-External Module
	0101  Casira with BlueCore2-Flash Module
	0102  Casira with BlueCore3-Multimedia Module
	0103  Casira with BlueCore3-Flash Module
	0104  Casira with BlueCore4-External Module
	0105  Casira with BlueCore4-Multimedia Module
	1000  Bluetooth Dongle (HID proxy mode)
	1010  Bluetooth Device
	1011  Bluetooth Device
	1012  Bluetooth Device
	ffff  USB Bluetooth Device in DFU State
0a13  Telebyte, Inc.
0a14  Spacelabs Medical, Inc.
0a15  Scalar Corp.
0a16  Trek Technology (S) PTE, Ltd
	1111  ThumbDrive
	8888  IBM USB Memory Key
	9988  Trek2000 TD-G2
0a17  Pentax Corp.
	0004  Optio 330
	0006  Optio S / S4
	0007  Optio 550
	0009  Optio 33WR
	000a  Optio 555
	000c  Optio 43WR (mass storage mode)
	000d  Optio 43WR
	0015  Optio S40/S5i
	003b  Optio 50 (mass storage mode)
	003d  Optio S55
	0041  Optio S5z
	0043  *ist DL
	0047  Optio S60
	0052  Optio 60 Digital Camera
	006e  K10D
	0070  K100D
	0093  K200D
	00a7  Optio E50
	1001  EI2000 Camera powered by Digita!
0a18  Heidelberger Druckmaschinen AG
0a19  Hua Geng Technologies, Inc.
0a21  Medtronic Physio Control Corp.
	8001  MMT-7305WW [Medtronic Minimed CareLink]
0a22  Century Semiconductor USA, Inc.
0a27  Datacard Group
	0102  SP35
0a2c  AK-Modul-Bus Computer GmbH
	0008  GPIO Ports
0a34  TG3 Electronics, Inc.
	0101  TG82tp
	0110  Deck 82-key backlit keyboard
0a35  Radikal Technologies
	002a  SAC - Software Assigned Controller
	008a  SAC Hub
0a39  Gilat Satellite Networks, Ltd
0a3a  PentaMedia Co., Ltd
	0163  KN-W510U 1.0 Wireless LAN Adapter
0a3c  NTT DoCoMo, Inc.
0a3d  Varo Vision
0a3f  Swissonic AG
0a43  Boca Systems, Inc.
0a46  Davicom Semiconductor, Inc.
	0268  ST268
	6688  ZT6688 Fast Ethernet Adapter
	8515  ADMtek ADM8515 NIC
	9000  DM9000E Fast Ethernet Adapter
	9601  DM9601 Fast Ethernet Adapter
0a47  Hirose Electric
0a48  I/O Interconnect
	3233  Multimedia Card Reader
	3239  Multimedia Card Reader
	3258  Dane Elec zMate SD Reader
	3259  Dane Elec zMate CF Reader
	5000  MediaGear xD-SM
	500a  Mass Storage Device
	500f  Mass Storage Device
	5010  Mass Storage Device
	5011  Mass Storage Device
	5014  Mass Storage Device
	5020  Mass Storage Device
	5021  Mass Storage Device
	5022  Mass Storage Device
	5023  Mass Storage Device
	5024  Mass Storage Device
	5025  Mass Storage Device
0a4a  Ploytec GmbH
0a4b  Fujitsu Media Devices, Ltd
0a4c  Computex Co., Ltd
	15d9  OPTICAL MOUSE
0a4d  Evolution Electronics, Ltd
	0064  MK-225 Driver
	0065  MK-225C Driver
	0066  MK-225C Driver
	0067  MK-425C Driver
	0078  MK-37 Driver
	0079  MK-37C Driver
	007a  MK-37C Driver
	008c  TerraTec MIDI MASTER
	008d  MK-249C Driver
	008e  MK-249C MIDI Keyboard
	008f  MK-449C Driver
	0090  Keystation 49e Driver
	0091  Keystation 61es Driver
	00a0  MK-361 Driver
	00a1  MK-361C Driver
	00a2  MK-361C Driver
	00a3  MK-461C MIDI Keyboard
	00b5  Keystation Pro 88 Driver
	00d2  E-Keys Driver
	00f0  UC-16 Driver
	00f1  X-Session Driver
	00f5  UC-33e MIDI Controller
0a4e  Steinberg Soft-und Hardware GmbH
0a4f  Litton Systems, Inc.
0a50  Mimaki Engineering Co., Ltd
0a51  Sony Electronics, Inc.
0a52  Jebsee Electronics Co., Ltd
0a53  Portable Peripheral Co., Ltd
	1000  Scanner
	2000  Q-Scan A6 Scanner
	2001  Q-Scan A6 Scanner
	2013  Media Drive A6 Scanner
	2014  Media Drive A6 Scanner
	2015  BizCardReader 600C
	2016  BizCardReader 600C
	202a  Scanshell-CSSN
	3000  Q-Scan A8 Scanner
	3002  Q-Scan A8 Reader
	3015  BizCardReader 300G
	302a  LM9832 - PA570 Mini Business Card Scanner [Targus]
	5001  BizCardReader 900C
0a5a  Electronics For Imaging, Inc.
0a5b  EAsics NV
0a5c  Broadcom Corp.
	0201  iLine10(tm) Network Adapter
	2000  Bluetooth Device
	2001  Bluetooth Device
	2009  BCM2035 Bluetooth
	200a  BCM2035 Bluetooth dongle
	200f  Bluetooth Controller
	201d  Bluetooth Device
	201e  IBM Integrated Bluetooth IV
	2020  Bluetooth dongle
	2021  BCM2035B3 Bluetooth Adapter
	2033  BCM2033 Bluetooth
	2035  BCM2035 Bluetooth
	2038  Blutonium Device
	2039  BCM2045 Bluetooth
	2045  Bluetooth Controller
	2046  Bluetooth Device
	2047  Bluetooth Device
	205e  Bluetooth Device
	2100  Bluetooth 2.0+eDR dongle
	2101  BCM2045 Bluetooth
	2102  ANYCOM Blue USB-200/250
	2110  BCM2045B (BDC-2) [Bluetooth Controller]
	2111  ANYCOM Blue USB-UHE 200/250
	2120  2045 Bluetooth 2.0 USB-UHE Device with trace filter
	2121  BCM2210 Bluetooth
	2122  Bluetooth 2.0+EDR dongle
	2123  Bluetooth dongle
	2130  2045 Bluetooth 2.0 USB-UHE Device with trace filter
	2131  2045 Bluetooth 2.0 Device with trace filter
	2145  BCM2045B (BDC-2.1) [Bluetooth Controller]
	2148  BCM92046DG-CL1ROM Bluetooth 2.1 Adapter
	2150  BCM2046 Bluetooth Device
	2151  Bluetooth
	2154  BCM92046DG-CL1ROM Bluetooth 2.1 UHE Dongle
	216c  BCM43142A0 Bluetooth Device
	216f  BCM20702A0 Bluetooth
	217d  HP Bluethunder
	217f  BCM2045B (BDC-2.1)
	2198  Bluetooth 3.0 Device
	219b  Bluetooth 2.1 Device
	21b1  HP Bluetooth Module
	21b4  BCM2070 Bluetooth 2.1 + EDR
	21b9  BCM2070 Bluetooth 2.1 + EDR
	21ba  BCM2070 Bluetooth 2.1 + EDR
	21bb  BCM2070 Bluetooth 2.1 + EDR
	21bc  BCM2070 Bluetooth 2.1 + EDR
	21bd  BCM2070 Bluetooth 2.1 + EDR
	21d7  BCM43142 Bluetooth 4.0
	21e1  HP Portable SoftSailing
	21e3  HP Portable Valentine
	21e6  BCM20702 Bluetooth 4.0 [ThinkPad]
	21e8  BCM20702A0 Bluetooth 4.0
	21f1  HP Portable Bumble Bee
	22be  BCM2070 Bluetooth 3.0 + HS
	4500  BCM2046B1 USB 2.0 Hub (part of BCM2046 Bluetooth)
	4502  Keyboard (Boot Interface Subclass)
	4503  Mouse (Boot Interface Subclass)
	5800  BCM5880 Secure Applications Processor
	5801  BCM5880 Secure Applications Processor with fingerprint swipe sensor
	5802  BCM5880 Secure Applications Processor with fingerprint touch sensor
	5803  BCM5880 Secure Applications Processor with secure keyboard
	5804  BCM5880 Secure Applications Processor with fingerprint swipe sensor
	6300  Pirelli Remote NDIS Device
	bd11  TiVo AG0100 802.11bg Wireless Adapter [Broadcom BCM4320]
	bd13  BCM4323 802.11abgn Wireless Adapter
	bd16  BCM4319 802.11bgn Wireless Adapter
	bd17  BCM43236 802.11abgn Wireless Adapter
	d11b  Eminent EM4045 [Broadcom 4320 USB]
0a5d  Diatrend Corp.
0a5f  Zebra
	0009  LP2844 Printer
	0081  GK420t Label Printer
	008b  HC100 wristbands Printer
	008c  ZP 450 Printer
	00d1  Zebra GC420d Label Printer
	930a  Printer
0a62  MPMan
	0010  MPMan MP-F40 MP3 Player
0a66  ClearCube Technology
0a67  Medeli Electronics Co., Ltd
0a68  Comaide Corp.
0a69  Chroma ate, Inc.
0a6b  Green House Co., Ltd
	0001  Compact Flash R/W with MP3 player
	000f  FlashDisk
0a6c  Integrated Circuit Systems, Inc.
0a6d  UPS Manufacturing
0a6e  Benwin
0a6f  Core Technology, Inc.
	0400  Xanboo
0a70  International Game Technology
0a71  VIPColor Technologies USA, Inc.
	0001  VP485 Printer
0a72  Sanwa Denshi
0a73  Mackie Designs
	0002  XD-2 [Spike]
0a7d  NSTL, Inc.
0a7e  Octagon Systems Corp.
0a80  Rexon Technology Corp., Ltd
0a81  Chesen Electronics Corp.
	0101  Keyboard
	0103  Keyboard
	0203  Mouse
	0205  PS/2 Keyboard+Mouse Adapter
	0701  USB Missile Launcher
	ff01  Wireless Missile Launcher
0a82  Syscan
	4600  TravelScan 460/464
0a83  NextComm, Inc.
0a84  Maui Innovative Peripherals
0a85  Idexx Labs
0a86  NITGen Co., Ltd
0a89  Aktiv
	0001  Guardant Stealth/Net
	0002  Guardant ID
	0003  Guardant Stealth 2
	0004  Rutoken
	0005  Guardant Fidus
	0006  Guardant Stealth 3
	0007  Guardant Stealth 2
	0008  Guardant Stealth 3 Sign/Time
	0009  Guardant Code
	000a  Guardant Sign Pro
	000b  Guardant Sign Pro HID
	000c  Guardant Stealth 3 Sign/Time
	000d  Guardant Code HID
	000f  Guardant System Firmware Update
	0020  Rutoken S
	0025  Rutoken lite
	0026  Rutoken lite HID
	002a  Rutoken Mass Storage
	002b  Guardant Mass Storage
	0030  Rutoken ECP
	0040  Rutoken ECP HID
	0060  Rutoken Magistra
	0061  Rutoken Magistra
	0069  Reader
	0080  Rutoken PinPad Ex
	0081  Rutoken PinPad In
	0082  Rutoken PinPad 2
0a8d  Picturetel
0a8e  Japan Aviation Electronics Industry, Ltd
	2011  Filter Driver For JAE XMC R/W
0a90  Candy Technology Co., Ltd
0a91  Globlink Technology, Inc.
	3801  Targus PAKP003 Mouse
0a92  EGO SYStems, Inc.
	0011  SYS WaveTerminal U2A
	0021  GIGAPort
	0031  GIGAPortAG
	0053  AudioTrak Optoplay
	0061  Waveterminal U24
	0071  MAYA EX7
	0091  Maya 44
	00b1  MAYA EX5
	1000  MIDI Mate
	1010  RoMI/O
	1020  M4U
	1030  M8U
	1090  KeyControl49
	10a0  KeyControl25
0a93  C Technologies AB
	0002  C-Pen 10
	0005  MyPen Light
	000d  Input Pen
	0010  C-Pen 20
	0a93  PayPen
0a94  Intersense
0aa3  Lava Computer Mfg., Inc.
0aa4  Develco Elektronik
0aa5  First International Digital
	0002  irock! 500 Series
	0801  MP3 Player
0aa6  Perception Digital, Ltd
	0101  Hercules Jukebox
	1501  Store 'n' Go HD Drive
0aa7  Wincor Nixdorf International GmbH
	0100  POS Keyboard, TA58P-USB
	0101  POS Keyboard, TA85P-USB
	0102  POS Keyboard, TA59-USB
	0103  POS Keyboard, TA60-USB
	0104  SNIkey Keyboard, SNIKey-KB-USB
	0200  Operator Display, BA63-USB
	0201  Operator Display, BA66-USB
	0202  Operator Display & Scanner, XiCheck-BA63
	0203  Operator Display & Scanner, XiCheck-BA66
	0204  Graphics Operator Display, BA63GV
	0300  POS Printer (printer class mode), TH210
	0301  POS Printer (native mode), TH210
	0302  POS Printer (printer class mode), TH220
	0303  POS Printer (native mode), TH220
	0304  POS Printer, TH230
	0305  Lottery Printer, XiPrintPlus
	0306  POS Printer (printer class mode), TH320
	0307  POS Printer (native mode), TH320
	0308  POS Printer (printer class mode), TH420
	0309  POS Printer (native mode), TH420
	030a  POS Printer, TH200B
	0400  Lottery Scanner, Xiscan S
	0401  Lottery Scanner, Xiscan 3
	0402  Programmable Magnetic Swipe Card Reader, MSRP-USB
	0500  IDE Adapter
	0501  Hub Printer Interface
	0502  Hub SNIKey Keyboard
	4304  Banking Printer TP07
	4305  Banking Printer TP07c
	4500  WN Central Special Electronics
0aa8  TriGem Computer, Inc.
	0060  TG 11Mbps WLAN Mini Adapter
	1001  DreamComboM4100
	3002  InkJet Color Printer
	8001  TG_iMON
	8002  TG_KLOSS
	a001  TG_X2
	a002  TGVFD_KLOSS
	ffda  iMON_VFD
0aa9  Baromtec Co.
	f01b  Medion MD 6242 MP3 Player
0aaa  Japan CBM Corp.
0aab  Vision Shape Europe SA
0aac  iCompression, Inc.
0aad  Rohde & Schwarz GmbH & Co. KG
	0003  NRP-Z21
	000c  NRP-Z11
	0013  NRP-Z22
	0014  NRP-Z23
	0015  NRP-Z24
	0016  NRP-Z51
	0017  NRP-Z52
	0018  NRP-Z55
	0019  NRP-Z56
	0021  NRP-Z91
	0023  NRP-Z81
	002c  NRP-Z31
	002d  NRP-Z37
	002f  NRP-Z27
	0051  NRP-Z28
	0052  NRP-Z98
	0062  NRP-Z92
	0070  NRP-Z57
	0083  NRP-Z85
	0095  NRP-Z86
0aae  NEC infrontia Corp. (Nitsuko)
0aaf  Digitalway Co., Ltd
0ab0  Arrow Strong Electronics Co., Ltd
0ab1  FEIG ELECTRONIC GmbH
	0002  OBID RFID-Reader
	0004  OBID classic-pro
0aba  Ellisys
	8001  Tracker 110 Protocol Analyzer
	8002  Explorer 200 Protocol Analyzer
0abe  Stereo-Link
	0101  SL1200 DAC
0abf  Diolan
	3370  I2C/SPI Adapter - U2C-12
0ac3  Sanyo Semiconductor Company Micro
0ac4  Leco Corp.
0ac5  I & C Corp.
0ac6  Singing Electrons, Inc.
0ac7  Panwest Corp.
0ac8  Z-Star Microelectronics Corp.
	0301  Web Camera
	0302  ZC0302 Webcam
	0321  Vimicro generic vc0321 Camera
	0323  Luxya WC-1200 USB 2.0 Webcam
	0328  A4Tech PK-130MG
	0336  Elecom UCAM-DLQ30
	301b  ZC0301 Webcam
	303b  ZC0303 Webcam
	305b  ZC0305 Webcam
	307b  USB 1.1 Webcam
	332d  Vega USB 2.0 Camera
	3343  Sirius USB 2.0 Camera
	3370  Traveler TV 6500 SF Dia-scanner
	3420  Venus USB2.0 Camera
	c001  Sony embedded vimicro Camera
	c002  Visual Communication Camera VGP-VCC1
	c302  Vega USB 2.0 Camera
	c303  Saturn USB 2.0 Camera
	c326  Namuga 1.3M Webcam
	c33f  Webcam
	c429  Lenovo ThinkCentre Web Camera
	c42d  Lenovo IdeaCentre Web Camera
0ac9  Micro Solutions, Inc.
	0000  Backpack CD-ReWriter
	0001  BACKPACK  2 Cable
	0010  BACKPACK
	0011  Backpack 40GB Hard Drive
	0110  BACKPACK
	0111  BackPack
	1234  BACKPACK
0aca  OPEN Networks Ltd
	1060  OPEN NT1 Plus II
0acc  Koga Electronics Co.
0acd  ID Tech
	0300  IDT1221U RS-232 Adapter
	0401  Spectrum III Hybrid Smartcard Reader
	0630  Spectrum III Mag-Only Insert Reader (SPT3-355 Series) USB-CDC
	0810  SecurePIN (IDPA-506100Y) PIN Pad
	2030  ValueMag Magnetic Stripe Reader
0ace  ZyDAS
	1201  ZD1201 802.11b
	1211  ZD1211 802.11g
	1215  ZD1211B 802.11g
	1221  ZD1221 802.11n
	1602  ZyXEL Omni FaxModem 56K
	1608  ZyXEL Omni FaxModem 56K UNO
	1611  ZyXEL Omni FaxModem 56K Plus
	2011  Virtual media for 802.11bg
	20ff  Virtual media for 802.11bg
	a211  ZD1211 802.11b/g Wireless Adapter
	b215  802.11bg
0acf  Intoto, Inc.
0ad0  Intellix Corp.
0ad1  Remotec Technology, Ltd
0ad2  Service & Quality Technology Co., Ltd
0ada  Data Encryption Systems Ltd.
	0005  DK2
0ae3  Allion Computer, Inc.
0ae4  Taito Corp.
0ae7  Neodym Systems, Inc.
0ae8  System Support Co., Ltd
0ae9  North Shore Circuit Design L.L.P.
0aea  SciEssence, LLC
0aeb  TTP Communications, Ltd
0aec  Neodio Technologies Corp.
	2101  SmartMedia Card Reader
	2102  CompactFlash Card Reader
	2103  MMC/SD Card Reader
	2104  MemoryStick Card Reader
	2201  SmartMedia+CompactFlash Card Reader
	2202  SmartMedia+MMC/SD Card Reader
	2203  SmartMedia+MemoryStick Card Reader
	2204  CompactFlash+MMC/SD Card Reader
	2205  CompactFlash+MemoryStick Card Reader
	2206  MMC/SD+MemoryStick Card Reader
	2301  SmartMedia+CompactFlash+MMC/SD Card Reader
	2302  SmartMedia+CompactFlash+MemoryStick Card Reader
	2303  SmartMedia+MMC/SD+MemoryStick Card Reader
	2304  CompactFlash+MMC/SD+MemoryStick Card Reader
	3016  MMC/SD+Memory Stick Card Reader
	3050  ND3050 8-in-1 Card Reader
	3060  1.1 FS Card Reader
	3101  MMC/SD Card Reader
	3102  MemoryStick Card Reader
	3201  MMC/SD+MemoryStick Card Reader
	3216  HS Card Reader
	3260  7-in-1 Card Reader
	5010  ND5010 Card Reader
0af0  Option
	5000  UMTS Card
	6000  GlobeTrotter 3G datacard
	6300  GT 3G Quad UMTS/GPRS Card
	6600  GlobeTrotter 3G+ datacard
	6711  GlobeTrotter Express 7.2 v2
	6971  Globetrotter HSDPA Modem
	7251  Globetrotter HSUPA Modem (aka iCON HSUPA E)
	7501  Globetrotter HSUPA Modem (icon 411 aka "Vodafone K3760")
	7601  Globetrotter MO40x 3G Modem (GTM 382)
	7701  Globetrotter HSUPA Modem (aka icon 451)
	d055  Globetrotter GI0505 [iCON 505]
0af6  Silver I Co., Ltd
0af7  B2C2, Inc.
	0101  Digital TV USB Receiver (DVB-S/T/C / ATSC)
0af9  Hama, Inc.
	0010  USB SightCam 100
	0011  Micro Innovations IC50C Webcam
0afa  DMC Co., Ltd.
	07d2  Controller Board for Projected Capacitive Touch Screen DUS3000
0afc  Zaptronix Ltd
0afd  Tateno Dennou, Inc.
0afe  Cummins Engine Co.
0aff  Jump Zone Network Products, Inc.
0b00  INGENICO
0b05  ASUSTek Computer, Inc.
	0001  MeMO Pad HD 7 (CD-ROM mode)
	1101  Mass Storage (UISDMC4S)
	1706  WL-167G v1 802.11g Adapter [Ralink RT2571]
	1707  WL-167G v1 802.11g Adapter [Ralink RT2571]
	1708  Mass Storage Device
	170b  Multi card reader
	170c  WL-159g 802.11bg
	170d  802.11b/g Wireless Network Adapter
	1712  BT-183 Bluetooth 2.0+EDR adapter
	1715  2045 Bluetooth 2.0 Device with trace filter
	1716  Bluetooth Device
	1717  WL169gE 802.11g Adapter [Broadcom 4320 USB]
	171b  A9T wireless 802.11bg
	171c  802.11b/g Wireless Network Adapter
	171f  My Cinema U3000 Mini [DiBcom DiB7700P]
	1723  WL-167G v2 802.11g Adapter [Ralink RT2571W]
	1724  RT2573
	1726  Laptop OLED Display
	172a  ASUS 802.11n Network Adapter
	172b  802.11n Network Adapter
	1731  802.11n Network Adapter
	1732  802.11n Network Adapter
	1734  ASUS AF-200
	173c  BT-183 Bluetooth 2.0
	173f  My Cinema U3100 Mini
	1742  802.11n Network Adapter
	1743  Xonar U1 Audio Station
	1751  BT-253 Bluetooth Adapter
	175b  Laptop OLED Display
	1760  802.11n Network Adapter
	1761  USB-N11 802.11n Network Adapter [Ralink RT2870]
	1774  Gobi Wireless Modem (QDL mode)
	1776  Gobi Wireless Modem
	1779  My Cinema U3100 Mini Plus [AF9035A]
	1784  USB-N13 802.11n Network Adapter (rev. A1) [Ralink RT3072]
	1786  USB-N10 802.11n Network Adapter [Realtek RTL8188SU]
	1788  BT-270 Bluetooth Adapter
	1791  WL-167G v3 802.11n Adapter [Realtek RTL8188SU]
	179d  USB-N53 802.11abgn Network Adapter [Ralink RT3572]
	179e  Eee Note EA800 (network mode)
	179f  Eee Note EA800 (tablet mode)
	17a0  Xonar U3 sound card
	17a1  Eee Note EA800 (mass storage mode)
	17ab  USB-N13 802.11n Network Adapter (rev. B1) [Realtek RTL8192CU]
	17ba  N10 Nano 802.11n Network Adapter [Realtek RTL8192CU]
	17c7  WL-330NUL
	17c9  USB-AC53 802.11a/b/g/n/ac Wireless Adapter [Broadcom BCM43526]
	17cb  Broadcom BCM20702A0 Bluetooth
	17d1  AC51 802.11a/b/g/n/ac Wireless Adapter [Mediatek MT7610/Ralink RT2870]
	4c80  Transformer Pad TF300TG
	4c90  Transformer Pad Infinity TF700
	4c91  Transformer Pad Infinity TF700 (Debug mode)
	4ca0  Transformer Pad TF701T
	4ca1  Transformer Pad TF701T (Debug mode)
	4d00  Transformer Prime TF201
	4d01  Transformer Prime TF201 (debug mode)
	4daf  Transformer Pad Infinity TF700 (Fastboot)
	5410  MeMO Pad HD 7 (MTP mode)
	5412  MeMO Pad HD 7 (PTP mode)
	550f  ASUS fonepad 7
	6101  Cable Modem
	620a  Remote NDIS Device
	b700  Broadcom Bluetooth 2.1
0b0b  Datamax-O'Neil
	106e  Datamax E-4304
0b0c  Todos AB
	0009  Todos Argos Mini II Smart Card Reader
	001e  e.dentifier2 (ABN AMRO electronic banking card reader NL)
	002e  C200 smartcard controller (Nordea card reader)
	003f  Todos C400 smartcard controller (Handelsbanken card reader)
	0050  Argos Mini II Smart Card Reader (CCID)
0b0d  ProjectLab
	0000  CenturyCD
0b0e  GN Netcom
	034c  Jabra UC Voice 750 MS
	0420  Jabra SPEAK 510
	094d  GN Netcom / Jabra REVO Wireless
	1022  Jabra PRO 9450, Type 9400BS (DECT Headset)
	2007  GN 2000 Stereo Corded Headset
	620c  Jabra BT620s
	9330  Jabra GN9330 Headset
0b0f  AVID Technology
0b10  Pcally
0b11  I Tech Solutions Co., Ltd
0b1e  Electronic Warfare Assoc., Inc. (EWA)
	8007  Blackhawk USB560-BP JTAG Emulator
0b1f  Insyde Software Corp.
0b20  TransDimension, Inc.
0b21  Yokogawa Electric Corp.
0b22  Japan System Development Co., Ltd
0b23  Pan-Asia Electronics Co., Ltd
0b24  Link Evolution Corp.
0b27  Ritek Corp.
0b28  Kenwood Corp.
0b2c  Village Center, Inc.
0b30  PNY Technologies, Inc.
	0006  SM Media-Shuttle Card Reader
0b33  Contour Design, Inc.
	0020  ShuttleXpress
	0700  RollerMouse Pro
0b37  Hitachi ULSI Systems Co., Ltd
0b38  Gear Head
	0003  Keyboard
	0010  107-Key Keyboard
0b39  Omnidirectional Control Technology, Inc.
	0001  Composite USB PS2 Converter
	0109  USB TO Ethernet
	0421  Serial
	0801  USB-Parallel Bridge
	0901  OCT To Fast Ethernet Converter
	0c03  LAN DOCK Serial Converter
0b3a  IPaxess
0b3b  Tekram Technology Co., Ltd
	0163  TL-WN320G 1.0 WLAN Adapter
	1601  Allnet 0193 802.11b Adapter
	1602  ZyXEL ZyAIR B200 802.11b Adapter
	1612  AIR.Mate 2@net 802.11b Adapter
	1613  802.11b Wireless LAN Adapter
	1620  Allnet Wireless Network Adapter [Envara WiND512]
	1630  QuickWLAN 802.11bg
	5630  802.11bg
	6630  ZD1211
0b3c  Olivetti Techcenter
	a010  Simple_Way Printer/Scanner/Copier
	c000  Olicard 100
	c700  Olicard 100 (Mass Storage mode)
0b3e  Kikusui Electronics Corp.
0b41  Hal Corp.
	0011  Crossam2+USB IR commander
0b43  Play.com, Inc.
	0003  PS2 Controller Converter
	0005  GameCube Adaptor
0b47  Sportbug.com, Inc.
0b48  TechnoTrend AG
	1003  Technotrend/Hauppauge USB-Nova
	1004  TT-PCline
	1005  Technotrend/Hauppauge USB-Nova
	1006  Technotrend/Hauppauge DEC3000-s
	1007  TT-micro plus Device
	1008  Technotrend/Hauppauge DEC2000-t
	1009  Technotrend/Hauppauge DEC2540-t
	3001  DVB-S receiver
	3002  DVB-C receiver
	3003  DVB-T receiver
	3004  TT TV-Stick
	3005  TT TV-Stick (8kB EEPROM)
	3006  TT-connect S-2400 DVB-S receiver
	3007  TT-connect S2-3600
	3008  TT-connect
	3009  TT-connect S-2400 DVB-S receiver (8kB EEPROM)
	300a  TT-connect S2-3650 CI
	300b  TT-connect C-3650 CI
	300c  TT-connect T-3650 CI
	300d  TT-connect CT-3650 CI
	300e  TT-connect C-2400
	3011  TT-connect S2-4600
	3012  TT-connect CT2-4650 CI
	3014  TT-TVStick CT2-4400
0b49  ASCII Corp.
	064f  Trance Vibrator
0b4b  Pine Corp. Ltd.
	0100  D'music MP3 Player
0b4d  Graphtec America, Inc.
	110a  Graphtec CC200-20
0b4e  Musical Electronics, Ltd
	6500  MP3 Player
	8028  MP3 Player
	8920  MP3 Player
0b50  Dumpries Co., Ltd
0b51  Comfort Keyboard Co.
	0020  Comfort Keyboard
0b52  Colorado MicroDisplay, Inc.
0b54  Sinbon Electronics Co., Ltd
0b56  TYI Systems, Ltd
0b57  Beijing HanwangTechnology Co., Ltd
0b59  Lake Communications, Ltd
0b5a  Corel Corp.
0b5f  Green Electronics Co., Ltd
0b60  Nsine, Ltd
0b61  NEC Viewtechnology, Ltd
0b62  Orange Micro, Inc.
	000b  Bluetooth Device
	0059  iBOT2 Webcam
0b63  ADLink Technology, Inc.
0b64  Wonderful Wire Cable Co., Ltd
0b65  Expert Magnetics Corp.
0b66  Cybiko Inc.
	0041  Xtreme
0b67  Fairbanks Scales
	555e  SCB-R9000
0b69  CacheVision
0b6a  Maxim Integrated Products
	a132  WUP-005 [Nintendo Wii U Pro Controller]
0b6f  Nagano Japan Radio Co., Ltd
0b70  PortalPlayer, Inc.
	00ba  iRiver H10 20GB
0b71  SHIN-EI Sangyo Co., Ltd
0b72  Embedded Wireless Technology Co., Ltd
0b73  Computone Corp.
0b75  Roland DG Corp.
0b79  Sunrise Telecom, Inc.
0b7a  Zeevo, Inc.
	07d0  Bluetooth Dongle
0b7b  Taiko Denki Co., Ltd
0b7c  ITRAN Communications, Ltd
0b7d  Astrodesign, Inc.
0b81  id3 Technologies
	0001  Biothentic II smartcard reader with fingerprint sensor
	0002  DFU-Enabled Devices (DFU)
	0012  BioPAD biometric module (DFU + CDC)
	0102  Certis V1 fingerprint reader
	0103  Certis V2 fingerprint reader
	0200  CL1356T / CL1356T5 / CL1356A smartcard readers (CCID)
	0201  CL1356T / CL1356T5 / CL1356A smartcard readers (DFU + CCID)
	0220  CL1356A FFPJP smartcard reader (CCID + HID)
	0221  CL1356A smartcard reader (DFU + CCID + HID)
0b84  Rextron Technology, Inc.
0b85  Elkat Electronics, Sdn., Bhd.
0b86  Exputer Systems, Inc.
	5100  XMC5100 Zippy Drive
	5110  XMC5110 Flash Drive
	5200  XMC5200 Zippy Drive
	5201  XMC5200 Zippy Drive
	5202  XMC5200 Zippy Drive
	5280  XMC5280 Storage Drive
	fff0  ISP5200 Debugger
0b87  Plus-One I & T, Inc.
0b88  Sigma Koki Co., Ltd, Technology Center
0b89  Advanced Digital Broadcast, Ltd
0b8c  SMART Technologies Inc.
	0001  Interactive Whiteboard Controller (SB6) (HID)
	00c3  Sympodium ID350
0b95  ASIX Electronics Corp.
	1720  10/100 Ethernet
	1780  AX88178
	1790  AX88179 Gigabit Ethernet
	7720  AX88772
	772a  AX88772A Fast Ethernet
	772b  AX88772B
	7e2b  AX88772B
0b96  Sewon Telecom
0b97  O2 Micro, Inc.
	7732  Smart Card Reader
	7761  Oz776 1.1 Hub
	7762  Oz776 SmartCard Reader
	7772  OZ776 CCID Smartcard Reader
0b98  Playmates Toys, Inc.
0b99  Audio International, Inc.
0b9b  Dipl.-Ing. Stefan Kunde
	4012  Reflex RC-controller Interface
0b9d  Softprotec Co.
0b9f  Chippo Technologies
0baf  U.S. Robotics
	00e5  USR6000
	00eb  USR1120 802.11b Adapter
	00ec  56K Faxmodem
	00f1  SureConnect ADSL ATM Adapter
	00f2  SureConnect ADSL Loader
	00f5  SureConnect ADSL ATM Adapter
	00f6  SureConnect ADSL Loader
	00f7  SureConnect ADSL ATM Adapter
	00f8  SureConnect ADSL Loader
	00f9  SureConnect ADSL ATM Adapter
	00fa  SureConnect ADSL Loader
	00fb  SureConnect ADSL Ethernet/USB Router
	0111  USR5420 802.11g Adapter [Broadcom 4320 USB]
	0118  U5 802.11g Adapter
	011b  Wireless MAXg Adapter [Broadcom 4320]
	0121  USR5423 802.11bg Wireless Adapter [ZyDAS ZD1211B]
	0303  USR5637 56K Faxmodem
	6112  FaxModem Model 5633
0bb0  Concord Camera Corp.
	0100  Sound Vision Stream
	5007  3340z/Rollei DC3100
0bb1  Infinilink Corp.
0bb2  Ambit Microsystems Corp.
	0302  U10H010 802.11b Wireless Adapter [Intersil PRISM 3]
	6098  USB Cable Modem
0bb3  Ofuji Technology
0bb4  HTC (High Tech Computer Corp.)
	0001  Android Phone via mass storage [Wiko Cink Peax 2]
	00ce  mmO2 XDA GSM/GPRS Pocket PC
	00cf  SPV C500 Smart Phone
	0a01  PocketPC Sync
	0a02  Himalaya GSM/GPRS Pocket PC
	0a03  PocketPC Sync
	0a04  PocketPC Sync
	0a05  PocketPC Sync
	0a06  PocketPC Sync
	0a07  Magician PocketPC SmartPhone / O2 XDA
	0a08  PocketPC Sync
	0a09  PocketPC Sync
	0a0a  PocketPC Sync
	0a0b  PocketPC Sync
	0a0c  PocketPC Sync
	0a0d  PocketPC Sync
	0a0e  PocketPC Sync
	0a0f  PocketPC Sync
	0a10  PocketPC Sync
	0a11  PocketPC Sync
	0a12  PocketPC Sync
	0a13  PocketPC Sync
	0a14  PocketPC Sync
	0a15  PocketPC Sync
	0a16  PocketPC Sync
	0a17  PocketPC Sync
	0a18  PocketPC Sync
	0a19  PocketPC Sync
	0a1a  PocketPC Sync
	0a1b  PocketPC Sync
	0a1c  PocketPC Sync
	0a1d  PocketPC Sync
	0a1e  PocketPC Sync
	0a1f  PocketPC Sync
	0a20  PocketPC Sync
	0a21  PocketPC Sync
	0a22  PocketPC Sync
	0a23  PocketPC Sync
	0a24  PocketPC Sync
	0a25  PocketPC Sync
	0a26  PocketPC Sync
	0a27  PocketPC Sync
	0a28  PocketPC Sync
	0a29  PocketPC Sync
	0a2a  PocketPC Sync
	0a2b  PocketPC Sync
	0a2c  PocketPC Sync
	0a2d  PocketPC Sync
	0a2e  PocketPC Sync
	0a2f  PocketPC Sync
	0a30  PocketPC Sync
	0a31  PocketPC Sync
	0a32  PocketPC Sync
	0a33  PocketPC Sync
	0a34  PocketPC Sync
	0a35  PocketPC Sync
	0a36  PocketPC Sync
	0a37  PocketPC Sync
	0a38  PocketPC Sync
	0a39  PocketPC Sync
	0a3a  PocketPC Sync
	0a3b  PocketPC Sync
	0a3c  PocketPC Sync
	0a3d  PocketPC Sync
	0a3e  PocketPC Sync
	0a3f  PocketPC Sync
	0a40  PocketPC Sync
	0a41  PocketPC Sync
	0a42  PocketPC Sync
	0a43  PocketPC Sync
	0a44  PocketPC Sync
	0a45  PocketPC Sync
	0a46  PocketPC Sync
	0a47  PocketPC Sync
	0a48  PocketPC Sync
	0a49  PocketPC Sync
	0a4a  PocketPC Sync
	0a4b  PocketPC Sync
	0a4c  PocketPC Sync
	0a4d  PocketPC Sync
	0a4e  PocketPC Sync
	0a4f  PocketPC Sync
	0a50  SmartPhone (MTP)
	0a51  SPV C400 / T-Mobile SDA GSM/GPRS Pocket PC
	0a52  SmartPhone Sync
	0a53  SmartPhone Sync
	0a54  SmartPhone Sync
	0a55  SmartPhone Sync
	0a56  SmartPhone Sync
	0a57  SmartPhone Sync
	0a58  SmartPhone Sync
	0a59  SmartPhone Sync
	0a5a  SmartPhone Sync
	0a5b  SmartPhone Sync
	0a5c  SmartPhone Sync
	0a5d  SmartPhone Sync
	0a5e  SmartPhone Sync
	0a5f  SmartPhone Sync
	0a60  SmartPhone Sync
	0a61  SmartPhone Sync
	0a62  SmartPhone Sync
	0a63  SmartPhone Sync
	0a64  SmartPhone Sync
	0a65  SmartPhone Sync
	0a66  SmartPhone Sync
	0a67  SmartPhone Sync
	0a68  SmartPhone Sync
	0a69  SmartPhone Sync
	0a6a  SmartPhone Sync
	0a6b  SmartPhone Sync
	0a6c  SmartPhone Sync
	0a6d  SmartPhone Sync
	0a6e  SmartPhone Sync
	0a6f  SmartPhone Sync
	0a70  SmartPhone Sync
	0a71  SmartPhone Sync
	0a72  SmartPhone Sync
	0a73  SmartPhone Sync
	0a74  SmartPhone Sync
	0a75  SmartPhone Sync
	0a76  SmartPhone Sync
	0a77  SmartPhone Sync
	0a78  SmartPhone Sync
	0a79  SmartPhone Sync
	0a7a  SmartPhone Sync
	0a7b  SmartPhone Sync
	0a7c  SmartPhone Sync
	0a7d  SmartPhone Sync
	0a7e  SmartPhone Sync
	0a7f  SmartPhone Sync
	0a80  SmartPhone Sync
	0a81  SmartPhone Sync
	0a82  SmartPhone Sync
	0a83  SmartPhone Sync
	0a84  SmartPhone Sync
	0a85  SmartPhone Sync
	0a86  SmartPhone Sync
	0a87  SmartPhone Sync
	0a88  SmartPhone Sync
	0a89  SmartPhone Sync
	0a8a  SmartPhone Sync
	0a8b  SmartPhone Sync
	0a8c  SmartPhone Sync
	0a8d  SmartPhone Sync
	0a8e  SmartPhone Sync
	0a8f  SmartPhone Sync
	0a90  SmartPhone Sync
	0a91  SmartPhone Sync
	0a92  SmartPhone Sync
	0a93  SmartPhone Sync
	0a94  SmartPhone Sync
	0a95  SmartPhone Sync
	0a96  SmartPhone Sync
	0a97  SmartPhone Sync
	0a98  SmartPhone Sync
	0a99  SmartPhone Sync
	0a9a  SmartPhone Sync
	0a9b  SmartPhone Sync
	0a9c  SmartPhone Sync
	0a9d  SmartPhone Sync
	0a9e  SmartPhone Sync
	0a9f  SmartPhone Sync
	0b03  Ozone Mobile Broadband
	0b04  Hermes / TyTN / T-Mobile MDA Vario II / O2 Xda Trion
	0b05  P3600
	0b06  Athena / Advantage x7500 / Dopod U1000 / T-Mobile AMEO
	0b0c  Elf / Touch / P3450 / T-Mobile MDA Touch / O2 Xda Nova / Dopod S1
	0b1f  Sony Ericsson XPERIA X1
	0b2f  Rhodium
	0b51  Qtek 8310 mobile phone [Tornado Noble]
	0bce  Vario MDA
	0c01  Dream / ADP1 / G1 / Magic / Tattoo
	0c02  Dream / ADP1 / G1 / Magic / Tattoo (Debug)
	0c03  Android Phone [Fairphone First Edition (FP1)]
	0c13  Diamond
	0c1f  Sony Ericsson XPERIA X1
	0c5f  Snap
	0c86  Sensation
	0c87  Desire (debug)
	0c8d  EVO 4G (debug)
	0c91  Vision
	0c94  Vision
	0c97  Legend
	0c99  Desire (debug)
	0c9e  Incredible
	0ca2  Desire HD (debug mode)
	0ca5  Android Phone [Evo Shift 4G]
	0cae  T-Mobile MyTouch 4G Slide [Doubleshot]
	0dea  M7_UL [HTC One]
	0f25  One M8
	0f64  Desire 601
	0ff8  Desire HD (Tethering Mode)
	0ff9  Desire / Desire HD / Hero / Thunderbolt (Charge Mode)
	0ffe  Desire HD (modem mode)
	0fff  Android Fastboot Bootloader
	2008  Android Phone via MTP [Wiko Cink Peax 2]
	200b  Android Phone via PTP [Wiko Cink Peax 2]
0bb5  Murata Manufacturing Co., Ltd
0bb6  Network Alchemy
0bb7  Joytech Computer Co., Ltd
0bb8  Hitachi Semiconductor and Devices Sales Co., Ltd
0bb9  Eiger M&C Co., Ltd
0bba  ZAccess Systems
0bbb  General Meters Corp.
0bbc  Assistive Technology, Inc.
0bbd  System Connection, Inc.
0bc0  Knilink Technology, Inc.
0bc1  Fuw Yng Electronics Co., Ltd
0bc2  Seagate RSS LLC
	0502  ST3300601CB-RK 300 GB External Hard Drive
	0503  ST3250824A [Barracuda 7200.9]
	2000  Storage Adapter V3 (TPP)
	2100  FreeAgent Go
	2200  FreeAgent Go FW
	2300  Expansion Portable
	2320  USB 3.0 bridge [Portable Expansion Drive]
	2321  Expansion Portable
	2340  FreeAgent External Hard Drive
	3000  FreeAgent Desktop
	3008  FreeAgent Desk 1TB
	3101  FreeAgent XTreme 640GB
	3312  SRD00F2 Expansion Desktop Drive (STBV)
	3320  SRD00F2 [Expansion Desktop Drive]
	3332  Expansion
	5020  FreeAgent GoFlex
	5021  FreeAgent GoFlex USB 2.0
	5030  FreeAgent GoFlex Upgrade Cable STAE104
	5031  FreeAgent GoFlex USB 3.0
	5070  FreeAgent GoFlex Desk
	5071  FreeAgent GoFlex Desk
	50a1  FreeAgent GoFlex Desk
	50a5  FreeAgent GoFlex Desk USB 3.0
	5121  FreeAgent GoFlex
	5161  FreeAgent GoFlex dock
	a003  Backup Plus
	a0a1  Backup Plus Desktop
	a0a4  Backup Plus Desktop Drive
	ab00  Slim Portable Drive
	ab20  Backup Plus Portable Drive
	ab21  Backup Plus Slim
	ab31  Backup Plus Desktop Drive (5TB)
0bc3  IPWireless, Inc.
	0001  UMTS-TDD (TD-CDMA) modem
0bc4  Microcube Corp.
0bc5  JCN Co., Ltd
0bc6  ExWAY, Inc.
0bc7  X10 Wireless Technology, Inc.
	0001  ActiveHome (ACPI-compliant)
	0002  Firecracker Interface (ACPI-compliant)
	0003  VGA Video Sender (ACPI-compliant)
	0004  X10 Receiver
	0005  Wireless Transceiver (ACPI-compliant)
	0006  Wireless Transceiver (ACPI-compliant)
	0007  Wireless Transceiver (ACPI-compliant)
	0008  Wireless Transceiver (ACPI-compliant)
	0009  Wireless Transceiver (ACPI-compliant)
	000a  Wireless Transceiver (ACPI-compliant)
	000b  Transceiver (ACPI-compliant)
	000c  Transceiver (ACPI-compliant)
	000d  Transceiver (ACPI-compliant)
	000e  Transceiver (ACPI-compliant)
	000f  Transceiver (ACPI-compliant)
0bc8  Telmax Communications
0bc9  ECI Telecom, Ltd
0bca  Startek Engineering, Inc.
0bcb  Perfect Technic Enterprise Co., Ltd
0bd7  Andrew Pargeter & Associates
	a021  Amptek DP4 multichannel signal analyzer
0bda  Realtek Semiconductor Corp.
	0103  USB 2.0 Card Reader
	0104  Mass Storage Device
	0106  Mass Storage Device
	0107  Mass Storage Device
	0108  Mass Storage Device
	0111  RTS5111 Card Reader Controller
	0113  Mass Storage Device
	0115  Mass Storage Device (Multicard Reader)
	0116  RTS5116 Card Reader Controller
	0117  Mass Storage Device
	0118  Mass Storage Device
	0119  Storage Device (SD card reader)
	0129  RTS5129 Card Reader Controller
	0138  RTS5138 Card Reader Controller
	0139  RTS5139 Card Reader Controller
	0151  Mass Storage Device (Multicard Reader)
	0152  Mass Storage Device
	0153  Mass Storage Device
	0156  Mass Storage Device
	0157  Mass Storage Device
	0158  USB 2.0 multicard reader
	0159  RTS5159 Card Reader Controller
	0161  Mass Storage Device
	0168  Mass Storage Device
	0169  Mass Storage Device
	0171  Mass Storage Device
	0176  Mass Storage Device
	0178  Mass Storage Device
	0179  RTL8188ETV Wireless LAN 802.11n Network Adapter
	0184  RTS5182 Card Reader
	0186  Card Reader
	0301  multicard reader
	1724  RTL8723AU 802.11n WLAN Adapter
	2831  RTL2831U DVB-T
	2832  RTL2832U DVB-T
	2838  RTL2838 DVB-T
	5401  RTL 8153 USB 3.0 hub with gigabit ethernet
	5730  HP 2.0MP High Definition Webcam
	5775  HP "Truevision HD" laptop camera
	8150  RTL8150 Fast Ethernet Adapter
	8151  RTL8151 Adapteon Business Mobile Networks BV
	8171  RTL8188SU 802.11n WLAN Adapter
	8172  RTL8191SU 802.11n WLAN Adapter
	8174  RTL8192SU 802.11n WLAN Adapter
	8176  RTL8188CUS 802.11n WLAN Adapter
	8178  RTL8192CU 802.11n WLAN Adapter
	8179  RTL8188EUS 802.11n Wireless Network Adapter
	817f  RTL8188RU 802.11n WLAN Adapter
	8187  RTL8187 Wireless Adapter
	8189  RTL8187B Wireless 802.11g 54Mbps Network Adapter
	8192  RTL8191SU 802.11n Wireless Adapter
	8193  RTL8192DU 802.11an WLAN Adapter
	8197  RTL8187B Wireless Adapter
	8198  RTL8187B Wireless Adapter
	8199  RTL8187SU 802.11g WLAN Adapter
	8812  RTL8812AU 802.11a/b/g/n/ac WLAN Adapter
0bdb  Ericsson Business Mobile Networks BV
	1000  BV Bluetooth Device
	1002  Bluetooth Device 1.2
	1049  C3607w Mobile Broadband Module
	1900  F3507g Mobile Broadband Module
	1902  F3507g v2 Mobile Broadband Module
	1904  F3607gw Mobile Broadband Module
	1905  F3607gw v2 Mobile Broadband Module
	1906  F3607gw v3 Mobile Broadband Module
	1909  F3307 v2 Mobile Broadband Module
	190a  F3307 Mobile Broadband Module
	190b  C3607w v2 Mobile Broadband Module
0bdc  Y Media Corp.
0bdd  Orange PCS
0be2  Kanda Tsushin Kogyo Co., Ltd
0be3  TOYO Corp.
0be4  Elka International, Ltd
0be5  DOME imaging systems, Inc.
0be6  Dong Guan Humen Wonderful Wire Cable Factory
0bed  MEI
	1100  CASHFLOW SC
	1101  Series 2000 Combo Acceptor
0bee  LTK Industries, Ltd
0bef  Way2Call Communications
0bf0  Pace Micro Technology PLC
0bf1  Intracom S.A.
	0001  netMod Driver Ver 2.4.17 (CAPI)
	0002  netMod Driver Ver 2.4 (CAPI)
	0003  netMod Driver Ver 2.4 (CAPI)
0bf2  Konexx
0bf6  Addonics Technologies, Inc.
	0103  Storage Device
	1234  Storage Device
	a000  Cable 205 (TPP)
	a001  Cable 205
	a002  IDE Bridge
0bf7  Sunny Giken, Inc.
0bf8  Fujitsu Siemens Computers
	1001  Fujitsu Pocket Loox 600 PDA
	1006  SmartCard Reader 2A
	1007  Connect2Air E-5400 802.11g Wireless Adapter
	1009  Connect2Air E-5400 D1700 802.11g Wireless Adapter [Intersil ISL3887]
	100c  Keyboard FSC KBPC PX
	100f  miniCard D2301 802.11bg Wireless Module [SiS 163U]
	1017  Keyboard KB SCR
0bfd  Kvaser AB
	0004  USBcan II
	000b  Leaf Light HS
	000e  Leaf SemiPro HS
0c04  MOTO Development Group, Inc.
0c05  Appian Graphics
0c06  Hasbro Games, Inc.
0c07  Infinite Data Storage, Ltd
0c08  Agate
	0378  Q 16MB Storage Device
0c09  Comjet Information System
	a5a5  Litto Version USB2.0
0c0a  Highpoint Technologies, Inc.
0c0b  Dura Micro, Inc. (Acomdata)
	27cb  6-in-1 Flash Reader and Writer
	27d7  Multi Memory reader/writer MD-005
	27da  Multi Memory reader/writer MD-005
	27dc  Multi Memory reader/writer MD-005
	27e7  3,5'' HDD case MD-231
	27ee  3,5'' HDD case MD-231
	2814  3,5'' HDD case MD-231
	2815  3,5'' HDD case MD-231
	281d  3,5'' HDD case MD-231
	5fab  Storage Adaptor
	a109  CF/SM Reader and Writer
	a10c  SD/MS Reader and Writer
	b001  USB 2.0 Mass Storage IDE adapter
	b004  MMC/SD Reader and Writer
0c12  Zeroplus
	0005  PSX Vibration Feedback Converter
	0030  PSX Vibration Feedback Converter
	700e  Logic Analyzer (LAP-C-16032)
	8801  Xbox Controller
	8802  Xbox Controller
	8809  Red Octane Ignition Xbox DDR Pad
	880a  Pelican Eclipse PL-2023
	8810  Xbox Controller
	9902  VibraX
0c15  Iris Graphics
0c16  Gyration, Inc.
	0002  RF Technology Receiver
	0003  RF Technology Receiver
	0008  RF Technology Receiver
	0080  eHome Infrared Receiver
	0081  eHome Infrared Receiver
0c17  Cyberboard A/S
0c18  SynerTek Korea, Inc.
0c19  cyberPIXIE, Inc.
0c1a  Silicon Motion, Inc.
0c1b  MIPS Technologies
0c1c  Hang Zhou Silan Electronics Co., Ltd
0c22  Tally Printer Corp.
0c23  Lernout + Hauspie
0c24  Taiyo Yuden
	0001  Bluetooth Adaptor
	0002  Bluetooth Device2
	0005  Bluetooth Device(BC04-External)
	000b  Bluetooth Device(BC04-External)
	000c  Bluetooth Adaptor
	000e  Bluetooth Device(BC04-External)
	000f  Bluetooth Device (V2.0+EDR)
	0010  Bluetooth Device(BC04-External)
	0012  Bluetooth Device(BC04-External)
	0018  Bluetooth Device(BC04-External)
	0019  Bluetooth Device
	0021  Bluetooth Device (V2.1+EDR)
	0c24  Bluetooth Device(SAMPLE)
	ffff  Bluetooth module with BlueCore in DFU mode
0c25  Sampo Corp.
	0310  Scream Cam
0c26  Prolific Technology Inc.
	0018  USB-Serial Controller [Icom Inc. OPC-478UC]
0c27  RFIDeas, Inc
	3bfa  pcProx Card Reader
0c2e  Metrologic Instruments
	0007  Metrologic MS7120 Barcode Scanner (IBM SurePOS mode)
	0200  MS7120 Barcode Scanner
	0204  Metrologic MS7120 Barcode Scanner (keyboard mode)
	0206  Metrologic MS4980 Barcode Scanner
	0700  Metrologic MS7120 Barcode Scanner (uni-directional serial mode)
	0720  Metrologic MS7120 Barcode Scanner (bi-directional serial mode)
	0b61  Vuquest 3310g
	0b6a  Vuquest 3310 Area-Imaging Scanner
	0b81  Barcode scanner Voyager 1400g Series
0c35  Eagletron, Inc.
0c36  E Ink Corp.
0c37  e.Digital
0c38  Der An Electric Wire & Cable Co., Ltd
0c39  IFR
0c3a  Furui Precise Component (Kunshan) Co., Ltd
0c3b  Komatsu, Ltd
0c3c  Radius Co., Ltd
0c3d  Innocom, Inc.
0c3e  Nextcell, Inc.
0c44  Motorola iDEN
	0021  iDEN P2k0 Device
	0022  iDEN P2k1 Device
	03a2  iDEN Smartphone
	41d9  i1 phone
0c45  Microdia
	0011  EBUDDY
	0520  MaxTrack Wireless Mouse
	1018  Compact Flash storage memory card reader
	1020  Mass Storage Reader
	1028  Mass Storage Reader
	1030  Mass Storage Reader
	1031  Sonix Mass Storage Device
	1032  Mass Storage Reader
	1033  Sonix Mass Storage Device
	1034  Mass Storage Reader
	1035  Mass Storage Reader
	1036  Mass Storage Reader
	1037  Sonix Mass Storage Device
	1050  CF Card Reader
	1058  HDD Reader
	1060  iFlash SM-Direct Card Reader
	1061  Mass Storage Reader
	1062  Mass Storage Reader
	1063  Sonix Mass Storage Device
	1064  Mass Storage Reader
	1065  Mass Storage Reader
	1066  Mass Storage Reader
	1067  Mass Storage Reader
	1158  A56AK
	184c  VoIP Phone
	6001  Genius VideoCAM NB
	6005  Sweex Mini Webcam
	6007  VideoCAM Eye
	6009  VideoCAM ExpressII
	600d  TwinkleCam USB camera
	6011  PC Camera (SN9C102)
	6019  PC Camera (SN9C102)
	6024  VideoCAM ExpressII
	6025  VideoCAM ExpressII
	6028  Typhoon Easycam USB 330K (older)
	6029  Triplex i-mini PC Camera
	602a  Meade ETX-105EC Camera
	602b  VideoCAM NB 300
	602c  Clas Ohlson TWC-30XOP Webcam
	602d  VideoCAM ExpressII
	602e  VideoCAM Messenger
	6030  VideoCAM ExpressII
	603f  VideoCAM ExpressII
	6040  CCD PC Camera (PC390A)
	606a  CCD PC Camera (PC390A)
	607a  CCD PC Camera (PC390A)
	607b  Win2 PC Camera
	607c  CCD PC Camera (PC390A)
	607e  CCD PC Camera (PC390A)
	6080  Audio (Microphone)
	6082  VideoCAM Look
	6083  VideoCAM Look
	608c  VideoCAM Look
	608e  VideoCAM Look
	608f  PC Camera (SN9C103 + OV7630)
	60a8  VideoCAM Look
	60aa  VideoCAM Look
	60ab  PC Camera
	60af  VideoCAM Look
	60b0  Genius VideoCam Look
	60c0  PC Camera with Mic (SN9C105)
	60c8  Win2 PC Camera
	60cc  PC Camera with Mic (SN9C105)
	60ec  PC Camera with Mic (SN9C105)
	60ef  Win2 PC Camera
	60fa  PC Camera with Mic (SN9C105)
	60fb  Composite Device
	60fc  PC Camera with Mic (SN9C105)
	60fe  Audio (Microphone)
	6108  Win2 PC Camera
	6122  PC Camera (SN9C110)
	6123  PC Camera (SN9C110)
	6128  PC Camera (SN9C325 + OM6802)
	612a  PC Camera (SN9C325)
	612c  PC Camera (SN9C110)
	612e  PC Camera (SN9C110)
	612f  PC Camera (SN9C110)
	6130  PC Camera (SN9C120)
	6138  Win2 PC Camera
	613a  PC Camera (SN9C120)
	613b  Win2 PC Camera
	613c  PC Camera (SN9C120)
	613e  PC Camera (SN9C120)
	6143  PC Camera (SN9C120 + SP80708)
	6240  PC Camera (SN9C201 + MI1300)
	6242  PC Camera (SN9C201 + MI1310)
	6243  PC Camera (SN9C201 + S5K4AAFX)
	6248  PC Camera (SN9C201 + OV9655)
	624b  PC Camera (SN9C201 + CX1332)
	624c  PC Camera (SN9C201 + MI1320)
	624e  PC Camera (SN9C201 + SOI968)
	624f  PC Camera (SN9C201 + OV9650)
	6251  PC Camera (SN9C201 + OV9650)
	6253  PC Camera (SN9C201 + OV9650)
	6260  PC Camera (SN9C201 + OV7670ISP)
	6262  PC Camera (SN9C201 + OM6802)
	6270  PC Camera (SN9C201 + MI0360/MT9V011 or MI0360SOC/MT9V111) U-CAM PC Camera NE878, Whitcom WHC017, ...
	627a  PC Camera (SN9C201 + S5K53BEB)
	627b  PC Camera (SN9C201 + OV7660)
	627c  PC Camera (SN9C201 + HV7131R)
	627f  PC Camera (SN9C201 + OV965x + EEPROM)
	6280  PC Camera with Microphone (SN9C202 + MI1300)
	6282  PC Camera with Microphone (SN9C202 + MI1310)
	6283  PC Camera with Microphone (SN9C202 + S5K4AAFX)
	6288  PC Camera with Microphone (SN9C202 + OV9655)
	628a  PC Camera with Microphone (SN9C202 + ICM107)
	628b  PC Camera with Microphone (SN9C202 + CX1332)
	628c  PC Camera with Microphone (SN9C202 + MI1320)
	628e  PC Camera with Microphone (SN9C202 + SOI968)
	628f  PC Camera with Microphone (SN9C202 + OV9650)
	62a0  PC Camera with Microphone (SN9C202 + OV7670ISP)
	62a2  PC Camera with Microphone (SN9C202 + OM6802)
	62b0  PC Camera with Microphone (SN9C202 + MI0360/MT9V011 or MI0360SOC/MT9V111)
	62b3  PC Camera with Microphone (SN9C202 + OV9655)
	62ba  PC Camera with Microphone (SN9C202 + S5K53BEB)
	62bb  PC Camera with Microphone (SN9C202 + OV7660)
	62bc  PC Camera with Microphone (SN9C202 + HV7131R)
	62be  PC Camera with Microphone (SN9C202 + OV7663)
	62c0  Sonix USB 2.0 Camera
	62e0  MSI Starcam Racer
	6300  PC Microscope camera
	6310  Sonix USB 2.0 Camera
	6340  Camera
	6341  Defender G-Lens 2577 HD720p Camera
	63e0  Sonix Integrated Webcam
	63f1  Integrated Webcam
	63f8  Sonix Integrated Webcam
	6409  Webcam
	6413  Integrated Webcam
	6417  Integrated Webcam
	6419  Integrated Webcam
	641d  1.3 MPixel Integrated Webcam
	643f  Dell Integrated HD Webcam
	644d  1.3 MPixel Integrated Webcam
	6480  Sonix 1.3 MP Laptop Integrated Webcam
	648b  Integrated Webcam
	64bd  Sony Visual Communication Camera
	7401  TEMPer Temperature Sensor
	7402  TEMPerHUM Temperature & Humidity Sensor
	7403  Foot Switch
	8000  DC31VC
	8006  Dual Mode Camera (8006 VGA)
	800a  Vivitar Vivicam3350B
0c46  WaveRider Communications, Inc.
0c4a  ALGE-TIMING GmbH
	0889  Timy
	088a  Timy 2
0c4b  Reiner SCT Kartensysteme GmbH
	0100  cyberJack e-com/pinpad
	0300  cyberJack pinpad(a)
	0400  cyberJack e-com(a)
	0401  cyberJack pinpad(a2)
	0500  cyberJack RFID standard dual interface smartcard reader
	0501  cyberJack RFID comfort dual interface smartcard reader
	0502  cyberJack compact
	0504  cyberJack go / go plus
	0505  cyberJack wave
	9102  cyberJack RFID basis contactless smartcard reader
0c4c  Needham's Electronics
	0021  EMP-21 Universal Programmer
0c52  Sealevel Systems, Inc.
	2101  SeaLINK+232
	2102  SeaLINK+485
	2103  SeaLINK+232I
	2104  SeaLINK+485I
	2211  SeaPORT+2/232 (Port 1)
	2212  SeaPORT+2/485 (Port 1)
	2213  SeaPORT+2 (Port 1)
	2221  SeaPORT+2/232 (Port 2)
	2222  SeaPORT+2/485 (Port 2)
	2223  SeaPORT+2 (Port 2)
	2411  SeaPORT+4/232 (Port 1)
	2412  SeaPORT+4/485 (Port 1)
	2413  SeaPORT+4 (Port 1)
	2421  SeaPORT+4/232 (Port 2)
	2422  SeaPORT+4/485 (Port 2)
	2423  SeaPORT+4 (Port 2)
	2431  SeaPORT+4/232 (Port 3)
	2432  SeaPORT+4/485 (Port 3)
	2433  SeaPORT+4 (Port 3)
	2441  SeaPORT+4/232 (Port 4)
	2442  SeaPORT+4/485 (Port 4)
	2443  SeaPORT+4 (Port 4)
	2811  SeaLINK+8/232 (Port 1)
	2812  SeaLINK+8/485 (Port 1)
	2813  SeaLINK+8 (Port 1)
	2821  SeaLINK+8/232 (Port 2)
	2822  SeaLINK+8/485 (Port 2)
	2823  SeaLINK+8 (Port 2)
	2831  SeaLINK+8/232 (Port 3)
	2832  SeaLINK+8/485 (Port 3)
	2833  SeaLINK+8 (Port 3)
	2841  SeaLINK+8/232 (Port 4)
	2842  SeaLINK+8/485 (Port 4)
	2843  SeaLINK+8 (Port 4)
	2851  SeaLINK+8/232 (Port 5)
	2852  SeaLINK+8/485 (Port 5)
	2853  SeaLINK+8 (Port 5)
	2861  SeaLINK+8/232 (Port 6)
	2862  SeaLINK+8/485 (Port 6)
	2863  SeaLINK+8 (Port 6)
	2871  SeaLINK+8/232 (Port 7)
	2872  SeaLINK+8/485 (Port 7)
	2873  SeaLINK+8 (Port 7)
	2881  SeaLINK+8/232 (Port 8)
	2882  SeaLINK+8/485 (Port 8)
	2883  SeaLINK+8 (Port 8)
	9020  SeaLINK+422
	a02a  SeaLINK+8 (Port 1+2)
	a02b  SeaLINK+8 (Port 3+4)
	a02c  SeaLINK+8 (Port 5+6)
	a02d  SeaLINK+8 (Port 7+8)
0c53  ViewPLUS, Inc.
0c54  Glory, Ltd
0c55  Spectrum Digital, Inc.
	0510  Spectrum Digital XDS510 JTAG Debugger
	0540  SPI540
	5416  TMS320C5416 DSK
	6416  TMS320C6416 DDB
0c56  Billion Bright, Ltd
0c57  Imaginative Design Operation Co., Ltd
0c58  Vidar Systems Corp.
0c59  Dong Guan Shinko Wire Co., Ltd
0c5a  TRS International Mfg., Inc.
0c5e  Xytronix Research & Design
0c60  Apogee Electronics Corp.
	0001  MiniMe
	0002  MiniDAC
	0003  ONE
	0004  GiO
	0007  Duet
	0009  Jam
	000a  Jam Bootloader
	000b  MiC
	000c  MiC Bootloader
	8007  Duet DFU Mode
0c62  Chant Sincere Co., Ltd
0c63  Toko, Inc.
0c64  Signality System Engineering Co., Ltd
0c65  Eminence Enterprise Co., Ltd
0c66  Rexon Electronics Corp.
0c67  Concept Telecom, Ltd
0c6a  ACS
	0005  Color 320 x 240 LCD Display Terminal with Touchscreen
0c6c  JETI Technische Instrumente GmbH
	04b2  Specbos 1201
0c70  MCT Elektronikladen
	0000  USB08 Development board
	0747  Eye Movement Recorder [Visagraph]/[ReadAlyzer]
0c72  PEAK System
	000c  PCAN-USB
	000d  PCAN Pro
0c74  Optronic Laboratories Inc.
	0002  OL 700-30 Goniometer
0c76  JMTek, LLC.
	0001  Mass Storage Controller
	0002  Mass Storage Controller
	0003  USBdisk
	0004  Mass Storage Controller
	0005  Transcend Flash disk
	0006  Transcend JetFlash
	0007  Mass Storage Device
	1600  Ion Quick Play LP turntable
	1605  SSS Headphone Set
	1607  audio controller
0c77  Sipix Group, Ltd
	1001  SiPix Web2
	1002  SiPix SC2100
	1010  SiPix Snap
	1011  SiPix Blink 2
	1015  SiPix CAMeleon
0c78  Detto Corp.
0c79  NuConnex Technologies Pte., Ltd
0c7a  Wing-Span Enterprise Co., Ltd
0c86  NDA Technologies, Inc.
0c88  Kyocera Wireless Corp.
	0021  Handheld
	17da  Qualcomm Kyocera CDMA Technologies MSM
0c89  Honda Tsushin Kogyo Co., Ltd
0c8a  Pathway Connectivity, Inc.
0c8b  Wavefly Corp.
0c8c  Coactive Networks
0c8d  Tempo
0c8e  Cesscom Co., Ltd
	6000  Luxian Series
0c8f  Applied Microsystems
0c94  Cryptera
	a000  EPP 1217
0c98  Berkshire Products, Inc.
	1140  USB PC Watchdog
0c99  Innochips Co., Ltd
0c9a  Hanwool Robotics Corp.
0c9b  Jobin Yvon, Inc.
0c9d  SemTek
	0170  3873 Manual Insert card reader
0ca2  Zyfer
0ca3  Sega Corp.
0ca4  ST&T Instrument Corp.
0ca5  BAE Systems Canada, Inc.
0ca6  Castles Technology Co., Ltd
	0010  EZUSB PC/SC Smart Card Reader
	0050  EZ220PU Reader Controller
	1077  Bludrive Family Smart Card Reader
	107e  Reader Controller
	2010  myPad110 PC/SC Smart Card Reader
	3050  EZ710 Smart Card Reader
0ca7  Information Systems Laboratories
0cad  Motorola CGISS
	9001  PowerPad Pocket PCDevice
0cae  Ascom Business Systems, Ltd
0caf  Buslink
	2507  Hi-Speed USB-to-IDE Bridge Controller
	2515  Flash Disk Embedded Hub
	2516  Flash Disk Security Device
	2517  Flash Disk Mass Storage Device
	25c7  Hi-Speed USB-to-IDE Bridge Controller
	3a00  Hard Drive
	3a20  Mass Storage Device
	3acd  Mass Storage Device
0cb0  Flying Pig Systems
0cb1  Innovonics, Inc.
0cb6  Celestix Networks, Pte., Ltd
0cb7  Singatron Enterprise Co., Ltd
0cb8  Opticis Co., Ltd
0cba  Trust Electronic (Shanghai) Co., Ltd
0cbb  Shanghai Darong Electronics Co., Ltd
0cbc  Palmax Technology Co., Ltd
	0101  Pocket PC P6C
	0201  Personal Digital Assistant
	0301  Personal Digital Assistant P6M+
	0401  Pocket PC
0cbd  Pentel Co., Ltd (Electronics Equipment Div.)
0cbe  Keryx Technologies, Inc.
0cbf  Union Genius Computer Co., Ltd
0cc0  Kuon Yi Industrial Corp.
0cc1  Given Imaging, Ltd
0cc2  Timex Corp.
0cc3  Rimage Corp.
0cc4  emsys GmbH
0cc5  Sendo
0cc6  Intermagic Corp.
0cc7  Kontron Medical AG
0cc8  Technotools Corp.
0cc9  BroadMAX Technologies, Inc.
0cca  Amphenol
0ccb  SKNet Co., Ltd
0ccc  Domex Technology Corp.
0ccd  TerraTec Electronic GmbH
	0012  PHASE 26
	0013  PHASE 26
	0014  PHASE 26
	0015  Flash Update for TerraTec PHASE 26
	0021  Cameo Grabster 200
	0023  Mystify Claw
	0028  Aureon 5.1 MkII
	0032  MIDI HUBBLE
	0035  Miditech Play'n Roll
	0036  Cinergy 250 Audio
	0037  Cinergy 250 Audio
	0038  Cinergy T DVB-T Receiver
	0039  Grabster AV 400
	003b  Cinergy 400
	003c  Grabster AV 250
	0042  Cinergy Hybrid T XS
	0043  Cinergy T XS
	004e  Cinergy T XS
	004f  Cinergy Analog XS
	0055  Cinergy T XE (Version 1, AF9005)
	005c  Cinergy T
	0069  Cinergy T XE (Version 2, AF9015)
	006b  Cinergy HT PVR (EU)
	0072  Cinergy Hybrid T
	0077  Aureon Dual USB
	0078  Cinergy T XXS
	0086  Cinergy Hybrid XE
	008e  Cinergy HTC XS
	0097  Cinergy T RC MKII
	0099  AfaTech 9015 [Cinergy T Stick Dual]
	00a5  Cinergy Hybrid Stick
	00a9  RTL2838 DVB-T COFDM Demodulator [TerraTec Cinergy T Stick Black]
	00b3  NOXON DAB/DAB+ Stick
	00e0  NOXON DAB/DAB+ Stick V2
	10a7  TerraTec G3
0cd4  Bang Olufsen
	0101  BeolinkPC2
0cd5  LabJack Corporation
	0003  U3
	0009  UE9
0cd7  NewChip S.r.l.
0cd8  JS Digitech, Inc.
	2007  Smart Card Reader/JSTU-9700
0cd9  Hitachi Shin Din Cable, Ltd
0cde  Z-Com
	0001  XI-750 802.11b Wireless Adapter [Atmel AT76C503A]
	0002  XI-725/726 Prism2.5 802.11b Adapter
	0003  Sagem 802.11b Dongle
	0004  Sagem 802.11b Dongle
	0005  XI-735 Prism3 802.11b Adapter
	0006  XG-300 802.11b Adapter
	0008  XG-703A 802.11g Wireless Adapter [Intersil ISL3887]
	0009  (ZD1211)IEEE 802.11b+g Adapter
	0011  ZD1211
	0012  AR5523
	0013  AR5523 driver (no firmware)
	0014  NB 802.11g Wireless LAN Adapter(3887A)
	0015  XG-705A 802.11g Wireless Adapter [Intersil ISL3887]
	0016  NB 802.11g Wireless LAN Adapter(3887A)
	0018  NB 802.11a/b/g Wireless LAN Adapter(3887A)
	001a  802.11bg
	001c  802.11b/g Wireless Network Adapter
	0020  AG-760A 802.11abg Wireless Adapter [ZyDAS ZD1211B]
	0022  802.11b/g/n Wireless Network Adapter
	0023  UB81 802.11bgn
	0025  802.11b/g/n USB Wireless Network Adapter
	0026  UB82 802.11abgn
	0027  Sphairon Homelink 1202 802.11n Wireless Adapter [Atheros AR9170]
0ce5  Validation Technologies International
	0003  Matrix
0ce9  Pico Technology
	1001  PicoScope3000 series PC Oscilloscope
	1007  PicoScope 2000 series PC Oscilloscope
	1008  PicoScope 5000 series PC Oscilloscope
	1009  PicoScope 4000 series PC Oscilloscope
	100e  PicoScope 6000 series PC Oscilloscope
	1012  PicoScope 3000A series PC Oscilloscope
	1016  PicoScope 2000A series PC Oscilloscope
	1018  PicoScope 4000A series PC Oscilloscope
	1200  PicoScope 2000 series PC Oscilloscope
	1201  PicoScope 3000 series PC Oscilloscope
	1202  PicoScope 4000 series PC Oscilloscope
	1203  PicoScope 5000 series PC Oscilloscope
	1204  PicoScope 6000 series PC Oscilloscope
	1211  PicoScope 3000 series PC Oscilloscope
	1212  PicoScope 4000 series PC Oscilloscope
0cf1  e-Conn Electronic Co., Ltd
0cf2  ENE Technology, Inc.
	6220  SD Card Reader (SG361)
	6225  SD card reader (UB6225)
	6230  SD Card Reader (UB623X)
	6250  SD card reader (UB6250)
0cf3  Atheros Communications, Inc.
	0001  AR5523
	0002  AR5523 (no firmware)
	0003  AR5523
	0004  AR5523 (no firmware)
	0005  AR5523
	0006  AR5523 (no firmware)
	1001  Thomson TG121N [Atheros AR9001U-(2)NG]
	1002  TP-Link TL-WN821N v2 / TL-WN822N v1 802.11n [Atheros AR9170]
	1006  TP-Link TL-WN322G v3 / TL-WN422G v2 802.11g [Atheros AR9271]
	1010  3Com 3CRUSBN275 802.11abgn Wireless Adapter [Atheros AR9170]
	20ff  AR7010 (no firmware)
	3000  AR3011 Bluetooth (no firmware)
	3002  AR3011 Bluetooth
	3004  AR3012 Bluetooth 4.0
	3005  AR3011 Bluetooth
	3008  Bluetooth (AR3011)
	7015  TP-Link TL-WN821N v3 / TL-WN822N v2 802.11n [Atheros AR7010+AR9287]
	9170  AR9170 802.11n
	9271  AR9271 802.11n
	b002  Ubiquiti WiFiStation 802.11n [Atheros AR9271]
	b003  Ubiquiti WiFiStationEXT 802.11n [Atheros AR9271]
0cf4  Fomtex Corp.
0cf5  Cellink Co., Ltd
0cf6  Compucable Corp.
0cf7  ishoni Networks
0cf8  Clarisys, Inc.
	0750  Claritel-i750 - vp
0cf9  Central System Research Co., Ltd
0cfa  Inviso, Inc.
0cfc  Minolta-QMS, Inc.
	2301  Magicolor 2300 DL
	2350  Magicolor 2350EN/3300
	3100  Magicolor 3100
	7300  Magicolor 5450/5550
0cff  SAFA MEDIA Co., Ltd.
	0320  SR-380N
0d06  telos EDV Systementwicklung GmbH
0d08  UTStarcom
	0602  DV007 [serial]
	0603  DV007 [storage]
0d0b  Contemporary Controls
0d0c  Astron Electronics Co., Ltd
0d0d  MKNet Corp.
0d0e  Hybrid Networks, Inc.
0d0f  Feng Shin Cable Co., Ltd
0d10  Elastic Networks
	0001  StormPort (WDM)
0d11  Maspro Denkoh Corp.
0d12  Hansol Electronics, Inc.
0d13  BMF Corp.
0d14  Array Comm, Inc.
0d15  OnStream b.v.
0d16  Hi-Touch Imaging Technologies Co., Ltd
	0001  PhotoShuttle
	0002  Photo Printer 730 series
	0004  Photo Printer 63xPL/PS
	0100  Photo Printer 63xPL/PS
	0102  Photo Printer 64xPS
	0103  Photo Printer 730 series
	0104  Photo Printer 63xPL/PS
	0105  Photo Printer 64xPS
	0200  Photo Printer 64xDL
0d17  NALTEC, Inc.
0d18  coaXmedia
0d19  Hank Connection Industrial Co., Ltd
0d28  NXP
	0204  LPC1768
0d32  Leo Hui Electric Wire & Cable Co., Ltd
0d33  AirSpeak, Inc.
0d34  Rearden Steel Technologies
0d35  Dah Kun Co., Ltd
0d3a  Posiflex Technologies, Inc.
	0206  Series 3xxx Cash Drawer
	0207  Series 3xxx Cash Drawer
	0500  Magnetic Stripe Reader
0d3c  Sri Cable Technology, Ltd
0d3d  Tangtop Technology Co., Ltd
	0001  HID Keyboard
	0040  PS/2 Adapter
0d3e  Fitcom, inc.
0d3f  MTS Systems Corp.
0d40  Ascor, Inc.
0d41  Ta Yun Terminals Industrial Co., Ltd
0d42  Full Der Co., Ltd
0d46  Kobil Systems GmbH
	2012  KAAN Standard Plus (Smartcard reader)
	3003  mIDentity Light / KAAN SIM III
	4000  mIDentity (mass storage)
	4001  mIDentity Basic/Classic (composite device)
	4081  mIDentity Basic/Classic (installationless)
0d48  Promethean Limited
	0001  ACTIVboard
	0004  ACTIVboard
	0100  Audio
0d49  Maxtor
	3000  Drive
	3010  3000LE Drive
	3100  Hi-Speed USB-IDE Bridge Controller
	3200  Personal Storage 3200
	5000  5000XT Drive
	5010  5000LE Drive
	5020  Mobile Hard Disk Drive
	7000  OneTouch
	7010  OneTouch
	7100  OneTouch II 300GB External Hard Disk
	7310  OneTouch 4
	7410  Mobile Hard Disk Drive (1TB)
	7450  Basics Portable USB Device
0d4a  NF Corp.
0d4b  Grape Systems, Inc.
0d4c  Tedas AG
0d4d  Coherent, Inc.
0d4e  Agere Systems Netherland BV
	047a  WLAN Card
	1000  Wireless Card Model 0801
	1001  Wireless Card Model 0802
0d4f  EADS Airbus France
0d50  Cleware GmbH
	0011  USB-Temp2 Thermometer
	0040  F4 foot switch
0d51  Volex (Asia) Pte., Ltd
0d53  HMI Co., Ltd
0d54  Holon Corp.
0d55  ASKA Technologies, Inc.
0d56  AVLAB Technology, Inc.
0d57  Solomon Microtech, Ltd
0d5c  SMC Networks, Inc.
	a001  SMC2662W (v1) EZ Connect 802.11b Wireless Adapter [Atmel AT76C503A]
	a002  SMC2662W v2 / SMC2662W-AR / Belkin F5D6050 [Atmel at76c503a]
0d5e  Myacom, Ltd
	2346  BT Digital Access adapter
0d5f  CSI, Inc.
0d60  IVL Technologies, Ltd
0d61  Meilu Electronics (Shenzhen) Co., Ltd
0d62  Darfon Electronics Corp.
	0003  Smartcard Reader
	0004  Keyboard
	001b  Keyboard
	001c  Benq X120 Internet Keyboard Pro
	0306  M530 Mouse
	0800  Magic Wheel
	2021  AM805 Keyboard
	2026  TECOM Bluetooth Device
	2050  Mouse
	2106  Dell L20U Multimedia Keyboard
	a100  Optical Mouse
0d63  Fritz Gegauf AG
0d64  DXG Technology Corp.
	0105  Dual Mode Digital Camera 1.3M
	0107  Horus MT-409 Camera
	0108  Dual Mode Digital Camera
	0202  Dual Mode Video Camera Device
	0303  DXG-305V Camera
	1001  SiPix Stylecam/UMAX AstraPix 320s
	1002  Fashion Cam 01 Dual-Mode DSC (Video Camera)
	1003  Fashion Cam Dual-Mode DSC (Controller)
	1021  D-Link DSC 350F
	1208  Dual Mode Still Camera Device
	2208  Mass Storage
	3105  Dual Mode Digital Camera Disk
	3108  Digicam Mass Storage Device
0d65  KMJP Co., Ltd
0d66  TMT
0d67  Advanet, Inc.
0d68  Super Link Electronics Co., Ltd
0d69  NSI
0d6a  Megapower International Corp.
0d6b  And-Or Logic
0d70  Try Computer Co., Ltd
0d71  Hirakawa Hewtech Corp.
0d72  Winmate Communication, Inc.
0d73  Hit's Communications, Inc.
0d76  MFP Korea, Inc.
0d77  Power Sentry/Newpoint
0d78  Japan Distributor Corp.
0d7a  MARX Datentechnik GmbH
	0001  CrypToken
0d7b  Wellco Technology Co., Ltd
0d7c  Taiwan Line Tek Electronic Co., Ltd
0d7d  Phison Electronics Corp.
	0100  PS1001/1011/1006/1026 Flash Disk
	0110  Gigabyte FlexDrive
	0120  Disk Pro 64MB
	0124  GIGABYTE Disk
	0240  I/O-Magic/Transcend 6-in-1 Card Reader
	110e  NEC uPD720121/130 USB-ATA/ATAPI Bridge
	1240  Apacer 6-in-1 Card Reader 2.0
	1270  Wolverine SixPac 6000
	1300  Flash Disk
	1320  PS2031 Flash Disk
	1400  Attache 256MB USB 2.0 Flash Drive
	1420  PS2044 Pen Drive
	1470  Vosonic X's-Drive II+ VP2160
	1620  USB Disk Pro
	1900  USB Thumb Drive
0d7e  American Computer & Digital Components
	2507  Hi-Speed USB-to-IDE Bridge Controller
	2517  Hi-Speed Mass Storage Device
	25c7  Hi-Speed USB-to-IDE Bridge Controller
0d7f  Essential Reality LLC
	0100  P5 Glove glove controller
0d80  H.R. Silvine Electronics, Inc.
0d81  TechnoVision
0d83  Think Outside, Inc.
0d87  Dolby Laboratories Inc.
0d89  Oz Software
0d8a  King Jim Co., Ltd
	0101  TEPRA PRO
0d8b  Ascom Telecommunications, Ltd
0d8c  C-Media Electronics, Inc.
	0001  Audio Device
	0002  Composite Device
	0003  Sound Device
	0006  Storm HP-USB500 5.1 Headset
	000c  Audio Adapter
	000d  Composite Device
	000e  Audio Adapter (Planet UP-100, Genius G-Talk)
	001f  CM108 Audio Controller
	0102  CM106 Like Sound Device
	0103  CM102-A+/102S+ Audio Controller
	0104  CM103+ Audio Controller
	0105  CM108 Audio Controller
	0107  CM108 Audio Controller
	010f  CM108 Audio Controller
	0115  CM108 Audio Controller
	0139  Multimedia Headset [Gigaware by Ignition L.P.]
	013c  CM108 Audio Controller
	0201  CM6501
	5000  Mass Storage Controller
	5200  Mass Storage Controller(0D8C,5200)
	b213  USB Phone CM109 (aka CT2000,VPT1000)
0d8d  Promotion & Display Technology, Ltd
	0234  V-234 Composite Device
	0550  V-550 Composite Device
	0551  V-551 Composite Device
	0552  V-552 Composite Device
	0651  V-651 Composite Device
	0652  V-652 Composite Device
	0653  V-653 Composite Device
	0654  V-654 Composite Device
	0655  V-655 Composite Device
	0656  V-656 Composite Device
	0657  V-657 Composite Device
	0658  V-658 Composite Device
	0659  V-659 Composite Device
	0660  V-660 Composite Device
	0661  V-661 Composite Device
	0662  V-662 Composite Device
	0850  V-850 Composite Device
	0851  V-851 Composite Device
	0852  V-852 Composite Device
	0901  V-901 Composite Device
	0902  V-902 Composite Device
	0903  V-903 Composite Device
	4754  Voyager DMP Composite Device
	bb00  Bloomberg Composite Device
	bb01  Bloomberg Composite Device
	bb02  Bloomberg Composite Device
	bb03  Bloomberg Composite Device
	bb04  Bloomberg Composite Device
	bb05  Bloomberg Composite Device
	fffe  Global Tuner Composite Device
	ffff  Voyager DMP Composite Device
0d8e  Global Sun Technology, Inc.
	0163  802.11g 54 Mbps Wireless Dongle
	1621  802.11b Wireless Adapter
	3762  Cohiba 802.11g Wireless Mini adapter [Intersil ISL3887]
	3763  802.11g Wireless dongle
	7100  802.11b Adapter
	7110  WL-210 / WU210P 802.11b Wireless Adapter [Atmel AT76C503A]
	7605  TRENDnet TEW-224UB 802.11b Wireless Adapter [Atmel AT76C503A]
	7801  AR5523
	7802  AR5523 (no firmware)
	7811  AR5523
	7812  AR5523 (no firmware)
	7a01  PRISM25 802.11b Adapter
0d8f  Pitney Bowes
0d90  Sure-Fire Electrical Corp.
0d96  Skanhex Technology, Inc.
	0000  Jenoptik JD350 video
	3300  SX330z Camera
	4100  SX410z Camera
	4102  MD 9700 Camera
	4104  Jenoptik JD-4100z3s
	410a  Medion 9801/Novatech SX-410z
	5200  SX-520z Camera
0d97  Santa Barbara Instrument Group
	0001  SBIG Astronomy Camera (without firmware)
	0101  SBIG Astronomy Camera (with firmware)
0d98  Mars Semiconductor Corp.
	0300  Avaya Wireless Card
	1007  Discovery Kids Digital Camera
0d99  Trazer Technologies, Inc.
0d9a  RTX Telecom AS
	0001  Bluetooth Device
0d9b  Tat Shing Electrical Co.
0d9c  Chee Chen Hi-Technology Co., Ltd
0d9d  Sanwa Supply, Inc.
0d9e  Avaya
	0300  Wireless Card
0d9f  Powercom Co., Ltd
	0001  Uninterruptible Power Supply
	0002  Black Knight PRO / WOW Uninterruptible Power Supply (Cypress HID->COM RS232)
	00a2  Imperial Uninterruptible Power Supply (HID PDC)
	00a3  Smart King PRO Uninterruptible Power Supply (HID PDC)
	00a4  WOW Uninterruptible Power Supply (HID PDC)
	00a5  Vanguard Uninterruptible Power Supply (HID PDC)
	00a6  Black Knight PRO Uninterruptible Power Supply (HID PDC)
0da0  Danger Research
0da1  Suzhou Peter's Precise Industrial Co., Ltd
0da2  Land Instruments International, Ltd
0da3  Nippon Electro-Sensory Devices Corp.
0da4  Polar Electro Oy
	0001  Interface
	0008  Loop
0da7  IOGear, Inc.
0da8  softDSP Co., Ltd
	0001  SDS 200A Oscilloscope
0dab  Cubig Group
	0100  DVR/CVR-M140 MP3 Player
0dad  Westover Scientific
0db0  Micro Star International
	1020  PC2PC WLAN Card
	1967  Bluetooth Dongle
	3713  Primo 73
	3801  Motorola Bluetooth 2.1+EDR Device
	4011  Medion Flash XL V2.0 Card Reader
	4023  Lexar Mobile Card Reader
	4600  802.11b/g Turbo Wireless Adapter
	5501  Mass Storage Device
	5502  Mass Storage Device
	5513  MP3 Player
	5515  MP3 Player
	5516  MP3 Player
	5580  Mega Sky 580 DVB-T Tuner [M902x]
	5581  Mega Sky 580 DVB-T Tuner [GL861]
	6823  UB11B/MS-6823 802.11b Wi-Fi adapter
	6826  IEEE 802.11g Wireless Network Adapter
	6855  Bluetooth Device
	6861  MSI-6861 802.11g WiFi adapter
	6865  RT2570
	6869  RT2570
	6874  RT2573
	6877  RT2573
	6881  Bluetooth Class I EDR Device
	688a  Bluetooth Class I EDR Device
	6899  802.11bgn 1T1R Mini Card Wireless Adapter
	6970  MS-6970 BToes Bluetooth adapter
	697a  Bluetooth Dongle
	6982  Medion Flash XL Card Reader
	a861  RT2573
	a874  RT2573
	a970  Bluetooth dongle
	a97a  Bluetooth EDR Device
	b970  Bluetooth EDR Device
	b97a  Bluetooth EDR Device
0db1  Wen Te Electronics Co., Ltd
0db2  Shian Hwi Plug Parts, Plastic Factory
0db3  Tekram Technology Co., Ltd
0db4  Chung Fu Chen Yeh Enterprise Corp.
0db5  Access IS
	0139  Barcode Module - CDC serial
	013a  Barcode Module - Virtual Keyboard
	013b  Barcode Module - HID
	0160  NFC and Smartcard Module (NSM)
0db7  ELCON Systemtechnik
	0002  Goldpfeil P-LAN
0dba  Digidesign
	1000  Mbox 1 [Mbox]
	3000  Mbox 2
0dbc  A&D Medical
	0003  AND Serial Cable [AND Smart Cable]
0dbe  Jiuh Shiuh Precision Industry Co., Ltd
0dbf  Jess-Link International
	0002  SmartDongle Security Key
	0200  HDD Storage Solution
	021b  USB-2.0 IDE Adapter
	0300  Storage Adapter
	0333  Storage Adapter
	0707  ZIV Drive
0dc0  G7 Solutions (formerly Great Notions)
0dc1  Tamagawa Seiki Co., Ltd
0dc3  Athena Smartcard Solutions, Inc.
	0801  ASEDrive III
	0802  ASEDrive IIIe
	1104  ASEDrive IIIe KB
	1701  ASEKey
	1702  ASEKey
0dc4  Macpower Peripherals, Ltd
	0040  Mass Storage Device
	0041  Mass Storage Device
	0042  Mass Storage Device
	0101  Hi-Speed Mass Storage Device
	0209  SK-3500 S2
	020a  Oyen Digital MiniPro 2.5" hard drive enclosure
0dc5  SDK Co., Ltd
0dc6  Precision Squared Technology Corp.
	2301  Wireless Touchpad Keyboard
0dc7  First Cable Line, Inc.
0dcd  NetworkFab Corp.
	0001  Remote Interface Adapter
	0002  High Bandwidth Codec
0dd0  Access Solutions
	1002  Triple Talk Speech Synthesizer
0dd1  Contek Electronics Co., Ltd
0dd2  Power Quotient International Co., Ltd
	0003  Mass Storage (P)
0dd3  MediaQ
0dd4  Custom Engineering SPA
0dd5  California Micro Devices
0dd7  Kocom Co., Ltd
0dd8  Netac Technology Co., Ltd
	1060  USB-CF-Card
	e007  OnlyDisk U222 Pendrive
	f607  OnlyDisk U208 1G flash drive [U-SAFE]
0dd9  HighSpeed Surfing
0dda  Integrated Circuit Solution, Inc.
	0001  Multi-Card Reader 6in1
	0002  Multi-Card Reader 7in1
	0003  Flash Disk
	0005  Internal Multi-Card Reader 6in1
	0008  SD single card reader
	0009  MS single card reader
	000a  MS+SD Dual Card Reader
	000b  SM single card reader
	0101  All-In-One Card Reader
	0102  All-In-One Card Reader
	0301  MP3 Player
	0302  Multi-Card MP3 Player
	1001  Multi-Flash Disk
	2001  Multi-Card Reader
	2002  Q018 default PID
	2003  Multi-Card Reader
	2005  Datalux DLX-1611 16in1 Card Reader
	2006  All-In-One Card Reader
	2007  USB to ATAPI bridge
	2008  All-In-One Card Reader
	2013  SD/MS Combo Card Reader
	2014  SD/MS Single Card Reader
	2023  card reader SD/MS DEMO board with ICSI brand name (MaskROM version)
	2024  card reader SD/MS DEMO board with Generic brand name (MaskROM version)
	2026  USB2.0 Card Reader
	2027  USB 2.0 Card Reader
	2315  UFD MP3 player (model 2)
	2318  UFD MP3 player (model 1)
	2321  UFD MP3 player
0ddb  Tamarack, Inc.
0ddd  Datelink Technology Co., Ltd
0dde  Ubicom, Inc.
0de0  BD Consumer Healthcare
0de7  USBmicro
	0191  U401 Interface card
	01a5  U421 interface card
	01c3  U451 relay interface card
0dea  UTECH Electronic (D.G.) Co., Ltd.
0ded  Novasonics
0dee  Lifetime Memory Products
	4010  Storage Adapter
0def  Full Rise Electronic Co., Ltd
0df4  NET&SYS
	0201  MNG-2005
0df6  Sitecom Europe B.V.
	0001  C-Media VOIP Device
	0004  Bluetooth 2.0 Adapter 100m
	0007  Bluetooth 2.0 Adapter 10m
	000b  Bluetooth 2.0 Adapter DFU
	000d  WL-168 Wireless Network Adapter 54g
	0017  WL-182 Wireless-N Network USB Card
	0019  Bluetooth 2.0 adapter 10m CN-512v2 001
	001a  Bluetooth 2.0 adapter 100m CN-521v2 001 
	002b  WL-188 Wireless Network 300N USB Adapter
	002c  WL-301 Wireless Network 300N USB Adapter
	002d  WL-302 Wireless Network 300N USB dongle 
	0036  WL-603 Wireless Adapter
	0039  WL-315 Wireless-N USB Adapter
	003b  WL-321 Wireless USB Gaming Adapter 300N
	003c  WL-323 Wireless-N USB Adapter
	003d  WL-324 Wireless USB Adapter 300N
	003e  WL-343 Wireless USB Adapter 150N X1
	003f  WL-608 Wireless USB Adapter 54g
	0040  WL-344 Wireless Adapter 300N X2 [Ralink RT3071]
	0041  WL-329 Wireless Dualband USB adapter 300N
	0042  WL-345 Wireless USB adapter 300N X3
	0045  WL-353 Wireless USB Adapter 150N Nano
	0047  WL-352v1 Wireless USB Adapter 300N 002
	0048  WL-349v1 Wireless Adapter 150N 002 [Ralink RT3070]
	0049  WL-356 Wireless Adapter 300N
	004a  WL-358v1 Wireless Micro USB Adapter 300N X3 002
	004b  WL-349v3 Wireless Micro Adapter 150N X1 [Realtek RTL8192SU]
	004c  WL-352 802.11n Adapter [Realtek RTL8191SU]
	0050  WL-349v4 Wireless Micro Adapter 150N X1 [Ralink RT3370]
	0056  LN-031 10/100/1000 Ethernet Adapter
	005d  WLA-2000 v1.001 WLAN [RTL8191SU]
	0060  WLA-4000 802.11bgn [Ralink RT3072]
	0062  WLA-5000 802.11abgn [Ralink RT3572]
	0072  AX88179 Gigabit Ethernet [Sitecom]
	061c  LN-028 Network USB 2.0 Adapter
	21f4  44 St Bluetooth Device
	2200  Sitecom bluetooth2.0 class 2 dongle CN-512
	2208  Sitecom bluetooth2.0 class 2 dongle CN-520
	2209  Sitecom bluetooth2.0 class 1 dongle CN-521
	3068  DC-104v2 ISDN Adapter [HFC-S]
	9071  WL-113 rev 1 Wireless Network USB Adapter
	9075  WL-117 Hi-Speed USB Adapter
	90ac  WL-172 Wireless Network USB Adapter 54g Turbo
	9712  WL-113 rev 2 Wireless Network USB Adapter
0df7  Mobile Action Technology, Inc.
	0620  MA-620 Infrared Adapter
	0700  MA-700 Bluetooth Adapter
	0720  MA-720 Bluetooth Adapter
	0722  Bluetooth Dongle
	0730  MA-730/MA-730G Bluetooth Adapter
	0800  Data Cable
	0820  Data Cable
	0900  MA i-gotU Travel Logger GPS
	1800  Generic Card Reader
	1802  Card Reader
0dfa  Toyo Communication Equipment Co., Ltd
0dfc  GeneralTouch Technology Co., Ltd
	0001  Touchscreen
0e03  Nippon Systemware Co., Ltd
0e08  Winbest Technology Co., Ltd
0e0b  Amigo Technology Inc.
	9031  802.11n Wireless USB Card
	9041  802.11n Wireless USB Card
0e0c  Gesytec
	0101  LonUSB LonTalk Network Adapter
0e0d  PicoQuant GmbH
	0003  PicoHarp 300
0e0f  VMware, Inc.
	0001  Device
	0002  Virtual USB Hub
	0003  Virtual Mouse
	0004  Virtual CCID
	0005  Virtual Mass Storage
	0006  Virtual Keyboard
	f80a  Smoker FX2
0e16  JMTek, LLC
0e17  Walex Electronic, Ltd
0e1a  Unisys
0e1b  Crewave
0e20  Pegasus Technologies Ltd.
	0101  NoteTaker
	0200  Seiko Instruments InkLink Handwriting System
0e21  Cowon Systems, Inc.
	0300  iAudio CW200
	0400  MP3 Player
	0500  iAudio M3
	0510  iAudio X5, subpack USB port
	0513  iAudio X5, side USB port
	0520  iAudio M5, side USB port
	0601  iAudio G3
	0681  iAUDIO E2
	0700  iAudio U3
	0751  iAudio 7
	0760  iAUDIO U5 / iAUDIO G2
	0800  Cowon D2 (UMS mode)
	0801  Cowon D2 (MTP mode)
	0910  iAUDIO 9
	0920  J3
0e22  Symbian Ltd.
0e23  Liou Yuane Enterprise Co., Ltd
0e25  VinChip Systems, Inc.
0e26  J-Phone East Co., Ltd
0e30  HeartMath LLC
0e34  Micro Computer Control Corp.
0e35  3Pea Technologies, Inc.
0e36  TiePie engineering
	0008  Handyscope HS3
	0009  Handyscope HS3 (br)
	000a  Handyscope HS4
	000b  Handyscope HS4 (br)
	000e  Handyscope HS4-DIFF
	000f  Handyscope HS4-DIFF (br)
	0010  Handyscope HS2
	0011  TiePieSCOPE HS805 (br)
	0012  TiePieSCOPE HS805
	0013  Handyprobe HP3
	0014  Handyprobe HP3
	0018  Handyprobe HP2
	001b  Handyscope HS5
	0042  TiePieSCOPE HS801
	00fd  USB To Parallel adapter
	00fe  USB To Parallel adapter
0e38  Stratitec, Inc.
0e39  Smart Modular Technologies, Inc.
	0137  Bluetooth Device
0e3a  Neostar Technology Co., Ltd
	1100  CW-1100 Wireless Network Adapter
0e3b  Mansella, Ltd
0e41  Line6, Inc.
	4147  TonePort GX
	414d  Pod HD500
	4156  POD HD Desktop
	4250  BassPODxt
	4252  BassPODxt Pro
	4642  BassPODxt Live
	4650  PODxt Live
	4750  GuitarPort
	5044  PODxt
	5050  PODxt Pro
	534d  SeaMonkey
0e44  Sun-Riseful Technology Co., Ltd.
0e48  Julia Corp., Ltd
	0100  CardPro SmartCard Reader
0e4a  Shenzhen Bao Hing Electric Wire & Cable Mfr. Co.
0e4c  Radica Games, Ltd
	1097  Gamester Controller
	2390  Games Jtech Controller
	7288  funkey reader
0e50  TechnoData Interware
	0002  Matrixlock Dongle (HID)
0e55  Speed Dragon Multimedia, Ltd
	110a  Tanic S110-SG1 + ISSC IS1002N [Slow Infra-Red (SIR) & Bluetooth 1.2 (Class 2) Adapter]
	110b  MS3303H USB-to-Serial Bridge
0e56  Kingston Technology Company, Inc.
	6021  K-PEX 100
0e5a  Active Co., Ltd
0e5b  Union Power Information Industrial Co., Ltd
0e5c  Bitland Information Technology Co., Ltd
	6118  LCD Device
	6119  remote receive and control device
	6441  C-Media Sound Device
0e5d  Neltron Industrial Co., Ltd
0e5e  Conwise Technology Co., Ltd.
	6622  CW6622
0e66  Hawking Technologies
	0001  HWUN1 Hi-Gain Wireless-300N Adapter w/ Upgradable Antenna [Ralink RT2870]
	0003  HWDN1 Hi-Gain Wireless-300N Dish Adapter [Ralink RT2870]
	0009  HWUN2 Hi-Gain Wireless-150N Adapter w/ Upgradable Antenna [Ralink RT2770]
	000b  HWDN2 Hi-Gain Wireless-150N Dish Adapter [Ralink RT2770]
	0013  HWUN3 Hi-Gain Wireless-N Adapter [Ralink RT3070]
	0015  HWDN2 Rev. E Hi-Gain Wireless-150N Dish Adapter [Realtek RTL8191SU]
	0017  HAWNU1 Hi-Gain Wireless-150N Network Adapter with Range Amplifier [Ralink RT3070]
	0018  Wireless-N Network Adapter [Ralink RT2870]
	400b  UF100 10/100 Network Adapter
	400c  UF100 Ethernet [pegasus2]
0e67  Fossil, Inc.
	0002  Wrist PDA
0e6a  Megawin Technology Co., Ltd
	0101  MA100 [USB-UART Bridge IC]
	030b  Truly Ergonomic Computer Keyboard (Device Firmware Update mode)
	030c  Truly Ergonomic Computer Keyboard
	6001  GEMBIRD Flexible keyboard KB-109F-B-DE
0e6f  Logic3
	0003  Freebird wireless Controller
	0005  Eclipse wireless Controller
	0006  Edge wireless Controller
	0128  Wireless PS3 Controller
0e70  Tokyo Electronic Industry Co., Ltd
0e72  Hsi-Chin Electronics Co., Ltd
0e75  TVS Electronics, Ltd
0e79  Archos, Inc.
	1106  Pocket Media Assistant - PMA400
	1204  Gmini XS 200
	1306  504 Portable Multimedia Player
	1330  5 Tablet
	1332  5 IMT
	1416  32 IT
	1417  A43 IT
	14ad  97 Titanium HD
	150e  80 G9
	3001  40 Titanium
0e7b  On-Tech Industry Co., Ltd
0e7e  Gmate, Inc.
	0001  Yopy 3000 PDA
	1001  YP3X00 PDA
0e82  Ching Tai Electric Wire & Cable Co., Ltd
0e83  Shin An Wire & Cable Co.
0e8c  Well Force Electronic Co., Ltd
0e8d  MediaTek Inc.
	0003  MT6227 phone
	0004  MT6227 phone
	0023  S103
	1806  Samsung SE-208 Slim Portable DVD Writer
	1836  Samsung SE-S084 Super WriteMaster Slim External DVD writer
	2000  MT65xx Preloader
	3329  Qstarz BT-Q1000XT
	763e  MT7630e Bluetooth Adapter
0e8f  GreenAsia Inc.
	0003  MaxFire Blaze2
	0012  Joystick/Gamepad
	0016  4 port USB 1.1 hub UH-174
	0020  USB to PS/2 Adapter
	0021  Multimedia Keyboard Controller
	0022  multimedia keyboard controller
	0201  SmartJoy Frag Xpad/PS2 adaptor
0e90  WiebeTech, LLC
	0100  Storage Adapter V1
0e91  VTech Engineering Canada, Ltd
0e92  C's Glory Enterprise Co., Ltd
0e93  eM Technics Co., Ltd
0e95  Future Technology Co., Ltd
0e96  Aplux Communications, Ltd
	c001  TRUST 380 USB2 SPACEC@M
0e97  Fingerworks, Inc.
	0908  Composite HID (Keyboard and Mouse)
0e98  Advanced Analogic Technologies, Inc.
0e99  Parallel Dice Co., Ltd
0e9a  TA HSING Industries, Ltd
0e9b  ADTEC Corp.
0e9c  Streamzap, Inc.
	0000  Streamzap Remote Control
0e9f  Tamura Corp.
0ea0  Ours Technology, Inc.
	2126  7-in-1 Card Reader
	2153  SD Card Reader Key
	2168  Transcend JetFlash 2.0 / Astone USB Drive / Intellegent Stick 2.0
	6803  OTI-6803 Flash Disk
	6808  OTI-6808 Flash Disk
	6828  OTI-6828 Flash Disk
	6858  OTi-6858 serial adapter
0ea6  Nihon Computer Co., Ltd
0ea7  MSL Enterprises Corp.
0ea8  CenDyne, Inc.
0ead  Humax Co., Ltd
0eb0  NovaTech
	9020  NovaTech NV-902W
	9021  RT2573
0eb1  WIS Technologies, Inc.
	6666  WinFast WalkieTV TV Loader
	6668  WinFast WalkieTV TV Loader
	7007  WinFast WalkieTV WDM Capture
0eb2  Y-S Electronic Co., Ltd
0eb3  Saint Technology Corp.
0eb7  Endor AG
0eb8  Mettler Toledo
	2200  Ariva Scale
	f000  PS60 Scale
0ebb  Thermo Fisher Scientific
	0002  FT-IR Spectrometer
0ebe  VWeb Corp.
0ebf  Omega Technology of Taiwan, Inc.
0ec0  LHI Technology (China) Co., Ltd
0ec1  Abit Computer Corp.
0ec2  Sweetray Industrial, Ltd
0ec3  Axell Co., Ltd
0ec4  Ballracing Developments, Ltd
0ec5  GT Information System Co., Ltd
0ec6  InnoVISION Multimedia, Ltd
0ec7  Theta Link Corp.
	1008  So., Show 301 Digital Camera
0ecd  Lite-On IT Corp.
	1400  CD\RW 40X
	a100  LDW-411SX DVD/CD Rewritable Drive
0ece  TaiSol Electronics Co., Ltd
0ecf  Phogenix Imaging, LLC
0ed1  WinMaxGroup
	6660  Flash Disk 64M-C
	6680  Flash Disk 64M-B
	7634  MP3 Player
0ed2  Kyoto Micro Computer Co., Ltd
0ed3  Wing-Tech Enterprise Co., Ltd
0ed5  Fiberbyte
	e000  USB-inSync Device
	f000  Fiberbyte USB-inSync Device
	f201  Fiberbyte USB-inSync DAQ-2500X
0eda  Noriake Itron Corp.
0edf  e-MDT Co., Ltd
	2060  FID irock! 100 Series
0ee0  Shima Seiki Mfg., Ltd
0ee1  Sarotech Co., Ltd
0ee2  AMI Semiconductor, Inc.
0ee3  ComTrue Technology Corp.
	1000  Image Tank 1.5
0ee4  Sunrich Technology, Ltd
	0690  SATA 3 Adapter
0eee  Digital Stream Technology, Inc.
	8810  Mass Storage Drive
0eef  D-WAV Scientific Co., Ltd
	0001  eGalax TouchScreen
	0002  Touchscreen Controller(Professional)
	7200  Touchscreen Controller
	a802  eGalaxTouch EXC7920
0ef0  Hitachi Cable, Ltd
0ef1  Aichi Micro Intelligent Corp.
0ef2  I/O Magic Corp.
0ef3  Lynn Products, Inc.
0ef4  DSI Datotech
0ef5  PointChips
	2202  Flash Disk
	2366  Flash Disk
0ef6  Yield Microelectronics Corp.
0ef7  SM Tech Co., Ltd (Tulip)
0efd  Oasis Semiconductor
0efe  Wem Technology, Inc.
0f03  Unitek UPS Systems
	0001  Alpha 1200Sx
0f06  Visual Frontier Enterprise Co., Ltd
0f08  CSL Wire & Plug (Shen Zhen) Co.
0f0c  CAS Corp.
0f0d  Hori Co., Ltd
	0011  Real Arcade Pro 3
0f0e  Energy Full Corp.
0f11  LD Didactic GmbH
	1000  CASSY-S
	1010  Pocket-CASSY
	1020  Mobile-CASSY
	1080  Joule and Wattmeter
	1081  Digital Multimeter P
	1090  UMI P
	1100  X-Ray Apparatus
	1101  X-Ray Apparatus
	1200  VideoCom
	2000  COM3LAB
	2010  Terminal Adapter
	2020  Network Analyser
	2030  Converter Control Unit
	2040  Machine Test System
0f12  Mars Engineering Corp.
0f13  Acetek Technology Co., Ltd
0f14  Ingenico
	0012  Vital'Act 3S
0f18  Finger Lakes Instrumentation
	0002  CCD
	0006  Focuser
	0007  Filter Wheel
	000a  ProLine CCD
	000b  Color Filter Wheel 4
	000c  PDF2
	000d  Guider
0f19  Oracom Co., Ltd
0f1b  Onset Computer Corp.
0f1c  Funai Electric Co., Ltd
0f1d  Iwill Corp.
0f21  IOI Technology Corp.
0f22  Senior Industries, Inc.
0f23  Leader Tech Manufacturer Co., Ltd
0f24  Flex-P Industries, Snd., Bhd.
0f2d  ViPower, Inc.
0f2e  Geniality Maple Technology Co., Ltd
0f2f  Priva Design Services
0f30  Jess Technology Co., Ltd
	001c  PS3 Guitar Controller Dongle
	0110  Dual Analog Rumble Pad
	0111  Colour Rumble Pad
	0208  Xbox & PC Gamepad
0f31  Chrysalis Development
0f32  YFC-BonEagle Electric Co., Ltd
0f37  Kokuyo Co., Ltd
0f38  Nien-Yi Industrial Corp.
0f39  TG3 Electronics
	0876  Keyboard [87 Francium Pro]
	1086  DK2108SZ Keyboard [Ducky Zero]
0f3d  Airprime, Incorporated
	0112  CDMA 1xEVDO PC Card, PC 5220
0f41  RDC Semiconductor Co., Ltd
0f42  Nital Consulting Services, Inc.
0f44  Polhemus
	ef11  Patriot (firmware not loaded)
	ef12  Patriot
	ff11  Liberty (firmware not loaded)
	ff12  Liberty
0f4b  St. John Technology Co., Ltd
0f4c  WorldWide Cable Opto Corp.
0f4d  Microtune, Inc.
	1000  Bluetooth Dongle
0f4e  Freedom Scientific
0f52  Wing Key Electrical Co., Ltd
0f53  Dongguan White Horse Cable Factory, Ltd
0f54  Kawai Musical Instruments Mfg. Co., Ltd
	0101  MP6 Stage Piano
0f55  AmbiCom, Inc.
0f5c  Prairiecomm, Inc.
0f5d  NewAge International, LLC
	9455  Compact Drive
0f5f  Key Technology Corp.
0f60  NTK, Ltd
0f61  Varian, Inc.
0f62  Acrox Technologies Co., Ltd
	1001  Targus Mini Trackball Optical Mouse
0f63  LeapFrog Enterprises
	0010  Leapster Explorer
	0022  Leap Reader
	0500  Fly Fusion
	0600  Leap Port Turbo
	0700  POGO
	0800  Didj
	0900  TAGSchool
	0a00  Leapster 2
	0b00  Crammer
	0c00  Tag Jr
	0d00  My Pal Scout
	0e00  Tag32
	0f00  Tag64
	1000  Kiwi16
	1100  Leapster L2x
	1111  Fly Fusion
	1300  Didj UK/France (Leapster Advance)
0f68  Kobe Steel, Ltd
0f69  Dionex Corp.
0f6a  Vibren Technologies, Inc.
0f6e  INTELLIGENT SYSTEMS
	0100  GameBoy Color Emulator
	0201  GameBoy Advance Flash Gang Writer
	0202  GameBoy Advance Capture
	0300  Gamecube DOL Viewer
	0400  NDS Emulator
	0401  NDS UIC
	0402  NDS Writer
	0403  NDS Capture
	0404  NDS Emulator (Lite)
0f73  DFI
0f78  Guntermann & Drunck GmbH
0f7c  DQ Technology, Inc.
0f7d  NetBotz, Inc.
0f7e  Fluke Corp.
0f88  VTech Holdings, Ltd
	3012  RT2570
	3014  ZD1211B
0f8b  Yazaki Corp.
0f8c  Young Generation International Corp.
0f8d  Uniwill Computer Corp.
0f8e  Kingnet Technology Co., Ltd
0f8f  Soma Networks
0f97  CviLux Corp.
0f98  CyberBank Corp.
0f9c  Hyun Won, Inc.
	0301  M-Any Premium DAH-610 MP3/WMA Player
	0332  mobiBLU DAH-1200 MP3/Ogg Player
0f9e  Lucent Technologies
0fa3  Starconn Electronic Co., Ltd
0fa4  ATL Technology
0fa5  Sotec Co., Ltd
0fa7  Epox Computer Co., Ltd
0fa8  Logic Controls, Inc.
0faf  Winpoint Electronic Corp.
0fb0  Haurtian Wire & Cable Co., Ltd
0fb1  Inclose Design, Inc.
0fb2  Juan-Chern Industrial Co., Ltd
0fb6  Heber Ltd
	3fc3  Firefly X10i I/O Board (with firmware)
	3fc4  Firefly X10i I/O Board (without firmware)
0fb8  Wistron Corp.
	0002  eHome Infrared Receiver
0fb9  AACom Corp.
0fba  San Shing Electronics Co., Ltd
0fbb  Bitwise Systems, Inc.
0fc1  Mitac Internatinal Corp.
0fc2  Plug and Jack Industrial, Inc.
0fc5  Delcom Engineering
	1222  I/O Development Board
0fc6  Dataplus Supplies, Inc.
0fca  Research In Motion, Ltd.
	0001  Blackberry Handheld
	0004  Blackberry Handheld
	0006  Blackberry Pearl
	0008  Blackberry Pearl
	8001  Blackberry Handheld
	8004  Blackberry
	8007  Blackberry Handheld
	8010  Blackberry Playbook (Connect to Windows mode)
	8011  Blackberry Playbook (Connect to Mac mode)
	8020  Blackberry Playbook (CD-Rom mode)
0fce  Sony Ericsson Mobile Communications AB
	0076  W910i (Multimedia mode)
	00af  V640i Phone [PTP Camera]
	00d4  C902 [MTP]
	00d9  C702 Phone
	0112  W995 Walkman Phone
	014e  J108i Cedar (MTP mode)
	015a  Xperia Pro [Media Transfer Protocol]
	0166  Xperia Mini Pro
	0167  ST15i (Xperia mini)
	0169  Xperia S
	0172  Xperia P
	0177  Xperia Ion [Mass Storage]
	01bb  D5803 [Xperia Z3 Compact] (MTP mode)
	0dde  Xperia Mini Pro Bootloader
	1010  WMC Modem
	10af  V640i Phone [PictBridge]
	10d4  C902 Phone [PictBridge]
	2105  W715 Phone
	2137  Xperia X10 mini (USB debug)
	2138  Xperia X10 mini pro (Debug)
	2149  Xperia X8 (debug)
	214e  J108i Cedar (Windows-driver mode)
	3137  Xperia X10 mini
	3138  Xperia X10 mini pro
	3149  Xperia X8
	514f  Xperia arc S [Adb-Enable Mode]
	5169  Xperia S [Adb-Enable Mode]
	5177  Xperia Ion [Debug Mode]
	518c  C1605 [Xperia E dual] MTD mode
	614f  Xperia X12 (debug mode)
	6166  Xperia Mini Pro
	618c  C1605 [Xperia E dual] MSC mode
	715a  Xperia Pro [Tethering]
	7166  Xperia Mini Pro (Tethering mode)
	7177  Xperia Ion [Tethering]
	8004  9000 Phone [Mass Storage]
	adde  C2005 (Xperia M dual) in service mode
	d008  V800-Vodafone 802SE Phone
	d016  K750i Phone
	d017  K608i Phone
	d019  VDC EGPRS Modem
	d025  520 WMC Data Modem
	d028  W800i
	d038  W850i Phone
	d039  K800i (phone mode)
	d041  K510i Phone
	d042  W810i Phone
	d043  V630i Phone
	d046  K610i Phone
	d065  W960i Phone (PC Suite)
	d076  W910i (Phone mode)
	d089  W580i Phone (mass storage)
	d0a1  K810
	d0af  V640i Phone
	d0cf  MD300 Mobile Broadband Modem
	d0d4  C902 Phone [Modem]
	d0e1  MD400 Mobile Broadband Modem
	d12a  U100i Yari Phone
	d12e  Xperia X10
	d14e  J108i Cedar (modem mode)
	e000  K810 (PictBridge mode)
	e039  K800i (msc mode)
	e042  W810i Phone
	e043  V630i Phone [Mass Storage]
	e075  K850i
	e076  W910i (Mass storage)
	e089  W580i Phone
	e090  W200 Phone (Mass Storage)
	e0a1  K810 (Mass Storage mode)
	e0a3  W660i
	e0af  V640i Phone [Mass Storage]
	e0d4  C902 Phone [Mass Storage] 
	e0ef  C905 Phone [Mass Storage]
	e0f3  W595
	e105  W705
	e112  W995 Phone (Mass Storage)
	e12e  X10i Phone
	e133  Vivaz
	e14e  J108i Cedar (mass-storage mode)
	e14f  Xperia Arc/X12
	e15a  Xperia Pro [Mass Storage Class]
	e161  Xperia Ray
	e166  Xperia Mini Pro
	e167  XPERIA mini
	e19b  C2005 [Xperia M dual] (Mass Storage)
	f0fa  MN800 / Smartwatch 2 (DFU mode)
0fcf  Dynastream Innovations, Inc.
	1003  ANT Development Board
	1004  ANTUSB Stick
	1006  ANT Development Board
	1008  ANTUSB2 Stick
	1009  ANTUSB-m Stick
0fd0  Tulip Computers B.V.
0fd1  Giant Electronics Ltd.
0fd2  Seac Banche
	0001  RDS 6000
0fd4  Tenovis GmbH & Co., KG
0fd5  Direct Access Technology, Inc.
0fd9  Elgato Systems GmbH
	0011  EyeTV Diversity
	0018  EyeTV Hybrid
	0020  EyeTV DTT Deluxe
	0021  EyeTV DTT
	002a  EyeTV Sat
	002c  EyeTV DTT Deluxe v2
	0033  Video Capture
	0037  Video Capture v2
0fda  Quantec Networks GmbH
	0100  quanton flight control
0fdc  Micro Plus
0fde  Oregon Scientific
	ca01  WMRS200 weather station
	ca05  CM160
0fe0  Osterhout Design Group
	0100  Bluetooth Mouse
	0101  Bluetooth IMU
	0200  Bluetooth Keypad
0fe4  IN-Tech Electronics, Ltd
0fe5  Greenconn (U.S.A.), Inc.
0fe6  Kontron (Industrial Computer Source / ICS Advent)
	8101  DM9601 Fast Ethernet Adapter
	811e  Parallel Adapter
	9700  DM9601 Fast Ethernet Adapter
0fe9  DVICO
	4020  TViX M-6500
	db00  FusionHDTV DVB-T (MT352+LgZ201) (uninitialized)
	db01  FusionHDTV DVB-T (MT352+LgZ201) (initialized)
	db10  FusionHDTV DVB-T (MT352+Thomson7579) (uninitialized)
	db11  FusionHDTV DVB-T (MT352+Thomson7579) (initialized)
	db78  FusionHDTV DVB-T Dual Digital 4 (ZL10353+xc2028/xc3028) (initialized)
0fea  United Computer Accessories
0feb  CRS Electronic Co., Ltd
0fec  UMC Electronics Co., Ltd
0fed  Access Co., Ltd
0fee  Xsido Corp.
0fef  MJ Research, Inc.
0ff6  Core Valley Co., Ltd
0ff7  CHI SHING Computer Accessories Co., Ltd
0ffc  Clavia DMI AB
	0021  Nord Stage 2
0ffd  EarlySense
	ff00  OEM
0fff  Aopen, Inc.
1000  Speed Tech Corp.
	153b  TerraTec Electronic GmbH
1001  Ritronics Components (S) Pte., Ltd
1003  Sigma Corp.
	0003  SD14
	0100  SD9/SD10
1004  LG Electronics, Inc.
	1fae  U8120 3G Cellphone
	6000  Various Mobile Phones
	6005  T5100
	6018  GM360/GD510/GW520/KP501
	618e  Ally/Optimus One/Vortex (debug mode)
	618f  Ally/Optimus One
	61c5  P880 / Charge only
	61c6  Vortex (msc)
	61cc  Optimus S
	61da  G2 Android Phone [tethering mode]
	61f1  Optimus Android Phone [LG Software mode]
	61f9  Optimus (Various Models) MTP Mode
	61fc  Optimus 3
	61fe  Optimus Android Phone [USB tethering mode]
	6300  G2/Optimus Android Phone
	631c  G2/Optimus Android Phone [MTP mode]
	631d  Optimus Android Phone (Camera/PTP Mode)
	631e  G2/Optimus Android Phone [Camera/PTP mode]
	631f  Optimus Android Phone (Charge Mode)
	633e  G2 Android Phone [MTP mode]
	6344  G2 Android Phone [tethering mode]
	6356  Optimus Android Phone [Virtual CD mode]
	6800  CDMA Modem
	7000  LG LDP-7024D(LD)USB
	91c8  P880 / USB tethering
	a400  Renoir (KC910)
1005  Apacer Technology, Inc.
	1001  MP3 Player
	1004  MP3 Player
	1006  MP3 Player
	b113  Handy Steno/AH123 / Handy Steno 2.0/HT203
	b223  CD-RW + 6in1 Card Reader Digital Storage / Converter
1006  iRiver, Ltd.
	3001  iHP-100
	3002  iHP-120/140 MP3 Player
	3003  H320/H340
	3004  H340 (mtp)
1009  Emuzed, Inc.
	000e  eHome Infrared Receiver
	0013  Angel MPEG Device
	0015  Lumanate Wave PAL SECAM DVBT Device
	0016  Lumanate Wave NTSC/ATSC Combo Device
100a  AV Chaseway, Ltd
	2402  MP3 Player
	2404  MP3 Player
	2405  MP3 Player
	2406  MP3 Player
	a0c0  MP3 Player
100b  Chou Chin Industrial Co., Ltd
100d  Netopia, Inc.
	3342  Cayman 3352 DSL Modem
	3382  3380 Series Network Interface
	6072  DSL Modem
	9031  Motorola 802.11n Dualband USB Wireless Adapter
	9032  Motorola 802.11n 5G USB Wireless Adapter
	cb01  Cayman 3341 Ethernet DSL Router
1010  Fukuda Denshi Co., Ltd
1011  Mobile Media Tech.
	0001  AccFast Mp3
1012  SDKM Fibres, Wires & Cables Berhad
1013  TST-Touchless Sensor Technology AG
1014  Densitron Technologies PLC
1015  Softronics Pty., Ltd
1016  Xiamen Hung's Enterprise Co., Ltd
1017  Speedy Industrial Supplies, Pte., Ltd
1019  Elitegroup Computer Systems (ECS)
	0c55  Flash Reader, Desknote UCR-61S2B
	0f38  Infrared Receiver
1020  Labtec
	0006  Wireless Keyboard
	000a  Wireless Optical Mouse
	0106  Wireless Optical Mouse
1022  Shinko Shoji Co., Ltd
1025  Hyper-Paltek
	005e  USB DVB-T device
	005f  USB DVB-T device
	0300  MP3 Player
	0350  MP3 Player
1026  Newly Corp.
1027  Time Domain
1028  Inovys Corp.
1029  Atlantic Coast Telesys
102a  Ramos Technology Co., Ltd
102b  Infotronic America, Inc.
102c  Etoms Electronics Corp.
	6151  Q-Cam Sangha CIF
	6251  Q-Cam VGA
102d  Winic Corp.
1031  Comax Technology, Inc.
1032  C-One Technology Corp.
1033  Nucam Corp.
	0068  3,5'' HDD case MD-231
1038  SteelSeries ApS
	0100  Ideazon Zboard
	1361  Ideazon Sensei
1039  devolo AG
	0824  1866 802.11bg [Texas Instruments TNETW1450]
	2140  dsl+ 1100 duo
103a  PSA
	f000  Actia Evo XS
103d  Stanton
	0100  ScratchAmp
	0101  ScratchAmp
1043  iCreate Technologies Corp.
	160f  Wireless Network Adapter
	4901  AV-836 Video Capture Device
	8006  Flash Disk 32-256 MB
	8012  Flash Disk 256 MB
1044  Chu Yuen Enterprise Co., Ltd
	7001  Gigabyte U7000 DVB-T tuner
	7002  Gigabyte U8000 DVB-T tuner
	7004  Gigabyte U7100 DVB-T tuner
	7005  Gigabyte U7200 DVB-T tuner [AF9035]
	7006  Gigabyte U6000 DVB-T tuner [em2863]
	8001  GN-54G
	8002  GN-BR402W
	8003  GN-WLBM101
	8004  GN-WLBZ101 802.11b Adapter
	8005  GN-WLBZ201 802.11b Adapter
	8006  GN-WBZB-M 802.11b Adapter
	8007  GN-WBKG
	8008  GN-WB01GS
	800a  GN-WI05GS
	800b  GN-WB30N 802.11n WLAN Card
	800c  GN-WB31N 802.11n USB WLAN Card
	800d  GN-WB32L 802.11n USB WLAN Card
1046  Winbond Electronics Corp. [hex]
	6694  Generic W6694 USB
	8901  Bluetooth Device
	9967  W9967CF/W9968CF Webcam IC
1048  Targus Group International
	2010  4-Port hub
104b  Mylex / Buslogic
104c  AMCO TEC International, Inc.
104d  Newport Corporation
	1003  Model-52 LED Light Source Power Supply and Driver
104f  WB Electronics
	0001  Infinity Phoenix
	0002  Smartmouse
	0003  FunProgrammer
	0004  Infinity Unlimited
	0006  Infinity Smart
	0007  Infinity Smart module
	0008  Infinity CryptoKey
	0009  RE-BL PlayStation 3 IR-to-Bluetooth converter
1050  Yubico.com
	0010  Yubikey (v1 or v2)
	0110  Yubikey NEO(-N) OTP
	0111  Yubikey NEO(-N) OTP+CCID
	0112  Yubikey NEO(-N) CCID
	0113  Yubikey NEO(-N) U2F
	0114  Yubikey NEO(-N) OTP+U2F
	0115  Yubikey NEO(-N) U2F+CCID
	0116  Yubikey NEO(-N) OTP+U2F+CCID
	0120  Yubikey Touch U2F Security Key
	0200  Gnubby U2F
	0211  Gnubby
	0401  Yubikey 4 OTP
	0402  Yubikey 4 U2F
	0403  Yubikey 4 OTP+U2F
	0404  Yubikey 4 CCID
	0405  Yubikey 4 OTP+CCID
	0406  Yubikey 4 U2F+CCID
	0407  Yubikey 4 OTP+U2F+CCID
	0410  Yubikey plus OTP+U2F
1053  Immanuel Electronics Co., Ltd
1054  BMS International Beheer N.V.
	5004  DSL 7420 Loader
	5005  DSL 7420 LAN Modem
1055  Complex Micro Interconnection Co., Ltd
1056  Hsin Chen Ent Co., Ltd
1057  ON Semiconductor
1058  Western Digital Technologies, Inc.
	0200  FireWire USB Combo
	0400  External HDD
	0500  hub
	0701  WD Passport (WDXMS)
	0702  WD Passport (WDXMS)
	0704  My Passport Essential (WDME)
	0705  My Passport Elite (WDML)
	070a  My Passport Essential (WDBAAA), My Passport for Mac (WDBAAB), My Passport Essential SE (WDBABM), My Passport SE for Mac (WDBABW)
	070b  My Passport Elite (WDBAAC)
	070c  My Passport Studio (WDBAAE)
	071a  My Passport Essential (WDBAAA)
	071d  My Passport Studio (WDBALG)
	0730  My Passport Essential (WDBACY)
	0732  My Passport Essential SE (WDBGYS)
	0740  My Passport Essential (WDBACY)
	0741  My Passport Ultra
	0742  My Passport Essential SE (WDBGYS)
	0748  My Passport (WDBKXH, WDBY8L)
	07a8  My Passport (WDBBEP), My Passport for Mac (WDBLUZ)
	0810  My Passport Ultra (WDBZFP)
	0820  My Passport Ultra (WDBMWV, WDBZFP)
	0830  My Passport Ultra (WDBZFP)
	0900  MyBook Essential External HDD
	0901  My Book Essential Edition (Green Ring) (WDG1U)
	0902  My Book Pro Edition (WDG1T)
	0903  My Book Premium Edition
	0910  My Book Essential Edition (Green Ring) (WDG1U)
	1001  Elements Desktop (WDE1U)
	1003  WD Elements Desktop (WDE1UBK)
	1010  Elements Portable (WDBAAR)
	1021  Elements Desktop (WDBAAU)
	1023  Elements SE Portable (WDBABV)
	1042  Elements SE Portable (WDBPCK)
	1048  Elements Portable (WDBU6Y)
	107c  Elements Desktop (WDBWLG)
	10a2  Elements SE Portable (WDBPCK)
	10a8  Elements Portable (WDBUZG)
	10b8  Elements Portable (WDBU6Y, WDBUZG)
	1100  My Book Essential Edition 2.0 (WDH1U)
	1102  My Book Home Edition (WDH1CS)
	1103  My Book Studio
	1104  My Book Mirror Edition (WDH2U)
	1105  My Book Studio II
	1110  My Book Essential (WDBAAF), My Book for Mac (WDBAAG)
	1111  My Book Elite (WDBAAH)
	1112  My Book Studio (WDBAAJ), My Book Studio LX (WDBACH)
	1123  My Book 3.0 (WDBABP)
	1130  My Book Essential (WDBACW)
	1140  My Book Essential (WDBACW)
	1230  My Book (WDBFJK0030HBK)
1059  Giesecke & Devrient GmbH
	000b  StarSign Bio Token 3.0
105b  Foxconn International, Inc.
	e065  BCM43142A0 Bluetooth module
105c  Hong Ji Electric Wire & Cable (Dongguan) Co., Ltd
105d  Delkin Devices, Inc.
105e  Valence Semiconductor Design, Ltd
105f  Chin Shong Enterprise Co., Ltd
1060  Easthome Industrial Co., Ltd
1063  Motorola Electronics Taiwan, Ltd [hex]
	1555  MC141555 Hub
	4100  SB4100 USB Cable Modem
1065  CCYU Technology
	0020  USB-DVR2 Dev Board
	2136  EasyDisk ED1064
106a  Loyal Legend, Ltd
106c  Curitel Communications, Inc.
	1101  CDMA 2000 1xRTT USB modem (HX-550C)
	1102  Packet Service
	1103  Packet Service Diagnostic Serial Port (WDM)
	1104  Packet Service Diagnostic Serial Port (WDM)
	1105  Composite Device
	1106  Packet Service Diagnostic Serial Port (WDM)
	1301  Composite Device
	1302  Packet Service Diagnostic Serial Port (WDM)
	1303  Packet Service
	1304  Packet Service
	1401  Composite Device
	1402  Packet Service
	1403  Packet Service Diagnostic Serial Port (WDM)
	1501  Packet Service
	1502  Packet Service Diagnostic Serial Port (WDM)
	1503  Packet Service
	1601  Packet Service
	1602  Packet Service Diagnostic Serial Port (WDM)
	1603  Packet Service
	2101  AudioVox 8900 Cell Phone
	2102  Packet Service
	2103  Packet Service Diagnostic Serial Port (WDM)
	2301  Packet Service
	2302  Packet Service Diagnostic Serial Port (WDM)
	2303  Packet Service
	2401  Packet Service Diagnostic Serial Port (WDM)
	2402  Packet Service
	2403  Packet Service Diagnostic Serial Port (WDM)
	2501  Packet Service
	2502  Packet Service Diagnostic Serial Port (WDM)
	2503  Packet Service
	2601  Packet Service
	2602  Packet Service Diagnostic Serial Port (WDM)
	2603  Packet Service
	3701  Broadband Wireless modem
	3702  Pantech PX-500
	3714  PANTECH USB MODEM [UM175]
	3716  UMW190 Modem
	3721  Option Beemo (GI0801) LTE surfstick
	3b14  Option Beemo (GI0801) LTE surfstick
	3eb4  Packet Service Diagnostic Serial Port (WDM)
	4101  Packet Service Diagnostic Serial Port (WDM)
	4102  Packet Service
	4301  Composite Device
	4302  Packet Service Diagnostic Serial Port (WDM)
	4401  Composite Device
	4402  Packet Service
	4501  Packet Service
	4502  Packet Service Diagnostic Serial Port (WDM)
	4601  Composite Device
	4602  Packet Service Diagnostic Serial Port (WDM)
	5101  Packet Service
	5102  Packet Service Diagnostic Serial Port (WDM)
	5301  Packet Service Diagnostic Serial Port (WDM)
	5302  Packet Service
	5401  Packet Service
	5402  Packet Service Diagnostic Serial Port (WDM)
	5501  Packet Service Diagnostic Serial Port (WDM)
	5502  Packet Service
	5601  Packet Service Diagnostic Serial Port (WDM)
	5602  Packet Service
	7101  Composite Device
	7102  Packet Service
	a000  Packet Service
	a001  Packet Service Diagnostic Serial Port (WDM)
	c100  Packet Service
	c200  Packet Service
	c500  Packet Service Diagnostic Serial Port (WDM)
	e200  Packet Service
106d  San Chieh Manufacturing, Ltd
106e  ConectL
106f  Money Controls
	0009  CT10x Coin Transaction
	000a  CR10x Coin Recycler
	000c  Xchange
1076  GCT Semiconductor, Inc.
	0031  Bluetooth Device
	0032  Bluetooth Device
	8002  LU150 LTE Modem [Yota LU150]
107b  Gateway, Inc.
	3009  eHome Infrared Transceiver
	55b2  WBU-110 802.11b Wireless Adapter [Intersil PRISM 3]
	55f2  WGU-210 802.11g Adapter [Intersil ISL3886]
107d  Arlec Australia, Ltd
107e  Midoriya Electric Co., Ltd
107f  KidzMouse, Inc.
1082  Shin-Etsukaken Co., Ltd
1083  Canon Electronics, Inc.
	161b  DR-2010C Scanner
	162c  P-150 Scanner
1084  Pantech Co., Ltd
108a  Chloride Power Protection
108b  Grand-tek Technology Co., Ltd
	0005  HID Keyboard/Mouse PS/2 Translator
108c  Robert Bosch GmbH
108e  Lotes Co., Ltd.
1099  Surface Optics Corp.
109a  DATASOFT Systems GmbH
109b  Hisense
	9118  Medion P4013 Mobile
109f  eSOL Co., Ltd
	3163  Trigem Mobile SmartDisplay84
	3164  Trigem Mobile SmartDisplay121
10a0  Hirotech, Inc.
10a3  Mitsubishi Materials Corp.
10a9  SK Teletech Co., Ltd
	1102  Sky Love Actually IM-U460K
	1104  Sky Vega IM-A650S
	1105  VEGA Android composite
	1106  VEGA Android composite
	1107  VEGA Android composite
	1108  VEGA Android composite
	1109  VEGA Android composite
	6021  SIRIUS alpha
	6031  Pantech Android composite
	6032  Pantech Android composite
	6033  Pantech Android composite
	6034  Pantech Android composite
	6035  Pantech Android composite
	6036  Pantech Android composite
	6037  Pantech Android composite
	6050  Pantech Android composite
	6051  Pantech Android composite
	6052  Pantech Android composite
	6053  Pantech Android composite
	6054  Pantech Android composite
	6055  Pantech Android composite
	6056  Pantech Android composite
	6057  Pantech Android composite
	6058  Pantech Android composite
	6059  Pantech Android composite
	6080  MHS291LVW LTE Modem [Verizon Jetpack 4G LTE Mobile Hotspot MHS291L] (Zero CD Mode)
	6085  MHS291LVW LTE Modem [Verizon Jetpack 4G LTE Mobile Hotspot MHS291L] (Modem Mode)
	7031  Pantech Android composite
	7032  Pantech Android composite
	7033  Pantech Android composite
	7034  Pantech Android composite
	7035  Pantech Android composite
	7036  Pantech Android composite
	7037  Pantech Android composite
10aa  Cables To Go
10ab  USI Co., Ltd
	1002  Bluetooth Device
	1003  BC02-EXT in DFU
	1005  Bluetooth Adptr
	1006  BC04-EXT in DFU
	10c5  Sony-Ericsson / Samsung DataCable
10ac  Honeywell, Inc.
10ae  Princeton Technology Corp.
10af  Liebert Corp.
	0000  UPS
	0001  PowerSure PSA UPS
	0002  PowerSure PST UPS
	0003  PowerSure PSP UPS
	0004  PowerSure PSI UPS
	0005  UPStation GXT 2U UPS
	0006  UPStation GXT UPS
	0007  Nfinity Power Systems UPS
	0008  PowerSure Interactive UPS
10b5  Comodo (PLX?)
	9060  Test Board
10b8  DiBcom
	0bb8  DiBcom USB DVB-T reference design (MOD300) (cold)
	0bb9  DiBcom USB DVB-T reference design (MOD300) (warm)
	0bc6  DiBcom USB2.0 DVB-T reference design (MOD3000P) (cold)
	0bc7  DiBcom USB2.0 DVB-T reference design (MOD3000P) (warm)
10bb  TM Technology, Inc.
10bc  Dinging Technology Co., Ltd
10bd  TMT Technology, Inc.
	1427  Ethernet
10bf  SmartHome
	0001  SmartHome PowerLinc
10c3  Universal Laser Systems, Inc.
	00a4  ULS PLS Series Laser Engraver Firmware Loader
	00a5  ULS Print Support
10c4  Cygnal Integrated Products, Inc.
	0002  F32x USBXpress Device
	0003  CommandIR
	8030  K4JRG Ham Radio devices
	8044  USB Debug Adapter
	804e  Software Bisque Paramount ME
	80a9  CP210x to UART Bridge Controller
	80ca  ATM2400 Sensor Device
	813f  tams EasyControl
	8149  West Mountain Radio Computerized Battery Analyzer
	814a  West Mountain Radio RIGblaster P&P
	814b  West Mountain Radio RIGtalk
	818a  Silicon Labs FM Radio Reference Design
	81e8  Zephyr BioHarness
	8460  Sangoma Wanpipe VoiceTime
	8461  Sangoma U100
	8477  Balluff RFID Reader
	8496  SiLabs Cypress FW downloader
	8497  SiLabs Cypress EVB
	8605  dilitronics ESoLUX solar lighting controller
	86bc  C8051F34x AudioDelay [AD-340]
	8789  C8051F34x Extender & EDID MGR [EMX-DVI]
	87be  C8051F34x HDMI Audio Extractor [EMX-HD-AUD]
	8863  C8051F34x Bootloader
	8897  C8051F38x HDMI Splitter [UHBX]
	8918  C8051F38x HDMI Audio Extractor [VSA-HA-DP]
	8973  C8051F38x HDMI Extender [UHBX-8X]
	89e1  C8051F38x HDMI Extender [UHBX-SW3-WP]
	ea60  CP210x UART Bridge / myAVR mySmartUSB light
	ea61  CP210x UART Bridge
	ea70  CP210x UART Bridge
	ea80  CP210x UART Bridge
10c5  Sanei Electric, Inc.
	819a  FM Radio
10c6  Intec, Inc.
10cb  Eratech
10cc  GBM Connector Co., Ltd
	1101  MP3 Player
10cd  Kycon, Inc.
10ce  Silicon Labs
	000e  Shinko/Sinfonia CHC-S2145
	ea6a  MobiData EDGE USB Modem
10cf  Velleman Components, Inc.
	2011  R-Engine MPEG2 encoder/decoder
	5500  8055 Experiment Interface Board (address=0)
	5501  8055 Experiment Interface Board (address=1)
	5502  8055 Experiment Interface Board (address=2)
	5503  8055 Experiment Interface Board (address=3)
10d1  Hottinger Baldwin Measurement
	0101  USB-Module for Spider8, CP32
	0202  CP22 - Communication Processor
	0301  CP42 - Communication Processor
10d2  RayComposer - R. Adams
	5243  RayComposer
10d4  Man Boon Manufactory, Ltd
10d5  Uni Class Technology Co., Ltd
	0004  PS/2 Converter
	5552  KVM Human Interface Composite Device (Keyboard/Mouse ports)
	55a2  2Port KVMSwitcher
10d6  Actions Semiconductor Co., Ltd
	0c02  BioniQ 1001 Tablet
	1000  MP3 Player
	1100  MPMan MP-Ki 128 MP3 Player/Recorder
	1101  D-Wave 2GB MP4 Player / AK1025 MP3/MP4 Player
	2200  Acer MP-120 MP3 player
	8888  ADFU Device
	ff51  ADFU Device
	ff61  MP4 Player
	ff66  Craig 2GB MP3/Video Player
10de  Authenex, Inc.
10df  In-Win Development, Inc.
	0500  iAPP CR-e500 Card reader
10e0  Post-Op Video, Inc.
10e1  CablePlus, Ltd
10e2  Nada Electronics, Ltd
10ec  Vast Technologies, Inc.
10f0  Nexio Co., Ltd
	2002  iNexio Touchscreen controller
10f1  Importek
	1a08  Internal Webcam
	1a1e  Laptop Integrated Webcam 1.3M
	1a2a  Laptop Integrated Webcam
10f5  Turtle Beach
	0200  Audio Advantage Roadie
10fb  Pictos Technologies, Inc.
10fd  Anubis Electronics, Ltd
	7e50  FlyCam Usb 100
	804d  Typhoon Webshot II Webcam [zc0301]
	8050  FlyCAM-USB 300 XP2
	de00  WinFast WalkieTV WDM Capture Driver.
10fe  Thrane & Thrane
	000c  TT-3750 BGAN-XL Radio Module
1100  VirTouch, Ltd
	0001  VTPlayer VTP-1 Braille Mouse
1101  EasyPass Industrial Co., Ltd
	0001  FSK Electronics Super GSM Reader
1108  Brightcom Technologies, Ltd
110a  Moxa Technologies Co., Ltd.
	1250  UPort 1250 2-Port RS-232/422/485
	1251  UPort 1250I 2-Port RS-232/422/485 with Isolation
	1410  UPort 1410 4-Port RS-232
	1450  UPort 1450 4-Port RS-232/422/485
	1451  UPort 1450I 4-Port RS-232/422/485 with Isolation
	1613  UPort 1610-16 16-Port RS-232
	1618  UPort 1610-8 8-Port RS-232
	1653  UPort 1650-16 16-Port RS-232/422/485
	1658  UPort 1650-8 8-Port RS-232/422/485
1110  Analog Devices Canada, Ltd (Allied Telesyn)
	5c01  Huawei MT-882 Remote NDIS Network Device
	6489  ADSL ETH/USB RTR
	9000  ADSL LAN Adapter
	9001  ADSL Loader
	900f  AT-AR215 DSL Modem
	9010  AT-AR215 DSL Modem
	9021  ADSL WAN Adapter
	9022  ADSL Loader
	9023  ADSL WAN Adapter
	9024  ADSL Loader
	9031  ADSL LAN Adapter
	9032  ADSL Loader
1111  Pandora International Ltd.
	8888  Evolution Device
1112  YM ELECTRIC CO., Ltd
1113  Medion AG
	a0a2  Active Sync device
111e  VSO Electric Co., Ltd
112a  RedRat
	0001  RedRat3 IR Transceiver
	0005  RedRat3II IR Transceiver
112e  Master Hill Electric Wire and Cable Co., Ltd
112f  Cellon International, Inc.
1130  Tenx Technology, Inc.
	0001  BlyncLight
	0002  iBuddy
	0202  Rocket Launcher
	6604  MCE IR-Receiver
	660c  Foot Pedal/Thermometer
	6806  Keychain photo frame
	c301  Digital Photo viewer [Wallet Pix]
	f211  TP6911 Audio Headset
1131  Integrated System Solution Corp.
	1001  KY-BT100 Bluetooth Adapter
	1002  Bluetooth Device
	1003  Bluetooth Device
	1004  Bluetooth Device
1132  Toshiba Corp., Digital Media Equipment [hex]
	4331  PDR-M4/M5/M70 Digital Camera
	4332  PDR-M60 Digital Camera
	4333  PDR-M2300/PDR-M700
	4334  PDR-M65
	4335  PDR-M61
	4337  PDR-M11
	4338  PDR-M25
1136  CTS Electronincs
	3131  CTS LS515
113c  Arin Tech Co., Ltd
113d  Mapower Electronics Co., Ltd
1141  V One Multimedia, Pte., Ltd
1142  CyberScan Technologies, Inc.
	0709  Cyberview High Speed Scanner
1145  Japan Radio Company
	0001  AirH PHONE AH-J3001V/J3002V
1146  Shimane SANYO Electric Co., Ltd.
1147  Ever Great Electric Wire and Cable Co., Ltd
114b  Sphairon Access Systems GmbH
	0110  Turbolink UB801R WLAN Adapter
	0150  Turbolink UB801RE Wireless 802.11g 54Mbps Network Adapter [RTL8187]
114c  Tinius Olsen Testing Machine Co., Inc.
114d  Alpha Imaging Technology Corp.
114f  Wavecom
	1234  Fastrack Xtend FXT001 Modem
115b  Salix Technology Co., Ltd.
1162  Secugen Corp.
1163  DeLorme Publishing, Inc.
	0100  Earthmate GPS (orig)
	0200  Earthmate GPS (LT-20, LT-40)
	2020  Earthmate GPS (PN-40)
1164  YUAN High-Tech Development Co., Ltd
	0300  ELSAVISION 460D
	0601  Analog TV Tuner
	0900  TigerBird BMP837 USB2.0 WDM Encoder
	0bc7  Digital TV Tuner
	521b  MC521A mini Card ATSC Tuner
	6601  Digital TV Tuner Card [RTL2832U]
1165  Telson Electronics Co., Ltd
1166  Bantam Interactive Technologies
1167  Salient Systems Corp.
1168  BizConn International Corp.
116e  Gigastorage Corp.
116f  Silicon 10 Technology Corp.
	0005  Flash Card Reader
	c108  Flash Card Reader
	c109  Flash Card Reader
1175  Shengyih Steel Mold Co., Ltd
117d  Santa Electronic, Inc.
117e  JNC, Inc.
1182  Venture Corp., Ltd
1183  Compaq Computer Corp. [hex] (Digital Dream ??)
	0001  DigitalDream l'espion XS
	19c7  ISDN TA
	4008  56k FaxModem
	504a  PJB-100 Personal Jukebox
1184  Kyocera Elco Corp.
1188  Bloomberg L.P.
1189  Acer Communications & Multimedia
	0893  EP-1427X-2 Ethernet Adapter [Acer]
118f  You Yang Technology Co., Ltd
1190  Tripace
1191  Loyalty Founder Enterprise Co., Ltd
1196  Yankee Robotics, LLC
	0010  Trifid Camera without code
	0011  Trifid Camera
1197  Technoimagia Co., Ltd
1198  StarShine Technology Corp.
1199  Sierra Wireless, Inc.
	0019  AC595U
	0021  AC597E
	0024  MC5727 CDMA modem
	0110  Composite Device
	0112  CDMA 1xEVDO PC Card, AirCard 580
	0120  AC595U
	0218  MC5720 Wireless Modem
	6467  MP Series Network Adapter
	6468  MP Series Network Adapter
	6469  MP Series Network Adapter
	6802  MC8755 Device
	6803  MC8765 Device
	6804  MC8755 Device
	6805  MC8765 Device
	6812  MC8775 Device
	6820  AC875 Device
	6832  MC8780 Device
	6833  MC8781 Device
	683a  MC8785 Device
	683c  Mobile Broadband 3G/UMTS (MC8790 Device)
	6850  AirCard 880 Device
	6851  AirCard 881 Device
	6852  AirCard 880E Device
	6853  AirCard 881E Device
	6854  AirCard 885 Device
	6856  ATT "USB Connect 881"
	6870  MC8780 Device
	6871  MC8781 Device
	6893  MC8777 Device
	68a3  MC8700 Modem
	68aa  4G LTE adapter
	9000  Gobi 2000 Wireless Modem (QDL mode)
	9001  Gobi 2000 Wireless Modem
	9002  Gobi 2000 Wireless Modem
	9003  Gobi 2000 Wireless Modem
	9004  Gobi 2000 Wireless Modem
	9005  Gobi 2000 Wireless Modem
	9006  Gobi 2000 Wireless Modem
	9007  Gobi 2000 Wireless Modem
	9008  Gobi 2000 Wireless Modem
	9009  Gobi 2000 Wireless Modem
	900a  Gobi 2000 Wireless Modem
	9055  Gobi 9x15 Multimode 3G/4G LTE Modem (NAT mode)
	9057  Gobi 9x15 Multimode 3G/4G LTE Modem (IP passthrough mode)
119a  ZHAN QI Technology Co., Ltd
119b  ruwido austria GmbH
	0400  Infrared Keyboard V2.01
11a0  Chipcon AS
	eb11  CC2400EB 2.0 ZigBee Sniffer
11a3  Technovas Co., Ltd
	8031  MP3 Player
	8032  MP3 Player
11aa  GlobalMedia Group, LLC
	1518  iREZ K2
11ab  Exito Electronics Co., Ltd
11ac  Nike
	6565  FuelBand
11b0  ATECH FLASH TECHNOLOGY
	6208  PRO-28U
11be  R&D International NV
	f0a0  Martin Maxxyz DMX
11c5  Inmax
	0521  IMT-0521 Smartcard Reader
11ca  VeriFone Inc
	0207  PIN Pad VX 810
	0220  PIN Pad VX 805
11db  Topfield Co., Ltd.
	1000  PVR
	1100  PVR
11e6  K.I. Technology Co. Ltd.
11f5  Siemens AG
	0001  SX1
	0003  Mobile phone USB cable
	0004  X75
	0005  SXG75/EF81
	0008  UMTS/HSDPA Data Card
	0101  RCU Connect
11f6  Prolific
	2001  Willcom WSIM
11f7  Alcatel (?)
	02df  Serial cable (v2) for TD-10 Mobile Phone
1203  TSC Auto ID Technology Co., Ltd
	0140  TTP-245C
1209  InterBiometrics
	1001  USB Hub
	1002  USB Relais
	1003  IBSecureCam-P
	1004  IBSecureCam-O
	1005  IBSecureCam-N
	1006  Mini IO-Board
	2000  Zygmunt Krynicki Lantern Brightness Sensor
	2048  Housedillon.com MRF49XA Transciever
	2222  LabConnect Signalgenerator
	2300  Keyboardio Keyboardio Model 01 Bootloader
	2301  Keyboardio Keyboardio Model 01
	2337  /Dev or SlashDev /Net
	3000  lloyd3000
	3333  LabConnect Digitalnetzteil
	5222  telavivmakers attami
	5a22  ikari_01 sd2snes
	7bd0  pokey9000 Tiny Bit Dingus
	abd0  tibounise ADB converter
	beef  Modal MC-USB
	c0f5  unethi PERswitch
	ca1c  KnightOS Hub
	ca1d  KnightOS MTP Device
	cafe  ii iigadget
	dada  Rebel Technology OWL
	dead  chaosfield.at AVR-Ruler
	fa11  moonglow OpenXHC
	feed  ProgramGyar AVR-IR Sender
120e  Hudson Soft Co., Ltd
120f  Magellan
	524e  RoadMate 1475T
	5260  Triton Handheld GPS Receiver (300/400/500/1500/2000)
1210  DigiTech
	0016  RP500 Guitar Multi-Effects Processor
	001b  RP155 Guitar Multi-Effects Processor
	001c  RP255 Guitar Multi-Effects Processor
121e  Jungsoft Co., Ltd
	3403  Muzio JM250 Audio Player
1221  Unknown manufacturer
	3234  Disk (Thumb drive)
1223  SKYCABLE ENTERPRISE. CO., LTD.
1228  Datapaq Limited
	0012  Q18 Data Logger
	0015  TPaq21/MPaq21 Datalogger
	584c  XL2 Logger
1230  Chipidea-Microelectronica, S.A.
1233  Denver Electronics
	5677  FUSB200 mp3 player
1234  Brain Actuated Technologies
	0000  Neural Impulse Actuator Prototype 1.0 [NIA]
	4321  Human Interface Device
	ed02  Emotiv EPOC Developer Headset Wireless Dongle
1235  Focusrite-Novation
	0001  ReMOTE Audio/XStation First Edition
	0002  Speedio
	0003  RemoteSL + ZeroSL
	0004  ReMOTE LE
	0005  XIOSynth [First Edition]
	0006  XStation
	0007  XIOSynth
	0008  ReMOTE SL Compact
	0009  nIO
	000a  Nocturn
	000b  ReMOTE SL MkII
	000c  ZeRO MkII
	000e  Launchpad
	0010  Saffire 6
	0011  Ultranova
	0012  Nocturn Keyboard
	0013  VRM Box
	0014  VRM Box Audio Class (2-out)
	0015  Dicer
	0016  Ultranova
	0018  Twitch
	0019  Impulse 25
	001a  Impulse 49
	001b  Impulse 61
	4661  ReMOTE25
	8000  Scarlett 18i6
	8002  Scarlett 8i6
	8006  Focusrite Scarlett 2i2
	8008  Saffire 6
	800a  Scarlett 2i4
	800c  Scarlett 18i20
	800e  iTrack Solo
	8010  Forte
	8012  Scarlett 6i6
	8014  Scarlett 18i8
1241  Belkin
	0504  Wireless Trackball Keyboard
	1111  Mouse
	1122  Typhoon Stream Optical Mouse USB+PS/2
	1155  Memorex Optical ScrollPro Mouse SE MX4600
	1166  MI-2150 Trust Mouse
	1177  Mouse [HT82M21A]
	1503  Keyboard
	1603  Keyboard
	f767  Keyboard
124a  AirVast
	168b  PRISM3 WLAN Adapter
	4017  PC-Chips 802.11b Adapter
	4023  WM168g 802.11bg Wireless Adapter [Intersil ISL3886]
	4025  IOGear GWU513 v2 802.11bg Wireless Adapter [Intersil ISL3887]
124b  Nyko (Honey Bee)
	4d01  Airflo EX Joystick
124c  MXI - Memory Experts International, Inc.
	3200  Stealth MXP 1GB
125c  Apogee Inc.
	0010  Alta series CCD
125f  A-DATA Technology Co., Ltd.
	312a  Superior S102
	312b  Superior S102 Pro
	a15a  DashDrive Durable HD710 portable HDD various size
	a22a  DashDrive Elite HE720 500GB
	a91a  Portable HDD CH91
	c08a  C008 Flash Drive
	c81a  Flash drive
	c93a  4GB Pen Drive
	c96a  C906 Flash Drive
	cb10  Dash Drive UV100
1260  Standard Microsystems Corp.
	ee22  SMC2862W-G v3 EZ Connect 802.11g Adapter [Intersil ISL3887]
1264  Covidien Energy-based Devices
1266  Pirelli Broadband Solutions
	6302  Fastweb DRG A226M ADSL Router
1267  Logic3 / SpectraVideo plc
	0103  G-720 Keyboard
	0201  A4Tech SWOP-3 Mouse
	0210  LG Optical Mouse 3D-310
	a001  JP260 PC Game Pad
	c002  Wireless Optical Mouse
126c  Aristocrat Technologies
126d  Bel Stewart
126e  Strobe Data, Inc.
126f  TwinMOS
	0163  Storage device (2gB thumb drive)
	1325  Mobile Disk
	2168  Mobile Disk III
	a006  G240 802.11bg
1274  Ensoniq
1275  Xaxero Marine Software Engineering, Ltd.
	0002  WeatherFax 2000 Demodulator
	0080  SkyEye Weather Satellite Receiver
1278  Starlight Xpress
	0105  SXV-M5
	0107  SXV-M7
	0109  SXV-M9
	0110  SXVF-H16
	0115  SXVF-H5
	0119  SXV-H9
	0135  SXVF-H35
	0136  SXVF-H36
	0200  SXV interface for paraller MX cameras
	0305  SXV-M5C
	0307  SXV-M7C
	0319  SXV-H9C
	0325  SXV-M25C
	0326  SXVR-M26C
	0507  Lodestar autoguider
	0517  CoStar
1283  zebris Medical GmbH
	0100  USB-RS232 Adaptor
	0110  CMS20
	0111  CMS 10
	0112  CMS 05
	0114  ARCUS digma PC-Interface
	0115  SAM Axioquick recorder
	0116  SAM Axioquick recorder
	0120  emed-X
	0121  emed-AT
	0130  PDM
	0150  CMS10GI (Golf)
1286  Marvell Semiconductor, Inc.
	00bc  Marvell JTAG Probe
	1fab  88W8338 [Libertas] 802.11g
	2001  88W8388 802.11a/b/g WLAN
	2006  88W8362 802.11n WLAN
	8001  BLOB boot loader firmware
1291  Qualcomm Flarion Technologies, Inc. / Leadtek Research, Inc.
	0010  FDM 2xxx Flash-OFDM modem
	0011  LR7F06/LR7F14 Flash-OFDM modem
1292  Innomedia
	0258  Creative Labs VoIP Blaster
1293  Belkin Components [hex]
	0002  F5U002 Parallel Port [uss720]
	2101  104-key keyboard
1294  RISO KAGAKU CORP.
	1320  Webmail Notifier
129b  CyberTAN Technology
	160b  Siemens S30853-S1031-R351 802.11g Wireless Adapter [Atheros AR5523]
	160c  Siemens S30853-S1038-R351 802.11g Wireless Adapter [Atheros AR5523]
	1666  TG54USB 802.11bg
	1667  802.11bg
	1828  Gigaset USB Adapter 300
12a7  Trendchip Technologies Corp.
12ab  Honey Bee Electronic International Ltd.
12b8  Zhejiang Xinya Electronic Technology Co., Ltd.
12b9  E28
12ba  Licensed by Sony Computer Entertainment America
	00ff  Rocksmith Guitar Adapter
	0100  RedOctane Guitar for PlayStation(R)3
	0120  RedOctane Drum Kit for PlayStation(R)3
	0200  Harmonix Guitar for PlayStation(R)3
	0210  Harmonix Drum Kit for PlayStation(R)3
12bd  Gembird
	d012  JPD Shockforce gamepad
12c4  Autocue Group Ltd
	0006  Teleprompter Two-button Hand Control (v1)
	0008  Teleprompter Foot Control (v1)
12cf  DEXIN
	0170  Tt eSPORTS BLACK Gaming mouse
12d1  Huawei Technologies Co., Ltd.
	1001  E169/E620/E800 HSDPA Modem
	1003  E220 HSDPA Modem / E230/E270/E870 HSDPA/HSUPA Modem
	1004  E220 (bis)
	1009  U120
	1010  ETS2252+ CDMA Fixed Wireless Terminal
	1021  U8520
	1035  U8120
	1037  Ideos
	1038  Ideos (debug mode)
	1039  Ideos (tethering mode)
	1404  EM770W miniPCI WCDMA Modem
	1406  E1750
	140b  EC1260 Wireless Data Modem HSD USB Card
	140c  E180v
	1412  EC168c
	1436  Broadband stick
	1446  Broadband stick (modem on)
	1465  K3765 HSPA
	14c3  K5005 Vodafone LTE/UMTS/GSM Modem/Networkcard
	14c8  K5005 Vodafone LTE/UMTS/GSM MOdem/Networkcard
	14c9  K3770 3G Modem
	14cf  K3772
	14d1  K3770 3G Modem (Mass Storage Mode)
	14db  E353/E3131
	14f1  Gobi 3000 HSPA+ Modem
	14fe  Modem (Mass Storage Mode)
	1501  Pulse
	1505  E398 LTE/UMTS/GSM Modem/Networkcard
	1506  Modem/Networkcard
	150a  E398 LTE/UMTS/GSM Modem/Networkcard
	1520  K3765 HSPA
	1521  K4505 HSPA+
	155a  R205 Mobile WiFi (CD-ROM mode)
	1575  K5150 LTE modem
	15ca  E3131 3G/UMTS/HSPA+ Modem (Mass Storage Mode)
	1805  AT&T Go Phone U2800A phone
	1c05  Broadband stick (modem on)
	1c0b  E173s 3G broadband stick (modem off)
	1c20  R205 Mobile WiFi (Charging)
	1d50  ET302s TD-SCDMA/TD-HSDPA Mobile Broadband
	1f01  E353/E3131 (Mass storage mode)
	1f16  K5150 LTE modem (Mass Storage Mode)
	380b  WiMAX USB modem(s)
12d2  LINE TECH INDUSTRIAL CO., LTD.
12d6  EMS Dr. Thomas Wuensche
	0444  CPC-USB/ARM7
	0888  CPC-USB/M16C
12d7  BETTER WIRE FACTORY CO., LTD.
12d8  Araneus Information Systems Oy
	0001  Alea I True Random Number Generator
12e6  Waldorf Music GmbH
	0013  Blofeld
12ef  Tapwave, Inc.
	0100  Tapwave Handheld [Tapwave Zodiac]
12f5  Dynamic System Electronics Corp.
12f7  Memorex Products, Inc.
	1a00  TD Classic 003B
	1e23  TravelDrive 2007 Flash Drive
12fd  AIN Comm. Technology Co., Ltd
	1001  AWU2000b 802.11b Stick
12ff  Fascinating Electronics, Inc.
	0101  Advanced RC Servo Controller
1307  Transcend Information, Inc.
	0163  256MB/512MB/1GB Flash Drive
	0165  2GB/4GB/8GB Flash Drive
	0190  Ut190 8 GB Flash Drive with MicroSD reader
	0310  SD/MicroSD CardReader [hama]
	0330  63-in-1 Multi-Card Reader/Writer
	0361  CR-75: 51-in-1 Card Reader/Writer [Sakar]
	1169  TS2GJF210 JetFlash 210 2GB
	1171  Fingerprint Reader
1308  Shuttle, Inc.
	0003  VFD Module
	c001  eHome Infrared Transceiver
1310  Roper
	0001  Class 1 Bluetooth Dongle
1312  ICS Electronics
1313  ThorLabs
	0010  LC1 Linear Camera (Jungo)
	0011  SP1 Spectrometer (Jungo)
	0012  SP2 Spectrometer (Jungo)
	0110  LC1 Linear Camera (VISA)
	0111  SP1 Spectrometer (VISA)
	0112  SP2 Spectrometer (VISA)
	8001  TXP-Series Slot (TXP5001, TXP5004)
	8012  BC106 Camera Beam Profiler
	8013  WFS10 Wavefront Sensor
	8017  BC206 Camera Beam Profiler
	8019  BP2 Multi Slit Beam Profiler
	8020  PM300 Optical Power Meter
	8021  PM300E Optical Power and Energy Meter
	8022  PM320E Optical Power and Energy Meter
	8030  ER100 Extinction Ratio Meter
	8070  PM100D
131d  Natural Point
	0155  TrackIR 3 Pro Head Tracker
	0156  TrackIR 4 Pro Head Tracker
132a  Envara Inc.
	1502  WiND 802.11abg / 802.11bg WLAN
132b  Konica Minolta
	0000  Dimage A2 Camera
	0001  Minolta DiMAGE A2 (ptp)
	0003  Dimage Xg Camera
	0006  Dimage Z2 Camera
	0007  Minolta DiMAGE Z2 (PictBridge mode)
	0008  Dimage X21 Camera
	000a  Dimage Scan Dual IV AF-3200 (2891)
	000b  Dimage Z10 Camera
	000d  Dimage X50 Camera [storage?]
	000f  Dimage X50 Camera [p2p?]
	0010  Dimage G600 Camera
	0012  Dimage Scan Elite 5400 II (2892)
	0013  Dimage X31 Camera
	0015  Dimage G530 Camera
	0017  Dimage Z3 Camera
	0018  Minolta DiMAGE Z3 (PictBridge mode)
	0019  Dimage A200 Camera
	0021  Dimage Z5 Camera
	0022  Minolta DiMAGE Z5 (PictBridge mode)
	002c  Dynax 5D camera
	2001  Magicolor 2400w
	2004  Magicolor 5430DL
	2005  Magicolor 2430 DL
	2029  Magicolor 5440DL
	2030  PagePro 1350E(N)
	2033  PagePro 1400W
	2043  Magicolor 2530DL
	2045  Magicolor 2500W
	2049  Magicolor 2490MF
133e  Kemper Digital GmbH
	0815  Virus TI Desktop
1342  Mobility
	0200  EasiDock 200 Hub
	0201  EasiDock 200 Keyboard and Mouse Port
	0202  EasiDock 200 Serial Port
	0203  EasiDock 200 Printer Port
	0204  Ethernet
	0304  EasiDock Ethernet
1343  Citizen Systems
	0003  CX / DNP DS40
	0004  CX-W / DNP DS80
	0005  CY / DNP DSRX
1345  Sino Lite Technology Corp.
	001c  Xbox Controller Hub
	6006  Defender Wireless Controller
1347  Moravian Instruments
	0400  G2CCD USB 1.1 obsolete
	0401  G2CCD-S with Sony ICX285 CCD
	0402  G2CCD2
	0403  G2/G3CCD-I KAI CCD
	0404  G2/G3/G4 CCD-F KAF CCD
	0405  Gx CCD-I CCD
	0406  Gx CCD-F CCD
	0410  G1-0400 CCD
	0411  G1-0800 CCD
	0412  G1-0300 CCD
	0413  G1-2000 CCD
	0414  G1-1400 CCD
1348  Katsuragawa Electric Co., Ltd.
134c  PanJit International Inc.
	0001  Touch Panel Controller
	0002  Touch Panel Controller
	0003  Touch Panel Controller
	0004  Touch Panel Controller
134e  Digby's Bitpile, Inc. DBA D Bit
1357  P&E Microcomputer Systems
	0089  OpenSDA - CDC Serial Port
	0503  USB-ML-12 HCS08/HCS12 Multilink
	0504  DEMOJM
135f  Control Development Inc.
	0110  Linear Spectrograph
	0111  Spectrograph - Renumerated
	0200  Linear Spectrograph
	0201  Spectrograph - Renumerated
	0240  MPP Spectrograph
1366  SEGGER
	0101  J-Link PLUS
136b  STEC
136e  Andor Technology Ltd.
	0014  Zyla 5.5 sCMOS camera
1370  Swissbit
	0323  Swissmemory cirrusWHITE
	6828  Victorinox Flash Drive
1371  CNet Technology Inc.
	0001  CNUSB-611AR Wireless Adapter-G [AT76C503]
	0002  CNUSB-611AR Wireless Adapter-G [AT76C503] (FiberLine WL-240U)
	0013  CNUSB-611 Wireless Adapter [AT76C505]
	0014  CNUSB-611 Wireless Adapter [AT76C505] (FiberLine WL-240U)
	5743  CNUSB-611 (D) Wireless Adapter [AT76C503]
	9022  CWD-854 [RT2573]
	9032  CWD-854 rev F
	9401  CWD-854 Wireless 802.11g 54Mbps Network Adapter [RTL8187]
1376  Vimtron Electronics Co., Ltd.
137b  SCAPS GmbH
	0002  SCAPS USC-2 Scanner Controller
1385  Netgear, Inc
	4250  WG111T
	4251  WG111T (no firmware)
	5f00  WPN111 RangeMax(TM) Wireless USB 2.0 Adapter
	5f01  WPN111 (no firmware)
	5f02  WPN111 (no firmware)
	6e00  WPNT121 802.11g 240Mbps Wireless Adapter [Airgo AGN300]
138a  Validity Sensors, Inc.
	0001  VFS101 Fingerprint Reader
	0005  VFS301 Fingerprint Reader
	0007  VFS451 Fingerprint Reader
	0008  VFS300 Fingerprint Reader
	0010  VFS Fingerprint sensor
	0011  VFS5011 Fingerprint Reader
	0017  Fingerprint Reader
	0018  Fingerprint scanner
	003c  VFS471 Fingerprint Reader
	003d  VFS491
	003f  VFS495 Fingerprint Reader
	0050  Swipe Fingerprint Sensor
138e  Jungo LTD
	9000  Raisonance S.A. STM32 ARM evaluation board
1390  TOMTOM B.V.
	0001  GO 520 T/GO 630/ONE XL (v9)
	5454  Blue & Me 2
	7474  GPS Sport Watch [Runner, Multi-Sport]
1391  IdealTEK, Inc.
	1000  URTC-1000
1395  Sennheiser Communications
	3556  USB Headset
1397  BEHRINGER International GmbH
	00bc  BCF2000
1398  Q-tec
	2103  USB 2.0 Storage Device
13ad  Baltech
	9999  Card reader
13b0  PerkinElmer Optoelectronics
	000a  Alesis Photon X25 MIDI Controller
13b1  Linksys
	000a  WUSB54G v2 802.11g Adapter [Intersil ISL3887]
	000b  WUSB11 v4.0 802.11b Adapter [ALi M4301]
	000c  WUSB54AG 802.11a/g Adapter [Intersil ISL3887]
	000d  WUSB54G v4 802.11g Adapter [Ralink RT2500USB]
	000e  WUSB54GS v1 802.11g Adapter [Broadcom 4320 USB]
	0011  WUSB54GP v4.0 802.11g Adapter [Ralink RT2500USB]
	0014  WUSB54GS v2 802.11g Adapter [Broadcom 4320 USB]
	0018  USB200M 10/100 Ethernet Adapter
	001a  HU200TS Wireless Adapter
	001e  WUSBF54G 802.11bg
	0020  WUSB54GC v1 802.11g Adapter [Ralink RT73]
	0022  WUSB54GX4 802.11g 240Mbps Wireless Adapter [Airgo AGN300]
	0023  WUSB54GR
	0024  WUSBF54G v1.1 802.11bg
	0026  WUSB54GSC v1 802.11g Adapter [Broadcom 4320 USB]
	0028  WUSB200 802.11g Adapter [Ralink RT2671]
	0029  WUSB300N 802.11bgn Wireless Adapter [Marvell 88W8362+88W8060]
	002f  AE1000 v1 802.11n [Ralink RT3572]
	0031  AM10 v1 802.11n [Ralink RT3072]
	0039  AE1200 802.11bgn Wireless Adapter [Broadcom BCM43235]
	003a  AE2500 802.11abgn Wireless Adapter [Broadcom BCM43236]
	003b  AE3000 802.11abgn (3x3) Wireless Adapter [Ralink RT3573]
	003e  AE6000 802.11a/b/g/n/ac Wireless Adapter [MediaTek MT7610U]
	003f  WUSB6300 802.11a/b/g/n/ac Wireless Adapter [Realtek RTL8812AU]
	13b1  WUSB200: Wireless-G Business Network Adapter with Rangebooster
13b2  Alesis
	0030  Multimix 8
13b3  Nippon Dics Co., Ltd.
13ba  PCPlay
	0001  Konig Electronic CMP-KEYPAD12 Numeric Keypad
	0017  PS/2 Keyboard+Mouse Adapter
	0018  Barcode PCP-BCG4209
13be  Ricoh Printing Systems, Ltd.
13ca  JyeTai Precision Industrial Co., Ltd.
13cf  Wisair Ltd.
	1200  Olidata Wireless Multimedia Adapter
13d0  Techsan Electronics Co., Ltd.
	2282  TechniSat DVB-PC TV Star 2
13d1  A-Max Technology Macao Commercial Offshore Co. Ltd.
	7019  MD 82288
	abe6  Wireless 802.11g 54Mbps Network Adapter [RTL8187]
13d2  Shark Multimedia
	0400  Pocket Ethernet [klsi]
13d3  IMC Networks
	3201  VisionDTV USB-Ter/HAMA USB DVB-T device cold
	3202  VisionDTV USB-Ter/HAMA USB DVB-T device warm
	3203  DTV-DVB UDST7020BDA DVB-S Box(DVBS for MCE2005)
	3204  DTV-DVB UDST7020BDA DVB-S Box(DVBS for MCE2005)
	3205  DNTV Live! Tiny USB2 BDA (No Remote)
	3206  DNTV Live! Tiny USB2 BDA (No Remote)
	3207  DTV-DVB UDST7020BDA DVB-S Box(DVBS for MCE2005)
	3208  DTV-DVB UDST7020BDA DVB-S Box(DVBS for MCE2005)
	3209  DTV-DVB UDST7022BDA DVB-S Box(Without HID)
	3211  DTV-DVB Hybrid Analog/Capture / Pinnacle PCTV 310e
	3212  DTV-DVB UDTT704C - DVBT/NTSC/PAL Driver(PCM4)
	3213  DTV-DVB UDTT704D - DVBT/NTSC/PAL Driver (PCM4)
	3214  DTV-DVB UDTT704F -(MiniCard) DVBT/NTSC/PAL Driver(Without HID)
	3215  DTV-DVB UDAT7240 - ATSC/NTSC/PAL Driver(PCM4)
	3216  DTV-DVB UDTT 7047-USB 2.0 DVB-T Driver
	3217  Digital-TV Receiver.
	3219  DTV-DVB UDTT7049 - DVB-T Driver(Without HID)
	3220  DTV-DVB UDTT 7047M-USB 2.0 DVB-T Driver
	3223  DNTV Live! Tiny USB2 BDA (No Remote)
	3224  DNTV Live! Tiny USB2 BDA (No Remote)
	3226  DigitalNow TinyTwin DVB-T Receiver
	3234  DVB-T FTA Half Minicard [RTL2832U]
	3236  DTV-DVB UDTT 7047A-USB 2.0 DVB-T Driver
	3237  DTV-DVB UDTT 704J - dual DVB-T Driver
	3239  DTV-DVB UDTT704D - DVBT/NTSC/PAL Driver(Without HID)
	3240  DTV-DVB UDXTTM6010 - A/D Driver(Without HID)
	3241  DTV-DVB UDXTTM6010 - A/D Driver(Without HID)
	3242  DTV-DVB UDAT7240LP - ATSC/NTSC/PAL Driver(Without HID)
	3243  DTV-DVB UDXTTM6010 - A/D Driver(Without HID)
	3244  DTV-DVB UDTT 7047Z-USB 2.0 DVB-T Driver
	3247  802.11 n/g/b Wireless LAN Adapter
	3249  Internal Bluetooth
	3262  802.11 n/g/b Wireless LAN USB Adapter
	3273  802.11 n/g/b Wireless LAN USB Mini-Card
	3274  DVB-T Dongle [RTL2832U]
	3282  DVB-T + GPS Minicard [RTL2832U]
	3284  Wireless LAN USB Mini-Card
	3304  Asus Integrated Bluetooth module [AR3011]
	3306  Mediao 802.11n WLAN [Realtek RTL8191SU]
	3315  Bluetooth module
	3362  Atheros AR3012 Bluetooth 4.0 Adapter
	3375  Atheros AR3012 Bluetooth 4.0 Adapter
	3392  Azurewave 43228+20702
	3394  Bluetooth
	3474  Atheros AR3012 Bluetooth
	5070  Webcam
	5111  Integrated Webcam
	5115  Integrated Webcam
	5116  Integrated Webcam
	5122  2M Integrated Webcam
	5126  PC Cam
	5130  Integrated Webcam
	5702  UVC VGA Webcam
	5710  UVC VGA Webcam
	5716  UVC VGA Webcam
	7020  DTV-DVB UDST7020BDA DVB-S Box(DVBS for MCE2005)
	7022  DTV-DVB UDST7022BDA DVB-S Box(Without HID)
13d7  Guidance Software, Inc.
	0001  T5 PATA forensic bridge
13dc  ALEREON, INC.
13dd  i.Tech Dynamic Limited
13e1  Kaibo Wire & Cable (Shenzhen) Co., Ltd.
13e5  Rane
	0001  SL-1
	0003  TTM 57SL
13e6  TechnoScope Co., Ltd.
13ea  Hengstler
	0001  C-56 Thermal Printer
13ec  Zydacron
	0006  HID Remote Control
13ee  MosArt
	0001  Optical Mouse
	0003  Optical Mouse
13fd  Initio Corporation
	0840  INIC-1618L SATA
	0841  Samsung SE-T084M DVD-RW
	1040  INIC-1511L PATA Bridge
	1340  Hi-Speed USB to SATA Bridge
	160f  RocketFish SATA Bridge [INIC-1611]
	1640  INIC-1610L SATA Bridge
	1669  INIC-1609PN
	1840  INIC-1608 SATA bridge
	1e40  INIC-1610P SATA bridge
13fe  Kingston Technology Company Inc.
	1a00  512MB/1GB Flash Drive
	1a23  512MB Flash Drive
	1d00  DataTraveler 2.0 1GB/4GB Flash Drive / Patriot Xporter 4GB Flash Drive
	1e00  Flash Drive 2 GB [ICIDU 2 GB]
	1e50  U3 Smart Drive
	1f00  Kingston DataTraveler / Patriot Xporter
	1f23  PS2232 flash drive controller
	2240  microSD card reader
	3100  2/4 GB stick
	3123  Verbatim STORE N GO 4GB
	3600  flash drive (4GB, EMTEC)
	3800  Rage XT Flash Drive
	3e00  Flash Drive
	4100  Flash drive
	5000  USB flash drive (32 GB SHARKOON Accelerate)
	5100  Flash Drive
1400  Axxion Group Corp.
1402  Bowe Bell & Howell
1403  Sitronix
	0001  Digital Photo Frame
1409  IDS Imaging Development Systems GmbH
	1000  generic (firmware not loaded yet)
	1485  uEye UI1485
140e  Telechips, Inc.
	b011  TCC780X-based player (USB Boot mode)
	b021  TCC77X-based players (USB Boot mode)
1410  Novatel Wireless
	1110  Merlin S620
	1120  Merlin EX720
	1130  Merlin S720
	1400  Merlin U730/U740 (Vodafone)
	1410  Merlin U740 (non-Vodafone)
	1430  Merlin XU870
	1450  Merlin X950D
	2110  Ovation U720/MCD3000
	2410  Expedite EU740
	2420  Expedite EU850D/EU860D/EU870D
	4100  U727
	4400  Ovation MC930D/MC950D
	9010  Expedite E362
	a001  Gobi Wireless Modem
	a008  Gobi Wireless Modem (QDL mode)
	b001  Ovation MC551
1415  Nam Tai E&E Products Ltd. or OmniVision Technologies, Inc.
	0000  Sony SingStar USBMIC
	0020  Sony Wireless SingStar
	2000  Sony Playstation Eye
1419  ABILITY ENTERPRISE CO., LTD.
1421  Sensor Technology
	0605  Sentech Camera
1429  Vega Technologies Industrial (Austria) Co.
142a  Thales E-Transactions
	0003  Artema Hybrid
	0005  Artema Modular
	0043  medCompact
142b  Arbiter Systems, Inc.
	03a5  933A Portable Power Sentinel
1430  RedOctane
	0150  wireless receiver for skylanders wii
	4734  Guitar Hero4 hub
	474b  Guitar Hero MIDI interface
1431  Pertech Resources, Inc.
1435  Wistron NeWeb
	0427  UR054g 802.11g Wireless Adapter [Intersil ISL3887]
	0711  UR055G 802.11bg
	0804  AR9170+AR9104 802.11abgn Wireless Adapter
	0826  AR5523
	0827  AR5523 (no firmware)
	0828  AR5523
	0829  AR5523 (no firmware)
1436  Denali Software, Inc.
143c  Altek Corporation
1443  Digilent
	0007  Development board JTAG
1446  X.J.GROUP
	6a73  Stamps.com Model 510 5LB Scale
	6a78  DYMO Endicia 75lb Digital Scale
1453  Radio Shack
	4026  26-183 Serial Cable
1456  Extending Wire & Cable Co., Ltd.
1457  First International Computer, Inc.
	5117  OpenMoko Neo1973 kernel usbnet (g_ether, CDC Ethernet) mode
	5118  OpenMoko Neo1973 Debug board (V2+)
	5119  OpenMoko Neo1973 u-boot cdc_acm serial port
	511a  HXD8 u-boot usbtty CDC ACM Mode
	511b  SMDK2440 u-boot usbtty CDC ACM mode
	511c  SMDK2443 u-boot usbtty CDC ACM mode
	511d  QT2410 u-boot usbtty CDC ACM mode
	5120  OpenMoko Neo1973 u-boot usbtty generic serial
	5121  OpenMoko Neo1973 kernel mass storage (g_storage) mode
	5122  OpenMoko Neo1973 / Neo Freerunner kernel cdc_ether USB network
	5123  OpenMoko Neo1973 internal USB CSR4 module
	5124  OpenMoko Neo1973 Bluetooth Device ID service
145f  Trust
	0106  Trust K56 V92 USB Modem
	013d  PC Camera (SN9C201 + OV7660)
	013f  Megapixel Auto Focus Webcam
	0142  WB-6250X Webcam
	015a  WB-8300X 2MP Webcam
	0161  15901 802.11bg Wireless Adapter [Realtek RTL8187L]
	0167  Widescreen 3MP Webcam
	0176  Isla Keyboard
1460  Tatung Co.
	9150  eHome Infrared Transceiver
1461  Staccato Communications
1462  Micro Star International
	5512  MegaStick-1 Flash Stick
	8807  DIGIVOX mini III [af9015]
1472  Huawei-3Com
	0007  Aolynk WUB300g [ZyDAS ZD1211]
	0009  Aolynk WUB320g
147a  Formosa Industrial Computing, Inc.
	e015  eHome Infrared Receiver
	e016  eHome Infrared Receiver
	e017  eHome Infrared Receiver
	e018  eHome Infrared Receiver
	e02c  Infrared Receiver
	e03a  eHome Infrared Receiver
	e03c  eHome Infrared Receiver
	e03d  2 Channel Audio
	e03e  Infrared Receiver [IR605A/Q]
147e  Upek
	1000  Biometric Touchchip/Touchstrip Fingerprint Sensor
	1001  TCS5B Fingerprint sensor
	1002  Biometric Touchchip/Touchstrip Fingerprint Sensor
	2016  Biometric Touchchip/Touchstrip Fingerprint Sensor
	2020  TouchChip Fingerprint Coprocessor (WBF advanced mode)
	3000  TCS1C EIM/Cypress Fingerprint sensor
	3001  TCS1C EIM/STM32 Fingerprint sensor
147f  Hama GmbH & Co., KG
1482  Vaillant
	1005  VRD PC-Interface
1484  Elsa AG [hex]
	1746  Ecomo 19H99 Monitor
	7616  Elsa Hub
1485  Silicom
	0001  U2E
	0002  Psion Gold Port Ethernet
1487  DSP Group, Ltd.
148e  EVATRONIX SA
148f  Ralink Technology, Corp.
	1000  Motorola BC4 Bluetooth 3.0+HS Adapter
	1706  RT2500USB Wireless Adapter
	2070  RT2070 Wireless Adapter
	2570  RT2570 Wireless Adapter
	2573  RT2501/RT2573 Wireless Adapter
	2671  RT2601/RT2671 Wireless Adapter
	2770  RT2770 Wireless Adapter
	2870  RT2870 Wireless Adapter
	3070  RT2870/RT3070 Wireless Adapter
	3071  RT3071 Wireless Adapter
	3072  RT3072 Wireless Adapter
	3370  RT3370 Wireless Adapter
	3572  RT3572 Wireless Adapter
	3573  RT3573 Wireless Adapter
	5370  RT5370 Wireless Adapter
	5372  RT5372 Wireless Adapter
	5572  RT5572 Wireless Adapter
	7601  MT7601U Wireless Adapter
	760b  MT7601U Wireless Adapter
	9020  RT2500USB Wireless Adapter
	9021  RT2501USB Wireless Adapter
1491  Futronic Technology Co. Ltd.
	0020  FS81 Fingerprint Scanner Module
1493  Suunto
	0010  Bluebird [Ambit]
	0019  Duck [Ambit2]
	001a  Colibri [Ambit2 S]
	001b  Emu [Ambit3 Peak]
	001c  Finch [Ambit3 Sport]
	001d  Greentit [Ambit2 R]
1497  Panstrong Company Ltd.
1498  Microtek International Inc.
	a090  DVB-T Tuner
149a  Imagination Technologies
	2107  DBX1 DSP core
14aa  WideView Technology Inc.
	0001  Avermedia AverTV DVBT USB1.1 (cold)
	0002  Avermedia AverTV DVBT USB1.1 (warm)
	0201  AVermedia/Yakumo/Hama/Typhoon DVB-T USB2.0 (cold)
	0221  WT-220U DVB-T dongle
	022b  WT-220U DVB-T dongle
	0301  AVermedia/Yakumo/Hama/Typhoon DVB-T USB2.0 (warm)
14ad  CTK Corporation
14ae  Printronix Inc.
14af  ATP Electronics Inc.
14b0  StarTech.com Ltd.
14b2  Ralink Technology, Corp.
	3a93  Topcom 802.11bg Wireless Adapter [Atheros AR5523]
	3a95  Toshiba WUS-G06G-JT 802.11bg Wireless Adapter [Atheros AR5523]
	3a98  Airlink101 AWLL4130 802.11bg Wireless Adapter [Atheros AR5523]
	3c02  Conceptronic C54RU v2 802.11bg Wireless Adapter [Ralink RT2571]
	3c05  rt2570 802.11g WLAN
	3c06  Conceptronic C300RU v1 802.11bgn Wireless Adapter [Ralink RT2870]
	3c07  802.11n adapter
	3c09  802.11n adapter
	3c22  Conceptronic C54RU v3 802.11bg Wireless Adapter [Ralink RT2571W]
	3c23  Airlink101 AWLL6080 802.11bgn Wireless Adapter [Ralink RT2870]
	3c24  NEC NP01LM 802.11abg Wireless Adapter [Ralink RT2571W]
	3c25  DrayTek Vigor N61 802.11bgn Wireless Adapter [Ralink RT2870]
	3c27  Airlink101 AWLL6070 802.11bgn Wireless Adapter [Ralink RT2770]
	3c28  Conceptronic C300RU v2 802.11bgn Wireless Adapter [Ralink RT2770]
	3c2b  NEC NP02LM 802.11bgn Wireless Adapter [Ralink RT3072]
	3c2c  Keebox W150NU 802.11bgn Wireless Adapter [Ralink RT3070]
14c0  Rockwell Automation, Inc.
14c2  Gemlight Computer, Ltd
	0250  Storage Adapter V2
	0350  Storage Adapter V2
14c8  Zytronic
14cd  Super Top
	1212  microSD card reader (SY-T18)
	121c  microSD card reader
	121f  microSD CardReader SY-T18
	123a  SD/MMC/RS-MMC Card Reader
	125c  SD card reader
	127b  SDXC Reader
	6116  M6116 SATA Bridge
	6600  M110E PATA bridge
	6700  Card Reader
	6900  Card Reader
	8123  SD MMC Reader
	8125  SD MMC Reader
14d8  JAMER INDUSTRIES CO., LTD.
14dd  Raritan Computer, Inc.
	1007  D2CIM-VUSB KVM connector
14e0  WiNRADiO Communications
	0501  WR-G528e 'CHEETAH'
14e1  Dialogue Technology Corp.
	5000  PenMount 5000 Touch Controller
14e5  SAIN Information & Communications Co., Ltd.
14ea  Planex Communications
	ab10  GW-US54GZ
	ab11  GU-1000T
	ab13  GW-US54Mini 802.11bg
14ed  Shure Inc.
	29b6  X2u Adapter
14f7  TechniSat Digital GmbH
	0001  SkyStar 2 HD CI
	0002  SkyStar 2 HD CI
	0003  CableStar Combo HD CI
	0004  AirStar TeleStick 2
	0500  DVB-PC TV Star HD
1500  Ellisys
1501  Pine-Tum Enterprise Co., Ltd.
1509  First International Computer, Inc.
	0a01  LI-3100 Area Meter
	0a02  LI-7000 CO2/H2O Gas Analyzer
	0a03  C-DiGit Blot Scanner
	9242  eHome Infrared Transceiver
1513  medMobile
	0444  medMobile
1514  Actel
	2003  FlashPro3 Programmer
	2004  FlashPro3 Programmer
	2005  FlashPro3 Programmer
1516  CompUSA
	1603  Flash Drive
	8628  Pen Drive
1518  Cheshire Engineering Corp.
	0001  HDReye High Dynamic Range Camera
	0002  HDReye (before firmware loads)
1519  Comneon
	0020  HSIC Device
	0443  Telit LN930
1520  Bitwire Corp.
1524  ENE Technology Inc
	6680  UTS 6680
1527  Silicon Portals
	0200  YAP Phone (no firmware)
	0201  YAP Phone
1529  UBIQUAM Co., Ltd.
	3100  CDMA 1xRTT USB Modem (U-100/105/200/300/520)
152a  Thesycon Systemsoftware & Consulting GmbH
	8350  NET Gmbh iCube Camera
	8400  INI DVS128
	840d  INI DAViS
	841a  INI DAViS FX3
152b  MIR Srl
	0001  spirobank II
	0002  spirolab III
	0003  MiniSpir
	0004  Oxi
	0005  spiros II
	0006  smiths spirobank II
	0007  smiths spirobank G-USB
	0008  smiths MiniSpir
	0009  spirobank G-USB
	000a  smiths Oxi
	000b  smiths spirolab III
	000c  chorus III
	000d  spirolab III Bw
	000e  spirolab III
	000f  easySpiro
	0010  Spirotel converter
	0011  spirobank
	0012  spiro3 Zimmer
	0013  spirotel serial
	0014  spirotel II
	0015  spirodoc
152d  JMicron Technology Corp. / JMicron USA Technology Corp.
	0539  JMS539/567 SuperSpeed SATA II/III 3.0G/6.0G Bridge
	0567  JMS567 SATA 6Gb/s bridge
	0770  Alienware Integrated Webcam
	2329  JM20329 SATA Bridge
	2335  ATA/ATAPI Bridge
	2336  Hard Disk Drive
	2337  ATA/ATAPI Bridge
	2338  JM20337 Hi-Speed USB to SATA & PATA Combo Bridge
	2339  JM20339 SATA Bridge
	2352  ATA/ATAPI Bridge
	2509  JMS539 SuperSpeed SATA II 3.0G Bridge
	2551  JMS551 SATA 3Gb/s bridge
	2566  JMS566 SATA 3Gb/s bridge
	2590  Seatay ATA/ATAPI Bridge
	3562  JMS567 SATA 6Gb/s bridge
	3569  JMS566 SATA 3Gb/s bridge
152e  LG (HLDS)
	2507  PL-2507 IDE Controller
	e001  GSA-5120D DVD-RW
1532  Razer USA, Ltd
	0001  RZ01-020300 Optical Mouse [Diamondback]
	0003  Krait Mouse
	0007  DeathAdder Mouse
	0013  Orochi mouse
	0015  Naga Mouse
	0016  DeathAdder Mouse
	0017  RZ01-0035 Laser Gaming Mouse [Imperator]
	001c  RZ01-0036 Optical Gaming Mouse [Abyssus]
	0024  Razer Mamba
	002e  RZ01-0058 Gaming Mouse [Naga]
	0036  RZ01-0075, Gaming Mouse [Naga Hex]
	0101  Copperhead Mouse
	0102  Tarantula Keyboard
	0109  Lycosa Keyboard
	0113  RZ07-0074 Gaming Keypad [Orbweaver]
	0300  RZ06-0063 Motion Sensing Controllers [Hydra]
153b  TerraTec Electronic GmbH
	1181  Cinergy S2 PCIe Dual Port 1
	1182  Cinergy S2 PCIe Dual Port 2
1546  U-Blox AG
	01a5  NL-402U
1547  SG Intec Ltd & Co KG
	1000  SG-Lock[U2]
154a  Celectronic GmbH
	8180  CARD STAR/medic2
154b  PNY
	0010  USB 2.0 Flash Drive
	0048  Flash Drive
	004d  8 GB Flash Drive
	0053  Flash Drive
	0057  32GB Micro Slide Attache Flash Drive
	005b  Flash Drive
	0062  Flash Drive
	007a  Classic Attache Flash Drive
	6545  FD Device
	fa05  Flash Drive
154d  ConnectCounty Holdings Berhad
154e  D&M Holdings, Inc. (Denon/Marantz)
	3000  Marantz RC9001 Remote Control
154f  SNBC CO., Ltd
1554  Prolink Microsystems Corp.
	5010  PV-D231U(RN)-F [PixelView PlayTV SBTVD Full-Seg]
1557  OQO
	0002  model 01 WiFi interface
	0003  model 01 Bluetooth interface
	0a80  Gobi Wireless Modem (QDL mode)
	7720  model 01+ Ethernet
	8150  model 01 Ethernet interface
1568  Sunf Pu Technology Co., Ltd
156f  Quantum Corporation
1570  ALLTOP TECHNOLOGY CO., LTD.
157b  Ketron SRL
157e  TRENDnet
	3006  TEW-444UB EU [TRENDnet]
	3007  TEW-444UB EU (no firmware)
	300a  TEW-429UB 802.11bg
	300b  TEW-429UB 802.11bg
	300c  TEW-429UF A1 802.11bg Wireless Adapter [ZyDAS ZD1211B]
	300d  TEW-429UB C1 802.11bg
	300e  SMC SMCWUSB-N 802.11bgn 2x2:2 Wireless Adapter [Ralink RT2870]
	3012  TEW-604UB 802.11bg Wireless Adapter [Atheros AR5523]
	3013  TEW-645UB 802.11bgn 1x2:2 Wireless Adapter [Ralink RT2770]
	3204  Allnet ALL0298 v2 802.11bg
	3205  Allnet ALL0283 [AR5523]
	3206  Allnet ALL0283 [AR5523](no firmware)
	3207  TEW-509UB A1 802.11abg Wireless Adapter [ZyDAS ZD1211]
	3208  TEW-509UB 1.1R 802.11abg Wireless Adapter
1582  Fiberline
	6003  WL-430U 802.11bg
1587  SMA Technologie AG
158d  Oakley Inc.
158e  JDS Uniphase Corporation (JDSU)
	0820  SmartPocket Class Device
1598  Kunshan Guoji Electronics Co., Ltd.
15a2  Freescale Semiconductor, Inc.
	0038  9S08JS Bootloader
	003b  USB2CAN Application for ColdFire DEMOJM board
	0042  OSBDM - Debug Port
	004f  i.MX28 SystemOnChip in RecoveryMode
	0052  i.MX50 SystemOnChip in RecoveryMode
	0054  i.MX 6Dual/6Quad SystemOnChip in RecoveryMode
	0061  i.MX 6Solo/6DualLite SystemOnChip in RecoveryMode
15a4  Afatech Technologies, Inc.
	1000  AF9015/AF9035 DVB-T stick
	1001  AF9015/AF9035 DVB-T stick
	1336  SDHC/MicroSD/MMC/MS/M2/CF/XD Flash Card Reader
	9015  AF9015 DVB-T USB2.0 stick
	9016  AF9015 DVB-T USB2.0 stick
15a8  Teams Power Limited
15a9  Gemtek
	0002  SparkLAN WL-682 802.11bg Wireless Adapter [Intersil ISL3887]
	0004  WUBR-177G [Ralink RT2571W]
	0006  Wireless 11n USB Adapter
	0010  802.11n USB Wireless Card
	0012  WUBR-208N 802.11abgn Wireless Adapter [Ralink RT2870]
	002d  WLTUBA-107 [Yota 4G LTE]
15aa  Gearway Electronics (Dong Guan) Co., Ltd.
15ad  VMware Inc.
15ba  Olimex Ltd.
	0003  OpenOCD JTAG
	0004  OpenOCD JTAG TINY
	002a  ARM-USB-TINY-H JTAG interface
	002b  ARM-USB-OCD-H JTAG+RS232
15c0  XL Imaging
	0001  2M pixel Microscope Camera
	0002  3M pixel Microscope Camera
	0003  1.3M pixel Microscope Camera (mono)
	0004  1.3M pixel Microscope Camera (colour)
	0005  3M pixel Microscope Camera (Mk 2)
	0006  2M pixel Microscope Camera (with capture button)
	0007  3M pixel Microscope Camera (with capture button)
	0008  1.3M pixel Microscope Camera (colour, with capture button)
	0009  1.3M pixel Microscope Camera (colour, with capture button)
	000a  2M pixel Microscope Camera (Mk 2)
	0010  1.3M pixel "Tinycam"
	0101  3M pixel Microscope Camera
15c2  SoundGraph Inc.
	0036  LC16M VFD Display/IR Receiver
	0038  GD01 MX LCD Display/IR Receiver
	0042  Antec Veris Multimedia Station E-Z IR Receiver
	ffda  iMON PAD Remote Controller
	ffdc  iMON PAD Remote Controller
15c5  Advance Multimedia Internet Technology Inc. (AMIT)
	0008  WL532U 802.11g Adapter
15c6  Laboratoires MXM
	1000  DigistimSP (cold)
	1001  DigistimSP (warm)
	1002  DigimapSP USB (cold)
	1003  DigimapSP USB (warm)
	1004  DigistimSP (cold)
	1005  DigistimSP (warm)
	1100  Odyssee (cold)
	1101  Odyssee (warm)
	1200  Digispy
15c8  KTF Technologies
	3201  EVER EV-W100/EV-W250
15c9  D-Box Technologies
15ca  Textech International Ltd.
	00c3  Mini Optical Mouse
	0101  MIDI Interface cable
	1806  MIDI Interface cable
15d5  Coulomb Electronics Ltd.
15d9  Trust International B.V.
	0a33  Optical Mouse
	0a37  Mouse
	0a41  MI-2540D [Optical mouse]
	0a4c  USB+PS/2 Optical Mouse
	0a4d  Optical Mouse
	0a4f  Optical Mouse
15dc  Hynix Semiconductor Inc.
15e0  Seong Ji Industrial Co., Ltd.
15e1  RSA
	2007  RSA SecurID (R) Authenticator
15e4  Numark
	0024  Mixtrack
	0140  ION VCR 2 PC / Video 2 PC
15e8  SohoWare
	9100  NUB100 Ethernet [pegasus]
	9110  10/100 USB Ethernet
15e9  Pacific Digital Corp.
	04ce  MemoryFrame MF-570
	1968  MemoryFrame MF-570
	1969  Digital Frame
15ec  Belcarra Technologies Corp.
15f4  HanfTek
	0001  HanfTek UMT-010 USB2.0 DVB-T (cold)
	0025  HanfTek UMT-010 USB2.0 DVB-T (warm)
1604  Tascam
	8000  US-428 Audio/Midi Controller (without fw)
	8001  US-428 Audio/Midi Controller
	8004  US-224 Audio/Midi Controller (without fw)
	8005  US-224 Audio/Midi Controller
	8006  US-122 Audio/Midi Interface (without fw)
	8007  US-122 Audio/Midi Interface
1606  Umax
	0002  Astra 1236U Scanner
	0010  Astra 1220U
	0030  Astra 1600U/2000U
	0050  Scanner
	0060  Astra 3400/3450
	0070  Astra 4400/4450
	0130  Astra 2100U
	0160  Astra 5400U
	0170  Uniscan D50
	0230  Astra 2200/2200SU
	0350  Astra 4800/4850 Scanner
	1030  Astra 4000U
	1220  Genesys Logic Scanner Controller NT5.0
	2010  AstraCam Digital Camera
	2020  AstraCam 1000
	2030  AstraCam 1800 Digital Camera
1608  Inside Out Networks [hex]
	0001  EdgePort/4 Serial Port
	0002  Edgeport/8
	0003  Rapidport/4
	0004  Edgeport/4
	0005  Edgeport/2
	0006  Edgeport/4i
	0007  Edgeport/2i
	0008  Edgeport/8
	000c  Edgeport/421
	000d  Edgeport/21
	000e  Edgeport/4
	000f  Edgeport/8
	0010  Edgeport/2
	0011  Edgeport/4
	0012  Edgeport/416
	0014  Edgeport/8i
	0018  Edgeport/412
	0019  Edgeport/412
	001a  Edgeport/2+2i
	0101  Edgeport/4
	0105  Edgeport/2
	0106  Edgeport/4i
	0107  Edgeport/2i
	010c  Edgeport/421
	010d  Edgeport/21
	0110  Edgeport/2
	0111  Edgeport/4
	0112  Edgeport/416
	0114  Edgeport/8i
	0201  Edgeport/4
	0203  Rapidport/4
	0204  Edgeport/4
	0205  Edgeport/2
	0206  Edgeport/4i
	0207  Edgeport/2i
	020c  Edgeport/421
	020d  Edgeport/21
	020e  Edgeport/4
	020f  Edgeport/8
	0210  Edgeport/2
	0211  Edgeport/4
	0212  Edgeport/416
	0214  Edgeport/8i
	0215  Edgeport/1
	0216  EPOS/44
	0217  Edgeport/42
	021a  Edgeport/2+2i
	021b  Edgeport/2c
	021c  Edgeport/221c
	021d  Edgeport/22c
	021e  Edgeport/21c
	021f  Edgeport/62
	0240  Edgeport/1
	0241  Edgeport/1i
	0242  Edgeport/4s
	0243  Edgeport/8s
	0244  Edgeport/8
	0245  Edgeport/22c
	0301  Watchport/P
	0302  Watchport/M
	0303  Watchport/W
	0304  Watchport/T
	0305  Watchport/H
	0306  Watchport/E
	0307  Watchport/L
	0308  Watchport/R
	0309  Watchport/A
	030a  Watchport/D
	030b  Watchport/D
	030c  Power Management Port
	030e  Power Management Port
	030f  Watchport/G
	0310  Watchport/Tc
	0311  Watchport/Hc
	1403  MultiTech Systems MT4X56 Modem
	1a17  Agilent Technologies (E6473)
160a  VIA Technologies, Inc.
	3184  VIA VNT-6656 [WiFi 802.11b/g USB Dongle]
160e  INRO
	0001  E2USBKey
1614  Amoi Electronics
	0404  WMA9109 UMTS Phone
	0600  Vodafone VDA GPS / Toschiba Protege G710
	0804  WP-S1 Phone
1617  Sony Corp.
	2002  NVX-P1 Personal Navigation System
1619  L & K Precision Technology Co., Ltd.
1621  Wionics Research
1628  Stonestreet One, Inc.
162a  Airgo Networks Inc.
162f  WiQuest Communications, Inc.
1630  2Wire, Inc.
	0005  802.11g Wireless Adapter [Intersil ISL3886]
	0011  PC Port 10 Mps Adapter
	ff81  802.11b Wireless Adapter [Lucent/Agere Hermes I]
1631  Good Way Technology
	6200  GWUSB2E
	c019  RT2573
1645  Entrega [hex]
	0001  1S Serial Port
	0002  2S Serial Port
	0003  1S25 Serial Port
	0004  4S Serial Port
	0005  E45 Ethernet [klsi]
	0006  Parallel Port
	0007  U1-SC25 SCSI
	0008  Ethernet
	0016  Bi-directional to Parallel Printer Converter
	0080  1 port to Serial Converter
	0081  1 port to Serial Converter
	0093  1S9 Serial Port
	8000  EZ-USB
	8001  1 port to Serial
	8002  2x Serial Port
	8003  1 port to Serial
	8004  2U4S serial/usb hub
	8005  Ethernet
	8080  1 port to Serial
	8081  1 port to Serial
	8093  PortGear Serial Port
1649  SofTec Microsystems
	0102  uDART In-Circuit Debugger
	0200  SpYder USBSPYDER08
164a  ChipX
164c  Matrix Vision GmbH
	0101  mvBlueFOX camera (no firmware)
	0103  mvBlueFOX camera
	0201  mvBlueLYNX-X intelligent camera (bootloader)
	0203  mvBlueLYNX-X intelligent camera
1657  Struck Innovative Systeme GmbH
	3150  SIS3150 USB2.0 to VME interface
165b  Frontier Design Group
	8101  Tranzport Control Surface
	fad1  Alphatrack Control Surface
165c  Kondo Kagaku
	0002  Serial Adapter
1660  Creatix Polymedia GmbH
1667  GIGA-TMS INC.
	0005  PCR330A RFID Reader (125 kHz, keyboard emulation)
1668  Actiontec Electronics, Inc. [hex]
	0009  Gateway
	0333  Modem
	0358  InternetPhoneWizard
	0405  Gateway
	0408  Prism2.5 802.11b Adapter
	0413  Gateway
	0421  Prism2.5 802.11b Adapter
	0441  IBM Integrated Bluetooth II
	0500  BTM200B BlueTooth Adapter
	1050  802UIG-1 802.11g Wireless Mini Adapter [Intersil ISL3887]
	1200  802AIN Wireless N Network Adapter [Atheros AR9170+AR9101]
	1441  IBM Integrated Bluetooth II
	2441  BMDC-2 IBM Bluetooth III w.56k
	3441  IBM Integrated Bluetooth III
	6010  Gateway
	6097  802.11b Wireless Adapter
	6106  802UI3(B) 802.11b Wireless Adapter [Intersil PRISM 3]
	7605  UAT1 Wireless Ethernet Adapter
1669  PiKRON Ltd. [hex]
	1001  uLan2USB Converter - PS1 protocol
166a  Clipsal
	0101  C-Bus Multi-room Audio Matrix Switcher
	0201  C-Bus Pascal Automation Controller
	0301  C-Bus Wireless PC Interface
	0303  C-Bus interface
	0304  C-Bus Black and White Touchscreen
	0305  C-Bus Spectrum Colour Touchscreen
	0401  C-Bus Architectural Dimmer
1677  China Huada Integrated Circuit Design (Group) Co., Ltd. (CIDC Group)
	0103  Token
1679  Total Phase
	2001  Beagle Protocol Analyzer
	2002  Cheetah SPI Host Adapter
1680  Golden Bridge Electech Inc.
	a332  DVB-T Dongle [RTL2832U]
1681  Prevo Technologies, Inc.
	0001  Tuner's Dashboard
	0002  Tubachron
1682  Maxwise Production Enterprise Ltd.
1684  Godspeed Computer Corp.
1685  Delock
	0200  Infrared adapter
1686  ZOOM Corporation
	0045  H4 Digital Recorder
1687  Kingmax Digital Inc.
	5289  FlashDisk
	6211  FlashDisk
	6213  FlashDisk
1688  Saab AB
1689  Razer USA, Ltd
	fd00  Onza Tournament Edition controller
168c  Atheros Communications
	0001  AR5523
	0002  AR5523 (no firmware)
1690  Askey Computer Corp. [hex]
	0001  Arcaze Gamepad
	0101  Creative Modem Blaster DE5670
	0102  V1456 VQE-R2 Modem [conexant]
	0103  1456 VQE-R3 Modem [conexant]
	0104  HCF V90 Data Fax RTAD Modem
	0107  HCF V.90 Data,Fax,RTAD Modem
	0109  MagicXpress V.90 Pocket Modem [conexant]
	0203  Voyager ADSL Modem Loader
	0204  Voyager ADSL Modem
	0205  DSL Modem
	0206  GlobeSpan ADSL WAN Modem
	0208  DSL Modem
	0209  Voyager 100 ADSL Modem
	0211  Globespan Virata ADSL LAN Modem
	0212  DSL Modem
	0213  HM121d DSL Modem
	0214  HM121d DSL Modem
	0215  Voyager 105 ADSL Modem
	0701  WLAN
	0710  SMCWUSBT-G
	0711  SMCWUSBT-G (no firmware)
	0712  AR5523
	0713  AR5523 (no firmware)
	0715  Name: Voyager 1055 Laptop 802.11g Adapter [Broadcom 4320]
	0722  RT2573
	0726  Wi-Fi Wireless LAN Adapter
	0740  802.11n Wireless LAN Card
	0901  Voyager 205 ADSL Router
	2000  naturaSign Pad Standard
	2001  naturaSign Pad Standard
	fe12  Bootloader
1696  Hitachi Video and Information System, Inc.
1697  VTec Test, Inc.
16a5  Shenzhen Zhengerya Cable Co., Ltd.
16a6  Unigraf
	3000  VTG-3xxx Video Test Generator family
	4000  VTG-4xxx Video Test Generator family
	5000  VTG-5xxx Video Test Generator family
	5001  VTG-5xxx Special (update) mode of VTG-5xxx family
16ab  Global Sun Technology
	7801  AR5523
	7802  AR5523 (no firmware)
	7811  AR5523
	7812  AR5523 (no firmware)
16ac  Dongguan ChingLung Wire & Cable Co., Ltd.
16b4  iStation
	0801  U43
16b5  Persentec, Inc.
	0002  Otto driving companion
16c0  Van Ooijen Technische Informatica
	03e8  free for internal lab use 1000
	03e9  free for internal lab use 1001
	03ea  free for internal lab use 1002
	03eb  free for internal lab use 1003
	03ec  free for internal lab use 1004
	03ed  free for internal lab use 1005
	03ee  free for internal lab use 1006
	03ef  free for internal lab use 1007
	03f0  free for internal lab use 1008
	03f1  free for internal lab use 1009
	0477  Teensy Rebootor
	0478  Teensy Halfkay Bootloader
	0479  Teensy Debug
	047a  Teensy Serial
	047b  Teensy Serial+Debug
	047c  Teensy Keyboard
	047d  Teensy Keyboard+Debug
	047e  Teensy Mouse
	047f  Teensy Mouse+Debug
	0480  Teensy RawHID
	0481  Teensy RawHID+Debug
	0482  Teensyduino Keyboard+Mouse+Joystick
	0483  Teensyduino Serial
	0484  Teensyduino Disk
	0485  Teensyduino MIDI
	0486  Teensyduino RawHID
	0487  Teensyduino Serial+Keyboard+Mouse+Joystick
	0488  Teensyduino Flight Sim Controls
	05dc  shared ID for use with libusb
	05dd  BlackcatUSB2
	05df  HID device except mice, keyboards, and joysticks
	05e1  Free shared USB VID/PID pair for CDC devices
	05e4  Free shared USB VID/PID pair for MIDI devices
	06b4  USB2LPT with 2 interfaces
	06b5  USB2LPT with 3 interfaces (native, HID, printer)
	074e  DSP-Weuffen USB-HPI-Programmer
	074f  DSP-Weuffen USB2-HPI-Programmer
	0762  Osmocom SIMtrace
	076b  OpenPCD 13.56MHz RFID Reader
	076c  OpenPICC 13.56MHz RFID Simulator (native)
	08ac  OpenBeacon USB stick
	08ca  Alpermann+Velte Universal Display
	08cb  Alpermann+Velte Studio Clock
	08cc  Alpermann+Velte SAM7S MT Boot Loader
	08cd  Alpermann+Velte SAM7X MT Boot Loader
	0a32  jbmedia Light-Manager Pro
	27d8  libusb-bound devices
	27d9  HID device except mice, keyboards, and joysticks
	27da  Mouse
	27db  Keyboard
	27dc  Joystick
	27dd  CDC-ACM class devices (modems)
	27de  MIDI class devices
	294a  Eye Movement Recorder [Visagraph]
	294b  Eye Movement Recorder [ReadAlyzer]
16ca  Wireless Cables, Inc.
	1502  Bluetooth Dongle
16cc  silex technology, Inc.
16d0  MCS
	0498  Braintechnology USB-LPS
	0504  RETRO Innovations ZoomFloppy
	054b  GrauTec ReelBox OLED Display (external)
	05be  EasyLogic Board
	06f9  Gabotronics Xminilab
	0753  Digistump DigiSpark
	075c  AB-1.x UAC1 [Audio Widget]
	075d  AB-1.x UAC2 [Audio Widget]
	080a  S2E1 Interface
	0870  Kaufmann Automotive GmbH, RKS+CAN Interface
16d1  Suprema Inc.
	0401  SUP-SFR400(A) BioMini Fingerprint Reader
16d3  Frontline Test Equipment, Inc.
16d5  AnyDATA Corporation
	6202  CDMA/UMTS/GPRS modem
	6501  CDMA 2000 1xRTT/EV-DO Modem
	6502  CDMA/UMTS/GPRS modem
	6603  ADU-890WH modem
16d6  JABLOCOM s.r.o.
	8000  GDP-04 desktop phone
	8001  EYE-02
	8003  GDP-04 modem
	8004  Bootloader
	8005  GDP-04i
	8007  BTP-06 modem
16d8  CMOTECH Co., Ltd.
	5141  CMOTECH CDMA Technologies modem
	5533  CCU-550 CDMA EV-DO modem
	5543  CDMA 2000 1xRTT/1xEVDO modem
	6280  CMOTECH CDMA Technologies modem
	6803  CNU-680 CDMA EV-DO modem
	8001  Gobi 2000 Wireless Modem (QDL mode)
	8002  Gobi 2000 Wireless Modem
16dc  Wiener, Plein & Baus
	0001  CC
	000b  VM
	0010  PL512 Power Supply System
	0011  MARATON Power Supply System
	0012  MPOD Multi Channel Power Supply System
	0015  CML Control, Measurement and Data Logging System
16df  King Billion Electronics Co., Ltd.
16f0  GN ReSound A/S
	0001  Speedlink Programming Interface
	0003  Airlink Wireless Programming Interface
16f5  Futurelogic Inc.
1706  BlueView Technologies, Inc.
1707  ARTIMI
170b  Swissonic
	0011  MIDI-USB 1x1
170d  Avnera
1711  Leica Microsystems
	0101  DFC-365FX camera
	3020  IC80 HD Camera
1724  Meyer Instruments (MIS)
	0115  PAXcam5
1725  Vitesse Semiconductor
1726  Axesstel, Inc.
	1000  wireless modem
	2000  wireless modem
	3000  wireless modem
172f  Waltop International Corp.
	0022  Tablet
	0024  Tablet
	0025  Tablet
	0026  Tablet
	0031  Slim Tablet 12.1"
	0032  Slim Tablet 5.8"
	0034  Slim Tablet 12.1"
	0038  Genius G-Pen F509
	0500  Media Tablet 14.1"
	0501  Media Tablet 10.6"
	0502  Sirius Battery Free Tablet
1733  Cellink Technology Co., Ltd
	0101  RF Wireless Optical Mouse OP-701
1736  CANON IMAGING SYSTEM TECHNOLOGIES INC.
1737  Linksys
	0039  USB1000 Gigabit Notebook Adapter
	0070  WUSB100 v1 RangePlus Wireless Network Adapter [Ralink RT2870]
	0071  WUSB600N v1 Dual-Band Wireless-N Network Adapter [Ralink RT2870]
	0073  WUSB54GC v2 802.11g Adapter [Realtek RTL8187B]
	0075  WUSB54GSC v2 802.11g Adapter [Broadcom 4326U]
	0077  WUSB54GC v3 802.11g Adapter [Ralink RT2070L]
	0078  WUSB100 v2 RangePlus Wireless Network Adapter [Ralink RT3070]
	0079  WUSB600N v2 Dual-Band Wireless-N Network Adapter [Ralink RT3572]
173d  QSENN
	0002  GP-K7000 keyboard
1740  Senao
	0100  EUB1200AC AC1200 DB Wireless Adapter [Realtek RTL8812AU]
	0600  EUB600v1 802.11abgn Wireless Adapter [Ralink RT3572]
	0605  LevelOne WUA-0605 N_Max Wireless USB Adapter
	0615  LevelOne WUA-0615 N_Max Wireless USB Adapter
	1000  NUB-350 802.11g Wireless Adapter [Intersil ISL3887]
	2000  NUB-8301 802.11bg
	3701  EUB-3701 EXT 802.11g Wireless Adapter [Ralink RT2571W]
	9603  RTL8188S WLAN Adapter
	9701  EnGenius 802.11n Wireless USB Adapter
	9702  EnGenius 802.11n Wireless USB Adapter
	9703  EnGenius 802.11n Wireless USB Adapter
	9705  EnGenius 802.11n Wireless USB Adapter
	9706  EUB9706 802.11n Wireless Adapter [Ralink RT3072]
	9801  EUB9801 802.11abgn Wireless Adapter [Ralink RT3572]
1743  General Atomics
1748  MQP Electronics
	0101  Packet-Master USB12
174c  ASMedia Technology Inc.
	1153  ASM2115 SATA 6Gb/s bridge
	2074  ASM1074 High-Speed hub
	3074  ASM1074 SuperSpeed hub
	5106  ASM1051 SATA 3Gb/s bridge
	5136  ASM1053 SATA 6Gb/s bridge
	55aa  ASM1051E SATA 6Gb/s bridge, ASM1053E SATA 6Gb/s bridge, ASM1153 SATA 3Gb/s bridge
174f  Syntek
	1105  SM-MS/Pro-MMC-XD Card Reader
	110b  HP Webcam
	1403  Integrated Webcam
	1404  USB Camera device, 1.3 MPixel Web Cam
	5212  USB 2.0 UVC PC Camera
	5a11  PC Camera
	5a31  Sonix USB 2.0 Camera
	5a35  Sonix 1.3MPixel USB 2.0 Camera
	6a31  Web Cam - Asus A8J, F3S, F5R, VX2S, V1S
	6a33  Web Cam - Asus F3SA, F9J, F9S
	6a51  2.0MPixel Web Cam - Asus Z96J, Z96S, S96S
	6a54  Web Cam
	6d51  2.0Mpixel Web Cam - Eurocom D900C
	8a12  Syntek 0.3MPixel USB 2.0 UVC PC Camera
	8a33  Syntek USB 2.0 UVC PC Camera
	a311  1.3MPixel Web Cam - Asus A3A, A6J, A6K, A6M, A6R, A6T, A6V, A7T, A7sv, A7U
	a312  1.3MPixel Web Cam
	a821  Web Cam - Packard Bell BU45, PB Easynote MX66-208W
	aa11  Web Cam
1753  GERTEC Telecomunicacoes Ltda.
	c901  PPC900 Pinpad Terminal
1756  ENENSYS Technologies
	0006  DiviPitch
1759  LucidPort Technology, Inc.
1761  ASUSTek Computer, Inc. (wrong ID)
	0b05  802.11n Network Adapter (wrong ID - swapped vendor and device)
1772  System Level Solutions, Inc.
1776  Arowana
	501c  300K CMOS Camera
177f  Sweex
	0004  MM004V5 Photo Key Chain (Digital Photo Frame) 1.5"
	0153  LW153 802.11n Adapter [ralink rt3070]
	0154  LW154 802.11bgn (1x1:1) Wireless Adapter [Realtek RTL8188SU]
	0313  LW313 802.11n Adapter [ralink rt2770 + rt2720]
1781  Multiple Vendors
	083e  MetaGeek Wi-Spy
	083f  MetaGeek Wi-Spy 2.4x
	0938  Iguanaworks USB IR Transceiver
	0a96  raphnet.net usb_game12
	0a97  raphnet.net SNES mouse adapter
	0a98  raphnet.net USBTenki
	0a99  raphnet.net NES
	0a9a  raphnet.net Gamecube/N64 controller
	0a9b  raphnet.net DB9Joy
	0a9c  raphnet.net Intellivision
	0a9d  raphnet.net 4nes4snes
	0a9e  raphnet.net Megadrive multitap
	0a9f  raphnet.net MultiDB9joy
	0c30  Telldus TellStick
	0c31  Telldus TellStick Duo
	0c9f  USBtiny
	1eef  OpenAPC SecuKey
	1ef0  E1701 Modular Controller Card
	1ef1  E1701 Modular Controller Card
1782  Spreadtrum Communications Inc.
1784  TopSeed Technology Corp.
	0001  eHome Infrared Transceiver
	0004  RF Combo Device
	0006  eHome Infrared Transceiver
	0007  eHome Infrared Transceiver
	0008  eHome Infrared Transceiver
	000a  eHome Infrared Transceiver
	0011  eHome Infrared Transceiver
1787  ATI AIB
1788  ShenZhen Litkconn Technology Co., Ltd.
1796  Printrex, Inc.
1797  JALCO CO., LTD.
1799  Thales Norway A/S
	7051  Belkin F5D7051 802.11g Adapter v1000 [Broadcom 4320]
	8051  Belkin F5D8051 v2 802.11bgn Wireless Adapter [Marvell 88W8362]
179d  Ricavision International, Inc.
	0010  Internal Infrared Transceiver
17a0  Samson Technologies Corp.
	0001  C01U condenser microphone
	0002  Q1U dynamic microphone
	0100  C03U multi-pattern microphone
	0101  UB1 boundary microphone
	0120  Meteorite condenser microphone
	0200  StudioDock monitors (internal hub)
	0201  StudioDock monitors (audio)
	0210  StudioGT monitors
	0301  Q2U handheld microphone with XLR
	0302  GoMic compact condenser microphone
	0303  C01U Pro condenser microphone
	0304  Q2U handheld mic with XLR
	0305  GoMic compact condenser mic
	0310  Meteor condenser microphone
17a4  Concept2
	0001  Performance Monitor 3
	0002  Performance Monitor 4
17a5  Advanced Connection Technology Inc.
17a7  MICOMSOFT CO., LTD.
17a8  Kamstrup A/S
	0001  Optical Eye/3-wire
	0005  M-Bus Master MultiPort 250D
17b3  Grey Innovation
	0004  Linux-USB Midi Gadget
17b5  Lunatone
	0010  MFT Sensor
17ba  SAURIS GmbH
	0001  SAU510-USB [no firmware]
	0510  SAU510-USB and SAU510-USB plus JTAG Emulators
	0511  SAU510-USB Iso Plus JTAG Emulator
	0520  SAU510-USB Nano JTAG Emulator
	1511  Onboard Emulator on SAUModule development kit
17c3  Singim International Corp.
17cc  Native Instruments
	041c  Audio 2 DJ
	0808  Maschine Controller
	0815  Audio Kontrol 1
	0839  Audio 4 DJ
	0d8d  Guitarrig Mobile
	1915  Session I/O
	1940  RigKontrol3
	1969  RigKontrol2
	1978  Audio 8 DJ
	2280  Medion MDPNA1500 in card reader mode
	2305  Traktor Kontrol X1
	4711  Kore Controller
	4712  Kore Controller 2
	baff  Traktor Kontrol S4
17cf  Hip Hing Cable & Plug Mfy. Ltd.
17d0  Sanford L.P.
17d3  Korea Techtron Co., Ltd.
17e9  DisplayLink
	0051  USB VGA Adaptor
	030b  HP T100
	0377  Plugable UD-160-A (M)
	0378  Plugable UGA-2K-A
	0379  Plugable UGA-125
	037a  Plugable UGA-165
	037b  Plugable USB-VGA-165
	037c  Plugable DC-125
	037d  Plugable USB2-HDMI-165
	410a  HDMI Adapter
	430a  HP Port Replicator (Composite Device)
	4312  S2340T
17eb  Cornice, Inc.
17ef  Lenovo
	1000  Hub
	1003  Integrated Smart Card Reader
	1004  Integrated Webcam
	1008  Hub
	100a  ThinkPad Mini Dock Plus Series 3
	304b  AX88179 Gigabit Ethernet [ThinkPad OneLink GigaLAN]
	3815  ChipsBnk 2GB USB Stick
	4802  Lenovo Vc0323+MI1310_SOC Camera
	4807  UVC Camera
	480c  Integrated Webcam
	480d  Integrated Webcam [R5U877]
	480e  Integrated Webcam [R5U877]
	480f  Integrated Webcam [R5U877]
	4810  Integrated Webcam [R5U877]
	4811  Integrated Webcam [R5U877]
	4812  Integrated Webcam [R5U877]
	4813  Integrated Webcam [R5U877]
	4814  Integrated Webcam [R5U877]
	4815  Integrated Webcam [R5U877]
	4816  Integrated Webcam
	481c  Integrated Webcam
	481d  Integrated Webcam
	6004  ISD-V4 Tablet Pen
	6007  Smartcard Keyboard
	6009  ThinkPad Keyboard with TrackPoint
	6014  Mini Wireless Keyboard N5901
	6025  ThinkPad Travel Mouse
	7203  Ethernet adapter [U2L 100P-Y1]
	7423  IdeaPad A1 Tablet
	7435  A789 (Mass Storage mode, with debug)
	743a  A789 (Mass Storage mode)
	7497  A789 (MTP mode)
	7498  A789 (MTP mode, with debug)
	749a  A789 (PTP mode)
	749b  A789 (PTP mode, with debug)
17f4  WaveSense
	aaaa  Jazz Blood Glucose Meter
17f5  K.K. Rocky
17f6  Unicomp, Inc
	0709  Model M Keyboard
1809  Advantech
	4604  USB-4604
	4761  USB-4761 Portable Data Acquisition Module
1822  Twinhan
	3201  VisionDTV USB-Ter/HAMA USB DVB-T device cold
	3202  VisionDTV USB-Ter/HAMA USB DVB-T device warm
1831  Gwo Jinn Industries Co., Ltd.
1832  Huizhou Shenghua Industrial Co., Ltd.
183d  VIVOphone
	0010  VoiceKey
1843  Vaisala
1849  ASRock Incorporation
1852  GYROCOM C&C Co., LTD
	7922  Audiotrak DR.DAC2 DX [GYROCOM C&C]
1854  Memory Devices Ltd.
185b  Compro
	3020  K100 Infrared Receiver
	3082  K100 Infrared Receiver v2
	d000  Compro Videomate DVB-U2000 - DVB-T USB cold
	d001  Compro Videomate DVB-U2000 - DVB-T USB warm
1861  Tech Technology Industrial Company
1862  Teridian Semiconductor Corp.
1870  Nexio Co., Ltd
	0001  iNexio Touchscreen controller
1871  Aveo Technology Corp.
	0101  UVC camera (Bresser microscope)
	0141  Camera
	0d01  USB2.0 Camera
1873  Navilock
	ee93  EasyLogger
187c  Alienware Corporation
	0511  AlienFX Mobile lighting
	0600  Dual Compatible Game Pad
187f  Siano Mobile Silicon
	0010  Stallar Board
	0100  Stallar Board
	0200  Nova A
	0201  Nova B
	0202  Nice
	0300  Vega
	0301  VeNice
1892  Vast Technologies, Inc.
1894  Topseed
	5632  Atek Tote Remote
	5641  TSAM-004 Presentation Remote
1897  Evertop Wire Cable Co.
189f  3Shape A/S
	0002  Legato2 3D Scanner
18a4  CSSN
	0001  Snapshell IDR
18a5  Verbatim, Ltd
	0214  Portable Hard Drive
	0216  External Hard Drive
	0218  External Hard Drive
	0224  Store 'n' Go Micro Plus
	0227  Pocket Hard Drive
	022b  Portable Hard Drive (Store'n'Go)
	0237  Portable Harddrive
	0243  Flash Drive (Store'n'Go)
	0302  Flash Drive
	0304  Store 'n' Go
	4123  Store N Go
18b1  Petalynx
	0037  Maxter Remote Control
18b4  e3C Technologies
	1001  DUTV007
	1002  EC168 (v5) based USB DVB-T receiver
	1689  DUTV009
	fffa  EC168 (v2) based USB DVB-T receiver
	fffb  EC168 (v3) based USB DVB-T receiver
18b6  Mikkon Technology Limited
18b7  Zotek Electronic Co., Ltd.
18c5  AMIT Technology, Inc.
	0002  CG-WLUSB2GO
	0008  CG-WLUSB2GNR Corega Wireless USB Adapter
	0012  CG-WLUSB10 Corega Wireless USB Adapter
18cd  Ecamm
	cafe  Pico iMage
18d1  Google Inc.
	0001  Onda V972 (storage access)
	0003  Android-powered device using AllWinner Technology SoC
	0006  Onda V972 MTP
	0008  Onda V972 PTP (camera)
	0d02  Celkon A88
	2d00  Android-powered device in accessory mode
	2d01  Android-powered device in accessory mode with ADB support
	4e11  Nexus One
	4e12  Nexus One (debug)
	4e13  Nexus One (tether)
	4e20  Nexus S (fastboot)
	4e21  Nexus S
	4e22  Nexus S (debug)
	4e24  Nexus S (tether)
	4e30  Galaxy Nexus (fastboot)
	4e40  Nexus 7 (fastboot)
	4e41  Nexus 7 (MTP)
	4e42  Nexus 7 (debug)
	4e43  Nexus 7 (PTP)
	4e44  Nexus 7 2012 (PTP)
	4ee0  Nexus 4 (bootloader)
	4ee1  Nexus Device (MTP)
	4ee2  Nexus Device (debug)
	4ee3  Nexus 4/5/7/10 (tether)
	4ee4  Nexus 4/5/7/10 (debug + tether)
	4ee5  Nexus 4 (PTP)
	4ee6  Nexus 4/5 (PTP + debug)
	7102  Toshiba Thrive tablet
	b004  Pandigital / B&N Novel 9" tablet
	d001  Nexus 4 (fastboot)
	d002  Nexus 4 (debug)
	d109  LG G2x MTP
	d10a  LG G2x MTP (debug)
18d5  Starline International Group Limited
18d9  Kaba
	01a0  B-Net 91 07
18dc  LKC Technologies, Inc.
18dd  Planon System Solutions Inc.
	1000  DocuPen RC800
18e3  Fitipower Integrated Technology Inc
	7102  Multi Card Reader (Internal)
	9101  All-in-1 Card Reader
	9102  Multi Card Reader
	9512  Webcam
18e8  Qcom
	6144  LR802UA 802.11b Wireless Adapter [ALi M4301AU]
	6196  RT2573
	6229  RT2573
	6232  Wireless 802.11g 54Mbps Network Adapter [RTL8187]
18ea  Matrox Graphics, Inc.
	0002  DualHead2Go [Analog Edition]
	0004  TripleHead2Go [Digital Edition]
18ec  Arkmicro Technologies Inc.
	3118  USB to IrDA adapter [ARK3116T]
	3188  ARK3188 UVC Webcam
	3299  Webcam Carrefour
	3366  Bresser Biolux NV
18fd  FineArch Inc.
1908  GEMBIRD
	1320  PhotoFrame PF-15-1
190d  Motorola GSG
1914  Alco Digital Devices Limited
1915  Nordic Semiconductor ASA
	000c  Wireless Desktop nRF24L01 CX-1766
	2233  Linksys WUSB11 v2.8 802.11b Adapter [Atmel AT76C505]
	2234  Linksys WUSB54G v1 OEM 802.11g Adapter [Intersil ISL3886]
	2235  Linksys WUSB54GP v1 OEM 802.11g Adapter [Intersil ISL3886]
	2236  Linksys WUSB11 v3.0 802.11b Adapter [Intersil PRISM 3]
1923  FitLinxx
	0002  Personal SyncPoint
1926  NextWindow
	0003  1900 HID Touchscreen
	0006  1950 HID Touchscreen
	0064  1950 HID Touchscreen
	0065  1950 HID Touchscreen
	0066  1950 HID Touchscreen
	0067  1950 HID Touchscreen
	0068  1950 HID Touchscreen
	0069  1950 HID Touchscreen
	0071  1950 HID Touchscreen
	0072  1950 HID Touchscreen
	0073  1950 HID Touchscreen
	0074  1950 HID Touchscreen
	0075  1950 HID Touchscreen
	0076  1950 HID Touchscreen
	0077  1950 HID Touchscreen
	0078  1950 HID Touchscreen
	0079  1950 HID Touchscreen
	007a  1950 HID Touchscreen
	007e  1950 HID Touchscreen
	007f  1950 HID Touchscreen
	0080  1950 HID Touchscreen
	0081  1950 HID Touchscreen
	0082  1950 HID Touchscreen
	0083  1950 HID Touchscreen
	0084  1950 HID Touchscreen
	0085  1950 HID Touchscreen
	0086  1950 HID Touchscreen
	0087  1950 HID Touchscreen
	0dc2  HID Touchscreen
192f  Avago Technologies, Pte.
	0000  Mouse
	0416  ADNS-5700 Optical Mouse Controller (3-button)
	0616  ADNS-5700 Optical Mouse Controller (5-button)
1930  Shenzhen Xianhe Technology Co., Ltd.
1931  Ningbo Broad Telecommunication Co., Ltd.
1934  Feature Integration Technology Inc. (Fintek)
	0602  F71610 or F71612 Consumer Infrared Receiver/Transceiver
	0702  Integrated Consumer Infrared Receiver/Transceiver
	5168  F71610A or F71612A Consumer Infrared Receiver/Transceiver
1941  Dream Link
	8021  WH1080 Weather Station / USB Missile Launcher
1943  Sensoray Co., Inc.
	2250  Model 2250 MPEG and JPEG Capture Card
	2253  Model 2253 Audio/Video Codec Card
	2255  Model 2255 4 Channel Capture Card
	2257  Model 2257 4 Channel Capture Card
	a250  Model 2250 MPEG and JPEG Capture Card (cold)
	a253  Model 2253 Audio/Video Codec Card (cold)
1949  Lab126, Inc.
	0002  Amazon Kindle
	0004  Amazon Kindle 3/4/Paperwhite
	0006  Kindle Fire
	0008  Amazon Kindle Fire HD 8.9"
194f  PreSonus Audio Electronics, Inc.
	0101  AudioBox 22 VSL
	0102  AudioBox 44 VSL
	0103  AudioBox 1818 VSL
	0301  AudioBox
1951  Hyperstone AG
1953  Ironkey Inc.
	0202  S200 2GB Rev. 1
1954  Radiient Technologies
195d  Itron Technology iONE
	7002  Libra-Q11 IR remote
	7006  Libra-Q26 / 1.0 Remote
	7777  Scorpius wireless keyboard
	7779  Scorpius-P20MT
1965  Uniden Corporation
	0016  HomePatrol-1
1967  CASIO HITACHI Mobile Communications Co., Ltd.
196b  Wispro Technology Inc.
1970  Dane-Elec Corp. USA
	0000  Z Mate 16GB
1975  Dongguan Guneetal Wire & Cable Co., Ltd.
1976  Chipsbrand Microelectronics (HK) Co., Ltd.
	6025  Flash Drive 512 MB
1977  T-Logic
	0111  TL203 MP3 Player and Voice Recorder
197d  Leuze electronic
	0222  BCL 508i
1989  Nuconn Technology Corp.
198f  Beceem Communications Inc.
	0210  BCS200 WiMAX Adapter
	0220  BCSM250 WiMAX Adapter
1990  Acron Precision Industrial Co., Ltd.
1995  Trillium Technology Pty. Ltd.
	3202  REC-ADPT-USB (recorder)
	3203  REC-A-ADPT-USB (recorder)
1996  PixeLINK
	3010  Camera Release 4
	3011  OEM Camera
	3012  e-ImageData Corp. ScanPro
199b  MicroStrain, Inc.
	3065  3DM-GX3-25 Orientation Sensor
199e  The Imaging Source Europe GmbH
	8101  DFx 21BU04 Camera
199f  Benica Corporation
19a8  Biforst Technology Inc.
19ab  Bodelin
	1000  ProScope HR
19af  S Life
	6611  Celestia VoIP Phone
19b2  Batronix
	0010  BX32 Batupo
	0011  BX32P Barlino
	0012  BX40 Bagero
	0013  BX48 Batego
19b4  Celestron
	0002  SkyScout Personal Planetarium
	0101  Handheld Digital Microscope 44302
19b5  B & W Group
19b6  Infotech Logistic, LLC
19b9  Data Robotics
	8d20  Drobo Elite
19c2  Futuba
	6a11  MDM166A Fluorescent Display
19ca  Mindtribe
	0001  Sandio 3D HID Mouse
19cf  Parrot SA
19d2  ZTE WCDMA Technologies MSM
	0001  CDMA Wireless Modem
	0002  MF632/ONDA ET502HS/MT505UP
	0007  TU25 WiMAX Adapter [Beceem BCS200]
	0031  MF110/MF627/MF636
	0063  K3565-Z HSDPA
	0064  MF627 AU
	0083  MF190
	0103  MF112
	0104  K4505-Z
	0146  MF 195E (HSPA+ Modem)
	0167  MF820 4G LTE
	0172  AX226 WIMAX MODEM (After Modeswitch)
	0325  LTE4G O2 ZTE MF821D LTE/UMTS/GSM Modem/Networkcard
	0326  LTE4G O2 ZTE MF821D LTE/UMTS/GSM Modem/Networkcard
	1008  K3570-Z
	1010  K3571-Z
	1017  K5006-Z vodafone LTE/UMTS/GSM Modem/Networkcard
	1018  K5006-Z vodafone LTE/UMTS/GSM Modem/Networkcard
	1203  MF691 [ T-Mobile webConnect Rocket 2.0]
	1217  MF652
	1218  MF652
	2000  MF627/MF628/MF628+/MF636+ HSDPA/HSUPA
	fff2  Gobi Wireless Modem (QDL mode)
	fff3  Gobi Wireless Modem
19db  KFI Printers
	02f1  NAUT324C
19e1  WeiDuan Electronic Accessory (S.Z.) Co., Ltd.
19e8  Industrial Technology Research Institute
19ef  Pak Heng Technology (Shenzhen) Co., Ltd.
19f7  RODE Microphones
	0001  Podcaster
19fa  Gampaq Co.Ltd
	0703  Steering Wheel
19ff  Dynex
	0102  1.3MP Webcam
	0201  Rocketfish Wireless 2.4G Laser Mouse
	0238  DX-WRM1401 Mouse
1a08  Bellwood International, Inc.
1a0a  USB-IF non-workshop
	badd  USB OTG Compliance test device
1a12  KES Co., Ltd.
1a1d  Veho
	0407  Mimi WiFi speakers
1a25  Amphenol East Asia Ltd.
1a2a  Seagate Branded Solutions
1a2c  China Resource Semico Co., Ltd
	0021  Keyboard
	0024  Multimedia Keyboard
1a32  Quanta Microsystems, Inc.
	0304  802.11n Wireless LAN Card
1a34  ACRUX
	0802  Gamepad
1a36  Biwin Technology Ltd.
1a40  Terminus Technology Inc.
	0101  Hub
	0201  FE 2.1 7-port Hub
1a41  Action Electronics Co., Ltd.
1a44  VASCO Data Security International
	0001  Digipass 905 SmartCard Reader
1a4a  Silicon Image
1a4b  SafeBoot International B.V.
1a5a  Tandberg Data
1a61  Abbott Diabetes Care
	3410  CoPilot System Cable
1a6a  Spansion Inc.
1a6d  SamYoung Electronics Co., Ltd
1a6e  Global Unichip Corp.
1a6f  Sagem Orga GmbH
1a72  Physik Instrumente
	1008  E-861 PiezoWalk NEXACT Controller
1a79  Bayer Health Care LLC
	6002  Contour
	7410  Contour Next
1a7b  Lumberg Connect  GmbH & Co. KG
1a7c  Evoluent
	0068  VerticalMouse 3
	0168  VerticalMouse 3 Wireless
	0191  VerticalMouse 4
1a81  Holtek Semiconductor, Inc.
	2203  Laser Gaming mouse
	2204  Optical Mouse
	2205  Laser Mouse
1a86  QinHeng Electronics
	5512  CH341 in EPP/MEM/I2C mode, EPP/I2C adapter
	5523  CH341 in serial mode, usb to serial port converter
	5584  CH341 in parallel mode, usb to printer port converter
	7523  HL-340 USB-Serial adapter
	752d  CH345 MIDI adapter
	7584  CH340S
	e008  HID-based serial adapater
1a89  Dynalith Systems Co., Ltd.
1a8b  SGS Taiwan Ltd.
1a8d  BandRich, Inc.
	1002  BandLuxe 3.5G HSDPA Adapter
	1009  BandLuxe 3.5G HSPA Adapter
	100d  4G LTE adapter
1a98  Leica Camera AG
1aa4  Data Drive Thru, Inc.
1aa5  UBeacon Technologies, Inc.
1aa6  eFortune Technology Corp.
1aad  KeeTouch
	0001  Touchscreen
1ab1  Rigol Technologies
	0588  DS1000 SERIES
1acb  Salcomp Plc
1acc  Midiplus Co, Ltd.
	0103  AudioLink plus 4x4 2.9.28
1ad1  Desay Wire Co., Ltd.
1ad4  APS
	0002  KM290-HRS
1adb  SEL C662 Serial Cable
1ae4  ic-design Reinhard Gottinger GmbH
1ae7  X-TENSIONS
	0381  VS-DVB-T 380U (af9015 based)
	2001  SpeedLink Snappy Mic webcam (SL-6825-SBK)
	9003  SpeedLink Vicious And Devine Laplace webcam, white (VD-1504-SWT)
	9004  SpeedLink Vicious And Devine Laplace webcam, black (VD-1504-SBK)
1aed  High Top Precision Electronic Co., Ltd.
1aef  Conntech Electronic (Suzhou) Corporation
1af1  Connect One Ltd.
1afe  A. Eberle GmbH & Co. KG
	0001  PQ Box 100
1b04  Meilhaus Electronic GmbH
	0630  ME-630
	0940  ME-94
	0950  ME-95
	0960  ME-96
	1000  ME-1000
	100a  ME-1000
	100b  ME-1000
	1400  ME-1400
	140a  ME-1400A
	140b  ME-1400B
	140c  ME-1400C
	140d  ME-1400D
	140e  ME-1400E
	14ea  ME-1400EA
	14eb  ME-1400EB
	1604  ME-1600/4U
	1608  ME-1600/8U
	160c  ME-1600/12U
	160f  ME-1600/16U
	168f  ME-1600/16U8I
	4610  ME-4610
	4650  ME-4650
	4660  ME-4660
	4661  ME-4660I
	4662  ME-4660
	4663  ME-4660I
	4670  ME-4670
	4671  ME-4670I
	4672  ME-4670S
	4673  ME-4670IS
	4680  ME-4680
	4681  ME-4680I
	4682  ME-4680S
	4683  ME-4680IS
	6004  ME-6000/4
	6008  ME-6000/8
	600f  ME-6000/16
	6014  ME-6000I/4
	6018  ME-6000I/8
	601f  ME-6000I/16
	6034  ME-6000ISLE/4
	6038  ME-6000ISLE/8
	603f  ME-6000ISLE/16
	6044  ME-6000/4/DIO
	6048  ME-6000/8/DIO
	604f  ME-6000/16/DIO
	6054  ME-6000I/4/DIO
	6058  ME-6000I/8/DIO
	605f  ME-6000I/16/DIO
	6074  ME-6000ISLE/4/DIO
	6078  ME-6000ISLE/8/DIO
	607f  ME-6000ISLE/16/DIO
	6104  ME-6100/4
	6108  ME-6100/8
	610f  ME-6100/16
	6114  ME-6100I/4
	6118  ME-6100I/8
	611f  ME-6100I/16
	6134  ME-6100ISLE/4
	6138  ME-6100ISLE/8
	613f  ME-6100ISLE/16
	6144  ME-6100/4/DIO
	6148  ME-6100/8/DIO
	614f  ME-6100/16/DIO
	6154  ME-6100I/4/DIO
	6158  ME-6100I/8/DIO
	615f  ME-6100I/16/DIO
	6174  ME-6100ISLE/4/DIO
	6178  ME-6100ISLE/8/DIO
	617f  ME-6100ISLE/16/DIO
	6259  ME-6200I/9/DIO
	6359  ME-6300I/9/DIO
	810a  ME-8100A
	810b  ME-8100B
	820a  ME-8200A
	820b  ME-8200B
1b0e  BLUTRONICS S.r.l.
	1078  BLUDRIVE II CCID
	1079  BLUDRIVE II CCID
	1080  WRITECHIP II CCID
1b1c  Corsair
	0890  Flash Padlock
	0a00  SP2500 Speakers
	0a60  Vengeance K60 Keyboard
	1a01  Flash Voyager GT
	1a0a  Survivor Stealth Flash Drive
	1a90  Flash Voyager GT
1b20  MStar Semiconductor, Inc.
1b22  WiLinx Corp.
1b26  Cellex Power Products, Inc.
1b27  Current Electronics Inc.
1b28  NAVIsis Inc.
1b32  Ugobe Life Forms, Inc.
	0064  Pleo robotic dinosaur
1b36  ViXS Systems, Inc.
1b3b  iPassion Technology Inc.
	2933  PC Camera/Webcam controller
	2935  PC Camera/Webcam controller
	2936  PC Camera/Webcam controller
	2937  PC Camera/Webcam controller
	2938  PC Camera/Webcam controller
	2939  PC Camera/Webcam controller
	2950  PC Camera/Webcam controller
	2951  PC Camera/Webcam controller
	2952  PC Camera/Webcam controller
	2953  PC Camera/Webcam controller
	2955  PC Camera/Webcam controller
	2956  PC Camera/Webcam controller
	2957  PC Camera/Webcam controller
	2958  PC Camera/Webcam controller
	2959  PC Camera/Webcam controller
	2960  PC Camera/Webcam controller
	2961  PC Camera/Webcam controller
	2962  PC Camera/Webcam controller
	2963  PC Camera/Webcam controller
	2965  PC Camera/Webcam controller
	2966  PC Camera/Webcam controller
	2967  PC Camera/Webcam controller
	2968  PC Camera/Webcam controller
	2969  PC Camera/Webcam controller
1b3f  Generalplus Technology Inc.
	0c52  808 Camera #9 (mass storage mode)
	2002  808 Camera #9 (web-cam mode)
1b47  Energizer Holdings, Inc.
	0001  CHUSB Duo Charger (NiMH AA/AAA USB smart charger)
1b48  Plastron Precision Co., Ltd.
1b52  ARH Inc.
	2101  FXMC Neural Network Controller
	2102  FXMC Neural Network Controller V2
	2103  FXMC Neural Network Controller V3
	4101  Passport Reader CLR device
	4201  Passport Reader PRM device
	4202  Passport Reader PRM extension device
	4203  Passport Reader PRM DSP device
	4204  Passport Reader PRMC device
	4205  Passport Reader CSHR device
	4206  Passport Reader PRMC V2 device
	4301  Passport Reader MRZ device
	4302  Passport Reader MRZ DSP device
	4303  Passport Reader CSLR device
	4401  Card Reader
	4501  Passport Reader RFID device
	4502  Passport Reader RFID AIG device
	6101  Neural Network Controller
	6202  Fingerprint Reader device
	6203  Fingerprint Scanner device
	8101  Camera V1
	8102  Recovery / Camera V2
	8103  Camera V3
1b59  K.S. Terminals Inc.
1b5a  Chao Zhou Kai Yuan Electric Co., Ltd.
1b65  The Hong Kong Standards and Testing Centre Ltd.
1b71  Fushicai
	3002  USBTV007 Video Grabber [EasyCAP]
1b72  ATERGI TECHNOLOGY CO., LTD.
1b73  Fresco Logic
	1000  xHC1 Controller
1b75  Ovislink Corp.
	3072  AirLive WN-360USB adapter
	8171  WN-370USB 802.11bgn Wireless Adapter [Realtek RTL8188SU]
	8187  AirLive WL-1600USB 802.11g Adapter [Realtek RTL8187L]
	9170  AirLive X.USB 802.11abgn [Atheros AR9170+AR9104]
	a200  AirLive WN-200USB wireless 11b/g/n dongle
1b76  Legend Silicon Corp.
1b80  Afatech
	c810  MC810 [af9015]
	d393  DVB-T receiver [RTL2832U]
	d396  UB396-T [RTL2832U]
	d397  DVB-T receiver [RTL2832U]
	d398  DVB-T receiver [RTL2832U]
	d700  FM Radio SnapMusic Mobile 700 (FM700)
	e297  Conceptronic DVB-T CTVDIGRCU V3.0
	e383  DVB-T UB383-T [af9015]
	e385  DVB-T UB385-T [af9015]
	e386  DVB-T UB385-T [af9015]
	e399  DVB-T KWorld PlusTV 399U [af9015]
	e39a  DVB-T395U [af9015]
	e39b  DVB-T395U [af9015]
	e401  Sveon STV22 DVB-T [af9015]
	e409  IT9137FN Dual DVB-T [KWorld UB499-2T]
1b86  Dongguan Guanshang Electronics Co., Ltd.
1b88  ShenMing Electron (Dong Guan) Co., Ltd.
1b8c  Altium Limited
1b8d  e-MOVE Technology Co., Ltd.
1b8e  Amlogic, Inc.
1b8f  MA LABS, Inc.
1b96  N-Trig
	0001  Duosense Transparent Electromagnetic Digitizer
1b98  YMax Communications Corp.
1b99  Shenzhen Yuanchuan Electronic
1ba1  JINQ CHERN ENTERPRISE CO., LTD.
1ba2  Lite Metals & Plastic (Shenzhen) Co., Ltd.
1ba4  Ember Corporation
	0001  InSight USB Link
1ba6  Abilis Systems
1ba8  China Telecommunication Technology Labs
1bad  Harmonix Music
	0002  Guitar for Xbox 360
	0003  Drum Kit for Xbox 360
1bae  Vuzix Corporation
	0002  VR920 Immersive Eyewear
1bbb  T & A Mobile Phones
	011e  Alcatel One Touch L100V / Telekom Speedstick LTE II
	f017  Alcatel One Touch L100V / Telekom Speedstick LTE II
1bc4  Ford Motor Co.
1bc5  AVIXE Technology (China) Ltd.
1bc7  Telit Wireless Solutions
	0020  HE863
	0021  HE910
	0023  HE910-D ECM
	1003  UC864-E
	1004  UC864-G
	1005  CC864-DUAL
	1006  CC864-SINGLE
	1010  DE910-DUAL
	1011  CE910-DUAL
	1200  LE920
1bce  Contac Cable Industrial Limited
1bcf  Sunplus Innovation Technology Inc.
	0005  Optical Mouse
	0007  Optical Mouse
	053a  Targa Silvercrest OMC807-C optische Funkmaus
	05c5  SPRF2413A [2.4GHz Wireless Keyboard/Mouse Receiver]
	05cf  Micro keyboard & mouse receiver
	0c31  SPIF30x Serial-ATA bridge
	2880  Dell HD Webcam
	2885  ASUS Webcam
	2888  HP Universal Camera
	28a2  Dell Integrated Webcam
	28a6  DELL XPS Integrated Webcam
	28ae  Laptop Integrated Webcam HD
	28bd  Dell Integrated HD Webcam
	2985  Laptop Integrated Webcam HD
	2b83  Laptop Integrated Webcam FHD
1bd0  Hangzhou Riyue Electronic Co., Ltd.
1bd5  BG Systems, Inc.
1bde  P-TWO INDUSTRIES, INC.
1bef  Shenzhen Tongyuan Network-Communication Cables Co., Ltd
1bf0  RealVision Inc.
1bf5  Extranet Systems Inc.
1bf6  Orient Semiconductor Electronics, Ltd.
1bfd  TouchPack
	1268  Touch Screen
	1368  Touch Screen
	1568  Capacitive Touch Screen
	1668  IR Touch Screen
	1688  Resistive Touch Screen
	2968  Touch Screen
	5968  Touch Screen
	6968  Touch Screen
1c02  Kreton Corporation
1c04  QNAP System Inc.
1c0c  Ionics EMS, Inc.
	0102  Plug Computer
1c0d  Relm Wireless
1c10  Lanterra Industrial Co., Ltd.
1c13  ALECTRONIC LIMITED
1c1a  Datel Electronics Ltd.
1c1b  Volkswagen of America, Inc.
1c1f  Goldvish S.A.
1c20  Fuji Electric Device Technology Co., Ltd.
1c21  ADDMM LLC
1c22  ZHONGSHAN CHIANG YU ELECTRIC CO., LTD.
1c26  Shanghai Haiying Electronics Co., Ltd.
1c27  HuiYang D & S Cable Co., Ltd.
1c29  Elster GmbH
	0001  ExMFE5 Simulator
	10fc  enCore device
1c31  LS Cable Ltd.
1c34  SpringCard
	7241  Prox'N'Roll RFID Scanner
1c37  Authorizer Technologies, Inc.
1c3d  NONIN MEDICAL INC.
1c3e  Wep Peripherals
1c40  EZPrototypes
	0533  TiltStick
	0534  i2c-tiny-usb interface
	0535  glcd2usb interface
	0536  Swiss ColorPAL
1c49  Cherng Weei Technology Corp.
1c4f  SiGma Micro
	0002  Keyboard TRACER Gamma Ivory
	0003  HID controller
	000e  Genius KB-120 Keyboard
	0026  Keyboard
	3000  Micro USB Web Camera
	3002  WebCam SiGma Micro
1c6b  Philips & Lite-ON Digital Solutions Corporation
	a222  DVD Writer Slimtype eTAU108
1c6c  Skydigital Inc.
1c73  AMT
	861f  Anysee E30 USB 2.0 DVB-T Receiver
1c77  Kaetat Industrial Co., Ltd.
1c78  Datascope Corp.
1c79  Unigen Corporation
1c7a  LighTuning Technology Inc.
	0801  Fingerprint Reader
1c7b  LUXSHARE PRECISION INDUSTRY (SHENZHEN) CO., LTD.
1c83  Schomaecker GmbH
	0001  RS150 V2
1c87  2N TELEKOMUNIKACE a.s.
1c88  Somagic, Inc.
	0007  SMI Grabber (EasyCAP DC60+ clone) (no firmware) [SMI-2021CBE]
	003c  SMI Grabber (EasyCAP DC60+ clone) [SMI-2021CBE]
1c89  HONGKONG WEIDIDA ELECTRON LIMITED
1c8e  ASTRON INTERNATIONAL CORP.
1c98  ALPINE ELECTRONICS, INC.
1c9e  OMEGA TECHNOLOGY
	6061  WL-72B 3.5G MODEM
1ca0  ACCARIO Inc.
1ca1  Symwave
	18ab  SATA bridge
1cac  Kinstone
	a332  C8 Webcam
	b288  C18 Webcam
1cb3  Aces Electronic Co., Ltd.
1cb4  OPEX CORPORATION
1cb6  IdeaCom Technology Inc.
	6681  IDC6681
1cbe  Luminary Micro Inc.
	00fd  In-Circuit Debug Interface
	00ff  Stellaris ROM DFU Bootloader
	0166  CANAL USB2CAN
1cbf  FORTAT SKYMARK INDUSTRIAL COMPANY
1cc0  PlantSense
1cca  NextWave Broadband Inc.
1ccd  Bodatong Technology (Shenzhen) Co., Ltd.
1cd4  adp corporation
1cd5  Firecomms Ltd.
1cd6  Antonio Precise Products Manufactory Ltd.
1cde  Telecommunications Technology Association (TTA)
1cdf  WonTen Technology Co., Ltd.
1ce0  EDIMAX TECHNOLOGY CO., LTD.
1ce1  Amphenol KAE
1cf1  Dresden Elektronik
	0001  Sensor Terminal Board
	0004  Wireless Handheld Terminal
	0017  deRFusbSniffer 2.4 GHz
	0018  deRFusb24E001
	0019  deRFusb14E001
	001a  deRFusb23E00
	001b  deRFusb13E00
	001c  deRFnode
	001d  deRFnode / gateway
	0022  deUSB level shifter
	0023  deRFusbSniffer Sub-GHz
	0025  deRFusb23E06
	0027  deRFusb13E06
1cfc  ANDES TECHNOLOGY CORPORATION
1cfd  Flextronics Digital Design Japan, LTD.
1d03  iCON
	0028  iCreativ MIDI Controller
1d07  Solid-Motion
1d08  NINGBO HENTEK DRAGON ELECTRONICS CO., LTD.
1d09  TechFaith Wireless Technology Limited
	1026  HSUPA Modem FLYING-LARK46-VER0.07 [Flying Angel]
1d0a  Johnson Controls, Inc. The Automotive Business Unit
1d0b  HAN HUA CABLE & WIRE TECHNOLOGY (J.X.) CO., LTD.
1d0f  Sonix Technology Co., Ltd.
1d14  ALPHA-SAT TECHNOLOGY LIMITED
1d17  C-Thru Music Ltd.
	0001  AXiS-49 Harmonic Table MIDI Keyboard
1d19  Dexatek Technology Ltd.
	1101  DK DVB-T Dongle
	1102  DK mini DVB-T Dongle
	1103  DK 5217 DVB-T Dongle
	6105  Video grabber
	8202  DK DVBC/T DONGLE
1d1f  Diostech Co., Ltd.
1d20  SAMTACK INC.
1d27  ASUS
1d34  Dream Cheeky
	0001  Dream Cheeky Fidget
	0004  Dream Cheeky Webmail Notifier
	0008  Dream Cheeky button
	000a  Dream Cheeky Mailbox Friends Alert
	000d  Dream Cheeky Big Red Button
	0013  Dream Cheeky LED Message Board
1d45  Touch
	1d45  Foxlink Optical touch sensor
1d4d  PEGATRON CORPORATION
	0002  Ralink RT2770/2720 802.11b/g/n Wireless LAN Mini-USB Device
	000c  Ralink RT3070 802.11b/g/n Wireless Lan USB Device
	000e  Ralink RT3070 802.11b/g/n Wireless Lan USB Device
1d50  OpenMoko, Inc.
	5119  GTA01/GTA02 U-Boot Bootloader
	602b  FPGALink
	6053  Darkgame Controller
1d57  Xenta
	0005  Wireless Receiver (Keyboard and Mouse)
	0006  Wireless Receiver (RC Laser Pointer)
	000c  Optical Mouse
	2400  Wireless Mouse Receiver
	32da  2.4GHz Receiver (Keyboard and Mouse)
	83d0  Click-mouse!
	ac01  Wireless Receiver (Keyboard and Mouse)
	ad02  SE340D PC Remote Control
	af01  AUVIO Universal Remote Receiver for PlayStation 3
1d5b  Smartronix, Inc.
1d6b  Linux Foundation
	0001  1.1 root hub
	0002  2.0 root hub
	0003  3.0 root hub
	0100  PTP Gadget
	0101  Audio Gadget
	0102  EEM Gadget
	0103  NCM (Ethernet) Gadget
	0104  Multifunction Composite Gadget
	0105  FunctionFS Gadget
	0200  Qemu Audio Device
1d90  Citizen
	201e  PPU-700
1de1  Actions Microelectronics Co.
	1101  Generic Display Device (Mass storage mode)
	c101  Generic Display Device
1e0e  Qualcomm / Option
	f000  iCON 210 UMTS Surfstick
1e10  Point Grey Research, Inc.
	2004  Sony 1.3MP 1/3" ICX445 IIDC video camera [Chameleon]
1e17  Mirion Technologies Dosimetry Services Division
	0001  instadose dosimeter
1e1d  Lumension Security
	0165  Secure Pen drive
1e1f  INVIA
1e29  Festo AG & Co. KG
	0101  CPX Adapter
	0102  CPX Adapter >=HW10.09 [CP2102]
	0401  iL3-TP [AT90USB646]
	0402  FTDI232 [EasyPort]
	0403  FTDI232 [EasyPort Mini]
	0404  FTDI232 [Netzteil-GL]
	0405  FTDI232 [MotorPrfstand]
	0406  STM32F103 [EasyKit]
	0407  LPC2378 [Robotino]
	0408  LPC2378 [Robotino-Arm]
	0409  LPC2378 [Robotino-Arm Bootloader]
	040a  LPC2378 [Robotino Bootloader]
	040b  LPC2378 [Robotino XT]
	040c  LPC2378 [Robotino XT Bootloader]
	040d  LPC2378 [Robotino 3]
	040e  LPC2378 [Robotino 3 Bootloader]
	0501  CP2102 [CMSP]
	0601  CMMP-AS
1e3d  Chipsbank Microelectronics Co., Ltd
	2093  CBM209x Flash Drive (OEM)
	4082  CBM4082 SD Card Reader
1e41  Cleverscope
	0001  CS328A PC Oscilloscope
1e4e  Cubeternet
	0100  WebCam
	0102  GL-UPC822 UVC WebCam
1e54  TypeMatrix
	2030  2030 USB Keyboard
1e68  TrekStor GmbH & Co. KG
	001b  DataStation maxi g.u
	0050  DataStation maxi light
1e71  NZXT
	0001  Avatar Optical Mouse
1e74  Coby Electronics Corporation
	2211  MP300
	2647  2 GB 2 Go Video MP3 Player [MP601-2G]
	2659  Coby 4GB Go Video MP3 Player [MP620-4G]
	4641  A8705 MP3/Video Player
	6511  MP705-8G MP3 player
	6512  MP705-4G
	7111  MP957 Music and Video Player
1e7d  ROCCAT
	2c24  Pyra Mouse (wired)
	2ced  Kone Mouse
	2cf6  Pyra Mouse (wireless)
	2d50  Kova+ Mouse
	2d51  Kone+ Mouse
	30d4  Arvo Keyboard
1ebb  NuCORE Technology, Inc.
1eda  AirTies Wireless Networks
	2012  Air2210 54 Mbps Wireless Adapter
	2210  Air2210 54 Mbps Wireless Adapter
	2310  Air2310 150 Mbps Wireless Adapter
	2410  Air2410 300 Mbps Wireless Adapter
1edb  Blackmagic design
	bd3b  Intensity Shuttle
1ee8  ONDA COMMUNICATION S.p.a.
	0014  MT833UP
1ef6  EADS Deutschland GmbH
	2233  Cassidian NH90 STTE
	5064  FDR Interface
	5523  Cassidian SSDC Adapter II
	5545  Cassidian SSDC Adapter III
	5648  RIU CSMU/BSD
	564a  Cassidian RIU CSMU/BSD Simulator
1f28  Cal-Comp
	0020  CDMA USB Modem A600
	0021  CD INSTALLER USB Device
1f3a  Onda (unverified)
	efe8  V972 tablet in flashing mode
1f44  The Neat Company
	0001  NM-1000 scanner
1f48  H-TRONIC GmbH
	0627  Data capturing system
	0628  Data capturing and control module
1f4d  G-Tek Electronics Group
	b803  Lifeview LV5TDLX DVB-T [RTL2832U]
	d220  Geniatech T220 DVB-T2 TV Stick
1f6f  Aliph
	0023  Jawbone Jambox
	8000  Jawbone Jambox - Updating
1f75  Innostor Technology Corporation
	0888  IS888 SATA Storage Controller
	0902  IS902 UFD controller
1f82  TANDBERG
	0001  PrecisionHD Camera
1f84  Alere, Inc.
1f87  Stantum
	0002  Multi-touch HID Controller
1f9b  Ubiquiti Networks, Inc.
	0241  AirView2-EXT
1fab  Samsung Opto-Electroncs Co., Ltd.
	104d  ES65
1fbd  Delphin Technology AG
	0001  Expert Key - Data aquisition system
1fc9  NXP Semiconductors
	0003  LPC1343
	010b  PR533
1fde  ILX Lightwave Corporation
	0001  UART Bridge
1fe7  Vertex Wireless Co., Ltd.
	1000  VW100 series CDMA EV-DO Rev.A modem
1ff7  CVT Electronics.Co.,Ltd
	0013  CVTouch Screen (HID)
	001a  Human Interface Device
1fff  Ideofy Inc.
2001  D-Link Corp.
	0001  DWL-120 WIRELESS ADAPTER
	0201  DHN-120 10Mb Home Phoneline Adapter
	1a00  DUB-E100 Fast Ethernet Adapter(rev.A) [ASIX AX88172]
	1a02  DUB-E100 Fast Ethernet Adapter(rev.C1) [ASIX AX88772]
	200c  10/100 Ethernet
	3200  DWL-120 802.11b Wireless Adapter(rev.E1) [Atmel at76c503a]
	3301  DWA-130 802.11n Wireless N Adapter(rev.C1) [Realtek RTL8192U]
	3306  DWL-G122 Wireless Adapter(rev.F1) [Realtek RTL8188SU]
	3308  DWA-121 802.11n Wireless N 150 Pico Adapter [Realtek RTL8188CUS]
	3309  DWA-135 802.11n Wireless N Adapter(rev.A1) [Realtek RTL8192CU]
	330a  DWA-133 802.11n Wireless N Adapter [Realtek RTL8192CU]
	3500  Elitegroup Computer Systems WLAN card WL-162
	3700  DWL-122 802.11b [Intersil Prism 3]
	3701  DWL-G120 Spinnaker 802.11g [Intersil ISL3886]
	3702  DWL-120 802.11b Wireless Adapter(rev.F) [Intersil ISL3871]
	3703  AirPlus G DWL-G122 Wireless Adapter(rev.A1) [Intersil ISL3880]
	3704  AirPlus G DWL-G122 Wireless Adapter(rev.A2) [Intersil ISL3887]
	3705  AirPlus G DWL-G120 Wireless Adapter(rev.C) [Intersil ISL3887]
	3761  IEEE 802.11g USB2.0 Wireless Network Adapter-PN
	3a00  DWL-AG132 [Atheros AR5523]
	3a01  DWL-AG132 (no firmware) [Atheros AR5523]
	3a02  DWL-G132 [Atheros AR5523]
	3a03  DWL-G132 (no firmware) [Atheros AR5523]
	3a04  DWL-AG122 [Atheros AR5523]
	3a05  DWL-AG122 (no firmware) [Atheros AR5523]
	3a80  AirPlus Xtreme G DWL-G132 Wireless Adapter
	3a81  predator Bootloader Download
	3a82  AirPremier AG DWL-AG132 Wireless Adapter
	3a83  predator Bootloader Download
	3b00  AirPlus DWL-120+ Wireless Adapter [Texas Instruments ACX100USB]
	3b01  WLAN Boot Device
	3c00  AirPlus G DWL-G122 Wireless Adapter(rev.B1) [Ralink RT2571]
	3c01  AirPlus AG DWL-AG122 Wireless Adapter
	3c02  AirPlus G DWL-G122 Wireless Adapter
	3c05  DUB-E100 Fast Ethernet Adapter(rev.B1) [ASIX AX88772]
	3c15  DWA-140 RangeBooster N Adapter(rev.B3) [Ralink RT5372]
	3c17  DWA-123 Wireless N 150 Adapter(rev.A1) [Ralink RT3370]
	3c19  DWA-125 Wireless N 150 Adapter(rev.A3) [Ralink RT5370]
	3c1a  DWA-160 802.11abgn Xtreme N Dual Band Adapter(rev.B2) [Ralink RT5572]
	3c1b  DWA-127 Wireless N 150 High-Gain Adapter(rev.A1) [Ralink RT3070]
	4000  DSB-650C Ethernet [klsi]
	4001  DSB-650TX Ethernet [pegasus]
	4002  DSB-650TX Ethernet [pegasus]
	4003  DSB-650TX-PNA Ethernet [pegasus]
	400b  10/100 Ethernet
	4102  10/100 Ethernet
	5100  DSL-200 ADSL ATM Modem
	5102  DSL-200 ADSL Loader
	5b00  Remote NDIS Network Device
	9414  Cable Modem
	9b00  Broadband Cable Modem Remote NDIS Device
	abc1  DSB-650 Ethernet [pegasus]
	f013  DLink 7 port USB2.0 Hub
	f103  DUB-H7 7-port USB 2.0 hub
	f10d  Accent Communications Modem
	f110  DUB-AV300 A/V Capture
	f111  DBT-122 Bluetooth adapter
	f112  DUB-T210 Audio Device
	f116  Formosa 2
	f117  Formosa 3
	f118  Formosa 4
2002  DAP Technologies
2003  detectomat
	ea61  dc3500
200c  Reloop
	100b  Play audio soundcard
2013  PCTV Systems
	0245  PCTV 73ESE
	0246  PCTV 74E
	0248  PCTV 282E
	024f  nanoStick T2 290e
2019  PLANEX
	3220  GW-US11S WLAN [Atmel AT76C503A]
	4901  GW-USSuper300 802.11bgn Wireless Adapter [Realtek RTL8191SU]
	4903  GW-USFang300 802.11abgn Wireless Adapter [Realtek RTL8192DU]
	4904  GW-USUltra300 802.11abgn Wireless Adapter [Realtek RTL8192DU]
	5303  GW-US54GXS 802.11bg
	5304  GWUS300 802.11n
	ab01  GW-US54HP
	ab24  GW-US300MiniS
	ab25  GW-USMini2N 802.11n Wireless Adapter [Ralink RT2870]
	ab28  GW-USNano
	ab29  GW-USMicro300
	ab2a  GW-USNano2 802.11n Wireless Adapter [Realtek RTL8188CUS]
	ab2b  GW-USEco300 802.11bgn Wireless Adapter [Realtek RTL8192CU]
	ab2c  GW-USDual300 802.11abgn Wireless Adapter [Realtek RTL8192DU]
	ab50  GW-US54Mini2
	c002  GW-US54SG
	c007  GW-US54GZL
	ed02  GW-USMM
	ed06  GW-US300MiniW 802.11bgn Wireless Adapter
	ed10  GW-US300Mini2
	ed14  GW-USMicroN
	ed16  GW-USMicroN2W 802.11bgn Wireless Adapter [Realtek RTL8188SU]
	ed17  GW-USValue-EZ 802.11n Wireless Adapter [Realtek RTL8188CUS]
	ed18  GW-USHyper300 / GW-USH300N 802.11bgn Wireless Adapter [Realtek RTL8191SU]
203d  Encore Electronics Inc.
	1480  ENUWI-N3 [802.11n Wireless N150 Adapter]
2040  Hauppauge
	0c80  Windham
	0c90  Windham
	1700  CataMount
	1800  Okemo A
	1801  Okemo B
	2000  Tiger Minicard
	2009  Tiger Minicard R2
	200a  Tiger Minicard
	2010  Tiger Minicard
	2011  WinTV MiniCard [Dell Digital TV Receiver]
	2019  Tiger Minicard
	2400  WinTV PVR USB2 (Model 24019)
	4700  WinTV Nova-S-USB2
	4902  HD PVR
	4903  HS PVR
	4982  HD PVR
	5500  Windham
	5510  Windham
	5520  Windham
	5530  Windham
	5580  Windham
	5590  Windham
	6500  WinTV HVR-900
	6502  WinTV HVR-900
	6503  WinTV HVR-930
	6513  WinTV HVR-980
	7050  Nova-T Stick
	7060  Nova-T Stick 2
	7070  Nova-T Stick 3
	7240  WinTV HVR-850
	8400  WinTV Nova-T-500
	9300  WinTV NOVA-T USB2 (cold)
	9301  WinTV NOVA-T USB2 (warm)
	9941  WinTV Nova-T-500
	9950  WinTV Nova-T-500
	b910  Windham
	b980  Windham
	b990  Windham
	c000  Windham
	c010  Windham
2047  Texas Instruments
	0200  MSP430 USB HID Bootstrap Loader
	0855  Invensense Embedded MotionApp HID Sensor
	0964  Inventio Software MSP430
2058  Nano River Technology
	2058  ViperBoard I2C, SPI, GPIO interface
2077  Taicang T&W Electronics Co. Ltd
	9002  W1M100 HSPA/WCDMA Module
2080  Barnes & Noble
	0001  nook
	0002  NOOKcolor
	0003  NOOK Simple Touch
	0004  NOOK Tablet
2086  SIMPASS
2087  Cando
	0a01  Multi Touch Panel
	0a02  Multi Touch Panel
	0b03  Multi Touch Panel
20a0  Clay Logic
	4123  IKALOGIC SCANALOGIC 2
	414a  MDE SPI Interface
	415a  OpenPilot
	415b  CopterControl
	415c  PipXtreme
20b1  XMOS Ltd
	10ad  XUSB Loader
	f7d1  XTAG2 - JTAG Adapter
20b3  Hanvon
	0a18  10.1 Touch screen overlay
20b7  Qi Hardware
	0713  Milkymist JTAG/serial
	1540  ben-wpan, AT86RF230-based
	1db5  IDBG in DFU mode
	1db6  IDBG in normal mode
	c25b  C2 Dongle
	cb72  ben-wpan, cntr
20ce  Minicircuits
	0012  RF Sythesizer 250-4200MHz model SSG-4000LH
	0021  RF Switch Matrix
	0022  I/O Controller
20df  Simtec Electronics
	0001  Entropy Key [UDEKEY01]
20f1  NET New Electronic Technology GmbH
	0101  iCube3 Camera
20f4  TRENDnet
	648b  TEW-648UBM 802.11n 150Mbps Micro Wireless N Adapter [Realtek RTL8188CUS]
20f7  XIMEA
	3001  Camera with CMOS sensor [MQ]
	3021  Camera with CCD sensor [MD]
	30b3  Camera with CMOS sensor in Vision mode [MQ]
	a003  Subminiature 5Mpix B/W Camera, MU9PM-MH
2100  RT Systems
	9e52  Yaesu VX-7
	9e54  CT29B Radio Cable
	9e57  RTS01 Radio Cable
	9e5d  K4Y Radio Cable
	9e5f  FT232RL [RTS05 Serial Cable]
2101  ActionStar
	0201  SIIG 4-to-2 Printer Switch
2109  VIA Labs, Inc.
	0700  VL700 SATA 3Gb/s bridge
	0701  VL701 SATA 3Gb/s bridge
	0810  VL81x Hub
	0811  Hub
	0812  VL812 Hub
	2811  Hub
	2812  VL812 Hub
	3431  Hub
	8110  Hub
2113  Softkinetic
	0137  DepthSense 311 (3D)
	0145  DepthSense 325
	8000  DepthSense 311 (Color)
2149  Advanced Silicon S.A.
	211b  Touchscreen Controller
	2703  TS58xxA/TC56xxA [CoolTouch]
2162  Creative (?)
	2031  Network Blaster Wireless Adapter
	500c  DE5771 Modem Blaster
	8001  Broadxent BritePort DSL Bridge 8010U
2184  GW Instek
	0005  GDS-3000 Oscilloscope
	0006  GDS-3000 Oscilloscope
	0011  AFG Function Generator (CDC)
21a1  Emotiv Systems Pty. Ltd.
	0001  EPOC Consumer Headset Wireless Dongle
21d6  Agecodagis SARL
	0002  Seismic recorder [Tellus]
2222  MacAlly
	0004  iWebKey Keyboard
	2520  Mini Tablet
	4050  AirStick joystick
2227  SAMWOO Enterprise
	3105  SKYDATA SKD-U100
2232  Silicon Motion
	1005  WebCam SCB-0385N
	1028  WebCam SC-03FFL11939N
	1029  WebCam SC-13HDL11939N
	1037  WebCam SC-03FFM12339N
2233  RadioShack Corporation
	6323  USB Electronic Scale
2237  Kobo Inc.
	4161  eReader White
225d  Morpho
	0001  FINGER VP Multimodal Biometric Sensor
	0008  CBM-E3 Fingerprint Sensor
	0009  CBM Fingerprint Sensor [CBM-V3]
	000a  MSO1300-E3 Fingerprint Sensor
	000b  MSO1300 Fingerprint Sensor [MSO1300-V3]
	000c  MSO1350-E3 Fingerprint Sensor & SmartCard Reader
	000d  MSO1350 Fingerprint Sensor & SmartCard Reader [MSO1350-V3]
	000e  MorphoAccess SIGMA Biometric Access Control Terminal
228d  8D Technologies inc.
	0001  Terminal Bike Key Reader
22a6  Pie Digital, Inc.
	ffff  PieKey "beta" 4GB model 4E4F41482E4F5247 (SM3251Q BB)
22b8  Motorola PCS
	0001  Wally 2.2 chipset
	0002  Wally 2.4 chipset
	0005  V.60c/V.60i GSM Phone
	0830  2386C-HT820
	0833  2386C-HT820 [Flash Mode]
	0850  Bluetooth Device
	1001  Patriot 1.0 (GSM) chipset
	1002  Patriot 2.0 chipset
	1005  T280e GSM/GPRS Phone
	1101  Patriot 1.0 (TDMA) chipset
	1801  Rainbow chipset flash
	2035  Bluetooth Device
	2805  GSM Modem
	2821  T720 GSM Phone
	2822  V.120e GSM Phone
	2823  Flash Interface
	2a01  MSM6050 chipset
	2a02  CDMA modem
	2a03  MSM6050 chipset flash
	2a21  V710 GSM Phone (P2K)
	2a22  V710 GSM Phone (AT)
	2a23  MSM6100 chipset flash
	2a41  MSM6300 chipset
	2a42  Usb Modem
	2a43  MSM6300 chipset flash
	2a61  E815 GSM Phone (P2K)
	2a62  E815 GSM Phone (AT)
	2a63  MSM6500 chipset flash
	2a81  MSM6025 chipset
	2a83  MSM6025 chipset flash
	2ac1  MSM6100 chipset
	2ac3  MSM6100 chipset flash
	2d78  XT300[SPICE]
	3001  A835/E1000 GSM Phone (P2K)
	3002  A835/E1000 GSM Phone (AT)
	3801  C350L/C450 (P2K)
	3802  C330/C350L/C450/EZX GSM Phone (AT)
	3803  Neptune LT chipset flash
	4001  OMAP 1.0 chipset
	4002  A920/A925 UMTS Phone
	4003  OMAP 1.0 chipset flash
	4008  OMAP 1.0 chipset RDL
	41d6  Droid X (Windows media mode)
	41d9  Droid/Milestone
	41db  Droid/Milestone (Debug mode)
	41de  Droid X (PC mode)
	4204  MPx200 Smartphone
	4214  MPc GSM
	4224  MPx220 Smartphone
	4234  MPc CDMA
	4244  MPx100 Smartphone
	4285  Droid X (Mass storage)
	4801  Neptune LTS chipset
	4803  Neptune LTS chipset flash
	4810  Triplet GSM Phone (storage)
	4901  Triplet GSM Phone (P2K)
	4902  Triplet GSM Phone (AT)
	4903  Neptune LTE chipset flash
	4a01  Neptune LTX chipset
	4a03  Neptune LTX chipset flash
	4a32  L6-imode Phone
	5801  Neptune ULS chipset
	5803  Neptune ULS chipset flash
	5901  Neptune VLT chipset
	5903  Neptune VLT chipset flash
	6001  Dalhart EZX
	6003  Dalhart flash
	6004  EZX GSM Phone (CDC Net)
	6006  MOTOROKR E6
	6008  Dalhart RDL
	6009  EZX GSM Phone (P2K)
	600a  Dalhart EZX config 17
	600b  Dalhart EZX config 18
	600c  EZX GSM Phone (USBLAN)
	6021  JUIX chipset
	6023  JUIX chipset flash
	6026  Flash RAM Downloader/miniOS
	6027  USBLAN
	604c  EZX GSM Phone (Storage)
	6101  Talon integrated chipset
	6401  Argon chipset
	6403  Argon chipset flash
	6415  ROKR Z6 (MTP mode)
	6604  Washington CDMA Phone
	6631  CDC Modem
	7001  Q Smartphone
	fe01  StarTAC III MS900
22b9  eTurboTouch Technology, Inc.
	0006  Touch Screen
22ba  Technology Innovation Holdings, Ltd
2304  Pinnacle Systems, Inc.
	0109  Studio PCTV USB (SECAM)
	0110  Studio PCTV USB (PAL)
	0111  Miro PCTV USB
	0112  Studio PCTV USB (NTSC) with FM radio
	0201  Systems MovieBox Device
	0204  MovieBox USB_B
	0205  DVC 150B
	0206  Systems MovieBox Deluxe Device
	0207  Dazzle DVC90 Video Device
	0208  Studio PCTV USB2
	020e  PCTV 200e
	020f  PCTV 400e BDA Device
	0210  Studio PCTV USB (PAL) with FM radio
	0212  Studio PCTV USB (NTSC)
	0213  500-USB Device
	0214  Studio PCTV USB (PAL) with FM radio
	0216  PCTV 60e
	0219  PCTV 260e
	021a  Dazzle DVC100 Audio Device
	021b  Dazzle DVC130/DVC170
	021d  Dazzle DVC130
	021e  Dazzle DVC170
	021f  PCTV Sat HDTV Pro BDA Device
	0222  PCTV Sat Pro BDA Device
	0223  DazzleTV Sat BDA Device
	0225  Remote Kit Infrared Transceiver
	0226  PCTV 330e
	0227  PCTV for Mac, HD Stick
	0228  PCTV DVB-T Flash Stick
	0229  PCTV Dual DVB-T 2001e
	022a  PCTV 160e
	022b  PCTV 71e [Afatech AF9015]
	0232  PCTV 170e
	0236  PCTV 72e [DiBcom DiB7000PC]
	0237  PCTV 73e [DiBcom DiB7000PC]
	023a  PCTV 801e
	023b  PCTV 801e SE
	023d  PCTV 340e
	023e  PCTV 340e SE
	0300  Studio Linx Video input cable (NTSC)
	0301  Studio Linx Video input cable (PAL)
	0302  Dazzle DVC120
	0419  PCTV Bungee USB (PAL) with FM radio
	061d  PCTV Deluxe (NTSC) Device
	061e  PCTV Deluxe (PAL) Device
2318  Shining Technologies, Inc. [hex]
	0011  CitiDISK Jr. IDE Enclosure
2341  Arduino SA
	0001  Uno (CDC ACM)
	0010  Mega 2560 (CDC ACM)
	003b  Serial Adapter (CDC ACM)
	003f  Mega ADK (CDC ACM)
	0042  Mega 2560 R3 (CDC ACM)
	0043  Uno R3 (CDC ACM)
	0044  Mega ADK R3 (CDC ACM)
	0045  Serial R3 (CDC ACM)
	8036  Leonardo (CDC ACM, HID)
2373  Pumatronix Ltda
	0001  5 MegaPixel Digital Still Camera [DSC5M]
2375  Digit@lway, Inc.
	0001  Digital Audio Player
2406  SANHO Digital Electronics Co., Ltd.
	6688  PD7X Portable Storage
2443  Aessent Technology Ltd
	00dc  aes220 FPGA Mini-Module
2478  Tripp-Lite
	2008  U209-000-R Serial Port
248a  Maxxter
	8366  Wireless Optical Mouse ACT-MUSW-002
249c  M2Tech s.r.l.
24e1  Paratronic
	3001  Adp-usb
	3005  Radius
2632  TwinMOS
	3209  7-in-1 Card Reader
2639  Xsens
	0001  MTi-10 IMU
	0002  MTi-20 VRU
	0003  MTi-30 AHRS
	0011  MTi-100 IMU
	0012  MTi-200 VRU
	0013  MTi-300 AHRS
	0017  MTi-G 7xx GNSS/INS
	0100  Body Pack
	0101  Awinda Station
	0102  Awinda Dongle
	0103  Sync Station
	0200  MTw
	d00d  Wireless Receiver
2650  Electronics For Imaging, Inc. [hex]
2659  Sundtek
	1101  TNT DVB-T/DAB/DAB+/FM
	1201  FM Transmitter/Receiver
	1202  MediaTV Analog/FM/DVB-T
	1203  MediaTV Analog/FM/DVB-T MiniPCIe
	1204  MediaTV Analog/FM/ATSC
	1205  SkyTV Ultimate V
	1206  MediaTV DVB-T MiniPCIe
	1207  Sundtek HD Capture
	1208  Sundtek SkyTV Ultimate III
	1209  MediaTV Analog/FM/ATSC MiniPCIe
	1210  MediaTV Pro III (EU)
	1211  MediaTV Pro III (US)
	1212  MediaTV Pro III MiniPCIe (EU)
	1213  MediaTV Pro III MiniPCIe (US)
2676  Basler AG
	ba02  ace
2730  Citizen
	200f  CT-S310 Label printer
2735  DigitalWay
	0003  MPIO HS100
	1001  MPIO FY200
	1002  MPIO FL100
	1003  MPIO FD100
	1004  MPIO HD200
	1005  MPIO HD300
	1006  MPIO FG100
	1007  MPIO FG130
	1008  MPIO FY300
	1009  MPIO FY400
	100a  MPIO FL300
	100b  MPIO HS200
	100c  MPIO FL350
	100d  MPIO FY500
	100e  MPIO FY500
	100f  MPIO FY600
	1012  MPIO FL400
	1013  MPIO HD400
	1014  MPIO HD400
	1016  MPIO FY700
	1017  MPIO FY700
	1018  MPIO FY800
	1019  MPIO FY800
	101a  MPIO FY900
	101b  MPIO FY900
	102b  MPIO FL500
	102c  MPIO FL500
	103f  MPIO FY570
	1040  MPIO FY570
	1041  MPIO FY670
	1042  MPIO FY670
	1043  HCT HMD-180A
	1044  HCT HMD-180A
2770  NHJ, Ltd
	0a01  ScanJet 4600 series
	905c  Che-Ez Snap SNAP-U/Digigr8/Soundstar TDC-35
	9060  A130
	9120  Che-ez! Snap / iClick Tiny VGA Digital Camera
	9130  TCG 501
	913c  Argus DC-1730
	9150  Mini Cam
	9153  iClick 5X
	915d  Cyberpix S-210S / Little Tikes My Real Digital Camera
	930b  CCD Webcam(PC370R)
	930c  CCD Webcam(PC370R)
27b8  ThingM
	01ed  blink(1)
2821  ASUSTek Computer Inc.
	0161  WL-161 802.11b Wireless Adapter [SiS 162U]
	160f  WL-160g 802.11g Wireless Adapter [Envara WiND512]
	3300  WL-140 / Hawking HWU36D 802.11b Wireless Adapter [Intersil PRISM 3]
2899  Toptronic Industrial Co., Ltd
	012c  Camera Device
289b  Dracal/Raphnet technologies
	0001  Gamecube/N64 controller v2.2
	0002  2nes2snes
	0003  4nes4snes
	0004  Gamecube/N64 controller v2.3
	0005  Saturn (Joystick mode)
	0006  Saturn (Mouse mode)
	0007  Famicom controller
	0008  Dreamcast (Joystick mode)
	0009  Dreamcast (Mouse mode)
	000a  Dreamcast (Keyboard mode)
	000b  Gamecube/N64 controller v2.9 (Keyboard mode)
	000c  Gamecube/N64 controller v2.9 (Joystick mode)
	0100  Dual-relay board
	0500  Energy meter
	0502  Precision barometer
2931  Jolla Oy
	0a01  Jolla Phone MTP
	0a02  Jolla Phone Developer
	0a05  Jolla PC connection
	0afe  Jolla charging only
2a03  dog hunter AG
	0001  Linino ONE (bootloader)
	0036  Arduino Leonardo (bootloader)
	0037  Arduino Micro (bootloader)
	0038  Arduino Robot Control (bootloader)
	0039  Arduino Robot Motor (bootloader)
	003a  Arduino Micro ADK rev3 (bootloader)
	003b  Arduino Serial
	003c  Arduino Explora (bootloader)
	003d  Arduino Due (usb2serial)
	003e  Arduino Due
	0041  Arduino Yun (bootloader)
	0042  Arduino Mega 2560 Rev3
	0043  Arduino Uno Rev3
	004d  Arduino Zero Pro (bootloader)
	8001  Linino ONE (CDC ACM)
	8036  Arduino Leonardo (CDC ACM)
	8037  Arduino Micro (CDC ACM)
	8038  Arduino Robot Control (CDC ACM)
	8039  Arduino Robot Motor (CDC ACM)
	803a  Arduino Micro ADK rev3 (CDC ACM)
	803c  Arduino Explora (CDC ACM)
	8041  Arduino Yun (CDC ACM)
	804d  Arduino Zero Pro (CDC ACM)
2a37  RTD Embedded Technologies, Inc.
	5110  UPS35110/UPS25110
2a45  Meizu Corp.
	0001  MX Phone (BICR)
	0c02  MX Phone (MTP & ADB)
	0c03  MX Phone (BICR & ADB)
	2008  MX Phone (MTP)
	200a  MX Phone (MTP & ACM & ADB)
	200b  MX Phone (PTP)
	200c  MX Phone (PTP & ADB)
	2012  MX Phone (MTP & ACM)
2c02  Planex Communications
	14ea  GW-US11H WLAN
2c1a  Dolphin Peripherals
	0000  Wireless Optical Mouse
2fb2  Fujitsu, Ltd
3125  Eagletron
	0001  TrackerPod Camera Stand
3136  Navini Networks
3176  Whanam Electronics Co., Ltd
3195  Link Instruments
	f190  MSO-19
	f280  MSO-28
	f281  MSO-28
3275  VidzMedia Pte Ltd
	4fb1  MonsterTV P2H
3333  InLine
	3333  2 port KVM switch model 60652K
3334  AEI
	1701  Fast Ethernet
3340  Yakumo
	043a  Mio A701 DigiWalker PPCPhone
	0e3a  Pocket PC 300 GPS SL / Typhoon MyGuide 3500
	a0a3  deltaX 5 BT (D) PDA
3344  Leaguer Microelectronics (LME)
	3744  OEM PC Remote
3504  Micro Star
	f110  Security Key
3538  Power Quotient International Co., Ltd
	0001  Travel Flash
	0015  Mass Storge Device
	0022  Hi-Speed Mass Storage Device
	0042  Cool Drive U339 Flash Disk
	0054  Flash Drive (2GB)
3579  DIVA
	6901  Media Reader
357d  Sharkoon
	7788  QuickPort XT
3636  InVibro
3838  WEM
	0001  5-in-1 Card Reader
3923  National Instruments Corp.
	12c0  DAQPad-6020E
	12d0  DAQPad-6507
	12e0  NI 4350
	12f0  NI 5102
	1750  DAQPad-6508
	17b0  USB-ISA-Bridge
	1820  DAQPad-6020E (68 pin I/O)
	1830  DAQPad-6020E (BNC)
	1f00  DAQPad-6024E
	1f10  DAQPad-6024E
	1f20  DAQPad-6025E
	1f30  DAQPad-6025E
	1f40  DAQPad-6036E
	1f50  DAQPad-6036E
	2f80  DAQPad-6052E
	2f90  DAQPad-6052E
	702b  GPIB-USB-B
	703c  USB-485 RS485 Cable
	709b  GPIB-USB-HS
	7254  NI MIO (data acquisition card) firmware updater
	729e  USB-6251 (OEM) data acquisition card
40bb  I-O Data
	0a09  USB2.0-SCSI Bridge USB2-SC
4101  i-rocks
	1301  IR-2510 usb phone
4102  iRiver, Ltd.
	1001  iFP-100 series mp3 player
	1003  iFP-300 series mp3 player
	1005  iFP-500 series mp3 player
	1007  iFP-700 series mp3/ogg vorbis player
	1008  iFP-800 series mp3/ogg vorbis player
	100a  iFP-1000 series mp3/ogg vorbis player
	1014  T20 series mp3/ogg vorbis player (ums firmware)
	1019  T30
	1034  T60
	1040  M1Player
	1041  E100 (ums)
	1101  iFP-100 series mp3 player (ums firmware)
	1103  iFP-300 series mp3 player (ums firmware)
	1105  iFP-500 series mp3 player (ums firmware)
	1113  T10 (alternate)
	1117  T10
	1119  T30 series mp3/ogg/wma player
	1141  E100 (mtp)
	2002  H10 6GB
	2101  H10 20GB (mtp)
	2102  H10 5GB (mtp)
	2105  H10 5/6GB (mtp)
413c  Dell Computer Corp.
	0000  DRAC 5 Virtual Keyboard and Mouse
	0001  DRAC 5 Virtual Media
	0058  Port Replicator
	1001  Keyboard Hub
	1002  Keyboard Hub
	1003  Keyboard Hub
	1005  Multimedia Pro Keyboard Hub
	2001  Keyboard HID Support
	2002  SK-8125 Keyboard
	2003  Keyboard
	2005  RT7D50 Keyboard
	2010  Keyboard
	2011  Multimedia Pro Keyboard
	2100  SK-3106 Keyboard
	2101  SmartCard Reader Keyboard
	2105  Model L100 Keyboard
	2106  Dell QuietKey Keyboard
	2500  DRAC4 Remote Access Card
	2513  internal USB Hub of E-Port Replicator
	3010  Optical Wheel Mouse
	3012  Optical Wheel Mouse
	3016  Optical 5-Button Wheel Mouse
	3200  Mouse
	4001  Axim X5
	4002  Axim X3
	4003  Axim X30
	4004  Axim Sync
	4005  Axim Sync
	4006  Axim Sync
	4007  Axim Sync
	4008  Axim Sync
	4009  Axim Sync
	4011  Axim X51v
	5103  AIO Printer A940
	5105  AIO Printer A920
	5107  AIO Printer A960
	5109  Photo AIO Printer 922
	5110  Photo AIO Printer 962
	5111  Photo AIO Printer 942
	5112  Photo AIO Printer 924
	5113  Photo AIO Printer 944
	5114  Photo AIO Printer 964
	5115  Photo AIO Printer 926
	5116  AIO Printer 946
	5117  Photo AIO Printer 966
	5118  AIO 810
	5124  Laser MFP 1815
	5128  Photo AIO 928
	5200  Laser Printer
	5202  Printing Support
	5203  Printing Support
	5210  Printing Support
	5211  1110 Laser Printer
	5220  Laser MFP 1600n
	5225  Printing Support
	5226  Printing Support
	5300  Laser Printer
	5400  Laser Printer
	5401  Laser Printer
	5513  WLA3310 Wireless Adapter [Intersil ISL3887]
	5601  Laser Printer 3100cn
	5602  Laser Printer 3000cn
	5631  Laser Printer 5100cn
	5905  Printing Support
	8000  BC02 Bluetooth Adapter
	8010  TrueMobile Bluetooth Module in
	8100  TrueMobile 1180 802.11b Adapter [Intersil PRISM 3]
	8102  TrueMobile 1300 802.11g Wireless Adapter [Intersil ISL3880]
	8103  Wireless 350 Bluetooth
	8104  Wireless 1450 Dual-band (802.11a/b/g) Adapter [Intersil ISL3887]
	8105  U2 in HID - Driver
	8106  Wireless 350 Bluetooth Internal Card in
	8110  Wireless 3xx Bluetooth Internal Card
	8111  Wireless 3xx Bluetooth Internal Card in
	8114  Wireless 5700 Mobile Broadband (CDMA EV-DO) Minicard Modem
	8115  Wireless 5500 Mobile Broadband (3G HSDPA) Minicard Modem
	8116  Wireless 5505 Mobile Broadband (3G HSDPA) Minicard Modem
	8117  Wireless 5700 Mobile Broadband (CDMA EV-DO) Expresscard Modem
	8118  Wireless 5510 Mobile Broadband (3G HSDPA) Expresscard Status Port
	8120  Bluetooth adapter
	8121  Eastfold in HID
	8122  Eastfold in DFU
	8123  eHome Infrared Receiver
	8124  eHome Infrared Receiver
	8126  Wireless 355 Bluetooth
	8127  Wireless 355 Module with Bluetooth 2.0 + EDR Technology.
	8128  Wireless 5700-Sprint Mobile Broadband (CDMA EV-DO) Mini-Card Status Port
	8129  Wireless 5700-Telus Mobile Broadband (CDMA EV-DO) Mini-Card Status Port
	8131  Wireless 360 Bluetooth 2.0 + EDR module.
	8133  Wireless 5720 VZW Mobile Broadband (EVDO Rev-A) Minicard GPS Port
	8134  Wireless 5720 Sprint Mobile Broadband (EVDO Rev-A) Minicard Status Port
	8135  Wireless 5720 TELUS Mobile Broadband (EVDO Rev-A) Minicard Diagnostics Port
	8136  Wireless 5520 Cingular Mobile Broadband (3G HSDPA) Minicard Diagnostics Port
	8137  Wireless 5520 Voda L Mobile Broadband (3G HSDPA) Minicard Status Port
	8138  Wireless 5520 Voda I Mobile Broadband (3G HSDPA) Minicard EAP-SIM Port
	8140  Wireless 360 Bluetooth
	8142  Mobile 360 in DFU
	8147  F3507g Mobile Broadband Module
	8156  Wireless 370 Bluetooth Mini-card
	8157  Integrated Keyboard
	8158  Integrated Touchpad / Trackstick
	8160  Wireless 365 Bluetooth
	8161  Integrated Keyboard
	8162  Integrated Touchpad [Synaptics]
	8171  Gobi Wireless Modem (QDL mode)
	8172  Gobi Wireless Modem
	8183  F3607gw Mobile Broadband Module
	8184  F3607gw v2 Mobile Broadband Module
	8185  Gobi 2000 Wireless Modem (QDL mode)
	8186  Gobi 2000 Wireless Modem
	8187  DW375 Bluetooth Module
	8501  Bluetooth Adapter
	9500  USB CP210x UART Bridge Controller [DW700]
	a001  Hub
	a005  Internal 2.0 Hub
	a700  Hub (in 1905FP LCD Monitor)
4146  USBest Technology
	9281  Iomega Micro Mini 128MB Flash Drive
	ba01  Intuix Flash Drive
4168  Targus
	1010  Wireless Compact Laser Mouse
4242  USB Design by Example
	4201  Buttons and Lights HID device
	4220  Echo 1 Camera
4255  GoPro
	1000  9FF2 [Digital Photo Display]
	2000  HD2-14 [Hero 2 Camera]
4317  Broadcom Corp.
	0700  U.S. Robotics USR5426 802.11g Adapter
	0701  U.S. Robotics USR5425 Wireless MAXg Adapter
	0711  Belkin F5D7051 v3000 802.11g
	0720  Dynex DX-BUSB
4348  WinChipHead
	5523  USB->RS 232 adapter with Prolifec PL 2303 chipset
	5537  13.56Mhz RFID Card Reader and Writer
	5584  CH34x printer adapter cable
4572  Shuttle, Inc.
	4572  Shuttle PN31 Remote
4586  Panram
	1026  Crystal Bar Flash Drive
4670  EMS Production
	9394  Game Cube USB Memory Adaptor 64M
4752  Miditech
	0011  Midistart-2
4757  GW Instek
	2009  PEL-2000 Series Electronic Load (CDC)
	2010  PEL-2000 Series Electronic Load (CDC)
4766  Aceeca
	0001  MEZ1000 RDA
4855  Memorex
	7288  Ultra Traveldrive 160G 2.5" HDD
4971  SimpleTech
	cb01  SP-U25/120G
	ce17  1TB SimpleDrive II USB External Hard Drive
4d46  Musical Fidelity
	0001  V-Link
	0002  V-DAC II
5032  Grandtec
	0bb8  Grandtec USB1.1 DVB-T (cold)
	0bb9  Grandtec USB1.1 DVB-T (warm)
	0fa0  Grandtec USB1.1 DVB-T (cold)
	0fa1  Grandtec USB1.1 DVB-T (warm)
5041  Linksys (?)
	2234  WUSB54G v1 802.11g Adapter [Intersil ISL3886]
	2235  WUSB54GP v1 802.11g Adapter [Intersil ISL3886]
50c2  Averatec (?)
	4013  WLAN Adapter
5173  Sweex
	1809  ZD1211
5219  I-Tetra
	1001  Cetus CDC Device
5345  Owon
	1234  PDS6062T Oscilloscope
534c  SatoshiLabs
	0001  Bitcoin Wallet [TREZOR]
5354  Meyer Instruments (MIS)
	0017  PAXcam2
544d  Transmeta Corp.
5543  UC-Logic Technology Corp.
	0002  SuperPen WP3325U Tablet
	0003  Tablet WP4030U
	0004  Tablet WP5540U
	0005  Tablet WP8060U
	0041  Genius PenSketch 6x8 Tablet
	0042  Tablet PF1209
	0064  Aiptek HyperPen 10000U
5555  Epiphan Systems Inc.
	1110  VGA2USB
	1120  KVM2USB
	2222  DVI2USB
	3333  VGA2USB Pro
	3337  KVM2USB Pro
	3340  VGA2USB LR
	3344  KVM2USB LR
	3411  DVI2USB Solo
	3422  DVI2USB Duo
55aa  OnSpec Electronic, Inc.
	0015  Hard Drive
	0102  SuperDisk
	0103  IDE Hard Drive
	0201  DDI to Reader-19
	1234  ATAPI Bridge
	a103  Sandisk SDDR-55 SmartMedia Card Reader
	b000  USB to CompactFlash Card Reader
	b004  OnSpec MMC/SD Reader/Writer
	b00b  USB to Memory Stick Card Reader
	b00c  USB to SmartMedia Card Reader
	b012  Mitsumi FA402M 8-in-2 Card Reader
	b200  Compact Flash Reader
	b204  MMC/ SD Reader
	b207  Memory Stick Reader
5654  Gotview
	ca42  MasterHD 3
5656  Uni-Trend Group Limited
	0832  UT2000/UT3000 Digital Storage Oscilloscope
595a  IRTOUCHSYSTEMS Co. Ltd.
	0001  Touchscreen
5986  Acer, Inc
	0100  Orbicam
	0101  USB2.0 Camera
	0102  Crystal Eye Webcam
	01a6  Lenovo Integrated Webcam
	01a7  Lenovo Integrated Webcam
	01a9  Lenovo Integrated Webcam
	0200  OrbiCam
	0203  BisonCam NB Pro 1300
	0241  BisonCam, NB Pro
	02d0  Lenovo Integrated Webcam [R5U877]
	03d0  Lenovo Integrated Webcam [R5U877]
59e3  Nonolith Labs
5a57  Zinwell
	0260  RT2570
	0280  802.11a/b/g/n USB Wireless LAN Card
	0282  802.11b/g/n USB Wireless LAN Card
	0283  802.11b/g/n USB Wireless LAN Card
	0284  802.11a/b/g/n USB Wireless LAN Card
	0290  ZW-N290 802.11n [Realtek RTL8192SU]
	5257  Metronic 495257 wifi 802.11ng
6000  Beholder International Ltd.
	dec0  TV Wander
	dec1  TV Voyage
601a  Ingenic Semiconductor Ltd.
	4740  XBurst Jz4740 boot mode
6189  Sitecom
	182d  USB 2.0 Ethernet
	2068  USB to serial cable (v2)
6244  LightingSoft AG
	0101  Intelligent Usb Dmx Interface SIUDI5A
	0201  Intelligent Usb Dmx Interface SIUDI5C
	0300  Intelligent Usb Dmx Interface SIUDI6 Firmware download
	0301  Intelligent Usb Dmx Interface SIUDI6C
	0302  Intelligent Usb Dmx Interface SIUDI6A
	0303  Intelligent Usb Dmx Interface SIUDI6D
	0400  Touch Sensitive Intelligent Control Keypad STICK1A
	0401  Touch Sensitive Intelligent Control Keypad STICK1A
	0410  Intelligent Usb Dmx Interface SIUDI7 Firmware Download
	0411  Intelligent Usb Dmx Interface SIUDI7A
	0420  Intelligent Usb Dmx Interface SIUDI8A Firmware Download
	0421  Intelligent Usb Dmx Interface SIUDI8A
	0430  Intelligent Usb Dmx Interface SIUDI8C Firmware Download
	0431  Intelligent Usb Dmx Interface SIUDI8C
	0440  Intelligent Usb Dmx Interface SIUDI9A Firmware Download
	0441  Intelligent Usb Dmx Interface SIUDI9A
	0450  Intelligent Usb Dmx Interface SIUDI9C Firmware Download
	0451  Intelligent Usb Dmx Interface SIUDI9C
	0460  Touch Sensitive Intelligent Control Keypad STICK2 Firmware download
	0461  Touch Sensitive Intelligent Control Keypad STICK2
	0470  Touch Sensitive Intelligent Control Keypad STICK1B Firmware download
	0471  Touch Sensitive Intelligent Control Keypad STICK1B
	0480  Touch Sensitive Intelligent Control Keypad STICK3 Firmware download
	0481  Touch Sensitive Intelligent Control Keypad STICK3
	0490  Intelligent Usb Dmx Interface SIUDI9D Firmware Download
	0491  Intelligent Usb Dmx Interface SIUDI9D
	0500  Touch Sensitive Intelligent Control Keypad STICK2B Firmware download
	0501  Touch Sensitive Intelligent Control Keypad STICK2B
6253  TwinHan Technology Co., Ltd
	0100  Ir reciver f. remote control
636c  CoreLogic, Inc.
6472  Unknown (Sony?)
	01c8  PlayStation Portable [Mass Storage]
6547  Arkmicro Technologies Inc.
	0232  ARK3116 Serial
6615  IRTOUCHSYSTEMS Co. Ltd.
	0001  Touchscreen
6666  Prototype product Vendor ID
	0667  WiseGroup Smart Joy PSX, PS-PC Smart JoyPad
	2667  JCOP BlueZ Smartcard reader
	8802  SmartJoy Dual Plus PS2 converter
	8804  WiseGroup SuperJoy Box 5
6677  WiseGroup, Ltd.
	8802  SmartJoy Dual Plus PS2 converter
	8811  Deluxe Dance Mat
6891  3Com
	a727  3CRUSB10075 802.11bg [ZyDAS ZD1211]
695c  Opera1
	3829  Opera1 DVB-S (warm state)
6993  Yealink Network Technology Co., Ltd.
	b001  VoIP Phone
6a75  Shanghai Jujo Electronics Co., Ltd
7104  CME (Central Music Co.)
	2202  UF5/UF6/UF7/UF8 MIDI Master Keyboard
726c  StackFoundry LLC
	2149  EntropyKing Random Number Generator
734c  TBS Technologies China
	5920  Q-Box II DVB-S2 HD
	5928  Q-Box II DVB-S2 HD
7373  Beijing STONE Technology Co. Ltd.
	5740  Intelligent TFT-LCD Module
7392  Edimax Technology Co., Ltd
	7711  EW-7711UTn nLite Wireless Adapter [Ralink RT2870]
	7717  EW-7717UN 802.11n Wireless Adapter [Ralink RT2870]
	7718  EW-7718UN 802.11n Wireless Adapter [Ralink RT2870]
	7722  EW-7722UTn 802.11n Wireless Adapter [Ralink RT307x]
	7811  EW-7811Un 802.11n Wireless Adapter [Realtek RTL8188CUS]
8086  Intel Corp.
	0001  AnyPoint (TM) Home Network 1.6 Mbps Wireless Adapter
	0044  CPU DRAM Controller
	0046  HD Graphics
	0100  Personal Audio Player 3000
	0101  Personal Audio Player 3000
	0110  Easy PC Camera
	0120  PC Camera CS120
	0180  WiMAX Connection 2400m
	0181  WiMAX Connection 2400m
	0182  WiMAX Connection 2400m
	0186  WiMAX Connection 2400m
	0188  WiMAX Connection 2400m
	0200  AnyPoint(TM) Wireless II Network 11Mbps Adapter [Atmel AT76C503A]
	0431  Intel Pro Video PC Camera
	0510  Digital Movie Creator
	0630  Pocket PC Camera
	0780  CS780 Microphone Input
	07d3  BLOB boot loader firmware
	0dad  Cherry MiniatureCard Keyboard
	1010  AnyPoint(TM) Home Network 10 Mbps Phoneline Adapter
	110a  Bluetooth Controller from (Ericsson P4A)
	110b  Bluetooth Controller from (Intel/CSR)
	1110  PRO/Wireless LAN Module
	1111  PRO/Wireless 2011B 802.11b Adapter [Intersil PRISM 2.5]
	1134  Hollister Mobile Monitor
	1139  In-Target Probe (ITP)
	1234  Prototype Reader/Writer
	1403  WiMAX Connection 2400m
	1405  WiMAX Connection 2400m
	1406  WiMAX Connection 2400m
	2448  82801 PCI Bridge
	3100  PRO/DSL 3220 Modem - WAN
	3101  PRO/DSL 3220 Modem
	3240  AnyPoint 3240 Modem - WAN
	3241  AnyPoint 3240 Modem
	8602  Miniature Card Slot
	9303  Intel 8x930Hx Hub
	9500  CE 9500 DVB-T
	9890  82930 Test Board
	beef  SCM Miniature Card Reader/Writer
	c013  Wireless HID Station
	f001  XScale PXA27x Bulverde flash
	f1a5  Z-U130 [Value Solid State Drive]
8087  Intel Corp.
	0020  Integrated Rate Matching Hub
	0024  Integrated Rate Matching Hub
80ee  VirtualBox
	0021  USB Tablet
8282  Keio
	3201  Retro Adapter
	3301  Retro Adapter Mouse
8341  EGO Systems, Inc.
	2000  Flashdisk
8564  Transcend Information, Inc.
	1000  JetFlash
	4000  RDF8
8644  Intenso GmbG
	8003  Micro Line
	800b  Micro Line (4GB)
8e06  CH Products, Inc.
	f700  DT225 Trackball
9016  Sitecom
	182d  WL-022 802.11b Adapter
9022  TeVii Technology Ltd.
	d630  DVB-S S630
	d650  DVB-S2 S650
	d660  DVB-S2 S660
9148  GeoLab, Ltd
# All of GeoLab's devices share the same ID 0004.
	0004  R3 Compatible Device
9710  MosChip Semiconductor
	7703  MCS7703 Serial Port Adapter
	7705  MCS7705 Parallel port adapter
	7715  MCS7715 Parallel and serial port adapter
	7717  MCS7717 3-port hub with serial and parallel adapter
	7720  MCS7720 Dual serial port adapter
	7730  MCS7730 10/100 Mbps Ethernet adapter
	7780  MCS7780 4Mbps Fast IrDA Adapter
	7830  MCS7830 10/100 Mbps Ethernet adapter
	7832  MCS7832 10/100 Mbps Ethernet adapter
	7840  MCS7820/MCS7840 2/4 port serial adapter
9849  Bestmedia CD Recordable GmbH & Co. KG
	0701  Platinum MyDrive HP
9999  Odeon
	0001  JAF Mobile Phone Flasher Interface
99fa  Grandtec
	8988  V.cap Camera Device
9ac4  J. Westhues
	4b8f  ProxMark-3 RFID Instrument
9e88  Marvell Semiconductor, Inc.
	9e8f  Plug Computer Basic [SheevaPlug]
a128  AnMo Electronics Corp. / Dino-Lite (?)
	0610  Dino-Lite Digital Microscope (SN9C201 + HV7131R)
	0611  Dino-Lite Digital Microscope (SN9C201 + HV7131R)
	0612  Dino-Lite Digital Microscope (SN9C120 + HV7131R)
	0613  Dino-Lite Digital Microscope (SN9C201 + HV7131R)
	0614  Dino-Lite Digital Microscope (SN9C201 + MI1310/MT9M111)
	0615  Dino-Lite Digital Microscope (SN9C201 + MI1310/MT9M111)
	0616  Dino-Lite Digital Microscope (SN9C120 + HV7131R)
	0617  Dino-Lite Digital Microscope (SN9C201 + MI1310/MT9M111)
	0618  Dino-Lite Digital Microscope (SN9C201 + HV7131R)
a168  AnMo Electronics Corporation
	0610  Dino-Lite Digital Microscope
	0611  Dino-Lite Digital Microscope
	0613  Dino-Lite Digital Microscope
	0614  Dino-Lite Pro Digital Microscope
	0615  Dino-Lite Pro Digital Microscope
	0617  Dino-Lite Pro Digital Microscope
	0618  Dino-Lite Digital Microscope
a600  Asix
	e110  OK1ZIA Davac 4.x
a727  3Com
	6893  3CRUSB20075 OfficeConnect Wireless 108Mbps 11g Adapter [Atheros AR5523]
	6895  AR5523
	6897  AR5523
aaaa  MXT
	8815  microSD CardReader
abcd  Unknown
	cdee  Petcam
b58e  Blue Microphones
	9e84  Yeti Stereo Microphone
c216  Card Device Expert Co., LTD
	0180  MSR90 MagStripe reader
c251  Keil Software, Inc.
	2710  ULink
cace  CACE Technologies Inc.
	0002  AirPCAP Classic 802.11 packet capture adapter
	0300  AirPcap NX [Atheros AR9001U-(2)NG]
cd12  SMART TECHNOLOGY INDUSTRIAL LTD.
d208  Ultimarc
	0310  Mini-PAC Arcade Control Interface
d209  Ultimarc
	0301  I-PAC Arcade Control Interface
	0501  Ultra-Stik Ultimarc Ultra-Stik Player 1
d904  LogiLink
	0003  Laser Mouse (ID0009A)
e4e4  Xorcom Ltd.
	1130  Astribank series
	1131  Astribank series
	1132  Astribank series
	1140  Astribank series
	1141  Astribank series
	1142  Astribank series
	1150  Astribank series
	1151  Astribank series
	1152  Astribank series
	1160  Astribank 2 series
	1161  Astribank 2 series
	1162  Astribank 2 series
eb03  MakingThings
	0920  Make Controller Kit
eb1a  eMPIA Technology, Inc.
	17de  KWorld V-Stream XPERT DTV - DVB-T USB cold
	17df  KWorld V-Stream XPERT DTV - DVB-T USB warm
	2571  M035 Compact Web Cam
	2710  SilverCrest Webcam
	2750  ECS Elitegroup G220 integrated Webcam
	2761  EeePC 701 integrated Webcam
	2776  Combined audio and video input device
	2800  Terratec Cinergy 200
	2801  GrabBeeX+ Video Encoder
	2863  Video Grabber
	2870  Pinnacle PCTV Stick
	2881  EM2881 Video Controller
	50a3  Gadmei UTV380 TV Box
	50a6  Gadmei UTV330 TV Box
	e355  KWorld DVB-T 355U Digital TV Dongle
eb2a  KWorld
ef18  SMART TECHNOLOGY INDUSTRIAL LTD.
f003  Hewlett Packard
	6002  PhotoSmart C500
f182  Leap Motion
	0003  Controller
f4ec  Atten Electronics / Siglent Technologies
	ee38  Digital Storage Oscilloscope
f4ed  Shenzhen Siglent Co., Ltd.
	ee37  SDG1010 Waveform Generator
	ee3a  SDG1010 Waveform Generator (TMC mode)
f766  Hama
	0001  PC-Gamepad "Greystorm"
fc08  Conrad Electronic SE
	0101  MIDI Cable UA0037
ffee  FNK Tech
	0100  Card Reader Controller RTS5101/RTS5111/RTS5116

# List of known device classes, subclasses and protocols

# Syntax:
# C class  class_name
#	subclass  subclass_name			<-- single tab
#		protocol  protocol_name		<-- two tabs

C 00  (Defined at Interface level)
C 01  Audio
	01  Control Device
	02  Streaming
	03  MIDI Streaming
C 02  Communications
	01  Direct Line
	02  Abstract (modem)
		00  None
		01  AT-commands (v.25ter)
		02  AT-commands (PCCA101)
		03  AT-commands (PCCA101 + wakeup)
		04  AT-commands (GSM)
		05  AT-commands (3G)
		06  AT-commands (CDMA)
		fe  Defined by command set descriptor
		ff  Vendor Specific (MSFT RNDIS?)
	03  Telephone
	04  Multi-Channel
	05  CAPI Control
	06  Ethernet Networking
	07  ATM Networking
	08  Wireless Handset Control
	09  Device Management
	0a  Mobile Direct Line
	0b  OBEX
	0c  Ethernet Emulation
		07  Ethernet Emulation (EEM)
C 03  Human Interface Device
	00  No Subclass
		00  None
		01  Keyboard
		02  Mouse
	01  Boot Interface Subclass
		00  None
		01  Keyboard
		02  Mouse
C 05  Physical Interface Device
C 06  Imaging
	01  Still Image Capture
		01  Picture Transfer Protocol (PIMA 15470)
C 07  Printer
	01  Printer
		00  Reserved/Undefined
		01  Unidirectional
		02  Bidirectional
		03  IEEE 1284.4 compatible bidirectional
		ff  Vendor Specific
C 08  Mass Storage
	01  RBC (typically Flash)
		00  Control/Bulk/Interrupt
		01  Control/Bulk
		50  Bulk-Only
	02  SFF-8020i, MMC-2 (ATAPI)
	03  QIC-157
	04  Floppy (UFI)
		00  Control/Bulk/Interrupt
		01  Control/Bulk
		50  Bulk-Only
	05  SFF-8070i
	06  SCSI
		00  Control/Bulk/Interrupt
		01  Control/Bulk
		50  Bulk-Only
C 09  Hub
	00  Unused
		00  Full speed (or root) hub
		01  Single TT
		02  TT per port
C 0a  CDC Data
	00  Unused
		30  I.430 ISDN BRI
		31  HDLC
		32  Transparent
		50  Q.921M
		51  Q.921
		52  Q.921TM
		90  V.42bis
		91  Q.932 EuroISDN
		92  V.120 V.24 rate ISDN
		93  CAPI 2.0
		fd  Host Based Driver
		fe  CDC PUF
		ff  Vendor specific
C 0b  Chip/SmartCard
C 0d  Content Security
C 0e  Video
	00  Undefined
	01  Video Control
	02  Video Streaming
	03  Video Interface Collection
C 58  Xbox
	42  Controller
C dc  Diagnostic
	01  Reprogrammable Diagnostics
		01  USB2 Compliance
C e0  Wireless
	01  Radio Frequency
		01  Bluetooth
		02  Ultra WideBand Radio Control
		03  RNDIS
	02  Wireless USB Wire Adapter
		01  Host Wire Adapter Control/Data Streaming
		02  Device Wire Adapter Control/Data Streaming
		03  Device Wire Adapter Isochronous Streaming
C ef  Miscellaneous Device
	01  ?
		01  Microsoft ActiveSync
		02  Palm Sync
	02  ?
		01  Interface Association
		02  Wire Adapter Multifunction Peripheral
	03  ?
		01  Cable Based Association
	05  USB3 Vision
C fe  Application Specific Interface
	01  Device Firmware Update
	02  IRDA Bridge
	03  Test and Measurement
		01  TMC
		02  USB488
C ff  Vendor Specific Class
	ff  Vendor Specific Subclass
		ff  Vendor Specific Protocol

# List of Audio Class Terminal Types

# Syntax:
# AT terminal_type  terminal_type_name

AT 0100  USB Undefined
AT 0101  USB Streaming
AT 01ff  USB Vendor Specific
AT 0200  Input Undefined
AT 0201  Microphone
AT 0202  Desktop Microphone
AT 0203  Personal Microphone
AT 0204  Omni-directional Microphone
AT 0205  Microphone Array
AT 0206  Processing Microphone Array
AT 0300  Output Undefined
AT 0301  Speaker
AT 0302  Headphones
AT 0303  Head Mounted Display Audio
AT 0304  Desktop Speaker
AT 0305  Room Speaker
AT 0306  Communication Speaker
AT 0307  Low Frequency Effects Speaker
AT 0400  Bidirectional Undefined
AT 0401  Handset
AT 0402  Headset
AT 0403  Speakerphone, no echo reduction
AT 0404  Echo-suppressing speakerphone
AT 0405  Echo-canceling speakerphone
AT 0500  Telephony Undefined
AT 0501  Phone line
AT 0502  Telephone
AT 0503  Down Line Phone
AT 0600  External Undefined
AT 0601  Analog Connector
AT 0602  Digital Audio Interface
AT 0603  Line Connector
AT 0604  Legacy Audio Connector
AT 0605  SPDIF interface
AT 0606  1394 DA stream
AT 0607  1394 DV stream soundtrack
AT 0700  Embedded Undefined
AT 0701  Level Calibration Noise Source
AT 0702  Equalization Noise
AT 0703  CD Player
AT 0704  DAT
AT 0705  DCC
AT 0706  MiniDisc
AT 0707  Analog Tape
AT 0708  Phonograph
AT 0709  VCR Audio
AT 070a  Video Disc Audio
AT 070b  DVD Audio
AT 070c  TV Tuner Audio
AT 070d  Satellite Receiver Audio
AT 070e  Cable Tuner Audio
AT 070f  DSS Audio
AT 0710  Radio Receiver
AT 0711  Radio Transmitter
AT 0712  Multitrack Recorder
AT 0713  Synthesizer

# List of HID Descriptor Types

# Syntax:
# HID descriptor_type  descriptor_type_name

HID 21  HID
HID 22  Report
HID 23  Physical

# List of HID Descriptor Item Types
# Note: 2 bits LSB encode data length following

# Syntax:
# R item_type  item_type_name

R 04  Usage Page
R 08  Usage
R 14  Logical Minimum
R 18  Usage Minimum
R 24  Logical Maximum
R 28  Usage Maximum
R 34  Physical Minimum
R 38  Designator Index
R 44  Physical Maximum
R 48  Designator Minimum
R 54  Unit Exponent
R 58  Designator Maximum
R 64  Unit
R 74  Report Size
R 78  String Index
R 80  Input
R 84  Report ID
R 88  String Minimum
R 90  Output
R 94  Report Count
R 98  String Maximum
R a0  Collection
R a4  Push
R a8  Delimiter
R b0  Feature
R b4  Pop
R c0  End Collection

# List of Physical Descriptor Bias Types

# Syntax:
# BIAS item_type  item_type_name

BIAS 0  Not Applicable
BIAS 1  Right Hand
BIAS 2  Left Hand
BIAS 3  Both Hands
BIAS 4  Either Hand

# List of Physical Descriptor Item Types

# Syntax:
# PHY item_type  item_type_name

PHY 00  None
PHY 01  Hand
PHY 02  Eyeball
PHY 03  Eyebrow
PHY 04  Eyelid
PHY 05  Ear
PHY 06  Nose
PHY 07  Mouth
PHY 08  Upper Lip
PHY 09  Lower Lip
PHY 0a  Jaw
PHY 0b  Neck
PHY 0c  Upper Arm
PHY 0d  Elbow
PHY 0e  Forearm
PHY 0f  Wrist
PHY 10  Palm
PHY 11  Thumb
PHY 12  Index Finger
PHY 13  Middle Finger
PHY 14  Ring Finger
PHY 15  Little Finger
PHY 16  Head
PHY 17  Shoulder
PHY 18  Hip
PHY 19  Waist
PHY 1a  Thigh
PHY 1b  Knee
PHY 1c  calf
PHY 1d  Ankle
PHY 1e  Foot
PHY 1f  Heel
PHY 20  Ball of Foot
PHY 21  Big Toe
PHY 22  Second Toe
PHY 23  Third Toe
PHY 24  Fourth Toe
PHY 25  Fifth Toe
PHY 26  Brow
PHY 27  Cheek

# List of HID Usages

# Syntax:
# HUT hi  _usage_page  hid_usage_page_name
#	hid_usage  hid_usage_name

HUT 00  Undefined
HUT 01  Generic Desktop Controls
	000  Undefined
	001  Pointer
	002  Mouse
	004  Joystick
	005  Gamepad
	006  Keyboard
	007  Keypad
	008  Multi-Axis Controller
	030  Direction-X
	031  Direction-Y
	032  Direction-Z
	033  Rotate-X
	034  Rotate-Y
	035  Rotate-Z
	036  Slider
	037  Dial
	038  Wheel
	039  Hat Switch
	03a  Counted Buffer
	03b  Byte Count
	03c  Motion Wakeup
	03d  Start
	03e  Select
	040  Vector-X
	041  Vector-Y
	042  Vector-Z
	043  Vector-X relative Body
	044  Vector-Y relative Body
	045  Vector-Z relative Body
	046  Vector
	080  System Control
	081  System Power Down
	082  System Sleep
	083  System Wake Up
	084  System Context Menu
	085  System Main Menu
	086  System App Menu
	087  System Menu Help
	088  System Menu Exit
	089  System Menu Select
	08a  System Menu Right
	08b  System Menu Left
	08c  System Menu Up
	08d  System Menu Down
	090  Direction Pad Up
	091  Direction Pad Down
	092  Direction Pad Right
	093  Direction Pad Left
HUT 02  Simulation Controls
	000  Undefined
	001  Flight Simulation Device
	002  Automobile Simulation Device
	003  Tank Simulation Device
	004  Spaceship Simulation Device
	005  Submarine Simulation Device
	006  Sailing Simulation Device
	007  Motorcycle Simulation Device
	008  Sports Simulation Device
	009  Airplane Simualtion Device
	00a  Helicopter Simulation Device
	00b  Magic Carpet Simulation Device
	00c  Bicycle Simulation Device
	020  Flight Control Stick
	021  Flight Stick
	022  Cyclic Control
	023  Cyclic Trim
	024  Flight Yoke
	025  Track Control
	0b0  Aileron
	0b1  Aileron Trim
	0b2  Anti-Torque Control
	0b3  Autopilot Enable
	0b4  Chaff Release
	0b5  Collective Control
	0b6  Dive Break
	0b7  Electronic Countermeasures
	0b8  Elevator
	0b9  Elevator Trim
	0ba  Rudder
	0bb  Throttle
	0bc  Flight COmmunications
	0bd  Flare Release
	0be  Landing Gear
	0bf  Toe Break
	0c0  Trigger
	0c1  Weapon Arm
	0c2  Weapons Select
	0c3  Wing Flaps
	0c4  Accelerator
	0c5  Brake
	0c6  Clutch
	0c7  Shifter
	0c8  Steering
	0c9  Turret Direction
	0ca  Barrel Elevation
	0cb  Drive Plane
	0cc  Ballast
	0cd  Bicylce Crank
	0ce  Handle Bars
	0cf  Front Brake
	0d0  Rear Brake
HUT 03  VR Controls
	000  Unidentified
	001  Belt
	002  Body Suit
	003  Flexor
	004  Glove
	005  Head Tracker
	006  Head Mounted Display
	007  Hand Tracker
	008  Oculometer
	009  Vest
	00a  Animatronic Device
	020  Stereo Enable
	021  Display Enable
HUT 04  Sport Controls
	000  Unidentified
	001  Baseball Bat
	002  Golf Club
	003  Rowing Machine
	004  Treadmill
	030  Oar
	031  Slope
	032  Rate
	033  Stick Speed
	034  Stick Face Angle
	035  Stick Heel/Toe
	036  Stick Follow Through
	038  Stick Type
	039  Stick Height
	047  Stick Temp
	050  Putter
	051  1 Iron
	052  2 Iron
	053  3 Iron
	054  4 Iron
	055  5 Iron
	056  6 Iron
	057  7 Iron
	058  8 Iron
	059  9 Iron
	05a  10 Iron
	05b  11 Iron
	05c  Sand Wedge
	05d  Loft Wedge
	05e  Power Wedge
	05f  1 Wood
	060  3 Wood
	061  5 Wood
	062  7 Wood
	063  9 Wood
HUT 05  Game Controls
	000  Undefined
	001  3D Game Controller
	002  Pinball Device
	003  Gun Device
	020  Point Of View
	021  Turn Right/Left
	022  Pitch Right/Left
	023  Roll Forward/Backward
	024  Move Right/Left
	025  Move Forward/Backward
	026  Move Up/Down
	027  Lean Right/Left
	028  Lean Forward/Backward
	029  Height of POV
	02a  Flipper
	02b  Secondary Flipper
	02c  Bump
	02d  New Game
	02e  Shoot Ball
	02f  Player
	030  Gun Bolt
	031  Gun Clip
	032  Gun Selector
	033  Gun Single Shot
	034  Gun Burst
	035  Gun Automatic
	036  Gun Safety
	037  Gamepad Fire/Jump
	038  Gamepad Fun
	039  Gamepad Trigger
HUT 07  Keyboard
	000  No Event
	001  Keyboard ErrorRollOver
	002  Keyboard POSTfail
	003  Keyboard Error Undefined
	004  A
	005  B
	006  C
	007  D
	008  E
	009  F
	00a  G
	00b  H
	00c  I
	00d  J
	00e  K
	00f  L
	010  M
	011  N
	012  O
	013  P
	014  Q
	015  R
	016  S
	017  T
	018  U
	019  V
	01a  W
	01b  X
	01c  Y
	01d  Z
	01e  1 and ! (One and Exclamation)
	01f  2 and @ (2 and at)
	020  3 and # (3 and Hash)
	021  4 and $ (4 and Dollar Sign)
	022  5 and % (5 and Percent Sign)
	023  6 and ^ (6 and circumflex)
	024  7 and & (Seven and Ampersand)
	025  8 and * (Eight and asterisk)
	026  9 and ( (Nine and Parenthesis Left)
	027  0 and ) (Zero and Parenthesis Right)
	028  Return (Enter)
	029  Escape
	02a  Delete (Backspace)
	02b  Tab
	02c  Space Bar
	02d  - and _ (Minus and underscore)
	02e  = and + (Equal and Plus)
	02f  [ and { (Bracket and Braces Left)
	030  ] and } (Bracket and Braces Right)
	031  \ and | (Backslash and Bar)
	032  # and ~ (Hash and Tilde, Non-US Keyboard near right shift)
	033  ; and : (Semicolon and Colon)
	034   and " (Accent Acute and Double Quotes)
	035   and ~ (Accent Grace and Tilde)
	036  , and < (Comma and Less)
	037  . and > (Period and Greater)
	038  / and ? (Slash and Question Mark)
	039  Caps Lock
	03a  F1
	03b  F2
	03c  F3
	03d  F4
	03e  F5
	03f  F6
	040  F7
	041  F8
	042  F9
	043  F10
	044  F11
	045  F12
	046  Print Screen
	047  Scroll Lock
	048  Pause
	049  Insert
	04a  Home
	04b  Page Up
	04c  Delete Forward (without Changing Position)
	04d  End
	04e  Page Down
	04f  Right Arrow
	050  Left Arrow
	051  Down Arrow
	052  Up Arrow
	053  Num Lock and Clear
	054  Keypad / (Division Sign)
	055  Keypad * (Multiplication Sign)
	056  Keypad - (Subtraction Sign)
	057  Keypad + (Addition Sign)
	058  Keypad Enter
	059  Keypad 1 and END
	05a  Keypad 2 and Down Arrow
	05b  Keypad 3 and Page Down
	05c  Keypad 4 and Left Arrow
	05d  Keypad 5 (Tactilei Raised)
	05f  Keypad 6 and Right Arrow
	060  Keypad 7 and Home
	061  Keypad 8 and Up Arrow
	062  Keypad 8 and Page Up
	063  Keypad . (decimal delimiter) and Delete
	064  \ and | (Backslash and Bar, UK and Non-US Keyboard near left shift)
	065  Keyboard Application (Windows Key for Win95 or Compose)
	066  Power (not a key)
	067  Keypad = (Equal Sign)
	068  F13
	069  F14
	06a  F15
	06b  F16
	06c  F17
	06d  F18
	06e  F19
	06f  F20
	070  F21
	071  F22
	072  F23
	073  F24
	074  Execute
	075  Help
	076  Menu
	077  Select
	078  Stop
	079  Again
	07a  Undo
	07b  Cut
	07c  Copy
	07d  Paste
	07e  Find
	07f  Mute
	080  Volume Up
	081  Volume Down
	082  Locking Caps Lock
	083  Locking Num Lock
	084  Locking Scroll Lock
	085  Keypad Comma
	086  Keypad Equal Sign (AS/400)
	087  International 1 (PC98)
	088  International 2 (PC98)
	089  International 3 (PC98)
	08a  International 4 (PC98)
	08b  International 5 (PC98)
	08c  International 6 (PC98)
	08d  International 7 (Toggle Single/Double Byte Mode)
	08e  International 8
	08f  International 9
	090  LANG 1 (Hangul/English Toggle, Korea)
	091  LANG 2 (Hanja Conversion, Korea)
	092  LANG 3 (Katakana, Japan)
	093  LANG 4 (Hiragana, Japan)
	094  LANG 5 (Zenkaku/Hankaku, Japan)
	095  LANG 6
	096  LANG 7
	097  LANG 8
	098  LANG 9
	099  Alternate Erase
	09a  SysReq/Attention
	09b  Cancel
	09c  Clear
	09d  Prior
	09e  Return
	09f  Separator
	0a0  Out
	0a1  Open
	0a2  Clear/Again
	0a3  CrSel/Props
	0a4  ExSel
	0e0  Control Left
	0e1  Shift Left
	0e2  Alt Left
	0e3  GUI Left
	0e4  Control Right
	0e5  Shift Right
	0e6  Alt Rigth
	0e7  GUI Right
HUT 08  LEDs
	000  Undefined
	001  NumLock
	002  CapsLock
	003  Scroll Lock
	004  Compose
	005  Kana
	006  Power
	007  Shift
	008  Do not disturb
	009  Mute
	00a  Tone Enabke
	00b  High Cut Filter
	00c  Low Cut Filter
	00d  Equalizer Enable
	00e  Sound Field ON
	00f  Surround On
	010  Repeat
	011  Stereo
	012  Sampling Rate Detect
	013  Spinning
	014  CAV
	015  CLV
	016  Recording Format Detect
	017  Off-Hook
	018  Ring
	019  Message Waiting
	01a  Data Mode
	01b  Battery Operation
	01c  Battery OK
	01d  Battery Low
	01e  Speaker
	01f  Head Set
	020  Hold
	021  Microphone
	022  Coverage
	023  Night Mode
	024  Send Calls
	025  Call Pickup
	026  Conference
	027  Stand-by
	028  Camera On
	029  Camera Off
	02a  On-Line
	02b  Off-Line
	02c  Busy
	02d  Ready
	02e  Paper-Out
	02f  Paper-Jam
	030  Remote
	031  Forward
	032  Reverse
	033  Stop
	034  Rewind
	035  Fast Forward
	036  Play
	037  Pause
	038  Record
	039  Error
	03a  Usage Selected Indicator
	03b  Usage In Use Indicator
	03c  Usage Multi Indicator
	03d  Indicator On
	03e  Indicator Flash
	03f  Indicator Slow Blink
	040  Indicator Fast Blink
	041  Indicator Off
	042  Flash On Time
	043  Slow Blink On Time
	044  Slow Blink Off Time
	045  Fast Blink On Time
	046  Fast Blink Off Time
	047  Usage Color Indicator
	048  Indicator Red
	049  Indicator Green
	04a  Indicator Amber
	04b  Generic Indicator
	04c  System Suspend
	04d  External Power Connected
HUT 09  Buttons
	000  No Button Pressed
	001  Button 1 (Primary)
	002  Button 2 (Secondary)
	003  Button 3 (Tertiary)
	004  Button 4
	005  Button 5
HUT 0a  Ordinal
	001  Instance 1
	002  Instance 2
	003  Instance 3
HUT 0b  Telephony
	000  Unassigned
	001  Phone
	002  Answering Machine
	003  Message Controls
	004  Handset
	005  Headset
	006  Telephony Key Pad
	007  Programmable Button
	020  Hook Switch
	021  Flash
	022  Feature
	023  Hold
	024  Redial
	025  Transfer
	026  Drop
	027  Park
	028  Forward Calls
	029  Alternate Function
	02a  Line
	02b  Speaker Phone
	02c  Conference
	02d  Ring Enable
	02e  Ring Select
	02f  Phone Mute
	030  Caller ID
	050  Speed Dial
	051  Store Number
	052  Recall Number
	053  Phone Directory
	070  Voice Mail
	071  Screen Calls
	072  Do Not Disturb
	073  Message
	074  Answer On/Offf
	090  Inside Dial Tone
	091  Outside Dial Tone
	092  Inside Ring Tone
	093  Outside Ring Tone
	094  Priority Ring Tone
	095  Inside Ringback
	096  Priority Ringback
	097  Line Busy Tone
	098  Recorder Tone
	099  Call Waiting Tone
	09a  Confirmation Tone 1
	09b  Confirmation Tone 2
	09c  Tones Off
	09d  Outside Ringback
	0b0  Key 1
	0b1  Key 2
	0b3  Key 3
	0b4  Key 4
	0b5  Key 5
	0b6  Key 6
	0b7  Key 7
	0b8  Key 8
	0b9  Key 9
	0ba  Key Star
	0bb  Key Pound
	0bc  Key A
	0bd  Key B
	0be  Key C
	0bf  Key D
HUT 0c  Consumer
	000  Unassigned
	001  Consumer Control
	002  Numeric Key Pad
	003  Programmable Buttons
	020  +10
	021  +100
	022  AM/PM
	030  Power
	031  Reset
	032  Sleep
	033  Sleep After
	034  Sleep Mode
	035  Illumination
	036  Function Buttons
	040  Menu
	041  Menu Pick
	042  Menu Up
	043  Menu Down
	044  Menu Left
	045  Menu Right
	046  Menu Escape
	047  Menu Value Increase
	048  Menu Value Decrease
	060  Data on Screen
	061  Closed Caption
	062  Closed Caption Select
	063  VCR/TV
	064  Broadcast Mode
	065  Snapshot
	066  Still
	080  Selection
	081  Assign Selection
	082  Mode Step
	083  Recall Last
	084  Enter Channel
	085  Order Movie
	086  Channel
	087  Media Selection
	088  Media Select Computer
	089  Media Select TV
	08a  Media Select WWW
	08b  Media Select DVD
	08c  Media Select Telephone
	08d  Media Select Program Guide
	08e  Media Select Video Phone
	08f  Media Select Games
	090  Media Select Messages
	091  Media Select CD
	092  Media Select VCR
	093  Media Select Tuner
	094  Quit
	095  Help
	096  Media Select Tape
	097  Media Select Cable
	098  Media Select Satellite
	099  Media Select Security
	09a  Media Select Home
	09b  Media Select Call
	09c  Channel Increment
	09d  Channel Decrement
	09e  Media Select SAP
	0a0  VCR Plus
	0a1  Once
	0a2  Daily
	0a3  Weekly
	0a4  Monthly
	0b0  Play
	0b1  Pause
	0b2  Record
	0b3  Fast Forward
	0b4  Rewind
	0b5  Scan Next Track
	0b6  Scan Previous Track
	0b7  Stop
	0b8  Eject
	0b9  Random Play
	0ba  Select Disc
	0bb  Enter Disc
	0bc  Repeat
	0bd  Tracking
	0be  Track Normal
	0bf  Slow Tracking
	0c0  Frame Forward
	0c1  Frame Back
	0c2  Mark
	0c3  Clear Mark
	0c4  Repeat from Mark
	0c5  Return to Mark
	0c6  Search Mark Forward
	0c7  Search Mark Backward
	0c8  Counter Reset
	0c9  Show Counter
	0ca  Tracking Increment
	0cb  Tracking Decrement
	0cc  Stop/Eject
	0cd  Play/Pause
	0ce  Play/Skip
	0e0  Volume
	0e1  Balance
	0e2  Mute
	0e3  Bass
	0e4  Treble
	0e5  Bass Boost
	0e6  Surround Mode
	0e7  Loudness
	0e8  MPX
	0e9  Volume Increment
	0ea  Volume Decrement
	0f0  Speed Select
	0f1  Playback Speed
	0f2  Standard Play
	0f3  Long Play
	0f4  Extended Play
	0f5  Slow
	100  Fan Enable
	101  Fan Speed
	102  Light Enable
	103  Light Illumination Level
	104  Climate Control Enable
	105  Room Temperature
	106  Security Enable
	107  Fire Alarm
	108  Police Alarm
	150  Balance Right
	151  Balance Left
	152  Bass Increment
	153  Bass Decrement
	154  Treble Increment
	155  Treble Decrement
	160  Speaker System
	161  Channel Left
	162  Channel Right
	163  Channel Center
	164  Channel Front
	165  Channel Center Front
	166  Channel Side
	167  Channel Surround
	168  Channel Low Frequency Enhancement
	169  Channel Top
	16a  Channel Unknown
	170  Sub-Channel
	171  Sub-Channel Increment
	172  Sub-Channel Decrement
	173  Alternative Audio Increment
	174  Alternative Audio Decrement
	180  Application Launch Buttons
	181  AL Launch Button Configuration Tool
	182  AL Launch Button Configuration
	183  AL Consumer Control Configuration
	184  AL Word Processor
	185  AL Text Editor
	186  AL Spreadsheet
	187  AL Graphics Editor
	188  AL Presentation App
	189  AL Database App
	18a  AL Email Reader
	18b  AL Newsreader
	18c  AL Voicemail
	18d  AL Contacts/Address Book
	18e  AL Calendar/Schedule
	18f  AL Task/Project Manager
	190  AL Log/Jounal/Timecard
	191  AL Checkbook/Finance
	192  AL Calculator
	193  AL A/V Capture/Playback
	194  AL Local Machine Browser
	195  AL LAN/Wan Browser
	196  AL Internet Browser
	197  AL Remote Networking/ISP Connect
	198  AL Network Conference
	199  AL Network Chat
	19a  AL Telephony/Dialer
	19b  AL Logon
	19c  AL Logoff
	19d  AL Logon/Logoff
	19e  AL Terminal Local/Screensaver
	19f  AL Control Panel
	1a0  AL Command Line Processor/Run
	1a1  AL Process/Task Manager
	1a2  AL Select Task/Application
	1a3  AL Next Task/Application
	1a4  AL Previous Task/Application
	1a5  AL Preemptive Halt Task/Application
	200  Generic GUI Application Controls
	201  AC New
	202  AC Open
	203  AC CLose
	204  AC Exit
	205  AC Maximize
	206  AC Minimize
	207  AC Save
	208  AC Print
	209  AC Properties
	21a  AC Undo
	21b  AC Copy
	21c  AC Cut
	21d  AC Paste
	21e  AC Select All
	21f  AC Find
	220  AC Find and Replace
	221  AC Search
	222  AC Go To
	223  AC Home
	224  AC Back
	225  AC Forward
	226  AC Stop
	227  AC Refresh
	228  AC Previous Link
	229  AC Next Link
	22b  AC History
	22c  AC Subscriptions
	22d  AC Zoom In
	22e  AC Zoom Out
	22f  AC Zoom
	230  AC Full Screen View
	231  AC Normal View
	232  AC View Toggle
	233  AC Scroll Up
	234  AC Scroll Down
	235  AC Scroll
	236  AC Pan Left
	237  AC Pan Right
	238  AC Pan
	239  AC New Window
	23a  AC Tile Horizontally
	23b  AC Tile Vertically
	23c  AC Format
HUT 0d  Digitizer
	000  Undefined
	001  Digitizer
	002  Pen
	003  Light Pen
	004  Touch Screen
	005  Touch Pad
	006  White Board
	007  Coordinate Measuring Machine
	008  3D Digitizer
	009  Stereo Plotter
	00a  Articulated Arm
	00b  Armature
	00c  Multiple Point Digitizer
	00d  Free Space Wand
	020  Stylus
	021  Puck
	022  Finger
	030  Tip Pressure
	031  Barrel Pressure
	032  In Range
	033  Touch
	034  Untouch
	035  Tap
	036  Quality
	037  Data Valid
	038  Transducer Index
	039  Tablet Function Keys
	03a  Program Change Keys
	03b  Battery Strength
	03c  Invert
	03d  X Tilt
	03e  Y Tilt
	03f  Azimuth
	040  Altitude
	041  Twist
	042  Tip Switch
	043  Secondary Tip Switch
	044  Barrel Switch
	045  Eraser
	046  Tablet Pick
	047  Confidence
	048  Width
	049  Height
	051  Contact ID
	052  Input Mode
	053  Device Index
	054  Contact Count
	055  Maximum Contact Number
HUT 0f  PID Page
	000  Undefined
	001  Physical Interface Device
	020  Normal
	021  Set Effect Report
	022  Effect Block Index
	023  Parameter Block Offset
	024  ROM Flag
	025  Effect Type
	026  ET Constant Force
	027  ET Ramp
	028  ET Custom Force Data
	030  ET Square
	031  ET Sine
	032  ET Triangle
	033  ET Sawtooth Up
	034  ET Sawtooth Down
	040  ET Spring
	041  ET Damper
	042  ET Inertia
	043  ET Friction
	050  Duration
	051  Sample Period
	052  Gain
	053  Trigger Button
	054  Trigger Repeat Interval
	055  Axes Enable
	056  Direction Enable
	057  Direction
	058  Type Specific Block Offset
	059  Block Type
	05A  Set Envelope Report
	05B  Attack Level
	05C  Attack Time
	05D  Fade Level
	05E  Fade Time
	05F  Set Condition Report
	060  CP Offset
	061  Positive Coefficient
	062  Negative Coefficient
	063  Positive Saturation
	064  Negative Saturation
	065  Dead Band
	066  Download Force Sample
	067  Isoch Custom Force Enable
	068  Custom Force Data Report
	069  Custom Force Data
	06A  Custom Force Vendor Defined Data
	06B  Set Custom Force Report
	06C  Custom Force Data Offset
	06D  Sample Count
	06E  Set Periodic Report
	06F  Offset
	070  Magnitude
	071  Phase
	072  Period
	073  Set Constant Force Report
	074  Set Ramp Force Report
	075  Ramp Start
	076  Ramp End
	077  Effect Operation Report
	078  Effect Operation
	079  Op Effect Start
	07A  Op Effect Start Solo
	07B  Op Effect Stop
	07C  Loop Count
	07D  Device Gain Report
	07E  Device Gain
	07F  PID Pool Report
	080  RAM Pool Size
	081  ROM Pool Size
	082  ROM Effect Block Count
	083  Simultaneous Effects Max
	084  Pool Alignment
	085  PID Pool Move Report
	086  Move Source
	087  Move Destination
	088  Move Length
	089  PID Block Load Report
	08B  Block Load Status
	08C  Block Load Success
	08D  Block Load Full
	08E  Block Load Error
	08F  Block Handle
	090  PID Block Free Report
	091  Type Specific Block Handle
	092  PID State Report
	094  Effect Playing
	095  PID Device Control Report
	096  PID Device Control
	097  DC Enable Actuators
	098  DC Disable Actuators
	099  DC Stop All Effects
	09A  DC Device Reset
	09B  DC Device Pause
	09C  DC Device Continue
	09F  Device Paused
	0A0  Actuators Enabled
	0A4  Safety Switch
	0A5  Actuator Override Switch
	0A6  Actuator Power
	0A7  Start Delay
	0A8  Parameter Block Size
	0A9  Device Managed Pool
	0AA  Shared Parameter Blocks
	0AB  Create New Effect Report
	0AC  RAM Pool Available
HUT 10  Unicode
HUT 14  Alphanumeric Display
	000  Undefined
	001  Alphanumeric Display
	020  Display Attributes Report
	021  ASCII Character Set
	022  Data Read Back
	023  Font Read Back
	024  Display Control Report
	025  Clear Display
	026  Display Enable
	027  Screen Saver Delay
	028  Screen Saver Enable
	029  Vertical Scroll
	02a  Horizontal Scroll
	02b  Character Report
	02c  Display Data
	02d  Display Status
	02e  Stat Not Ready
	02f  Stat Ready
	030  Err Not a loadable Character
	031  Err Font Data Cannot Be Read
	032  Cursur Position Report
	033  Row
	034  Column
	035  Rows
	036  Columns
	037  Cursor Pixel Positioning
	038  Cursor Mode
	039  Cursor Enable
	03a  Cursor Blink
	03b  Font Report
	03c  Font Data
	03d  Character Width
	03e  Character Height
	03f  Character Spacing Horizontal
	040  Character Spacing Vertical
	041  Unicode Character Set
HUT 80  USB Monitor
	001  Monitor Control
	002  EDID Information
	003  VDIF Information
	004  VESA Version
HUT 81  USB Monitor Enumerated Values
HUT 82  Monitor VESA Virtual Controls
	001  Degauss
	010  Brightness
	012  Contrast
	016  Red Video Gain
	018  Green Video Gain
	01a  Blue Video Gain
	01c  Focus
	020  Horizontal Position
	022  Horizontal Size
	024  Horizontal Pincushion
	026  Horizontal Pincushion Balance
	028  Horizontal Misconvergence
	02a  Horizontal Linearity
	02c  Horizontal Linearity Balance
	030  Vertical Position
	032  Vertical Size
	034  Vertical Pincushion
	036  Vertical Pincushion Balance
	038  Vertical Misconvergence
	03a  Vertical Linearity
	03c  Vertical Linearity Balance
	040  Parallelogram Balance (Key Distortion)
	042  Trapezoidal Distortion (Key)
	044  Tilt (Rotation)
	046  Top Corner Distortion Control
	048  Top Corner Distortion Balance
	04a  Bottom Corner Distortion Control
	04c  Bottom Corner Distortion Balance
	056  Horizontal Moire
	058  Vertical Moire
	05e  Input Level Select
	060  Input Source Select
	06c  Red Video Black Level
	06e  Green Video Black Level
	070  Blue Video Black Level
	0a2  Auto Size Center
	0a4  Polarity Horizontal Sychronization
	0a6  Polarity Vertical Synchronization
	0aa  Screen Orientation
	0ac  Horizontal Frequency in Hz
	0ae  Vertical Frequency in 0.1 Hz
	0b0  Settings
	0ca  On Screen Display (OSD)
	0d4  Stereo Mode
HUT 84  Power Device Page
	000  Undefined
	001  iName
	002  Present Status
	003  Changed Status
	004  UPS
	005  Power Supply
	010  Battery System
	011  Battery System ID
	012  Battery
	013  Battery ID
	014  Charger
	015  Charger ID
	016  Power Converter
	017  Power Converter ID
	018  Outlet System
	019  Outlet System ID
	01a  Input
	01b  Input ID
	01c  Output
	01d  Output ID
	01e  Flow
	01f  Flow ID
	020  Outlet
	021  Outlet ID
	022  Gang
	023  Gang ID
	024  Power Summary
	025  Power Summary ID
	030  Voltage
	031  Current
	032  Frequency
	033  Apparent Power
	034  Active Power
	035  Percent Load
	036  Temperature
	037  Humidity
	038  Bad Count
	040  Config Voltage
	041  Config Current
	042  Config Frequency
	043  Config Apparent Power
	044  Config Active Power
	045  Config Percent Load
	046  Config Temperature
	047  Config Humidity
	050  Switch On Control
	051  Switch Off Control
	052  Toggle Control
	053  Low Voltage Transfer
	054  High Voltage Transfer
	055  Delay Before Reboot
	056  Delay Before Startup
	057  Delay Before Shutdown
	058  Test
	059  Module Reset
	05a  Audible Alarm Control
	060  Present
	061  Good
	062  Internal Failure
	063  Voltage out of range
	064  Frequency out of range
	065  Overload
	066  Over Charged
	067  Over Temperature
	068  Shutdown Requested
	069  Shutdown  Imminent
	06a  Reserved
	06b  Switch On/Off
	06c  Switchable
	06d  Used
	06e  Boost
	06f  Buck
	070  Initialized
	071  Tested
	072  Awaiting Power
	073  Communication Lost
	0fd  iManufacturer
	0fe  iProduct
	0ff  iSerialNumber
HUT 85  Battery System Page
	000  Undefined
	001  SMB Battery Mode
	002  SMB Battery Status
	003  SMB Alarm Warning
	004  SMB Charger Mode
	005  SMB Charger Status
	006  SMB Charger Spec Info
	007  SMB Selector State
	008  SMB Selector Presets
	009  SMB Selector Info
	010  Optional Mfg. Function 1
	011  Optional Mfg. Function 2
	012  Optional Mfg. Function 3
	013  Optional Mfg. Function 4
	014  Optional Mfg. Function 5
	015  Connection to SMBus
	016  Output Connection
	017  Charger Connection
	018  Battery Insertion
	019  Use Next
	01a  OK to use
	01b  Battery  Supported
	01c  SelectorRevision
	01d  Charging Indicator
	028  Manufacturer Access
	029  Remaining Capacity Limit
	02a  Remaining Time Limit
	02b  At Rate
	02c  Capacity Mode
	02d  Broadcast To Charger
	02e  Primary Battery
	02f  Charge Controller
	040  Terminate Charge
	041  Terminate Discharge
	042  Below Remaining Capacity Limit
	043  Remaining Time Limit Expired
	044  Charging
	045  Discharging
	046  Fully Charged
	047  Fully Discharged
	048  Conditioning Flag
	049  At Rate OK
	04a  SMB Error Code
	04b  Need Replacement
	060  At Rate Time To Full
	061  At Rate Time To Empty
	062  Average Current
	063  Max Error
	064  Relative State Of Charge
	065  Absolute State Of Charge
	066  Remaining Capacity
	067  Full Charge Capacity
	068  Run Time To Empty
	069  Average Time To Empty
	06a  Average Time To Full
	06b  Cycle Count
	080  Batt. Pack Model Level
	081  Internal Charge Controller
	082  Primary Battery Support
	083  Design Capacity
	084  Specification Info
	085  Manufacturer Date
	086  Serial Number
	087  iManufacturerName
	088  iDeviceName
	089  iDeviceChemistry
	08a  Manufacturer Data
	08b  Rechargeable
	08c  Warning Capacity Limit
	08d  Capacity Granularity 1
	08e  Capacity Granularity 2
	08f  iOEMInformation
	0c0  Inhibit Charge
	0c1  Enable Polling
	0c2  Reset To Zero
	0d0  AC Present
	0d1  Battery Present
	0d2  Power Fail
	0d3  Alarm Inhibited
	0d4  Thermistor Under Range
	0d5  Thermistor Hot
	0d6  Thermistor Cold
	0d7  Thermistor Over Range
	0d8  Voltage Out Of Range
	0d9  Current Out Of Range
	0da  Current Not Regulated
	0db  Voltage Not Regulated
	0dc  Master Mode
	0f0  Charger Selector Support
	0f1  Charger Spec
	0f2  Level 2
	0f3  Level 3
HUT 86  Power Pages
HUT 87  Power Pages
HUT 8c  Bar Code Scanner Page (POS)
HUT 8d  Scale Page (POS)
HUT 90  Camera Control Page
HUT 91  Arcade Control Page
HUT f0  Cash Device
	0f1  Cash Drawer
	0f2  Cash Drawer Number
	0f3  Cash Drawer Set
	0f4  Cash Drawer Status
HUT ff  Vendor Specific

# List of Languages

# Syntax:
# L language_id  language_name
#	dialect_id  dialect_name

L 0001  Arabic
	01  Saudi Arabia
	02  Iraq
	03  Egypt
	04  Libya
	05  Algeria
	06  Morocco
	07  Tunesia
	08  Oman
	09  Yemen
	0a  Syria
	0b  Jordan
	0c  Lebanon
	0d  Kuwait
	0e  U.A.E
	0f  Bahrain
	10  Qatar
L 0002  Bulgarian
L 0003  Catalan
L 0004  Chinese
	01  Traditional
	02  Simplified
	03  Hongkong SAR, PRC
	04  Singapore
	05  Macau SAR
L 0005  Czech
L 0006  Danish
L 0007  German
	01  German
	02  Swiss
	03  Austrian
	04  Luxembourg
	05  Liechtenstein
L 0008  Greek
L 0009  English
	01  US
	02  UK
	03  Australian
	04  Canadian
	05  New Zealand
	06  Ireland
	07  South Africa
	08  Jamaica
	09  Carribean
	0a  Belize
	0b  Trinidad
	0c  Zimbabwe
	0d  Philippines
L 000a  Spanish
	01  Castilian
	02  Mexican
	03  Modern
	04  Guatemala
	05  Costa Rica
	06  Panama
	07  Dominican Republic
	08  Venzuela
	09  Colombia
	0a  Peru
	0b  Argentina
	0c  Ecuador
	0d  Chile
	0e  Uruguay
	0f  Paraguay
	10  Bolivia
	11  El Salvador
	12  Honduras
	13  Nicaragua
	14  Puerto Rico
L 000b  Finnish
L 000c  French
	01  French
	02  Belgian
	03  Canadian
	04  Swiss
	05  Luxembourg
	06  Monaco
L 000d  Hebrew
L 000e  Hungarian
L 000f  Idelandic
L 0010  Italian
	01  Italian
	02  Swiss
L 0011  Japanese
L 0012  Korean
	01  Korean
L 0013  Dutch
	01  Dutch
	02  Belgian
L 0014  Norwegian
	01  Bokmal
	02  Nynorsk
L 0015  Polish
L 0016  Portuguese
	01  Portuguese
	02  Brazilian
L 0017  forgotten
L 0018  Romanian
L 0019  Russian
L 001a  Serbian
	01  Croatian
	02  Latin
	03  Cyrillic
L 001b  Slovak
L 001c  Albanian
L 001d  Swedish
	01  Swedish
	02  Finland
L 001e  Thai
L 001f  Turkish
L 0020  Urdu
	01  Pakistan
	02  India
L 0021  Indonesian
L 0022  Ukrainian
L 0023  Belarusian
L 0024  Slovenian
L 0025  Estonian
L 0026  Latvian
L 0027  Lithuanian
	01  Lithuanian
L 0028  forgotten
L 0029  Farsi
L 002a  Vietnamese
L 002b  Armenian
L 002c  Azeri
	01  Cyrillic
	02  Latin
L 002d  Basque
L 002e  forgotten
L 002f  Macedonian
L 0036  Afrikaans
L 0037  Georgian
L 0038  Faeroese
L 0039  Hindi
L 003e  Malay
	01  Malaysia
	02  Brunei Darassalam
L 003f  Kazak
L 0041  Awahili
L 0043  Uzbek
	01  Latin
	02  Cyrillic
L 0044  Tatar
L 0045  Bengali
L 0046  Punjabi
L 0047  Gujarati
L 0048  Oriya
L 0049  Tamil
L 004a  Telugu
L 004b  Kannada
L 004c  Malayalam
L 004d  Assamese
L 004e  Marathi
L 004f  Sanskrit
L 0057  Konkani
L 0058  Manipuri
L 0059  Sindhi
L 0060  Kashmiri
	02  India
L 0061  Nepali
	02  India

# HID Descriptor bCountryCode
# HID Specification 1.11 (2001-06-27) page 23
#
# Syntax:
# HCC country_code keymap_type

HCC 00  Not supported
HCC 01  Arabic
HCC 02  Belgian
HCC 03  Canadian-Bilingual
HCC 04  Canadian-French
HCC 05  Czech Republic
HCC 06  Danish
HCC 07  Finnish
HCC 08  French
HCC 09  German
HCC 10  Greek
HCC 11  Hebrew
HCC 12  Hungary
HCC 13  International (ISO)
HCC 14  Italian
HCC 15  Japan (Katakana)
HCC 16  Korean
HCC 17  Latin American
HCC 18  Netherlands/Dutch
HCC 19  Norwegian
HCC 20  Persian (Farsi)
HCC 21  Poland
HCC 22  Portuguese
HCC 23  Russia
HCC 24  Slovakia
HCC 25  Spanish
HCC 26  Swedish
HCC 27  Swiss/French
HCC 28  Swiss/German
HCC 29  Switzerland
HCC 30  Taiwan
HCC 31  Turkish-Q
HCC 32  UK
HCC 33  US
HCC 34  Yugoslavia
HCC 35  Turkish-F

# List of Video Class Terminal Types

# Syntax:
# VT terminal_type  terminal_type_name

VT 0100  USB Vendor Specific
VT 0101  USB Streaming
VT 0200  Input Vendor Specific
VT 0201  Camera Sensor
VT 0202  Sequential Media
VT 0300  Output Vendor Specific
VT 0301  Generic Display
VT 0302  Sequential Media
VT 0400  External Vendor Specific
VT 0401  Composite Video
VT 0402  S-Video
VT 0403  Component Video
`
