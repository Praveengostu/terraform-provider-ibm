data "ibm_pi_instance_volumes" "voldata"
{
pi_instance_name="${var.instancename}"
pi_cloud_instance_id="${var.powerinstanceid}"
}



output "bootvolume"
{
value="${data.ibm_pi_instance_volumes.voldata.bootvolumeid}"
}

output "instance_volumes"
{
value="${data.ibm_pi_instance_volumes.voldata.instance_volumes}"
}
