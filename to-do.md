# To-Do

- Set up GORM models []

  - User []

    - email: string
    - uuid: uuid

  - Creator []

    - name: string
    - uuid: uuid

  - Writing []
    - title: string
    - description: string
    - tags (including genre): []string
    - genres: []string
    - writing type: string (poetry, fiction, non-fiction)
    - fiction category: string (short story, novellette, novella, novel)
    - non-fiction category: string
  - Tag []
  - Genre []
  - Creator_Donations []
  - FreeCreate_Donations []
