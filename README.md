Cowboy
=======

The good, the bad, and the ugly.

## Watch `cowboy` in action!

![watch cowboy in action!](https://raw.github.com/vishaltelangre/cowboy/master/static/sneak_peak.gif)

## Usage

From any Slack channel, just type `/<command> [search terms]`.

For example:

```
/imdb casablanca
/excuse
```

## Available commands
- Movie details lookup on IMDb (URL: http://cowboy-slack.herokuapp.com/movie.slack)
- Get excuses to spit on your boss' face (URL: http://cowboy-slack.herokuapp.com/excuse.slack)
- More coming soon...

## TODO
- DDG search
- Weather forecast
- Simple calculations
- Wiki lookup

## Integrate with your Slack team

1. Go to your channel
2. Click on **Integrations**.
3. Scroll all the way down to **DIY Integrations & Customizations section**.
4. For example, to add above movie lookup command, click on **Add** next to **Slash Commands**.
  - Command: `/imdb` or whatever you like
  - URL: `http://cowboy-slack.herokuapp.com/movie.slack`
  - Method: `POST`
  - For the **Autocomplete help text**, check to show the command in autocomplete list.
    - Description: `Show movie details from IMDb`
    - Usage hint: `[movie]`
  - Descriptive Label: `Movie Lookup`
  5. Do same for other available commands.

## Wanna hack?

Follow [this](https://devcenter.heroku.com/articles/getting-started-with-go) tutorial to setup this project locally, and optionally deploy it on Heroku. This project uses `foreman` utility to spin up/off server, `Godeps` to manage third-party libraries.

## Wanna Contribute?

- Please use the [issue tracker](https://github.com/vishaltelangre/cowboy/issues) to report any bugs or file feature requests.

## Thankings
- This project is inspired from @karan's [overflow](https://github.com/karan/slack-overflow), but is way more powerful!
- Movie details are retrieved from http://www.omdbapi.com/. Thanks to the creator of this site.
- Source of funny programmer's excuses: http://www.programmerexcuses.com/

## Copyright and License

Copyright (c) 2015, Vishal Telangre. All Rights Reserved.

This project is licenced under the [MIT License](LICENSE.md).