Cowboy
=======

The good, the bad, and the ugly.

## Watch `cowboy` in action!

![watch cowboy in action!](https://raw.github.com/vishaltelangre/cowboy/master/static/sneak_peak.gif?v=aug14)

## Usage

From any Slack channel, just type `/<command> [search terms]`.

For example:

```
/imdb casablanca
/excuse
/producthunt_posts 2
```

## Available commands

- Movie details lookup on IMDb (URL: http://cowboy-slack.herokuapp.com/movie.slack)
- Get fine excuses to convince your boss (URL: http://cowboy-slack.herokuapp.com/excuse.slack)
- Get all featured/trending products from Product Hunt (URL: http://cowboy-slack.herokuapp.com/producthunt/posts.slack)

## More coming soon...

- HN/best
- DDG search
- Weather forecast
- Simple calculations
- Define word
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

## Important Note

- If you are setting up your slack commands by relying on `http://cowboy-slack.herokuapp.com` server, then FYI, I am using Heroku's free dyno, which goes asleep after 20 minutes of inactivity. So if a command doesn't work for your first time, try it again after 5 seconds, as server would be awaken and wouldn't respond and that initial request times out, but subsequent requests will work quickly without a retry.
- A free Heroku dyno could stay awake for 18 hours a day -- if this limit gets exceeded, Heroku shuts that dyno for 6 hours to recharge it. In such a case the server won't be able to respond any request for 6 hours.
- If you want these commands respond all the time consistently, setup cowboy on your own premium server, follow next section on howto.

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

<a target='_blank' rel='nofollow' href='https://app.codesponsor.io/link/PfwgcRiC73ERAe1WTDUo4DmM/vishaltelangre/cowboy'>
  <img alt='Sponsor' width='888' height='68' src='https://app.codesponsor.io/embed/PfwgcRiC73ERAe1WTDUo4DmM/vishaltelangre/cowboy.svg' />
</a>
