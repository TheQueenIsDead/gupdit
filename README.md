# Gupdit


## What? 

**G**o **Up**dates via G**it**. (G~~o~~Upd~~atesviaG~~it)

## Huh?

Ever developed an internal tool for an organisation?

Gone through the pain of setting up versioning and releases?

Promoted the new versions and encouraged your colleagues to update via their preferred update mechanism?

And then suffered an outage due to a code generation tool being ran that was several versions behind and then failed to
include graceful shutdown logic for process interrupts?

Yeah, me neither.

## What does this actually do?

Gupdit is a go library that can be configured to verify whether the version
of an installed binary is up to date based on the latest semver tag in a git repository.

### Example Usage