package leetcode

import (
	"testing"
)

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				S: "apple",
			},
			want: "applemaa",
		},
		{
			name: "test2",
			args: args{
				S: "I speak Goat Latin",
			},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{
			name: "test3",
			args: args{
				S: "The quick brown fox jumped over the lazy dog",
			},
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
		{
			name: "test4",
			args: args{
				S: "Each word consists of lowercase and uppercase letters only",
			},
			want: "Eachmaa ordwmaaa onsistscmaaaa ofmaaaaa owercaselmaaaaaa andmaaaaaaa uppercasemaaaaaaaa etterslmaaaaaaaaa onlymaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.S); got != tt.want {
				t.Errorf("toGoatLatin() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
