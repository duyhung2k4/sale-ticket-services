try {
    const buildHelm = await $`helm upgrade --install iam -f k8s/values.yaml -n dev k8s`;
    console.log(buildHelm.stdout)
} catch (error) {
    console.log("ERROR: ", error);
}