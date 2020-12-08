import * as cdk from "@aws-cdk/core";
import * as eks from "@aws-cdk/aws-eks";
import * as ecr from "@aws-cdk/aws-ecr";

export class MikrodStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const simpleServiceRepo = new ecr.Repository(this, "SimpleServiceRepository", {
      repositoryName: "simple-service",
    });

    const complexServiceRepo = new ecr.Repository(this, "ComplexServiceRepository", {
      repositoryName: "complex-service",
    });

    const cluster = new eks.Cluster(this, "MikrodCluster", {
      version: eks.KubernetesVersion.V1_18,
      clusterName: "MikrodCluster",
      defaultCapacity: 2,
      outputClusterName: true,
    });
  }
}
