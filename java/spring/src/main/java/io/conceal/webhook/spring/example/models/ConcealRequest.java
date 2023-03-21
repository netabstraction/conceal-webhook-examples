package io.conceal.webhook.spring.example.models;

public class ConcealRequest {
    private String event;
    private String host;
    private String sourceType;
    private String companyId;
    private String companyName;
    private String userEmail;
    private String userId;
    private String url;
    private int count;
    private Decision decision;
    private String finalDecision;
    private long timeStamp;
    public ConcealRequest() {
    }
    public String getEvent() {
        return event;
    }
    public void setEvent(String event) {
        this.event = event;
    }
    public String getHost() {
        return host;
    }
    public void setHost(String host) {
        this.host = host;
    }
    public String getSourceType() {
        return sourceType;
    }
    public void setSourceType(String sourceType) {
        this.sourceType = sourceType;
    }
    public String getCompanyId() {
        return companyId;
    }
    public void setCompanyId(String companyId) {
        this.companyId = companyId;
    }
    public String getCompanyName() {
        return companyName;
    }
    public void setCompanyName(String companyName) {
        this.companyName = companyName;
    }
    public String getUserEmail() {
        return userEmail;
    }
    public void setUserEmail(String userEmail) {
        this.userEmail = userEmail;
    }
    public String getUserId() {
        return userId;
    }
    public void setUserId(String userId) {
        this.userId = userId;
    }
    public String getUrl() {
        return url;
    }
    public void setUrl(String url) {
        this.url = url;
    }
    public int getCount() {
        return count;
    }
    public void setCount(int count) {
        this.count = count;
    }
    public Decision getDecision() {
        return decision;
    }
    public void setDecision(Decision decision) {
        this.decision = decision;
    }
    public String isFinalDecision() {
        return finalDecision;
    }
    public void setFinalDecision(String finalDecision) {
        this.finalDecision = finalDecision;
    }
    public long getTimeStamp() {
        return timeStamp;
    }
    public void setTimeStamp(long timeStamp) {
        this.timeStamp = timeStamp;
    }
    @Override
    public int hashCode() {
        final int prime = 31;
        int result = 1;
        result = prime * result + ((event == null) ? 0 : event.hashCode());
        result = prime * result + ((host == null) ? 0 : host.hashCode());
        result = prime * result + ((sourceType == null) ? 0 : sourceType.hashCode());
        result = prime * result + ((companyId == null) ? 0 : companyId.hashCode());
        result = prime * result + ((companyName == null) ? 0 : companyName.hashCode());
        result = prime * result + ((userEmail == null) ? 0 : userEmail.hashCode());
        result = prime * result + ((userId == null) ? 0 : userId.hashCode());
        result = prime * result + ((url == null) ? 0 : url.hashCode());
        result = prime * result + count;
        result = prime * result + ((decision == null) ? 0 : decision.hashCode());
        result = prime * result + ((finalDecision == null) ? 0 : finalDecision.hashCode());
        result = prime * result + (int) (timeStamp ^ (timeStamp >>> 32));
        return result;
    }
    @Override
    public boolean equals(Object obj) {
        if (this == obj)
            return true;
        if (obj == null)
            return false;
        if (getClass() != obj.getClass())
            return false;
        ConcealRequest other = (ConcealRequest) obj;
        if (event == null) {
            if (other.event != null)
                return false;
        } else if (!event.equals(other.event))
            return false;
        if (host == null) {
            if (other.host != null)
                return false;
        } else if (!host.equals(other.host))
            return false;
        if (sourceType == null) {
            if (other.sourceType != null)
                return false;
        } else if (!sourceType.equals(other.sourceType))
            return false;
        if (companyId == null) {
            if (other.companyId != null)
                return false;
        } else if (!companyId.equals(other.companyId))
            return false;
        if (companyName == null) {
            if (other.companyName != null)
                return false;
        } else if (!companyName.equals(other.companyName))
            return false;
        if (userEmail == null) {
            if (other.userEmail != null)
                return false;
        } else if (!userEmail.equals(other.userEmail))
            return false;
        if (userId == null) {
            if (other.userId != null)
                return false;
        } else if (!userId.equals(other.userId))
            return false;
        if (url == null) {
            if (other.url != null)
                return false;
        } else if (!url.equals(other.url))
            return false;
        if (count != other.count)
            return false;
        if (decision == null) {
            if (other.decision != null)
                return false;
        } else if (!decision.equals(other.decision))
            return false;
        if (finalDecision != other.finalDecision)
            return false;
        if (timeStamp != other.timeStamp)
            return false;
        return true;
    }
    @Override
    public String toString() {
        return "ConcealRequest [event=" + event + ", host=" + host + ", sourceType=" + sourceType + ", companyId="
                + companyId + ", companyName=" + companyName + ", userEmail=" + userEmail + ", userId=" + userId
                + ", url=" + url + ", count=" + count + ", decision=" + decision + ", finalDecision=" + finalDecision
                + ", timeStamp=" + timeStamp + "]";
    }
}
