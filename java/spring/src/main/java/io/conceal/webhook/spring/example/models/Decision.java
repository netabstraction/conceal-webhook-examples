package io.conceal.webhook.spring.example.models;

public class Decision {
    private String enforceTls;
    private String noIp;
    public Decision() {
    }
    public String isEnforceTls() {
        return enforceTls;
    }
    public void setEnforceTls(String enforceTls) {
        this.enforceTls = enforceTls;
    }
    public String isNoIp() {
        return noIp;
    }
    public void setNoIp(String noIp) {
        this.noIp = noIp;
    }
    @Override
    public String toString() {
        return "Decision [enforceTls=" + enforceTls + ", noIp=" + noIp + "]";
    }
    @Override
    public int hashCode() {
        final int prime = 31;
        int result = 1;
        result = prime * result + ((enforceTls == null) ? 0 : enforceTls.hashCode());
        result = prime * result + ((noIp == null) ? 0 : noIp.hashCode());
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
        Decision other = (Decision) obj;
        if (enforceTls == null) {
            if (other.enforceTls != null)
                return false;
        } else if (!enforceTls.equals(other.enforceTls))
            return false;
        if (noIp == null) {
            if (other.noIp != null)
                return false;
        } else if (!noIp.equals(other.noIp))
            return false;
        return true;
    }
}
