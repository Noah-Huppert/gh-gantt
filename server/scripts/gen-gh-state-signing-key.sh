#!/usr/bin/env bash
#?
# gen-gh-state-signing-key - Generate ed25519 key pair which is used to sign the state field in a GitHub auth request
#
# USAGE
#	gen-gh-state-signing-key
#
# BEHAVIOR
#	Generates an ed25519 keypair and saves the result in a .env file under the environment variable keys:
#		- APP_GH_STATE_SIGNING_PRIV_KEY
#		- APP_GH_STATE_SIGNING_PUB_KEY
#
# REQUIREMENTS
#	Requires a version of ssh-keygen which supports ed25519 keypairs.
#?

priv_key_path="/tmp/gh_state_signing"
pub_key_path="$priv_key_path.pub"


echo "######################"
echo "# Generating Keypair #"
echo "######################"

if ! ssh-keygen \
		-t ed25519 \
		-f "$priv_key_path" \
		-P ""; then
	echo "Error: failed to generate keypair" &>2
	exit 1
fi

echo "##################"
echo "# Saving Keypair #"
echo "##################"

# Save into .env file
echo "export APP_GH_STATE_SIGNING_PRIV_KEY=\"$(cat $priv_key_path)\"" >> .env
if [[ "$?" != "0" ]]; then
	echo "Error: Failed to save private key in .env file" >&2
	exit 1
fi

echo "export APP_GH_STATE_SIGNING_PUB_KEY=\"$(cat $pub_key_path)\"" >> .env
if [[ "$?" != "0" ]]; then
	echo "Error: Failed to save public key in .env file" >&2
	exit 1
fi

# Remove from disk
if ! rm "$priv_key_path"; then
	echo "Error: Failed to remove private key from disk" >&2
	exit 1
fi

if ! rm "$pub_key_path"; then
	echo "Error: Failed to remove public key from disk" >&2
	exit 1
fi

echo "===== Success ====="
