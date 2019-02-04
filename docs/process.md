# Development Process Document

## API Implementation

### High Priority

- https://wakatime.com/developers#projects
- https://wakatime.com/developers#heartbeats
- https://wakatime.com/developers#durations
- https://wakatime.com/developers#stats
    - `catrgories` means `Coding`, `Designing` and other activities. For now we ignore this field, as most editors actually only cares about `Coding`.
    - `editor`: in CLI, `editor` is not a dedicated field but part of the User-Agent, which we would not parse. For not just let `editor` field be empty.
    - `operating_systems`: in CLI, `operating_systems` is not a dedicated field but part of the User-Agent, which we would not parse. For not just let `operating_systems` field be empty. 
    - `dependencies`: a HeartBeat can hold multiple dependencies. Due to implementation issues, for now we do not save it to the InfluxDB, nor do we count them in stat, and let `dependencies` field be empty. 
- https://wakatime.com/developers#users

### Medium Priority

- https://wakatime.com/developers#summaries
- https://wakatime.com/developers#user_agents
- https://wakatime.com/developers#commits

### Low Priority

- https://wakatime.com/developers#goals
- https://wakatime.com/developers#leaders
- https://wakatime.com/developers#private_leaderboards
- https://wakatime.com/developers#private_leaderboards_leaders
- https://wakatime.com/developers#team_member_durations
- https://wakatime.com/developers#team_member_summaries
- https://wakatime.com/developers#team_members
- https://wakatime.com/developers#teams
