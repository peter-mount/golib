// Banner is an implementation of SYSVbanner in GO.
//
// Unlike other code this is under the following copyright from SYSVbanner.c
//
// *****************************************************************
// *
// * SYSVbanner.c
// *
// * This is a PD version of the SYS V banner program (at least I think
// * it is compatible to SYS V) which I wrote to use with the clock
// * program written by:
// **     DCF, Inc.
// **     14623 North 49th Place
// **     Scottsdale, AZ 85254
// * and published in the net comp.sources.misc newsgroup in early July
// * since the BSD banner program works quite differently.
// *
// * There is no copyright or responsibility accepted for the use
// * of this software.
// *
// * Brian Wallis, brw@jim.odr.oz, 4 July 1988
// *
// *****************************************************************
//
// The version I based this on is here: https://sources.debian.org/src/sysvbanner/1:1.0-16/sysvbanner.c/
//
// This package provides a type BannerText and 2 functions to generate it.
//
// Banner( string ) will generate BannerText from up to the first 10 characters
// in string, matching the output of sysvbanner.
//
// BannerLen( string, len ) is the same but with a customisable length.
// For 80 character wide lines then len=10 matches that of Banner().
// For 132 character wide lines then use len=16.
//
// The returned type BannerText has one function String() which returns the
// text as a single string with '\n' between each line.
//
// Alternatively you can use it as a slice which has 8 entries, one per line.
//
// So to print out a banner text you can one of the following, each one
// generates the same result:
//
// fmt.Println( banner.Banner( "MyBanner" ) )
//
// fmt.Println( banner.Banner( "MyBanner" ).String() )
//
// for _, s := range banner.Banner( "MyBanner" ) { fmt.Println( s ) }
//
// myBanner := banner.Banner( "MyBanner" )
// fmt.Println( myBanner[0] )
// fmt.Println( myBanner[1] )
// fmt.Println( myBanner[2] )
// fmt.Println( myBanner[3] )
// fmt.Println( myBanner[4] )
// fmt.Println( myBanner[5] )
// fmt.Println( myBanner[6] )
// fmt.Println( myBanner[7] )
//
// Final note: element 7 (the last one) is always "" and not the same length as
// the other elements (0..6).
package banner

import (
  "strings"
)

