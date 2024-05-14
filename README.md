# Supastore

A Simple CLI to interact with your supabase storage

This is the v.0.1 of this, developed for quick use.
Check out the `dev` branch for developments on a pretty UI and more

## Installation

For now, the easiest way to install this is through go

```bash
go install github.com/newtoallofthis123/supastore
```

After this, setup the `~/.config/supastore/.env` file and change the variables 
in the example.
> Note: The secret key should be able to bypass RLS

You can also directly set the environment variables in your shell.
With this done, run `supastore version` to check :)

## Mini Doc

> This can be accessed using supastore help

Supastore: Easily interact with your supabase storage
general-usuage: supastore bucket-id command filenames...

Commands:
`download fileNameInSupabase <downloadName>`: Download a file
`upload files...` : Upload files to supabase
`url filename`: Get the public url of a file
`info`: Information about the storage bucket
`list`: List the store contents
`init`: Initialize Store if it doesn't exist

## License

This project is licensed using the MIT License. For more details, checkout the [LICENSE](LICENSE) file.
