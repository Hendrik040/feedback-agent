### Planning 

We need an agent that runs continuesly, takes in audio and text data as input and analysis it. 

- if it is something positive but it into one bucket (cudos & carrots)
- If it is something negative -> send it down a different analysis chain 
    - Negative feedback that has no real feedback or not actionable items that can be extracted from it has to go into one sub-bucket 

```mermaid
flowchart TD
    %% --- FRONTEND SECTION ---
    subgraph A[Frontend UI]
        A1[Feedback Submission Form]
        A2[Feedback Board]
        A3[Upvote Button]
        A4[Admin Dashboard]
    end

    %% --- BACKEND SECTION ---
    subgraph B[Backend Services]
        B1[API Gateway]
        B2[Feedback Controller]
        B3[Sentiment Analysis Service]
        B4[Database Service]
        B5[Scheduler / Agent]
    end

    %% --- DATA STORAGE ---
    subgraph C[Database]
        C1[(feedback)]
        C2[(votes)]
        C3[(users)]
        C4[(sentiment_logs)]
    end

    %% --- EXTERNAL SERVICES ---
    subgraph D[Integrations]
        D1[LLM / Sentiment API]
        D2[Jira / Notion / Slack]
    end

    %% --- CONNECTIONS ---
    A1 -->|Submit Feedback| B1
    A3 -->|Upvote| B1
    A4 -->|View/Filter| B1

    B1 --> B2
    B2 -->|Save| C1
    B2 -->|Fetch votes| C2
    B2 -->|Trigger Sentiment| B3
    B3 -->|Analyze text| D1
    D1 -->|Return sentiment + confidence| B3
    B3 -->|Store Result| C4
    B3 -->|Tag Feedback| C1ß

    B2 --> B4
    B4 --> C1
    B4 --> C2
    B4 --> C3
    B4 --> C4

    B5 -->|Weekly Digest / Auto-Tagging| B2
    B5 -->|Export / Notify| D2

    A2 -->|Display data| C1
    A4 -->|Visualize insights| C4

    %% --- SENTIMENT BUCKETS ---
    subgraph E[Sentiment Buckets]
        E1[Kudos & Carrots (Positive)]
        E2[Critiques (Negative)]
        E3[Neutral / Context]
        E4[Non-actionable Sub-Bucket]
    end

    C1 -->|Classify via Sentiment| E1 & E2 & E3
    E2 --> E4
```