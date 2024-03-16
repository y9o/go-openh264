#
# Reads simple enums, structs, and defines from header files.
#
# perl ./h2go.pl
# go tool cgo -godefs _codec.go > codec_api.go
#
use strict;
my $inc_dir = "inc_openh264";

open my $fh,">","_codec.go" or die $!;
print $fh <<TXT;
// Code generated by h2go.pl; DO NOT EDIT.
package openh264

// #include "$inc_dir/codec_api.h"
import "C"
TXT

h2go("$inc_dir/codec_api.h",$fh);
h2go("$inc_dir/codec_app_def.h",$fh);
h2go("$inc_dir/codec_def.h",$fh);
h2go("$inc_dir/codec_ver.h",$fh);
close $fh;

sub h2go{
	my($file,$fh)=@_;
	my $header = do {
		local $/ = undef;
		open my $fh, "<", $file or die "$file $!";
		<$fh>;
	};

	#Delete comment.
	$header =~ s!/\*.*?\*/!!sg;
	$header =~ s!//.*!!g;

	print $fh "//file:$file\n";
	while($header =~ /(typedef)?\s*(enum|struct\s*\w*)\s*({(?:(?>[^{}]+)|(?3))*})(.*?;)/sg){
		my($typedef,$type,$value,$names)=($1,$2,$3,$4);
		if($type eq "enum"){
			my($name,$type,@list)=enum($1,$2,$3,$4);
			if($name ne ""){
				print $fh "//type $name int\n";
			}
			if(@list){
				print $fh "const(\n";
				for my $val (@list){
					print $fh "\t",ucfirst($val)," = C.$val\n";
				}
				print $fh ")\n";
			}
		}elsif(index($type,"struct")==0){
			my $n=0;
			for my $name (split(/,/,$names)){
				$name =~ s/\s|;$//g;
				if($name && $name !~ /^\*/){
					print $fh "type ",ucfirst($name)," C.$name\n";
					$n++;
				}
			}
			if($n == 0 && $type =~ /struct\s+(\w+)/){
				print $fh "type ",ucfirst($1)," C.$1\n"
			}
		}
	}
	my @define;
	while($header =~ /^#define\s+(\w+)\s+\(?\s*(-?(?:0x)?[\d\.,]+L?)\s*\)?\s*$/igm){
			push @define ,[$1,$2];
	}
	if(@define){
		print $fh "//#define\n";
		print $fh "const(\n";
		for my $d (@define){
			print $fh "\t",ucfirst($d->[0])," = ", $d->[1] ,"\n";
		}
		print $fh ")\n";

	}
}

sub enum{
	my ($typedef,$enum,$values,$name)=@_;
	$name =~ s/\s|;//g;
	$values=~s/{|}//g;
	my @list;
	for my $val(split(/,/,$values)){
		$val =~ s/=.*//;
		$val =~ s/\s//g;
		if($val){
			push @list,$val;
		}
	}
	return ($name,"int",@list);
}
