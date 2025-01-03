<p>Packages:</p>
<ul>
<li>
<a href="#metal.provider.extensions.gardener.cloud%2fv1alpha1">metal.provider.extensions.gardener.cloud/v1alpha1</a>
</li>
</ul>
<h2 id="metal.provider.extensions.gardener.cloud/v1alpha1">metal.provider.extensions.gardener.cloud/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains the metal provider API resources.</p>
</p>
Resource Types:
<ul><li>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>
</li><li>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig</a>
</li></ul>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig
</h3>
<p>
<p>CloudProfileConfig contains provider-specific configuration that is embedded into Gardener&rsquo;s <code>CloudProfile</code>
resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
metal.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>CloudProfileConfig</code></td>
</tr>
<tr>
<td>
<code>machineImages</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.MachineImages">
[]MachineImages
</a>
</em>
</td>
<td>
<p>MachineImages is the list of machine images that are understood by the controller. It maps
logical names and versions to provider-specific identifiers.</p>
</td>
</tr>
<tr>
<td>
<code>regionConfigs</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.RegionConfig">
[]RegionConfig
</a>
</em>
</td>
<td>
<p>RegionConfigs is the list of supported regions.</p>
</td>
</tr>
<tr>
<td>
<code>machineTypes</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.MachineType">
[]MachineType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig
</h3>
<p>
<p>ControlPlaneConfig contains configuration settings for the control plane.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
metal.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>ControlPlaneConfig</code></td>
</tr>
<tr>
<td>
<code>cloudControllerManager</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CloudControllerManagerConfig">
CloudControllerManagerConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudControllerManager contains configuration settings for the cloud-controller-manager.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerConfig</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.LoadBalancerConfig">
LoadBalancerConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LoadBalancerConfig contains configuration settings for the shoot loadbalancing.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.AddressesFromNetworks">AddressesFromNetworks
</h3>
<p>
<p>AddressesFromNetworks is a reference to a network resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>key</code></br>
<em>
string
</em>
</td>
<td>
<p>Key is the name of metadata key for the network.</p>
</td>
</tr>
<tr>
<td>
<code>subnetRef</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.SubnetRef">
SubnetRef
</a>
</em>
</td>
<td>
<p>SubnetRef is a reference to the IP subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.BgpPeer">BgpPeer
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CalicoBgpConfig">CalicoBgpConfig</a>)
</p>
<p>
<p>BgpPeer contains configuration for BGPPeer resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>peerIP</code></br>
<em>
string
</em>
</td>
<td>
<p>PeerIP contains IP address of BGP peer followed by an optional port number to peer with.</p>
</td>
</tr>
<tr>
<td>
<code>asNumber</code></br>
<em>
int
</em>
</td>
<td>
<p>ASNumber contains the AS number of the BGP peer.</p>
</td>
</tr>
<tr>
<td>
<code>nodeSelector</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeSelector is a key-value pair to select nodes that should have this peering.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.CalicoBgpConfig">CalicoBgpConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.LoadBalancerConfig">LoadBalancerConfig</a>)
</p>
<p>
<p>CalicoBgpConfig contains BGP configuration settings for calico.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>asNumber</code></br>
<em>
int
</em>
</td>
<td>
<p>ASNumber is the default AS number used by a node.</p>
</td>
</tr>
<tr>
<td>
<code>serviceLoadBalancerIPs</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs.</p>
</td>
</tr>
<tr>
<td>
<code>serviceExternalIPs</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs.</p>
</td>
</tr>
<tr>
<td>
<code>serviceClusterIPs</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeer</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.BgpPeer">
[]BgpPeer
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>BGPPeer contains configuration for BGPPeer resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.CloudControllerManagerConfig">CloudControllerManagerConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig</a>)
</p>
<p>
<p>CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>featureGates</code></br>
<em>
map[string]bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>FeatureGates contains information about enabled feature gates.</p>
</td>
</tr>
<tr>
<td>
<code>networking</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CloudControllerNetworking">
CloudControllerNetworking
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Networking contains configuration settings for CCM networking.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.CloudControllerNetworking">CloudControllerNetworking
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CloudControllerManagerConfig">CloudControllerManagerConfig</a>)
</p>
<p>
<p>CloudControllerNetworking contains configuration settings for CCM networking.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>configureNodeAddresses</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigureNodeAddresses enables the configuration of node addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.IgnitionConfig">IgnitionConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.WorkerConfig">WorkerConfig</a>)
</p>
<p>
<p>IgnitionConfig contains ignition settings.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>raw</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Raw contains an inline ignition config, which is merged with the config from the os extension.</p>
</td>
</tr>
<tr>
<td>
<code>secretRef</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SecretRef is a reference to a secret containing the ignition config.</p>
</td>
</tr>
<tr>
<td>
<code>override</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Override configures, if ignition keys set by the os-extension are overridden
by extra ignition.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig
</h3>
<p>
<p>InfrastructureConfig infrastructure configuration resource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networks</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.Networks">
[]Networks
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Networks is the metal specific network configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.InfrastructureStatus">InfrastructureStatus
</h3>
<p>
<p>InfrastructureStatus contains information about created infrastructure resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.LoadBalancerConfig">LoadBalancerConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig</a>)
</p>
<p>
<p>LoadBalancerConfig contains configuration settings for the shoot loadbalancing.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metallbConfig</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.MetallbConfig">
MetallbConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>MetallbConfig contains configuration settings for metallb.</p>
</td>
</tr>
<tr>
<td>
<code>calicoBgpConfig</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CalicoBgpConfig">
CalicoBgpConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CalicoBgpConfig contains configuration settings for calico.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.MachineImage">MachineImage
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.WorkerStatus">WorkerStatus</a>)
</p>
<p>
<p>MachineImage is a mapping from logical names and versions to metal-specific identifiers.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the logical name of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version is the logical version of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>image</code></br>
<em>
string
</em>
</td>
<td>
<p>Image is the path to the image.</p>
</td>
</tr>
<tr>
<td>
<code>architecture</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Architecture is the CPU architecture of the machine image.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.MachineImageVersion">MachineImageVersion
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.MachineImages">MachineImages</a>)
</p>
<p>
<p>MachineImageVersion contains a version and a provider-specific identifier.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version is the version of the image.</p>
</td>
</tr>
<tr>
<td>
<code>image</code></br>
<em>
string
</em>
</td>
<td>
<p>Image is the path to the image.</p>
</td>
</tr>
<tr>
<td>
<code>architecture</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Architecture is the CPU architecture of the machine image.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.MachineImages">MachineImages
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>)
</p>
<p>
<p>MachineImages is a mapping from logical names and versions to provider-specific identifiers.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the logical name of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>versions</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.MachineImageVersion">
[]MachineImageVersion
</a>
</em>
</td>
<td>
<p>Versions contains versions and a provider-specific identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.MachineType">MachineType
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>serverLabels</code></br>
<em>
map[string]string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.MetallbConfig">MetallbConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.LoadBalancerConfig">LoadBalancerConfig</a>)
</p>
<p>
<p>MetallbConfig contains configuration settings for metallb.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipAddressPool</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPAddressPool contains IP address pools for metallb.</p>
</td>
</tr>
<tr>
<td>
<code>enableSpeaker</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableSpeaker enables the metallb speaker.</p>
</td>
</tr>
<tr>
<td>
<code>enableL2Advertisement</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableL2Advertisement enables L2 advertisement.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.Networks">Networks
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig</a>)
</p>
<p>
<p>Networks holds information about the Kubernetes and infrastructure networks.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the name for this CIDR.</p>
</td>
</tr>
<tr>
<td>
<code>cidr</code></br>
<em>
string
</em>
</td>
<td>
<p>CIDR is the workers subnet range to create.</p>
</td>
</tr>
<tr>
<td>
<code>id</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ID is the ID for the workers&rsquo; subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.RegionConfig">RegionConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>)
</p>
<p>
<p>RegionConfig is the definition of a region.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the name of a region.</p>
</td>
</tr>
<tr>
<td>
<code>server</code></br>
<em>
string
</em>
</td>
<td>
<p>Server is the server endpoint of this region.</p>
</td>
</tr>
<tr>
<td>
<code>certificateAuthorityData</code></br>
<em>
[]byte
</em>
</td>
<td>
<p>CertificateAuthorityData is the CA data of the region server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.SubnetRef">SubnetRef
</h3>
<p>
(<em>Appears on:</em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.AddressesFromNetworks">AddressesFromNetworks</a>)
</p>
<p>
<p>SubnetRef is a reference to the IP subnet.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the network.</p>
</td>
</tr>
<tr>
<td>
<code>apiGroup</code></br>
<em>
string
</em>
</td>
<td>
<p>APIGroup is the group of the IP pool</p>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
<em>
string
</em>
</td>
<td>
<p>Kind is the kind of the IP pool</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.WorkerConfig">WorkerConfig
</h3>
<p>
<p>WorkerConfig contains configuration settings for the worker nodes.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extraIgnition</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.IgnitionConfig">
IgnitionConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ExtraIgnition contains additional Ignition for Worker nodes.</p>
</td>
</tr>
<tr>
<td>
<code>extraServerLabels</code></br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ExtraServerLabels is a map of additional labels that are applied to the ServerClaim for Server selection.</p>
</td>
</tr>
<tr>
<td>
<code>addressesFromNetworks</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.*github.com/ironcore-dev/gardener-extension-provider-metal/pkg/apis/metal/v1alpha1.AddressesFromNetworks">
[]*github.com/ironcore-dev/gardener-extension-provider-metal/pkg/apis/metal/v1alpha1.AddressesFromNetworks
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>AddressesFromNetworks is a list of references to Network resources that should be used to assign IP addresses to the worker nodes.</p>
</td>
</tr>
<tr>
<td>
<code>metaData</code></br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>MedaData is a key-value map of additional data which should be passed to the Machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="metal.provider.extensions.gardener.cloud/v1alpha1.WorkerStatus">WorkerStatus
</h3>
<p>
<p>WorkerStatus contains information about created worker resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>machineImages</code></br>
<em>
<a href="#metal.provider.extensions.gardener.cloud/v1alpha1.MachineImage">
[]MachineImage
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>MachineImages is a list of machine images that have been used in this worker. Usually, the extension controller
gets the mapping from name/version to the provider-specific machine image data in its componentconfig. However, if
a version that is still in use gets removed from this componentconfig it cannot reconcile anymore existing <code>Worker</code>
resources that are still using this version. Hence, it stores the used versions in the provider status to ensure
reconciliation is possible.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <a href="https://github.com/ahmetb/gen-crd-api-reference-docs">gen-crd-api-reference-docs</a>
</em></p>
