package io.conceal.webhook.spring.example.models;

public class Decision {
    private boolean enforceTls;
    private boolean noIp;
    public Decision() {
    }
    public boolean isEnforceTls() {
        return enforceTls;
    }
    public void setEnforceTls(boolean enforceTls) {
        this.enforceTls = enforceTls;
    }
    public boolean isNoIp() {
        return noIp;
    }
    public void setNoIp(boolean noIp) {
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
        result = prime * result + (enforceTls ? 1231 : 1237);
        result = prime * result + (noIp ? 1231 : 1237);
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
        if (enforceTls != other.enforceTls)
            return false;
        if (noIp != other.noIp)
            return false;
        return true;
    }
}