var glyphs = []string {
  "         ###  ### ###  # #   ##### ###   #  ##     ###  ",
  "         ###  ### ###  # #  #  #  ## #  #  #  #    ###   ",
  "         ###   #   # ########  #   ### #    ##      #   ",
  "          #            # #   #####    #    ###     #    ",
  "                     #######   #  #  # ####   # #       ",
  "         ###           # #  #  #  # #  # ##    #        ",
  "         ###           # #   ##### #   ### #### #       ",

  "   ##    ##                                            #",
  "  #        #   #   #    #                             # ",
  " #          #   # #     #                            #  ",
  " #          # ### ### #####   ###   #####           #   ",
  " #          #   # #     #     ###           ###    #    ",
  "  #        #   #   #    #      #            ###   #     ",
  "   ##    ##                   #             ###  #      ",

  "  ###     #    #####  ##### #      ####### ##### #######",
  " #   #   ##   #     ##     ##    # #      #     ##    # ",
  "# #   # # #         #      ##    # #      #          #  ",
  "#  #  #   #    #####  ##### ####### ##### ######    #   ",
  "#   # #   #   #            #     #       ##     #  #    ",
  " #   #    #   #      #     #     # #     ##     #  #    ",
  "  ###   ##### ####### #####      #  #####  #####   #    ",

  " #####  #####    #     ###      #           #     ##### ",
  "#     ##     #  # #    ###     #             #   #     #",
  "#     ##     #   #            #     #####     #        #",
  " #####  ######         ###   #                 #     ## ",
  "#     #      #   #     ###    #     #####     #     #   ",
  "#     ##     #  # #     #      #             #          ",
  " #####  #####    #     #        #           #       #   ",

  " #####    #   ######  ##### ###### ############## ##### ",
  "#     #  # #  #     ##     ##     ##      #      #     #",
  "# ### # #   # #     ##      #     ##      #      #      ",
  "# # # ##     ####### #      #     ######  #####  #  ####",
  "# #### ########     ##      #     ##      #      #     #",
  "#     ##     ##     ##     ##     ##      #      #     #",
  " ##### #     #######  ##### ###### ########       ##### ",

  "#     #  ###        ##    # #      #     ##     ########",
  "#     #   #         ##   #  #      ##   ####    ##     #",
  "#     #   #         ##  #   #      # # # ## #   ##     #",
  "#######   #         ####    #      #  #  ##  #  ##     #",
  "#     #   #   #     ##  #   #      #     ##   # ##     #",
  "#     #   #   #     ##   #  #      #     ##    ###     #",
  "#     #  ###   ##### #    # ########     ##     ########",

  "######  ##### ######  ##### ########     ##     ##     #",
  "#     ##     ##     ##     #   #   #     ##     ##  #  #",
  "#     ##     ##     ##         #   #     ##     ##  #  #",
  "###### #     #######  #####    #   #     ##     ##  #  #",
  "#      #   # ##   #        #   #   #     # #   # #  #  #",
  "#      #    # #    # #     #   #   #     #  # #  #  #  #",
  "#       #### ##     # #####    #    #####    #    ## ## ",

  "#     ##     ######## ##### #       #####    #          ",
  " #   #  #   #      #  #      #          #   # #         ",
  "  # #    # #      #   #       #         #  #   #        ",
  "   #      #      #    #        #        #               ",
  "  # #     #     #     #         #       #               ",
  " #   #    #    #      #          #      #               ",
  "#     #   #   ####### #####       # #####        #######",

  "  ###                                                   ",
  "  ###     ##   #####   ####  #####  ###### ######  #### ",
  "   #     #  #  #    # #    # #    # #      #      #    #",
  "    #   #    # #####  #      #    # #####  #####  #     ",
  "        ###### #    # #      #    # #      #      #  ###",
  "        #    # #    # #    # #    # #      #      #    #",
  "        #    # #####   ####  #####  ###### #       #### ",

  "                                                        ",
  " #    #    #        # #    # #      #    # #    #  #### ",
  " #    #    #        # #   #  #      ##  ## ##   # #    #",
  " ######    #        # ####   #      # ## # # #  # #    #",
  " #    #    #        # #  #   #      #    # #  # # #    #",
  " #    #    #   #    # #   #  #      #    # #   ## #    #",
  " #    #    #    ####  #    # ###### #    # #    #  #### ",

  "                                                        ",
  " #####   ####  #####   ####   ##### #    # #    # #    #",
  " #    # #    # #    # #         #   #    # #    # #    #",
  " #    # #    # #    #  ####     #   #    # #    # #    #",
  " #####  #  # # #####       #    #   #    # #    # # ## #",
  " #      #   #  #   #  #    #    #   #    #  #  #  ##  ##",
  " #       ### # #    #  ####     #    ####    ##   #    #",

  "                       ###     #     ###   ##    # # # #",
  " #    #  #   # ###### #        #        # #  #  # # # # ",
  "  #  #    # #      #  #        #        #     ## # # # #",
  "   ##      #      #  ##                 ##        # # # ",
  "   ##      #     #    #        #        #        # # # #",
  "  #  #     #    #     #        #        #         # # # ",
  " #    #    #   ######  ###     #     ###         # # # #",
}

type BannerText []string

// NewBanner converts the first 10 characters of a string into large letters.
// The return value is a slice of 8 strings, one per row.
//
// The 10 limit matches that of the SYSVbanner/Debian Banner utility.
//
func Banner( s string ) BannerText {
  return BannerLen( s, 10 )
}

// NewBannerLen converts the first l characters of a string into large letters.
// The return value is a slice of 8 strings, one per row.
//
// If l <= 0 then the entire string is used regardless of it's length.
//
// As a guide: l=10 for 80 characters, 16 for 132.
func BannerLen( s string, l int ) BannerText {
  str := s
  if l>0 && l<len(s) {
    str = str[:l]
  }

  lines := make([][]rune,7)
  for _, ch := range str {
    if ch < 32 {
      ch = 32
    } else if ch > 127 {
      ch = 127
    }
    a := int(ch-' ')
    b := a/8*7
    c := a%8*7
    for i:=0; i<7; i++ {
      lines[i] = append( lines[i], []rune(glyphs[b+i][c:c+7])... )
      lines[i] = append( lines[i], ' ' )
    }
  }

  var ret []string
  for _, ls := range lines {
    ret = append( ret, string(ls) )
  }
  ret = append( ret, "" )
  return ret
}

// String returns the banner as a single string
func (b BannerText) String() string {
  return strings.Join(b,"\n")
}
