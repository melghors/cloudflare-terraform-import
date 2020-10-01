**No longer maintained, there is [an official Cloudflare importer tool](https://github.com/cloudflare/cf-terraforming) to import existing resources into terraform.**

# quick'n dirty cloudflare to terraform importer for DNS records

(Only tested with A, CAA, CNAME, MX, TXT records so far) 

## usage

### import all zones of your cf account:

export CF_API_EMAIL="example@gmail.com"                                  
export CF_API_KEY="your-api-key"
./cti zones all

zones
├── zones.sh
└── zones.tf

### import dns records of specific zone:

export CF_API_EMAIL="example@gmail.com"                                  
export CF_API_KEY="your-api-key"
./cti import test123.com

test123.com
├── records-test123.com.sh
└── records-test123.com.tf

## build & run

Requires Go >= 1.11

```
$ git clone git@github.com:emgag/cloudflare-terraform-import.git
$ go get && make
$ export CF_API_EMAIL="user@example.org" 
$ export CF_API_KEY="..."
$ ./cti import example.org
``` 

* Copy contents of `import.tf` to your terraform config.
* Review 
* Run `import.sh` to import records into state
