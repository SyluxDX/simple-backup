# simple-backup

A simple go script to backup small files. 

Had option to only backup file there were changes on the target file

#### Configuration file

```json
{
	"sourceFolder": "/data",
	"files": [
        "files1.txt",
        "files2.txt"
	],
	"backupFolder": "backups",
	"logFolder": "logs",
	"exampleFormat": "2006-01-02 15:04:05",
	"backupFormat": "20060102_150405",
	"backupOnlyChanges": true
}
```
| Field | Description |
|---|---|
| sourceFolder | Folder containng the files to be backed up |
| files | List with filenames to be backup |
| backupFolder | Folder where the backups file will be stored |
| logFolder | Folder where the logs will be written |
| exampleFormat | This field shows Golang timestamp format string, it is unsued |
| backupFormat | Datetime format that is appened to backup filename |
| backupOnlyChanges | True/false Flag, indicates if the backup is to occur only if the file has changed |