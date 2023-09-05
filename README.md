# permissura

An opinionated (and slightly better way) to manage your Hasura permissions.

By managing your permisions rules externally, you jail-break it and gain new
powers: YAML anchors and comments for readability; the ability to choose
whatever directory structure suits you; etc.

## Getting started

### Installing `permissura`

```shell
go install github.com/kevinmichaelchen/permissura@latest
```

### Writing your permissions rules

For a realistic example, see [**employee.yaml**][employee-rule].

In our example, we store our rule in [**./examples/policies**][policies], but
the CLI doesn't care how you name and structure your directories and
subdirectories. You just point it at a directory and it recursively parses out
any YAML files it finds. It's the actual content and structure of the YAML files
that matters.

[policies]: ./examples/policies
[employee-rule]: ./examples/policies/employee.yaml

### Syncing permissions to Hasura

As you write new permissions logic externally, you'll want to make sure you're
syncing them back into Hasura's GraphQL Engine.

You can do so with `permissura`:

```shell
permissura \
  --default-source default \
  --dir ./examples/policies \
  --debug \
  sync
```

Don't forget to [export][export] your metadata after you sync it. This is
necessary since the Hasura GraphQL Engine [does not have
access][gh-issue-md-fs-access] to your local filesystem.

[gh-issue-md-fs-access]: https://github.com/hasura/graphql-engine/issues/8272
[export]:
  https://hasura.io/docs/latest/hasura-cli/commands/hasura_metadata_export/

## FAQ

### How does it work?

We read permissions rules from YAML files stored locally on your machine.

We parse them into Hasura-compatible JSON and then upload them via Hasura's
[REST API][hasura-permissions-api].

### Why a separate directory?

Hasura manages permissions in large, cumbersome files with lots of other
metadata that won't be germane if you're only wanting to get a holistic view of
authorization rules. Moreover, Hasura only lets you see rules as they pertain to
one specific table: it's impossible AFAIK to see a Hasura role as it applies to
_all_ tables.

Managing permissions YAML outside of Hasura's metadata directory has a few other
advantages. Any YAML comments you write to explain your business rules won't be
overwritten by Hasura. You can also use YAML anchors to DRY things up. Both of
these features vastly improve readability.

### How can I view my rules when they have anchors in them?

Sometimes when your YAML is really DRY and you've got anchors, it can be hard to
fully read your permissions logic.

You can "explode" the anchors using [`yq`][yq], [`yj`][yj], and [`jq`][jq]:

```shell
yq '.rules[0] | explode(.)' examples/policies/employee.yaml | yj | jq
```

### What does the future of this project look like?

I'd like to open-source this tool, as it could provide value to other Hasura
users, it could validate our use case, and it fully acknowledges that we benefit
from lots of other open-source command-line utilities.

At some point, it could even be a contribution to Hasura's open-source ecosystem
of [CLI plugins][hasura-plugins].

[hasura-permissions-api]:
  https://hasura.io/docs/latest/api-reference/metadata-api/permission
[hasura-plugins]:
  https://hasura.io/docs/latest/hasura-cli/commands/hasura_plugins/
[yj]: https://github.com/sclevine/yj
[yq]: https://github.com/mikefarah/yq
[jq]: https://github.com/jqlang/jq
