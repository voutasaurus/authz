# authz

An tool for ACME that returns the authz URL associated with a given challenge URL.

# Install

  go get -u github.com/voutasaurus/authz
  
# Usage

  $ authz "https://acme-v01.api.letsencrypt.org/acme/challenge/XXXXXXXX/iiiiiiii"
  https://acme-v01.api.letsencrypt.org/acme/authz/XXXXXXXX

# Why

Useful for deactivating pending authz only given the challenge url.
Trying to do the same in a bash pipeline was doing my head in.

# Support

Use at your own risk. PRs welcome.
