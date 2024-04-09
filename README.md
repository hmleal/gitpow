# GitPow

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Overview

GitPow is a PowerShell command-line interface that seamlessly integrates Git
functionality with PowerShell by offering Git status summaries and tab
completion.

### Why GitPow?

PowerShell lacks native integration with Git, making it challenging for new
users to readily adopt the terminal as their primary Git tool.

This project draws heavy inspiration from Posh-Git, which excels in providing
similar functionality. I highly recommend checking out their repository.

However, we've chosen to rewrite the functionality for several reasons:

- **Performance:** Posh-Git can be slow, particularly in large Git repositories.
- **Accessibility:** Posh-Git is written in PowerShell, limiting contributions
from developers more comfortable with other languages.
- **Review and Merge:** Historically, Posh-Git has experienced delays in
reviewing and merging pull requests.

By addressing these issues, `GitPow` aims to provide a faster, more accessible,
and easier-to-contribute-to alternative for integrating Git with PowerShell.

## Git summary status

By default GitPow dosen't add anything into your setup, you have to integrate it
manually with your powershell.

```powershell
function prompt {
    $cmdLine = "gitpow status"
    "PS $($executionContext.SessionState.Path.CurrentLocation)$(Invoke-Expression $cmdLine)$('>' * ($nestedPromptLevel + 1)) ";
}
```

Please refer to the official powershell documentation for more information:

- https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_prompts?view=powershell-7.4

## License

This project is licensed under the [MIT License](LICENSE). You are free to use,
modify, and distribute the dataset for both commercial and non-commercial
purposes, subject to the conditions of the license.
