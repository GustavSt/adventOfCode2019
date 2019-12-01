param(
    [Parameter(Mandatory = $true)]
    [string]
    $day
)
$source = "./dayx/*"
$destination = "./day$day"
write-host $destination
if(! (Test-Path $destination)){
    New-Item -ItemType "Directory" -Path $destination -Force
    Copy-Item -Force -Recurse -Verbose $source -Destination $destination

    Rename-Item -Path "$destination/dayx.go" -NewName "day$day.go"

    (Get-Content "$destination/task.go") -replace '0', $day | Set-Content "$destination/task.go"
    write-host "Scaffold for day $day"
} else {
    Write-host "Day $day already exist"
}
